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

// Operations interface defines a set of predicates to satisfy rule operations.
type Operations[R any] interface {
	// Exists creates a criterion for an existential predicate.
	Exists(attr string) policy.Criterion[R]
	// CompareStr creates a criterion for a binary predicate over strings.
	CompareStr(lattr string, rattr string, op Operator[string]) policy.Criterion[R]
	// CompareInt creates a criterion for a binary predicate over integers.
	CompareInt(lattr string, rattr string, op Operator[int64]) policy.Criterion[R]
	// FoldAny creates a disjunctive criterion for a binary predicate over a list of strings.
	FoldAny(attr string, list []string, op Operator[string]) policy.Criterion[R]
	// FoldAll creates a conjunctive criterion for a binary predicate over a list of strings.
	FoldAll(attr string, list []string, op Operator[string]) policy.Criterion[R]
	// RegExp creates a criterion for a regular-expression predicate.
	RegExp(attr string, re string) policy.Criterion[R]
}
