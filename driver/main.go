//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/linkedin/goavro"
	"github.com/sysflow-telemetry/sf-apis/go/converter"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/goutils/logger"
	"github.ibm.com/sysflow/sf-processor/driver/pipeline"
)

// Driver constants
const (
	SockFile   = "/var/run/sysflow.sock"
	BuffSize   = 16384
	OOBuffSize = 1024
	PluginDir  = "../resources/plugins"
)

type inputType int

const (
	file inputType = iota
	socket
	winlog
)

func (it inputType) String() string {
	return [...]string{"file", "socket", "winlog"}[it]
}

func getFiles(filename string) ([]string, error) {
	var fls []string
	if fi, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	} else if fi.IsDir() {
		logger.Trace.Println("File is a directory")
		var files []os.FileInfo
		var err error
		if files, err = ioutil.ReadDir(filename); err != nil {
			return nil, err
		}
		for _, file := range files {
			f := filename + "/" + file.Name()
			logger.Trace.Println("File in Directory: " + f)
			fls = append(fls, f)
		}
		if len(fls) == 0 {
			return nil, errors.New("No files present in directory: " + filename)
		}

	} else {
		fls = append(fls, filename)
	}
	logger.Trace.Printf("Number of files in list: %d\n", len(fls))
	return fls, nil
}

func processInputFile(path string, pluginDir string, config string) {
	channel, pipeline, wg, channels, hdlers, err := pipeline.LoadPipeline(pluginDir, config)
	if err != nil {
		logger.Error.Println("pipeline error:", err)
		return
	}
	logger.Trace.Printf("Loaded %d stages\n", len(pipeline))
	logger.Trace.Printf("Loaded %d channels\n", len(channels))
	logger.Trace.Printf("Loaded %d hdlrs\n", len(hdlers))
	sfChannel := channel.(*plugins.SFChannel)

	records := sfChannel.In

	logger.Trace.Println("Loading file: ", flag.Arg(0))

	sfobjcvter := converter.NewSFObjectConverter()

	files, err := getFiles(flag.Arg(0))
	if err != nil {
		logger.Error.Println("files error: ", err)
		return
	}
	for _, fn := range files {
		logger.Trace.Println("Loading file: " + fn)
		f, err := os.Open(fn)
		if err != nil {
			logger.Error.Println("file open error: ", err)
			return
		}
		reader := bufio.NewReader(f)
		sreader, err := goavro.NewOCFReader(reader)
		if err != nil {
			logger.Error.Println("reader error: ", err)
			return
		}

		for sreader.Scan() {
			datum, err := sreader.Read()
			if err != nil {
				logger.Error.Println("datum reading error: ", err)
				return
			}

			records <- sfobjcvter.ConvertToSysFlow(datum)
		}
		f.Close()
	}
	logger.Trace.Println("Closing main channel")
	close(records)
	wg.Wait()
}

func processInputStream(path string, pluginDir string, config string) {
	if err := os.RemoveAll(path); err != nil {
		logger.Error.Println("remove error:", err)
		return
	}

	l, err := net.ListenUnix("unixpacket", &net.UnixAddr{path, "unixpacket"})
	if err != nil {
		logger.Error.Println("listen error:", err)
		return
	}
	defer l.Close()

	channel, pipeline, wg, channels, hdlers, err := pipeline.LoadPipeline(pluginDir, config)
	if err != nil {
		logger.Error.Println("pipeline error:", err)
		return
	}
	logger.Trace.Printf("Loaded %d stages\n", len(pipeline))
	logger.Trace.Printf("Loaded %d channels\n", len(channels))
	logger.Trace.Printf("Loaded %d hdlrs\n", len(hdlers))

	sFlow := sfgo.NewSysFlow()
	deser, err := compiler.CompileSchemaBytes([]byte(sFlow.Schema()), []byte(sFlow.Schema()))
	if err != nil {
		logger.Error.Println("compiler error:", err)
		return
	}

	sfChannel := channel.(*plugins.SFChannel)
	records := sfChannel.In

	for {
		buf := make([]byte, BuffSize)
		oobuf := make([]byte, OOBuffSize)
		reader := bytes.NewReader(buf)
		conn, err := l.AcceptUnix()
		if err != nil {
			logger.Error.Println("accept error:", err)
			break
		}
		for {
			sFlow = sfgo.NewSysFlow()
			_, _, flags, _, err := conn.ReadMsgUnix(buf[:], oobuf[:])
			if err != nil {
				logger.Error.Println("read error:", err)
				break
			}
			if flags == 0 {
				reader.Reset(buf)
				err = vm.Eval(reader, deser, sFlow)
				if err != nil {
					logger.Error.Println("deserialize:", err)
				}
				records <- sFlow
			} else {
				logger.Error.Println("Flag error ReadMsgUnix:", flags)
			}
		}
	}
	logger.Trace.Println("Closing main channel")
	close(records)
	wg.Wait()
}

func main() {
	// setup arg parsing
	inputType := flag.String("input", file.String(), fmt.Sprintf("Input type {%s|%s|%s}", file, socket, winlog))
	cpuprofile := flag.String("cpuprofile", "", "Write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "Write memory profile to `file`")
	configFile := flag.String("config", "pipeline.json", "Path to pipeline configuration file")
	logLevel := flag.String("log", "info", "Log level {trace|info|warn|error}")
	pluginDir := flag.String("plugdir", PluginDir, "Dynamic plugins directory")
	version := flag.Bool("version", false, "Outputs version information")

	flag.Usage = func() {
		fmt.Println("Usage: sfprocessor [[-version]|[-input <value>] [-log <value>] [-plugdir <value>] path]")
		fmt.Println()
		fmt.Println("Positional arguments:")
		fmt.Println("  path string\n\tInput path")
		fmt.Println()
		fmt.Println("Arguments:")
		flag.PrintDefaults()
		fmt.Println()
	}

	// parse args and validade positional args
	flag.Parse()
	if !*version && flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	if *version {
		hdr := sfgo.NewSFHeader()
		hdr.SetDefault(0)
		schemaVersion := hdr.Version
		fmt.Printf("Version: %s+%s, Avro Schema Version: %v, Export Schema Version: %v\n", Version, BuildNumber, schemaVersion, JSONSchemaVersion)
		os.Exit(0)
	}

	// retrieve positional args
	path := flag.Arg(0)

	// initialize logger
	logger.InitLoggers(logger.GetLogLevelFromValue(*logLevel))

	// CPU profiling
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// process input
	switch *inputType {
	case file.String():
		processInputFile(path, *pluginDir, *configFile)
		break
	case socket.String():
		processInputStream(path, *pluginDir, *configFile)
		break
	default:
		logger.Error.Println("Unrecognized input type: ", *inputType)
		os.Exit(1)
	}

	// memory profiling
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
