package main

import (
	"bufio"
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"log"
	"os"
	"sync"
)

const chanSize = 100000

func main() {
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

	f, _ := os.Open(os.Args[1])
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
