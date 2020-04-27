package syslogger

import (
	"sync"

	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.ibm.com/sysflow/sf-processor/common/logger"
	"github.ibm.com/sysflow/sf-processor/core/sfpe/engine"
)

// Syslogger defines a syslogger plugin.
type Syslogger struct{}

// NewSyslogger creates a new plugin instance.
func NewSyslogger() sp.SFProcessor {
	return new(Syslogger)
}

// Init initializes the plugin with a configuration map and cache.
func (s *Syslogger) Init(conf map[string]string, tables interface{}) error {
	return nil
}

// Process implements the main interface of the plugin.
func (s *Syslogger) Process(ch interface{}, wg *sync.WaitGroup) {
	cha := ch.(*engine.OccurenceChannel)
	record := cha.In
	logger.Trace.Println("Syslogger channel capacity:", cap(record))
	defer wg.Done()
	logger.Trace.Println("Starting Syslogger")
	for {
		fc, ok := <-record
		if !ok {
			logger.Trace.Println("Channel closed. Shutting down.")
			break
		}
		logger.Trace.Printf("Matched rules: %v", *fc)
	}
	logger.Trace.Println("Exiting Syslogger")
}

// SetOutChan sets the output channel of the plugin.
func (s *Syslogger) SetOutChan(ch interface{}) {}

// Cleanup tears down plugin resources.
func (s *Syslogger) Cleanup() {}

// This function is not run when module is used as a plugin.
func main() {}
