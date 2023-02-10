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

type Operations struct {
	strOps source.StrOps
	intOps source.IntOps[int64]
}

func NewOperations() source.Operations[*Record] {
	return &Operations{strOps: source.StrOps{}, intOps: source.IntOps[int64]{}}
}

// Exists creates a criterion for an existential predicate.
func (op *Operations) Exists(attr string) policy.Criterion[*Record] {
	m := Mapper.Map(attr)
	p := func(r *Record) bool { return !reflect.ValueOf(m(r)).IsZero() }
	return policy.Criterion[*Record]{Pred: p}
}

// Compare creates a criterion for a binary predicate.
func (op *Operations) Compare(lattr string, rattr string, operator source.Operator) policy.Criterion[*Record] {
	switch operator {
	case source.Lt, source.LEq, source.Gt, source.GEq:
		return op.compareInt(lattr, rattr, operator)
	}
	return op.compareStr(lattr, rattr, operator)
}

// compareStr creates a criterion for a binary predicate over strings.
func (op *Operations) compareStr(lattr string, rattr string, operator source.Operator) policy.Criterion[*Record] {
	ml := Mapper.MapStr(lattr)
	mr := Mapper.MapStr(rattr)
	o, _ := op.strOps.OpFunc(operator)
	p := func(r *Record) bool { return compareStr(ml(r), mr(r), o) }
	return policy.Criterion[*Record]{Pred: p}
}

// compareInt creates a criterion for a binary predicate over integers.
func (op *Operations) compareInt(lattr string, rattr string, operator source.Operator) policy.Criterion[*Record] {
	ml := Mapper.MapInt(lattr)
	mr := Mapper.MapInt(rattr)
	o, _ := op.intOps.OpFunc(operator)
	p := func(r *Record) bool { return compareInt(ml(r), mr(r), o) }
	return policy.Criterion[*Record]{Pred: p}
}

// FoldAny creates a disjunctive criterion for a binary predicate over a list of strings.
func (op *Operations) FoldAny(attr string, list []string, operator source.Operator) policy.Criterion[*Record] {
	m := Mapper.MapStr(attr)
	o, _ := op.strOps.OpFunc(operator)
	p := func(r *Record) bool {
		for _, v := range list {
			if compareStr(m(r), v, o) {
				return true
			}
		}
		return false
	}
	return policy.Criterion[*Record]{Pred: p}
}

// FoldAll creates a conjunctive criterion for a binary predicate over a list of strings.
func (op *Operations) FoldAll(attr string, list []string, operator source.Operator) policy.Criterion[*Record] {
	m := Mapper.MapStr(attr)
	o, _ := op.strOps.OpFunc(operator)
	p := func(r *Record) bool {
		for _, v := range list {
			if !compareStr(m(r), v, o) {
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
func compareStr(l string, r string, op source.OpFunc[string]) bool {
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
func compareInt(l int64, r int64, op source.OpFunc[int64]) bool {
	return op(l, r)
}
