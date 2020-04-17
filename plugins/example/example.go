package main

import (
	"fmt"
	"sync"

	hdl "github.com/sysflow-telemetry/sf-apis/go/handlers"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
)

// Example defines an example plugin.
type Example struct{}

// NewExample creates a new plugin instance.
func NewExample() sp.SFProcessor {
	return new(Example)
}

// Process implements the main interface of the plugin.
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

// SetOutChan sets the output channel of the plugin.
func (s *Example) SetOutChan(ch interface{}) {
}

// Cleanup tears down plugin resources.
func (s *Example) Cleanup() {}

// This function is not run when module is used as a plugin.
func main() {}
