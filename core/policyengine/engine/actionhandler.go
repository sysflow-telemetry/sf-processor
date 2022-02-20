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
)

// Prototype of an action function
type ActionFunc func(r *Record) error

type ActionMap map[string]ActionFunc

// Action interface for user-defined actions
type Action interface {
	GetName() string
	GetFunc() ActionFunc
}

const ActionSym = "Action"

// Registers an action function
func registerAction(reg ActionMap, name string, f ActionFunc) {
	if _, ok := reg[name]; ok {
		logger.Warn.Println("Re-declaration of action '" + name + "'")
	}
	reg[name] = f
}

// LoadActions loads user-defined actions from path
func (ah *ActionHandler) loadUserActions(dir string) {
	ah.UserDefinedActions = make(ActionMap)
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
			action, ok := sym.(Action)
			if !ok {
				logger.Error.Println("Action symbol loaded from " + path + " must implement Action interface")
				continue
			}

			name := action.GetName()
			logger.Info.Println("Registering user-defined action '" + name + "'")
			registerAction(ah.UserDefinedActions, name, action.GetFunc())
		}
	}
}

type ActionHandler struct {
	// Map of registered actions
	BuiltInActions     ActionMap
	UserDefinedActions ActionMap
}

func NewActionHandler(conf Config) *ActionHandler {
	ah := new(ActionHandler)

	// Register built-in actions
	ah.registerBuiltIns()

	// Load user-defined actions
	ah.loadUserActions(conf.ActionDir)

	return ah
}

// CheckActions checks whether actions rules definitions have known implementations.
func (ah *ActionHandler) CheckActions(rules []Rule) {
	for _, r := range rules {
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
func (ah *ActionHandler) HandleActions(rule Rule, r *Record) {
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

// Registers built-in actions
func (ah *ActionHandler) registerBuiltIns() {
	ah.BuiltInActions = make(ActionMap)
}
