package sfpe

import (
	"errors"
	"sync"

	hdl "github.com/sysflow-telemetry/sf-apis/go/handlers"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.ibm.com/sysflow/sf-processor/common/ioutils"
	"github.ibm.com/sysflow/sf-processor/common/logger"
	"github.ibm.com/sysflow/sf-processor/core/cache"
	"github.ibm.com/sysflow/sf-processor/core/sfpe/engine"
)

// PolicyEngine defines a driver for the Policy Engine plugin.
type PolicyEngine struct {
	pi     engine.PolicyInterpreter
	tables *cache.SFTables
	ch     chan *engine.Record
}

// NewPolicyEngine constructs a new Policy Engine plugin.
func NewPolicyEngine() sp.SFProcessor {
	return new(PolicyEngine)
}

// NewEventChan creates a new event record channel instance.
func NewEventChan(size int) interface{} {
	return &engine.RecordChannel{In: make(chan *engine.Record, size)}
}

// Init initializes the plugin.
func (s *PolicyEngine) Init(conf map[string]string, tables interface{}) error {
	s.pi = engine.NewPolicyInterpreter(conf)
	s.tables = tables.(*cache.SFTables)
	if path, ok := conf[engine.PoliciesConfigKey]; ok {
		logger.Trace.Println("Loading policies from: ", path)
		paths, err := ioutils.ListFilePaths(path, ".yaml")
		if err == nil {
			if len(paths) == 0 {
				return errors.New("No policy files with extension .yaml found in path: " + path)
			}
			return s.pi.Compile(paths...)
		}
		return errors.New("Error while listing policies: " + err.Error())
	}
	return errors.New("Configuration tag 'policies' missing from policy engine plugin settings")
}

// Process implements the main loop of the plugin.
func (s *PolicyEngine) Process(ch interface{}, wg *sync.WaitGroup) {
	in := ch.(*hdl.FlatChannel).In
	defer wg.Done()
	logger.Trace.Println("Starting policy engine with capacity: ", cap(in))
	for {
		if fc, ok := <-in; ok {
			if match, r := s.pi.Process(true, engine.NewRecord(*fc, s.tables)); match {
				s.ch <- r
			}
		} else {
			logger.Trace.Println("Input channel closed. Shutting down.")
			break
		}
	}
	logger.Trace.Println("Exiting policy engine")
	s.Cleanup()
}

// SetOutChan sets the output channel of the plugin.
func (s *PolicyEngine) SetOutChan(ch interface{}) {
	s.ch = (ch.(*engine.RecordChannel)).In
}

// Cleanup clean up the plugin resources.
func (s *PolicyEngine) Cleanup() {
	close(s.ch)
}
