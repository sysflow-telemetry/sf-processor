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

// Prefilter defines a prefilter object
type Prefilter struct{}

func NewPrefilter() source.Prefilter[*Record] {
	return &Prefilter{}
}

func (s *Prefilter) IsApplicable(r *Record, rule policy.Rule[*Record]) bool {
	if rule.Prefilter == nil || len(rule.Prefilter) == 0 {
		return true
	}
	rtype := Mapper.MapStr(SF_TYPE)(r)
	for _, pf := range rule.Prefilter {
		if rtype == pf {
			return true
		}
	}
	return false
}
