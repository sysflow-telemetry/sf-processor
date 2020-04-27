package main

import (
	"sync"

	hdl "github.com/sysflow-telemetry/sf-apis/go/handlers"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.ibm.com/sysflow/sf-processor/common/logger"
)

// Example defines an example plugin.
type Example struct{}

// NewExample creates a new plugin instance.
func NewExample() sp.SFProcessor {
	return new(Example)
}

// Init initializes the plugin with a configuration map and cache.
func (s *Example) Init(conf map[string]string, tables interface{}) error {
	return nil
}

// Process implements the main interface of the plugin.
func (s *Example) Process(ch interface{}, wg *sync.WaitGroup) {
	cha := ch.(*hdl.FlatChannel)
	record := cha.In
	logger.Trace.Println("Example channel capacity:", cap(record))
	defer wg.Done()
	logger.Trace.Println("Starting Example")
	for {
		fc, ok := <-record
		if !ok {
			logger.Trace.Println("Channel closed. Shutting down.")
			break
		}
		logger.Trace.Println(fc)
	}
	logger.Trace.Println("Exiting Example")
}

// SetOutChan sets the output channel of the plugin.
func (s *Example) SetOutChan(ch interface{}) {}

// Cleanup tears down plugin resources.
func (s *Example) Cleanup() {}

// This function is not run when module is used as a plugin.
func main() {}
