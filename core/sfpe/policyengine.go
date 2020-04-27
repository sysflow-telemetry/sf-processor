package sfpe

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	hdl "github.com/sysflow-telemetry/sf-apis/go/handlers"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.ibm.com/sysflow/sf-processor/common/logger"
	"github.ibm.com/sysflow/sf-processor/core/cache"
	"github.ibm.com/sysflow/sf-processor/core/sfpe/engine"
)

// PolicyEngine defines a driver for the Policy Engine plugin.
type PolicyEngine struct {
	pi     engine.PolicyInterpreter
	tables *cache.SFTables
	ch     chan *engine.Occurence
}

// NewPolicyEngine constructs a new Policy Engine plugin.
func NewPolicyEngine() sp.SFProcessor {
	pe := new(PolicyEngine)
	return pe
}

// NewOccurenceChan creates a new occurence channel instance.
func NewOccurenceChan(size int) interface{} {
	return &engine.OccurenceChannel{In: make(chan *engine.Occurence, size)}
}

// Init initializes the plugin.
func (s *PolicyEngine) Init(conf map[string]string, tables interface{}) error {
	s.tables = tables.(*cache.SFTables)
	if filename, ok := conf["policies"]; ok {
		logger.Trace.Println("Loading policies from: " + filename)
		if fi, err := os.Stat(filename); os.IsNotExist(err) {
			return err
		} else if fi.IsDir() {
			var files []os.FileInfo
			var err error
			if files, err = ioutil.ReadDir(filename); err != nil {
				return err
			}
			var fls []string
			for _, file := range files {
				if filepath.Ext(file.Name()) == ".yaml" {
					f := filename + "/" + file.Name()
					fls = append(fls, f)
				}
			}
			if len(fls) == 0 {
				return errors.New("No policy files with extension .yaml present in directory: " + filename)
			}
			s.pi.Compile(fls...)
		} else {
			s.pi.Compile(filename)
		}
	} else {
		return errors.New("policies tag missing from policy engine plugin")
	}
	return nil
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
		r := engine.NewRecord(*fc, s.tables)
		match, rlist := s.pi.Process(true, r)
		if match {
			//logger.Trace.Printf("Matched rules: %v", rlist)
			s.ch <- engine.NewOccurence(r, rlist)
		}
	}
	logger.Trace.Println("Exiting policy engine")
	s.Cleanup()
}

// SetOutChan sets the output channel of the plugin.
func (s *PolicyEngine) SetOutChan(ch interface{}) {
	s.ch = (ch.(*engine.OccurenceChannel)).In
}

// Cleanup clean up the plugin resources.
func (s *PolicyEngine) Cleanup() {
	close(s.ch)
}
