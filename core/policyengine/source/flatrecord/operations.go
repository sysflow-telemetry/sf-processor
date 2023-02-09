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

// CompareStr creates a criterion for a binary predicate over strings.
func (op *Operations) CompareStr(lattr string, rattr string, operator source.Operator[string]) policy.Criterion[*Record] {
	ml := Mapper.MapStr(lattr)
	mr := Mapper.MapStr(rattr)
	p := func(r *Record) bool { return compareStr(ml(r), mr(r), operator) }
	return policy.Criterion[*Record]{Pred: p}
}

// CompareInt creates a criterion for a binary predicate over integers.
func (op *Operations) CompareInt(lattr string, rattr string, operator source.Operator[int64]) policy.Criterion[*Record] {
	ml := Mapper.MapInt(lattr)
	mr := Mapper.MapInt(rattr)
	p := func(r *Record) bool { return compareInt(ml(r), mr(r), operator) }
	return policy.Criterion[*Record]{Pred: p}
}

// FoldAny creates a disjunctive criterion for a binary predicate over a list of strings.
func (op *Operations) FoldAny(attr string, list []string, operator source.Operator[string]) policy.Criterion[*Record] {
	m := Mapper.MapStr(attr)
	p := func(r *Record) bool {
		for _, v := range list {
			if compareStr(m(r), v, operator) {
				return true
			}
		}
		return false
	}
	return policy.Criterion[*Record]{Pred: p}
}

// FoldAll creates a conjunctive criterion for a binary predicate over a list of strings.
func (op *Operations) FoldAll(attr string, list []string, operator source.Operator[string]) policy.Criterion[*Record] {
	m := Mapper.MapStr(attr)
	p := func(r *Record) bool {
		for _, v := range list {
			if !compareStr(m(r), v, operator) {
				return false
			}
		}
		return true
	}
	return policy.Criterion[*Record]{Pred: p}
}

// RegExp creates a criterion for a regular-expression predicate.
func (op *Operations) RegExp(attr string, re string) policy.Criterion[*Record] {
	return policy.False[*Record]()
}

// compareStr compares two string values based on an operator.
func compareStr(l string, r string, op source.Operator[string]) bool {
	lattrs := strings.Split(l, common.LISTSEP)
	rattrs := strings.Split(r, common.LISTSEP)
	for _, lattr := range lattrs {
		for _, rattr := range rattrs {
			if op(lattr, rattr) {
				return true
			}
		}
	}
	return false
}

// compareInt compares two int64 values based on an operator.
func compareInt(l int64, r int64, op source.Operator[int64]) bool {
	return op(l, r)
}
