package sfpe

import (
	"sync"

	hdl "github.com/sysflow-telemetry/sf-apis/go/handlers"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.ibm.com/sysflow/sf-processor/common/logger"
	"github.ibm.com/sysflow/sf-processor/driver/core/sfpe/engine"
)

// PolicyEngine defines a driver for the Policy Engine plugin.
type PolicyEngine struct {
	pi engine.PolicyInterpreter
}

// NewPolicyEngine constructs a new Policy Engine plugin.
//func NewPolicyEngine(paths ...string) sp.SFProcessor {
func NewPolicyEngine() sp.SFProcessor {
	pe := new(PolicyEngine)
	pe.pi.Compile("../../../tests/policies/macro_test.yaml") // Fix: pass paths from constructor args.
	//pe.pi.Compile(paths...)
	return pe
}

// Process implements the main loop of the plugin.
func (s *PolicyEngine) Process(ch interface{}, wg *sync.WaitGroup) {
	cha := ch.(*hdl.FlatChannel)
	record := cha.In
	logger.Trace.Println("Policy engine capacity: ", cap(record))
	defer wg.Done()
	logger.Trace.Println("Starting policy engine")
	for {
		fc, ok := <-record
		if !ok {
			logger.Trace.Println("Channel closed. Shutting down.")
			break
		}
		match, rlist := s.pi.Process(true, *fc)
		if match {
			logger.Trace.Printf("Matched rules: %v", rlist)
		}
	}
	logger.Trace.Println("Exiting policy engine")
}

// SetOutChan sets the output channel of the plugin.
func (s *PolicyEngine) SetOutChan(ch interface{}) {}

// Cleanup clean up the plugin resources.
func (s *PolicyEngine) Cleanup() {}
