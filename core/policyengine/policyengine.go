//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package policyengine implements a plugin for a rules engine for telemetry records.
package policyengine

import (
	"errors"
	"sync"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/cache"
	"github.com/sysflow-telemetry/sf-processor/core/flattener"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/monitor"
)

const (
	pluginName  string = "policyengine"
	channelName string = "eventchan"
)

// PolicyEngine defines a driver for the Policy Engine plugin.
type PolicyEngine struct {
	pi            *engine.PolicyInterpreter
	tables        *cache.SFTables
	outCh         []chan *engine.Record
	filterOnly    bool
	bypass        bool
	config        engine.Config
	policyMonitor monitor.PolicyMonitor
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

func (s *PolicyEngine) compilePolicies(dir string) error {
	logger.Info.Println("Loading policies from: ", dir)
	paths, err := ioutils.ListFilePaths(dir, ".yaml")
	s.pi = engine.NewPolicyInterpreter(s.config)
	if err == nil {
		if len(paths) == 0 {
			return errors.New("no policy files with extension .yaml found in path: " + dir)
		}
		return s.pi.Compile(paths...)
	}
	return err
}

// Init initializes the plugin.
func (s *PolicyEngine) Init(conf map[string]interface{}) error {
	config, err := engine.CreateConfig(conf)
	if err != nil {
		return err
	}
	s.config = config
	s.tables = cache.GetInstance()

	if s.config.Mode == engine.BypassMode {
		logger.Trace.Println("Setting policy engine in bypass mode")
		s.bypass = true
		return nil
	}

	if s.config.PoliciesPath == sfgo.Zeros.String {
		return errors.New("Configuration tag 'policies' missing from policy engine plugin settings") 
	}
	if s.config.Mode == engine.FilterMode {
		logger.Trace.Println("Setting policy engine in filter mode")
		s.filterOnly = true
	} else {
		logger.Trace.Println("Setting policy engine in alert mode")
	}

	if s.config.Monitor == engine.NoneType {
		err = s.compilePolicies(s.config.PoliciesPath)
		if err != nil {
			logger.Error.Printf("Unable to compile local policies from directory %s, %v", s.config.PoliciesPath, err)
			return err
		}
	} else {
		pm, err := monitor.NewPolicyMonitor(s.config)
		if err != nil {
			logger.Error.Printf("Unable to load policy monitor %s, %v", config.Monitor.String(), err)
			return err
		}
		s.policyMonitor = pm
		s.policyMonitor.CheckForPolicyUpdate()
		select {
		case s.pi = <-s.policyMonitor.GetInterpreterChan():
			logger.Info.Printf("Loaded policy engine from policy monitor %s.", s.config.Monitor.String())
		default:
			logger.Error.Printf("No policy engine available for plugin.  Please check error logs for details.")
			return errors.New("no policy engine available for plugin.  Please check error logs for details")
		}
	}
	return nil
}

// Process implements the main loop of the plugin.
func (s *PolicyEngine) Process(ch interface{}, wg *sync.WaitGroup) {
	in := ch.(*flattener.FlatChannel).In
	defer wg.Done()
	logger.Trace.Println("Starting policy engine with capacity: ", cap(in))
	out := func(r *engine.Record) {
		for _, c := range s.outCh {
			c <- r
		}
	}
	start := time.Now()
	expiration := start.Add(20 * time.Second)
	if s.policyMonitor != nil {
		s.policyMonitor.StartMonitor()
	}

	for {
		if fc, ok := <-in; ok {
			if s.bypass {
				out(engine.NewRecord(*fc, s.tables))
			} else if s.policyMonitor != nil {
				now := time.Now()
				if now.After(expiration) {
					select {
					case s.pi = <-s.policyMonitor.GetInterpreterChan():
						logger.Info.Println("Updated policy engine in main policy engine thread.")
					default:
					}
					expiration = now.Add(20 * time.Second)
				}
				s.pi.ProcessAsync(true, s.filterOnly, engine.NewRecord(*fc, s.tables), out)
			} else {
				s.pi.ProcessAsync(true, s.filterOnly, engine.NewRecord(*fc, s.tables), out)
			}
		} else {
			logger.Trace.Println("Input channel closed. Shutting down.")
			break
		}
	}
}

// SetOutChan sets the output channel of the plugin.
func (s *PolicyEngine) SetOutChan(ch []interface{}) {
	for _, c := range ch {
		s.outCh = append(s.outCh, (c.(*engine.RecordChannel)).In)
	}
}

// Cleanup clean up the plugin resources.
func (s *PolicyEngine) Cleanup() {
	logger.Trace.Println("Exiting ", pluginName)
	if s.outCh != nil {
		for _, c := range s.outCh {
			close(c)
		}
	}
	if s.policyMonitor != nil {
		s.policyMonitor.StopMonitor()
	}
}
