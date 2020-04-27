package syslogger

import (
	"sync"

	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.ibm.com/sysflow/sf-processor/common/logger"
	"github.ibm.com/sysflow/sf-processor/core/sfpe/engine"
)

const (
	maxBuffer int64 = 100
)

// Syslogger defines a syslogger plugin.
type Syslogger struct {
	occs    []*engine.Occurence
	counter int64
}

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
			s.exportOffenses()
			logger.Trace.Println("Channel closed. Shutting down.")
			break
		}
		s.counter++
		s.occs = append(s.occs, fc)
		if s.counter > maxBuffer {
			s.exportOffenses()
			s.counter = 0
		}
		// var rlist []string
		// for _, rule := range fc.Rules {
		// 	rlist = append(rlist, rule.Name)
		// }
		// logger.Trace.Printf("\033[1;36m%v\033[0m \033[1;34m%v\033[0m", rlist, fc.Record.Fr)
	}
	logger.Trace.Println("Exiting Syslogger")
}

func (s *Syslogger) exportOffenses() {
	offenses := CreateOffenses(s.occs)
	for _, o := range offenses {
		logger.Trace.Printf("\033[1;34m%v\033[0m\n", o.ToJSONStr())
	}
}

// SetOutChan sets the output channel of the plugin.
func (s *Syslogger) SetOutChan(ch interface{}) {}

// Cleanup tears down plugin resources.
func (s *Syslogger) Cleanup() {}

// This function is not run when module is used as a plugin.
func main() {}
