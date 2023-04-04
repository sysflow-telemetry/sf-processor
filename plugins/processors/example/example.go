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
//
package main

import (
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/flattener"
)

const (
	pluginName string = "example"
)

// Plugin exports a symbol for this plugin.
var Plugin Example

// Example defines an example plugin.
type Example struct{}

// NewExample creates a new plugin instance.
func NewExample() plugins.SFProcessor {
	return new(Example)
}

// GetName returns the plugin name.
func (s *Example) GetName() string {
	return pluginName
}

// Init initializes the plugin with a configuration map.
func (s *Example) Init(conf map[string]interface{}) error {
	return nil
}

// Register registers plugin to plugin cache.
func (s *Example) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(pluginName, NewExample)
}

// Process implements the main interface of the plugin.
func (s *Example) Process(ch []interface{}, wg *sync.WaitGroup) {
	cha := ch[0].(*flattener.FlatChannel)
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
		if fc.Ints[sfgo.SYSFLOW_IDX][sfgo.SF_REC_TYPE] == sfgo.PROC_EVT {
			logger.Info.Printf("Process Event: %s, %d", fc.Strs[sfgo.SYSFLOW_IDX][sfgo.PROC_EXE_STR], fc.Ints[sfgo.SYSFLOW_IDX][sfgo.EV_PROC_TID_INT])
		}
	}
	logger.Trace.Println("Exiting Example")
}

// SetOutChan sets the output channel of the plugin.
func (s *Example) SetOutChan(ch []interface{}) {}

// Cleanup tears down plugin resources.
func (s *Example) Cleanup() {}

// This function is not run when module is used as a plugin.
func main() {}
