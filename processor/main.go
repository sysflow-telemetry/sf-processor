package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/sf-processor/common/logger"
)

const (
	chanSize = 100000
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
	var handler Printer
	var wg sync.WaitGroup
	processor := NewSysFlowProc(handler)
	records := make(chan *sfgo.SysFlow, chanSize)
	wg.Add(1)
	go processor.process(records, &wg)

	sFlow := sfgo.NewSysFlow()
	deser, err := compiler.CompileSchemaBytes([]byte(sFlow.Schema()), []byte(sFlow.Schema()))
	if err != nil {
		logger.Error.Println("compiler error: ", err)
	}

	f, _ := os.Open(flag.Arg(0))
	reader := bufio.NewReader(f)
	sreader, err := container.NewReader(reader)
	if err != nil {
		logger.Error.Println("reader error: ", err)
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
	close(records)
	wg.Wait()
}

func processInputStream(path string) {
	// TODO
	return
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
