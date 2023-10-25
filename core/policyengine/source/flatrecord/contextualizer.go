//
// Copyright (C) 2023 IBM Corporation.
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

// Package flatrecord implements a flatrecord source for the policy compilers.
package flatrecord

import (
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

type Contextualizer struct{}

func NewContextualizer() source.Contextualizer[*Record] {
	return &Contextualizer{}
}

func (s *Contextualizer) AddRules(r *Record, rules ...policy.Rule[*Record]) {
	r.Ctx.AddRules(rules...)
}

func (s *Contextualizer) GetRules(r *Record) []policy.Rule[*Record] {
	return r.Ctx.GetRules()
}

func (s *Contextualizer) AddTags(r *Record, tags ...string) {
	r.Ctx.AddTags(tags...)
}

func (s *Contextualizer) GetTags(r *Record) []string {
	return r.Ctx.GetTags()
}
