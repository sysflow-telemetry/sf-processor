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
package ground

import (
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.ibm.com/sysflow/sf-processor/core/policyengine/engine"
)

const pluginName = "ground"

// Grounder defines a plugin to consume any incoming messages and do nothing.
type Grounder struct{}

// NewGrounder constructs a new Ground plugin.
func NewGrounder() plugins.SFProcessor {
	return new(Grounder)
}

// GetName returns the plugin name.
func (_ *Grounder) GetName() string {
	return pluginName
}

// Register registers plugin to plugin cache.
func (_ *Grounder) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(pluginName, NewGrounder)
}

// Init initializes the plugin.
func (_ *Grounder) Init(conf map[string]string) error {
	return nil
}

// Process implements the main loop of the plugin.
func (g *Grounder) Process(inputCh interface{}, wg *sync.WaitGroup) {
	defer g.Cleanup()
	defer wg.Done()
	switch ch := inputCh.(type) {
	case *plugins.SFChannel:
		for _ = range ch.In { // consume and do nothing
		}
	case *engine.RecordChannel:
		for _ = range ch.In { // consume and do nothing
		}
	default:
		logger.Error.Fatalf("Invalid input channel type to ground: %T", inputCh)
	}
	logger.Trace.Println("Exiting pipeline ground")
}

// SetOutChan sets the output channel of the plugin.
func (g *Grounder) SetOutChan(ch interface{}) {}

// Cleanup clean up the plugin resources.
func (g *Grounder) Cleanup() {}
