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

// Package engine implements a rules engine for telemetry records.
package engine

// ActionHandler type
type ActionHandler struct {
	conf Config
}

// NewActionHandler creates a new handler.
func NewActionHandler(conf Config) ActionHandler {
	return ActionHandler{conf}
}

// HandleActionAsync handles actions defined in rule.
func (s ActionHandler) HandleActionAsync(rule Rule, r *Record, out func(r *Record)) {
	s.HandleAction(rule, r)
	//out(r)
}

// HandleAction handles actions defined in rule.
func (s ActionHandler) HandleAction(rule Rule, r *Record) {
	for _, a := range rule.Actions {
		switch a {
		case Tag:
			fallthrough
		case Alert:
			fallthrough
		default:
			r.Ctx.AddRule(rule)
		}
	}
}
