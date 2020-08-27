//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
	"strconv"
	"sync"

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
	ChanSize   = 100000
	SockFile   = "/var/run/sysflow.sock"
	BuffSize   = 16384
	OOBuffSize = 1024
	PluginDir  = "../resources/plugins"
)

type inputType int

const (
	file inputType = iota
	socket
)

func (it inputType) String() string {
	return [...]string{"file", "socket"}[it]
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
	channel, pipeline, wg, channels, hdlers, err := LoadPipeline(pluginDir, config)
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

	channel, pipeline, wg, channels, hdlers, err := LoadPipeline(pluginDir, config)
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

// setManifestInfo sets manifest attributes to plugins configuration items.
func setManifestInfo(conf *pipeline.Config) {
	addGlobalConfigItem(conf, VersionKey, Version)
	addGlobalConfigItem(conf, JSONSchemaVersionKey, JSONSchemaVersion)
	addGlobalConfigItem(conf, BuildNumberKey, BuildNumber)
}

// addGlobalConfigItem adds a config item to all processors in the pipeline.
func addGlobalConfigItem(conf *pipeline.Config, k string, v interface{}) {
	for _, c := range conf.Pipeline {
		if _, ok := c[pipeline.ProcConfig]; ok {
			if s, ok := v.(string); ok {
				c[k] = s
			} else if i, ok := v.(int); ok {
				c[k] = strconv.Itoa(i)
			}
		}
	}
}

// LoadPipeline sets up the an edge processing pipeline based on configuration settings.
func LoadPipeline(pluginDir string, config string) (interface{}, []plugins.SFProcessor, *sync.WaitGroup, []interface{}, []plugins.SFHandler, error) {
	pl := pipeline.NewPluginCache(config)
	wg := new(sync.WaitGroup)

	var processors []plugins.SFProcessor
	var channels []interface{}
	var hdlrs []plugins.SFHandler

	if err := pl.LoadPlugins(pluginDir); err != nil {
		logger.Error.Println("Unable to load dynamic plugins: ", err)
		return nil, nil, wg, nil, nil, err
	}

	conf, err := pl.GetConfig()
	if err != nil {
		logger.Error.Println("Unable to load pipeline config: ", err)
		return nil, nil, wg, nil, nil, err
	}
	setManifestInfo(conf)
	var in interface{}
	var out interface{}
	var first interface{}
	for idx, p := range conf.Pipeline {
		hdler := false
		var hdl plugins.SFHandler
		if val, ok := p[pipeline.HdlConfig]; ok {
			hdl, err = pl.GetHandler(val)
			if err != nil {
				logger.Error.Println(err)
				return nil, nil, wg, nil, nil, err
			}
			hdlrs = append(hdlrs, hdl)
			xType := fmt.Sprintf("%T", hdl)
			logger.Trace.Println(xType)
			hdler = true
		}
		var prc plugins.SFProcessor
		if val, ok := p[pipeline.ProcConfig]; ok {
			prc, err = pl.GetProcessor(val, hdl, hdler)
			if err != nil {
				logger.Error.Println(err)
				return nil, nil, wg, nil, nil, err
			}
			tp := fmt.Sprintf("%T", prc)
			logger.Trace.Println(tp)
			err = prc.Init(p)
			if err != nil {
				logger.Error.Println(err)
				return nil, nil, wg, nil, nil, err
			}
		} else {
			logger.Error.Println("processor or handler tag must exist in plugin config")
			return nil, nil, wg, nil, nil, err
		}
		if v, o := p[pipeline.InChanConfig]; o {
			in, err = pl.GetChan(v, ChanSize)
			channels = append(channels, in)
			chp := fmt.Sprintf("%T", in)
			logger.Trace.Println(chp)
		} else {
			logger.Error.Println("in tag must exist in plugin config")
			return nil, nil, wg, nil, nil, errors.New("in tag must exist in plugin config")
		}
		if v, o := p[pipeline.OutChanConfig]; o {
			out, err = pl.GetChan(v, ChanSize)
			chp := fmt.Sprintf("%T", out)
			channels = append(channels, out)
			logger.Trace.Println(chp)
			prc.SetOutChan(out)
		}
		processors = append(processors, prc)
		wg.Add(1)
		go prc.Process(in, wg)
		if idx == 0 {
			first = in
		}
	}
	return first, processors, wg, channels, hdlrs, nil
}

func main() {
	// setup arg parsing
	inputType := flag.String("input", file.String(), fmt.Sprintf("Input type {%s|%s}", file, socket))
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
