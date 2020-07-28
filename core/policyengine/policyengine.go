//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
package policyengine

import (
	"errors"
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.ibm.com/sysflow/goutils/ioutils"
	"github.ibm.com/sysflow/goutils/logger"
	"github.ibm.com/sysflow/sf-processor/core/cache"
	"github.ibm.com/sysflow/sf-processor/core/policyengine/engine"
)

const (
	pluginName  string = "policyengine"
	channelName string = "eventchan"
)

// PolicyEngine defines a driver for the Policy Engine plugin.
type PolicyEngine struct {
	pi         engine.PolicyInterpreter
	tables     *cache.SFTables
	outCh      chan *engine.Record
	filterOnly bool
	config     engine.Config
}

// NewPolicyEngine constructs a new Policy Engine plugin.
func NewPolicyEngine() plugins.SFProcessor {
	return new(PolicyEngine)
}

// GetName returns the plugin name.
func (s *PolicyEngine) GetName() string {
	return pluginName
}

// NewEventChan creates a new event record channel instance.
func NewEventChan(size int) interface{} {
	return &engine.RecordChannel{In: make(chan *engine.Record, size)}
}

// Register registers plugin to plugin cache.
func (s *PolicyEngine) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(pluginName, NewPolicyEngine)
	pc.AddChannel(channelName, NewEventChan)
}

// Init initializes the plugin.
func (s *PolicyEngine) Init(conf map[string]string) error {
	config, err := engine.CreateConfig(conf)
	if err != nil {
		return err
	}
	s.config = config
	s.pi = engine.NewPolicyInterpreter(s.config)
	s.tables = cache.GetInstance()
	if s.config.Mode == engine.FilterMode {
		logger.Trace.Println("Setting policy engine in filter mode")
		s.filterOnly = true
	}
	logger.Trace.Println("Loading policies from: ", config.PoliciesPath)
	paths, err := ioutils.ListFilePaths(config.PoliciesPath, ".yaml")
	if err == nil {
		if len(paths) == 0 {
			return errors.New("No policy files with extension .yaml found in path: " + config.PoliciesPath)
		}
		return s.pi.Compile(paths...)
	}
	return errors.New("Error while listing policies: " + err.Error())
}

// Process implements the main loop of the plugin.
func (s *PolicyEngine) Process(ch interface{}, wg *sync.WaitGroup) {
	in := ch.(*plugins.FlatChannel).In
	defer wg.Done()
	logger.Trace.Println("Starting policy engine with capacity: ", cap(in))
	out := func(r *engine.Record) { s.outCh <- r }
	for {
		if fc, ok := <-in; ok {
			s.pi.ProcessAsync(true, s.filterOnly, engine.NewRecord(*fc, s.tables), out)
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
	s.outCh = (ch.(*engine.RecordChannel)).In
}

// Cleanup clean up the plugin resources.
func (s *PolicyEngine) Cleanup() {
	close(s.outCh)
}
