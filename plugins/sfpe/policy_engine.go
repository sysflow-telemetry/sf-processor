package sfpe

import (
	hdl "github.com/sysflow-telemetry/sf-apis/go/handlers"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.ibm.com/sysflow/sf-processor/common/logger"
	"sync"
)

type PolicyEng struct {
}

func NewPolicyEngine() sp.SFProcessor {
	return new(PolicyEng)
}

func (s *PolicyEng) Process(ch interface{}, wg *sync.WaitGroup) {
	cha := ch.(*hdl.FlatChannel)
	record := cha.In
	logger.Trace.Println("Policy engine capacity:", cap(record))
	defer wg.Done()
	logger.Trace.Println("Starting Policy Engine")
	for {
		fc, ok := <-record
		if !ok {
			logger.Trace.Println("Channel closed. Shutting down.")
			break
		}
		logger.Trace.Println(fc)
	}
	logger.Trace.Println("Exiting PolicyEng")
}

func (s *PolicyEng) SetOutChan(ch interface{}) {
}

func (s *PolicyEng) Cleanup() {
}
