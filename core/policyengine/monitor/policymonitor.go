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

// Package monitor implements a policy monitor for the policy engine.
package monitor

import (
	"errors"

	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
)

// PolicyMonitor is an interface representing policy monitor objects.
// Currently the interface supports a local directory policy monitor.
type PolicyMonitor[R any] interface {
	GetInterpreterChan() chan *engine.PolicyInterpreter[R]
	StartMonitor() error
	StopMonitor() error
	CheckForPolicyUpdate() error
}

// NewPolicyMonitor creates a new policy monitor based on the engine configuration.
func NewPolicyMonitor[R any](config engine.Config, createInter func() (*engine.PolicyInterpreter[R], error), out func(R)) (PolicyMonitor[R], error) {
	if config.Monitor == engine.LocalType {
		return NewLocalPolicyMonitor(config, createInter, out)
	}
	return nil, errors.New("Policy monitor of type: " + config.Monitor.String() + " is not supported.")
}
