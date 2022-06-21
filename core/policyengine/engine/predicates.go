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

import (
	"fmt"
	"reflect"
	"strings"
)

// Predicate defines the type of a functional predicate.
type Predicate func(*Record) bool

// True defines a functional predicate that always returns true.
var True = Criterion{func(r *Record) bool { return true }}

// False defines a functional predicate that always returns false.
var False = Criterion{func(r *Record) bool { return false }}

// Criterion defines an interface for functional predicate operations.
type Criterion struct {
	Pred Predicate
}

// Eval evaluates a functional predicate.
func (c Criterion) Eval(r *Record) bool {
	return c.Pred(r)
}

// And computes the conjunction of two functional predicates.
func (c Criterion) And(cr Criterion) Criterion {
	var p Predicate = func(r *Record) bool { return c.Eval(r) && cr.Eval(r) }
	return Criterion{p}
}

// Or computes the conjunction of two functional predicates.
func (c Criterion) Or(cr Criterion) Criterion {
	var p Predicate = func(r *Record) bool { return c.Eval(r) || cr.Eval(r) }
	return Criterion{p}
}

// Not computes the negation of the function predicate.
func (c Criterion) Not() Criterion {
	var p Predicate = func(r *Record) bool { return !c.Eval(r) }
	return Criterion{p}
}

// All derives the conjuctive clause of all predicates in a slice of predicates.
func All(criteria []Criterion) Criterion {
	all := True
	for _, c := range criteria {
		all = all.And(c)
	}
	return all
}

// Any derives the disjuntive clause of all predicates in a slice of predicates.
func Any(criteria []Criterion) Criterion {
	any := False
	for _, c := range criteria {
		any = any.Or(c)
	}
	return any
}

// Exists creates a criterion for an existential predicate.
func Exists(attr string) Criterion {
	m := Mapper.Map(attr)
	p := func(r *Record) bool { return !reflect.ValueOf(m(r)).IsZero() }
	return Criterion{p}
}

// Eq creates a criterion for an equality predicate.
func Eq(lattr string, rattr string) Criterion {
	ml := Mapper.MapStr(lattr)
	mr := Mapper.MapStr(rattr)
	p := func(r *Record) bool { return eval(ml(r), mr(r), ops.eq) }
	return Criterion{p}
}

// NEq creates a criterion for an inequality predicate.
func NEq(lattr string, rattr string) Criterion {
	return Eq(lattr, rattr).Not()
}

// Ge creates a criterion for a greater-or-equal predicate.
func Ge(lattr string, rattr string) Criterion {
	ml := Mapper.MapInt(lattr)
	mr := Mapper.MapInt(rattr)
	p := func(r *Record) bool { return ml(r) >= mr(r) }
	return Criterion{p}
}

// Gt creates a criterion for a greater-than predicate.
func Gt(lattr string, rattr string) Criterion {
	ml := Mapper.MapInt(lattr)
	mr := Mapper.MapInt(rattr)
	p := func(r *Record) bool { return ml(r) > mr(r) }
	return Criterion{p}
}

// Le creates a criterion for a lower-or-equal predicate.
func Le(lattr string, rattr string) Criterion {
	return Gt(lattr, rattr).Not()
}

// Lt creates a criterion for a lower-than predicate.
func Lt(lattr string, rattr string) Criterion {
	return Ge(lattr, rattr).Not()
}

// StartsWith creates a criterion for a starts-with predicate.
func StartsWith(lattr string, rattr string) Criterion {
	ml := Mapper.MapStr(lattr)
	mr := Mapper.MapStr(rattr)
	p := func(r *Record) bool { return eval(ml(r), mr(r), ops.startswith) }
	return Criterion{p}
}

// EndsWith creates a criterion for a ends-with predicate.
func EndsWith(lattr string, rattr string) Criterion {
	ml := Mapper.MapStr(lattr)
	mr := Mapper.MapStr(rattr)
	p := func(r *Record) bool { return eval(ml(r), mr(r), ops.endswith) }
	return Criterion{p}
}

// Contains creates a criterion for a contains predicate.
func Contains(lattr string, rattr string) Criterion {
	ml := Mapper.MapStr(lattr)
	mr := Mapper.MapStr(rattr)
	p := func(r *Record) bool { return eval(ml(r), mr(r), ops.contains) }
	return Criterion{p}
}

// IContains creates a criterion for a case-insensitive contains predicate.
func IContains(lattr string, rattr string) Criterion {
	ml := Mapper.MapStr(lattr)
	mr := Mapper.MapStr(rattr)
	p := func(r *Record) bool { return eval(ml(r), mr(r), ops.icontains) }
	return Criterion{p}
}

// In creates a criterion for a list-inclusion predicate.
func In(attr string, list []string) Criterion {
	m := Mapper.MapStr(attr)
	p := func(r *Record) bool {
		for _, v := range list {
			if eval(m(r), v, ops.eq) {
				return true
			}
		}
		return false
	}
	return Criterion{p}
}

// PMatch creates a criterion for a list-pattern-matching predicate.
func PMatch(attr string, list []string) Criterion {
	m := Mapper.MapStr(attr)
	p := func(r *Record) bool {
		for _, v := range list {
			if eval(m(r), v, ops.contains) {
				return true
			}
		}
		return false
	}
	return Criterion{p}
}

// operator type.
type operator func(string, string) bool

// operators struct.
type operators struct {
	eq         operator
	contains   operator
	icontains  operator
	startswith operator
	endswith   operator
}

// ops defines boolean comparison operators over strings.
var ops = operators{
	eq:         func(l string, r string) bool { return l == r },
	contains:   func(l string, r string) bool { return strings.Contains(l, r) },
	icontains:  func(l string, r string) bool { return strings.Contains(strings.ToLower(l), strings.ToLower(r)) },
	startswith: func(l string, r string) bool { return strings.HasPrefix(l, r) },
	endswith:   func(l string, r string) bool { return strings.HasSuffix(l, r) },
}

// Eval evaluates a boolean operator over two predicates.
func eval(l interface{}, r interface{}, op operator) bool {
	lattrs := strings.Split(fmt.Sprintf("%v", l), LISTSEP)
	rattrs := strings.Split(fmt.Sprintf("%v", r), LISTSEP)
	for _, lattr := range lattrs {
		for _, rattr := range rattrs {
			if op(lattr, rattr) {
				return true
			}
		}
	}
	return false
}
