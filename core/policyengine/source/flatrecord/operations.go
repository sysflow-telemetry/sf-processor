//
// Copyright (C) 2022 IBM Corporation.
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

// Package flatrecord implements a flatten record source for the rules policy.
package flatrecord

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/sysflow-telemetry/sf-processor/core/policyengine/common"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

type Operations struct{}

func NewOperations() source.Operations[*Record] {
	return &Operations{}
}

// Exists creates a criterion for an existential predicate.
func (op *Operations) Exists(attr string) policy.Criterion[*Record] {
	m := Mapper.Map(attr)
	p := func(r *Record) bool { return !reflect.ValueOf(m(r)).IsZero() }
	return policy.Criterion[*Record]{Pred: p}
}

// Eq creates a criterion for an equality predicate.
func (op *Operations) Eq(lattr string, rattr string) policy.Criterion[*Record] {
	ml := Mapper.MapStr(lattr)
	mr := Mapper.MapStr(rattr)
	p := func(r *Record) bool { return eval(ml(r), mr(r), ops.eq) }
	return policy.Criterion[*Record]{Pred: p}
}

// NEq creates a criterion for an inequality predicate.
func (op *Operations) NEq(lattr string, rattr string) policy.Criterion[*Record] {
	return op.Eq(lattr, rattr).Not()
}

// Ge creates a criterion for a greater-or-equal predicate.
func (op *Operations) GEq(lattr string, rattr string) policy.Criterion[*Record] {
	ml := Mapper.MapInt(lattr)
	mr := Mapper.MapInt(rattr)
	p := func(r *Record) bool { return ml(r) >= mr(r) }
	return policy.Criterion[*Record]{Pred: p}
}

// Gt creates a criterion for a greater-than predicate.
func (op *Operations) Gt(lattr string, rattr string) policy.Criterion[*Record] {
	ml := Mapper.MapInt(lattr)
	mr := Mapper.MapInt(rattr)
	p := func(r *Record) bool { return ml(r) > mr(r) }
	return policy.Criterion[*Record]{Pred: p}
}

// Le creates a criterion for a lower-or-equal predicate.
func (op *Operations) LEq(lattr string, rattr string) policy.Criterion[*Record] {
	return op.Gt(lattr, rattr).Not()
}

// Lt creates a criterion for a lower-than predicate.
func (op *Operations) Lt(lattr string, rattr string) policy.Criterion[*Record] {
	return op.GEq(lattr, rattr).Not()
}

// StartsWith creates a criterion for a starts-with predicate.
func (op *Operations) StartsWith(lattr string, rattr string) policy.Criterion[*Record] {
	ml := Mapper.MapStr(lattr)
	mr := Mapper.MapStr(rattr)
	p := func(r *Record) bool { return eval(ml(r), mr(r), ops.startswith) }
	return policy.Criterion[*Record]{Pred: p}
}

// EndsWith creates a criterion for a ends-with predicate.
func (op *Operations) EndsWith(lattr string, rattr string) policy.Criterion[*Record] {
	ml := Mapper.MapStr(lattr)
	mr := Mapper.MapStr(rattr)
	p := func(r *Record) bool { return eval(ml(r), mr(r), ops.endswith) }
	return policy.Criterion[*Record]{Pred: p}
}

// Contains creates a criterion for a contains predicate.
func (op *Operations) Contains(lattr string, rattr string) policy.Criterion[*Record] {
	ml := Mapper.MapStr(lattr)
	mr := Mapper.MapStr(rattr)
	p := func(r *Record) bool { return eval(ml(r), mr(r), ops.contains) }
	return policy.Criterion[*Record]{Pred: p}
}

// IContains creates a criterion for a case-insensitive contains predicate.
func (op *Operations) IContains(lattr string, rattr string) policy.Criterion[*Record] {
	ml := Mapper.MapStr(lattr)
	mr := Mapper.MapStr(rattr)
	p := func(r *Record) bool { return eval(ml(r), mr(r), ops.icontains) }
	return policy.Criterion[*Record]{Pred: p}
}

// In creates a criterion for a list-inclusion predicate.
func (op *Operations) In(attr string, list []string) policy.Criterion[*Record] {
	m := Mapper.MapStr(attr)
	p := func(r *Record) bool {
		for _, v := range list {
			if eval(m(r), v, ops.eq) {
				return true
			}
		}
		return false
	}
	return policy.Criterion[*Record]{Pred: p}
}

// PMatch creates a criterion for a list-pattern-matching predicate.
func (op *Operations) PMatch(attr string, list []string) policy.Criterion[*Record] {
	m := Mapper.MapStr(attr)
	p := func(r *Record) bool {
		for _, v := range list {
			if eval(m(r), v, ops.contains) {
				return true
			}
		}
		return false
	}
	return policy.Criterion[*Record]{Pred: p}
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
	lattrs := strings.Split(fmt.Sprintf("%v", l), common.LISTSEP)
	rattrs := strings.Split(fmt.Sprintf("%v", r), common.LISTSEP)
	for _, lattr := range lattrs {
		for _, rattr := range rattrs {
			if op(lattr, rattr) {
				return true
			}
		}
	}
	return false
}
