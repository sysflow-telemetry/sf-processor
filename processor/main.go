package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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

func main() {
	// setup arg parsing
	inputType := flag.String("input", file.String(), fmt.Sprintf("Input type {%s|%s}", file, socket))

	flag.Usage = func() {
		fmt.Println("Usage: sysprocessor path")
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
		os.Exit(0)
	}

	// retrieve positional args
	path := flag.Arg(0)

	// Initialize logger
	logging.InitLoggers(logging.)

	fmt.Printf("path: %s, type: %s\n", path, *inputType)

	switch(*inputType){
	case file:
		break
	case socket:
		break:
	default:
		fmt.
	}

	os.Exit(0)
	var handler Printer
	var wg sync.WaitGroup
	processor := NewSysFlowProc(handler)
	records := make(chan *sfgo.SysFlow, chanSize)
	wg.Add(1)
	go processor.process(records, &wg)

	sFlow := sfgo.NewSysFlow()
	deser, err := compiler.CompileSchemaBytes([]byte(sFlow.Schema()), []byte(sFlow.Schema()))
	if err != nil {
		log.Fatal("compiler error:", err)
	}

	f, _ := os.Open(flag.Arg(0))
	reader := bufio.NewReader(f)
	sreader, err := container.NewReader(reader)
	if err != nil {
		log.Fatal("reader error:", err)
	}
	i := 0
	for {
		sFlow = sfgo.NewSysFlow()
		err = vm.Eval(sreader, deser, sFlow)
		if err != nil {
			log.Printf("deserialize: %s\n", err)
			break
		}
		i++
		records <- sFlow
	}
	close(records)
	wg.Wait()
	log.Printf("The number of sysflow objects parsed is: %d\n", i)
}
