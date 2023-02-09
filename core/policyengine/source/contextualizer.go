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

import (
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
)

type Contextualizer[R any] interface {
	// AddRule adds one or more rules matching a record.
	AddRules(r R, rules ...policy.Rule[R])
	// GetRules retrieves the list of stored rules associated with a record.
	GetRules(r R) []policy.Rule[R]
	// Adds one or more tags to a record.
	AddTags(r R, tags ...string)
	// GetTags retrieves the list of tags associated with a record.
	GetTags(r R) []string
}

// DefaultContextualizer is a default contextualizer object.
type DefaultContextualizer[R any] struct{}

func NewDefaultContextualizer[R any]() Contextualizer[R] {
	return &DefaultContextualizer[R]{}
}

// AddRule adds one or more rules matching a record.
func (s *DefaultContextualizer[R]) AddRules(r R, rules ...policy.Rule[R]) {}

// GetRules retrieves the list of stored rules associated with a record.
func (s *DefaultContextualizer[R]) GetRules(r R) []policy.Rule[R] { return nil }

// Adds one or more tags to a record.
func (s *DefaultContextualizer[R]) AddTags(r R, tags ...string) {}

// GetTags retrieves the list of tags associated with a record.
func (s *DefaultContextualizer[R]) GetTags(r R) []string { return nil }
