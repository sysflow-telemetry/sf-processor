//
// Copyright (C) 2021 IBM Corporation.
//
// Authors:
// Andreas Schade <san@zurich.ibm.com
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

// Package engine implements a rules engine for telemetry records.
package engine

import (
	"plugin"

	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
)

// Prototype of an action function
type ActionFunc[R any] func(r R) error

type ActionMap[R any] map[string]ActionFunc[R]

// Action interface for user-defined actions
type Action[R any] interface {
	GetName() string
	GetFunc() ActionFunc[R]
}

const ActionSym = "Action"

type ActionHandler[R any] struct {
	// Map of registered actions
	BuiltInActions     ActionMap[R]
	UserDefinedActions ActionMap[R]
}

func NewActionHandler[R any](conf Config) *ActionHandler[R] {
	ah := new(ActionHandler[R])

	// Register built-in actions
	ah.registerBuiltIns()

	// Load user-defined actions
	ah.loadUserActions(conf.ActionDir)

	return ah
}

// Registers built-in actions
func (ah *ActionHandler[R]) registerBuiltIns() {
	ah.BuiltInActions = make(ActionMap[R])
}

// LoadActions loads user-defined actions from path
func (ah *ActionHandler[R]) loadUserActions(dir string) {
	ah.UserDefinedActions = make(ActionMap[R])
	if paths, err := ioutils.ListFilePaths(dir, ".so"); err == nil {
		var plug *plugin.Plugin
		for _, path := range paths {
			logger.Info.Println("Loading user-defined action from file " + path)
			if plug, err = plugin.Open(path); err != nil {
				logger.Error.Println(err.Error())
				continue
			}
			sym, err := plug.Lookup(ActionSym)
			if err != nil {
				logger.Error.Println(err.Error())
				continue
			}
			action, ok := sym.(Action[R])
			if !ok {
				logger.Error.Println("Action symbol loaded from " + path + " must implement Action interface")
				continue
			}

			// Registers an action function
			name := action.GetName()
			logger.Info.Println("Registering user-defined action '" + name + "'")
			if _, ok := ah.UserDefinedActions[name]; ok {
				logger.Warn.Println("Re-declaration of action '" + name + "'")
			}
			ah.UserDefinedActions[name] = action.GetFunc()
		}
	}
}

// CheckActions checks whether actions rules definitions have known implementations.
func (ah *ActionHandler[R]) CheckActions(rules []policy.Rule[R]) {
	for _, r := range rules {
		if r.Actions == nil {
			continue
		}
		for _, a := range r.Actions {
			if _, ok := ah.BuiltInActions[a]; !ok {
				if _, ok = ah.UserDefinedActions[a]; !ok {
					logger.Warn.Printf("Unknown action identifier '%s' found in rule '%s'", a, r.Name)
				}
			}
		}
	}
}

// HandleAction handles actions defined in rule.
func (ah *ActionHandler[R]) HandleActions(rule policy.Rule[R], r R) {
	if rule.Actions == nil {
		return
	}
	for _, a := range rule.Actions {
		action, ok := ah.BuiltInActions[a]
		if !ok {
			action, ok = ah.UserDefinedActions[a]
		}
		if !ok {
			continue
		}
		if err := action(r); err != nil {
			logger.Error.Println("Error in action: " + err.Error())
		}
	}
}
