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

// Operations interface defines a set of predicates to satisfy rule operations.
type Operations[R any] interface {
	// Exists creates a criterion for an existential predicate.
	Exists(attr string) Criterion[R]
	// Eq creates a criterion for an equal predicate.
	Eq(lattr string, rattr string) Criterion[R]
	// NEq creates a criterion for a not-equal predicate.
	NEq(lattr string, rattr string) Criterion[R]
	// GEq creates a criterion for a greater-or-equal predicate.
	GEq(lattr string, rattr string) Criterion[R]
	// Gt creates a criterion for a greater-than predicate.
	Gt(lattr string, rattr string) Criterion[R]
	// LEq creates a criterion for a lower-or-equal predicate.
	LEq(lattr string, rattr string) Criterion[R]
	// Lt creates a criterion for a lower-than predicate.
	Lt(lattr string, rattr string) Criterion[R]
	// StartsWith creates a criterion for a starts-with predicate.
	StartsWith(lattr string, rattr string) Criterion[R]
	// EndsWith creates a criterion for a ends-with predicate.
	EndsWith(lattr string, rattr string) Criterion[R]
	// Contains creates a criterion for a contains predicate.
	Contains(lattr string, rattr string) Criterion[R]
	// IContains creates a criterion for a case-insensitive contains predicate.
	IContains(lattr string, rattr string) Criterion[R]
	// In creates a criterion for a list-inclusion predicate.
	In(attr string, list []string) Criterion[R]
	// PMatch creates a criterion for a list-pattern-matching predicate.
	PMatch(attr string, list []string) Criterion[R]
}
