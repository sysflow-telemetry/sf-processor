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

// Package policy implements input policy translation for the rules engine.
package policy

// EnrichmentTag denotes the type for enrichment tags.
type EnrichmentTag interface{}

// TODO: consider moving the priority const definitions to the falco and sigma packages, if a one to one mapping to the languages priority is desired.
// Priority denotes the type for rule priority.
type Priority int

// Priority enumeration.
const (
	Low Priority = iota
	Medium
	High
)

// String returns the string representation of a priority instance.
func (p Priority) String() string {
	return [...]string{"low", "medium", "high"}[p]
}

// Rule type
type Rule[R any] struct {
	Name      string
	Desc      string
	Condition Criterion[R]
	Actions   []string
	Tags      []EnrichmentTag
	Priority  Priority
	Prefilter []string
	Enabled   bool
	IsAlert   bool
}

// Filter type
type Filter[R any] struct {
	Name      string
	Condition Criterion[R]
	Enabled   bool
}
