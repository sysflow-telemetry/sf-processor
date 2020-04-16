package main

import (
	"fmt"
	"sync"

	hdl "github.com/sysflow-telemetry/sf-apis/go/handlers"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
)

type Example struct{}

func NewExample() sp.SFProcessor {
	return new(Example)
}

func (s *Example) Process(ch interface{}, wg *sync.WaitGroup) {
	cha := ch.(*hdl.FlatChannel)
	record := cha.In
	fmt.Println("Example channel capacity:", cap(record))
	defer wg.Done()
	fmt.Println("Starting Example")
	for {
		fc, ok := <-record
		if !ok {
			fmt.Println("Channel closed. Shutting down.")
			break
		}
		fmt.Println(fc)
	}
	fmt.Println("Exiting Example")
}

func (s *Example) SetOutChan(ch interface{}) {
}

func (s *Example) Cleanup() {
}
