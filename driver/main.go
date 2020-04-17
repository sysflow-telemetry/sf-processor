package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/sysflow-telemetry/sf-apis/go/handlers"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/sf-processor/common/logger"
	"github.ibm.com/sysflow/sf-processor/driver/pipeline"
)

const (
	CHAN_SIZE    = 100000
	SOCK_FILE    = "/var/run/sysflow.sock"
	BUFF_SIZE    = 16384
	OO_BUFF_SIZE = 1024
)

type inputType int

const (
	file inputType = iota
	socket
)

func (it inputType) String() string {
	return [...]string{"file", "socket"}[it]
}

func processInputFile(path string) {
	/*var handler Printer
	var wg sync.WaitGroup
	processor := NewSysFlowProc(handler)
	records := make(chan *sfgo.SysFlow, CHAN_SIZE)
	wg.Add(1)
	go processor.process(records, &wg)
	*/

	channel, pipeline, wg, channels, hdlers, err := LoadPipeline()
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
	f, err := os.Open(flag.Arg(0))
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
			logger.Error.Println("deserialize: ", err)
			break
		}
		records <- sFlow
	}
	logger.Trace.Println("Closing main channel")
	close(records)
	wg.Wait()
}

func processInputStream(path string) {
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

	channel, pipeline, wg, channels, hdlers, err := LoadPipeline()
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
	/*var wg sync.WaitGroup
	wg.Add(1)
	var handler Printer
	processor := NewSysFlowProc(handler)
	records := make(chan *sfgo.SysFlow, CHAN_SIZE)
	go processor.process(records, &wg)
	*/

	sfChannel := channel.(*sp.SFChannel)

	records := sfChannel.In

	for {
		buf := make([]byte, BUFF_SIZE)
		oobuf := make([]byte, OO_BUFF_SIZE)
		reader := bytes.NewReader(buf)
		conn, err := l.AcceptUnix()
		if err != nil {
			logger.Error.Println("accept error:", err)
			close(records)
			return
		}
		for {
			sFlow = sfgo.NewSysFlow()
			//n, _ ,  err := conn.ReadFromUnix(buf[:])
			_, _, flags, _, err := conn.ReadMsgUnix(buf[:], oobuf[:])
			if err != nil {
				logger.Error.Println("read error:", err)
				close(records)
				break
			}
			if flags == 0 {
				reader.Reset(buf)
				//println("Server Received: ", string(buf[0:n]))
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

func LoadPipeline() (interface{}, []sp.SFProcessor, *sync.WaitGroup, []interface{}, []handlers.SFHandler, error) {
	pl := pipeline.NewPluginCache()
	wg := new(sync.WaitGroup)
	var processors []sp.SFProcessor
	var channels []interface{}
	var hdlrs []handlers.SFHandler
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
			err = prc.Init(p)
			if err != nil {
				logger.Error.Println(err)
				return nil, nil, wg, nil, nil, err
			}

		} else {
			logger.Error.Println("processor or handler tag must exist in plugin config")
			return nil, nil, wg, nil, nil, err
		}

		if v, o := p["in"]; o {
			in, err = pl.GetChan(mod, v, CHAN_SIZE)
			channels = append(channels, in)
			chp := fmt.Sprintf("%T", in)
			logger.Trace.Println(chp)
		} else {
			logger.Error.Println("in tag must exist in plugin config")
			return nil, nil, wg, nil, nil, errors.New("in tag must exist in plugin config")
		}
		if v, o := p["out"]; o {
			out, err = pl.GetChan(mod, v, CHAN_SIZE)
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

	// process input
	switch *inputType {
	case file.String():
		processInputFile(path)
		break
	case socket.String():
		processInputStream(path)
		break
	default:
		logger.Error.Println("Unrecognized input type: ", *inputType)
		os.Exit(1)
	}
}
