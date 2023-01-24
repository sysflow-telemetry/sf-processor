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

// Predicate defines the type of a functional predicate.
type Predicate[R any] func(R) bool

// Criterion defines an interface for functional predicate operations.
type Criterion[R any] struct {
	Pred Predicate[R]
}

// Eval evaluates a functional predicate.
func (c Criterion[R]) Eval(r R) bool {
	return c.Pred(r)
}

// And computes the conjunction of two functional predicates.
func (c Criterion[R]) And(cr Criterion[R]) Criterion[R] {
	var p Predicate[R] = func(r R) bool { return c.Eval(r) && cr.Eval(r) }
	return Criterion[R]{p}
}

// Or computes the conjunction of two functional predicates.
func (c Criterion[R]) Or(cr Criterion[R]) Criterion[R] {
	var p Predicate[R] = func(r R) bool { return c.Eval(r) || cr.Eval(r) }
	return Criterion[R]{p}
}

// Not computes the negation of the function predicate.
func (c Criterion[R]) Not() Criterion[R] {
	var p Predicate[R] = func(r R) bool { return !c.Eval(r) }
	return Criterion[R]{p}
}

// True defines a functional predicate that always returns true.
func True[R any]() Criterion[R] { return Criterion[R]{Pred: func(r R) bool { return true }} }

// False defines a functional predicate that always returns false.
func False[R any]() Criterion[R] { return Criterion[R]{Pred: func(r R) bool { return false }} }

// All derives the conjuctive clause of all predicates in a slice of predicates.
func All[R any](criteria []Criterion[R]) Criterion[R] {
	all := True[R]()
	for _, c := range criteria {
		all = all.And(c)
	}
	return all
}

// Any derives the disjuntive clause of all predicates in a slice of predicates.
func Any[R any](criteria []Criterion[R]) Criterion[R] {
	any := False[R]()
	for _, c := range criteria {
		any = any.Or(c)
	}
	return any
}
