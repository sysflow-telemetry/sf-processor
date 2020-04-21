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
	"github.ibm.com/sysflow/sf-processor/core/sfpe/engine"
)

// PolicyEngine defines a driver for the Policy Engine plugin.
type PolicyEngine struct {
	pi engine.PolicyInterpreter
}

// NewPolicyEngine constructs a new Policy Engine plugin.
func NewPolicyEngine() sp.SFProcessor {
	pe := new(PolicyEngine)
	return pe
}

// Init initializes the plugin.
func (s *PolicyEngine) Init(conf map[string]string, tables interface{}) error {
	var filename string = ""
	if v, o := conf["policies"]; o {
		filename = v
	} else {
		return errors.New("policies tag missing from policy engine plugin")
	}
	logger.Trace.Println("Loading policies from: " + filename)
	var fi os.FileInfo
	var e error
	if fi, e = os.Stat(filename); os.IsNotExist(e) {
		return e
	}
	if fi.IsDir() {
		files, err := ioutil.ReadDir(filename)
		if err != nil {
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
		logger.Trace.Println("Compiling policies from: " + filename)
		s.pi.Compile(filename)
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
