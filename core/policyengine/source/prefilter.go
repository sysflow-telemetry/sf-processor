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

package source

import "github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"

// Prefilter interface
type Prefilter[R any] interface {
	IsApplicable(r R, rule policy.Rule[R]) bool
}

// DefaultPrefilter defines a prefilter object to be used as a default prefilter.
type DefaultPrefilter[R any] struct{}

func NewDefaultPrefilter[R any]() Prefilter[R] {
	return &DefaultPrefilter[R]{}
}

func (s *DefaultPrefilter[R]) IsApplicable(r R, rule policy.Rule[R]) bool {
	return true
}
