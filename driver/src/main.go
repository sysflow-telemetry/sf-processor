package main

import (
	"log"
	"os"

	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

const chanSize = 100000

func main() {
	var handler Printer
	processor := NewSysFlowProc(handler)
	records := make(chan *sfgo.SysFlow, chanSize)
	go processor.process(records)

	sFlow := sfgo.NewSysFlow()
	deser, err := compiler.CompileSchemaBytes([]byte(sFlow.Schema()), []byte(sFlow.Schema()))
	if err != nil {
		log.Fatal("compiler error:", err)
	}

	reader, _ := os.Open("../../tests/traces/tcp.sf")
	for {
		sFlow = sfgo.NewSysFlow()
		err = vm.Eval(reader, deser, sFlow)
		if err != nil {
			log.Fatal("deserialize:", err)
		}
		records <- sFlow
	}
}
