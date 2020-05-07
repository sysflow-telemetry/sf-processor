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
	"sync"

	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/sysflow-telemetry/sf-apis/go/handlers"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/sf-processor/common/logger"
	"github.ibm.com/sysflow/sf-processor/core/cache"
	"github.ibm.com/sysflow/sf-processor/driver/pipeline"
)

// Driver constants
const (
	ChanSize   = 100000
	SockFile   = "/var/run/sysflow.sock"
	BuffSize   = 16384
	OOBuffSize = 1024
	CacheSize  = 2
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

func processInputFile(path string, config string) {
	channel, pipeline, wg, channels, hdlers, err := LoadPipeline(config)
	if err != nil {
		logger.Error.Println("pipeline error:", err)
		return
	}
	logger.Trace.Printf("Loaded %d stages\n", len(pipeline))
	logger.Trace.Printf("Loaded %d channels\n", len(channels))
	logger.Trace.Printf("Loaded %d hdlrs\n", len(hdlers))
	sfChannel := channel.(*sp.SFChannel)

	records := sfChannel.In

	sFlow := sfgo.NewSysFlow()
	deser, err := compiler.CompileSchemaBytes([]byte(sFlow.Schema()), []byte(sFlow.Schema()))
	if err != nil {
		logger.Error.Println("compiler error: ", err)
	}
	logger.Trace.Println("Loading file: ", flag.Arg(0))

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
		sreader, err := container.NewReader(reader)
		if err != nil {
			logger.Error.Println("reader error: ", err)
			return
		}

		for {
			sFlow = sfgo.NewSysFlow()
			err = vm.Eval(sreader, deser, sFlow)
			if err != nil {
				if err.Error() != "EOF" {
					logger.Error.Println("deserialize: ", err)
				}
				break
			}
			records <- sFlow
		}
		f.Close()
	}
	logger.Trace.Println("Closing main channel")
	close(records)
	wg.Wait()
}

func processInputStream(path string, config string) {
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

	channel, pipeline, wg, channels, hdlers, err := LoadPipeline(config)
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

	sfChannel := channel.(*sp.SFChannel)
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

// LoadPipeline sets up the an edge processing pipeline based on configuration settings.
func LoadPipeline(config string) (interface{}, []sp.SFProcessor, *sync.WaitGroup, []interface{}, []handlers.SFHandler, error) {
	pl := pipeline.NewPluginCache(config)
	wg := new(sync.WaitGroup)
	var processors []sp.SFProcessor
	var channels []interface{}
	var hdlrs []handlers.SFHandler
	tables := cache.NewSFTables(CacheSize)
	conf, err := pl.GetConfig()
	if err != nil {
		logger.Error.Println("Unable to load pipeline config: ", err)
		return nil, nil, wg, nil, nil, err
	}
	var in interface{}
	var out interface{}
	var first interface{}
	for idx, p := range conf.Pipeline {
		hdler := false
		var hdl handlers.SFHandler
		mod := p["mod"]
		if val, ok := p["handler"]; ok {
			hdl, err = pl.GetHandler(mod, val)
			if err != nil {
				logger.Error.Println(err)
				return nil, nil, wg, nil, nil, err
			}
			hdlrs = append(hdlrs, hdl)
			xType := fmt.Sprintf("%T", hdl)
			logger.Trace.Println(xType)
			hdler = true
		}
		var prc sp.SFProcessor
		if val, ok := p["processor"]; ok {
			prc, err = pl.GetProcessor(mod, val, hdl, hdler)
			if err != nil {
				logger.Error.Println(err)
				return nil, nil, wg, nil, nil, err
			}
			tp := fmt.Sprintf("%T", prc)
			logger.Trace.Println(tp)
			err = prc.Init(p, tables)
			if err != nil {
				logger.Error.Println(err)
				return nil, nil, wg, nil, nil, err
			}
		} else {
			logger.Error.Println("processor or handler tag must exist in plugin config")
			return nil, nil, wg, nil, nil, err
		}
		if v, o := p["in"]; o {
			in, err = pl.GetChan(mod, v, ChanSize)
			channels = append(channels, in)
			chp := fmt.Sprintf("%T", in)
			logger.Trace.Println(chp)
		} else {
			logger.Error.Println("in tag must exist in plugin config")
			return nil, nil, wg, nil, nil, errors.New("in tag must exist in plugin config")
		}
		if v, o := p["out"]; o {
			out, err = pl.GetChan(mod, v, ChanSize)
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
	configFile := flag.String("config", "/usr/local/sf-processor/conf/pipeline.json", "Path to pipeline configuration file")

	flag.Usage = func() {
		fmt.Println("Usage: sysprocessor [-input <value>] path")
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
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	// retrieve positional args
	path := flag.Arg(0)

	// Initialize logger
	logger.InitLoggers(logger.TRACE)

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
		processInputFile(path, *configFile)
		break
	case socket.String():
		processInputStream(path, *configFile)
		break
	default:
		logger.Error.Println("Unrecognized input type: ", *inputType)
		os.Exit(1)
	}

	// Memory profiling
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
