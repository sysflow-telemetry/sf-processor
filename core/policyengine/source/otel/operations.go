//
// Copyright (C) 2024 IBM Corporation.
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

// Package otel implements an open telemetry backend for the policy compilers.
package otel

import (
	"regexp"

	"github.com/pkg/errors"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

type Operations struct{}

func NewOperations() source.Operations[*ResourceLogs] {
	return &Operations{}
}

func getAllAttributes(rl *ResourceLogs) []*KeyValue {
	allAttrs := rl.Resource.Attributes
	scopeLogs := rl.ScopeLogs

	for _, sl := range scopeLogs {
		//scope logs attributes
		if sl.Scope != nil && sl.Scope.Attributes != nil {
			allAttrs = append(allAttrs, sl.Scope.Attributes...)
		}

		// sl.Log records have their own attributes as well
		for _, lr := range sl.LogRecords {
			allAttrs = append(allAttrs, lr.Attributes...)
		}
	}
	return allAttrs
}

func (ops *Operations) Exists(attr string) (policy.Criterion[*ResourceLogs], error) {
	f := func(rl *ResourceLogs) bool {
		allAttrs := getAllAttributes(rl)

		for _, a := range allAttrs {
			if a.Key == attr {
				return true
			}
		}
		return false
	}
	return policy.Criterion[*ResourceLogs]{Pred: f}, nil
}

func (ops *Operations) compareHelper(lattr string, rattr string, op source.Operator, kvs []*KeyValue) bool {
	for _, a := range kvs {
		if a.Key != lattr {
			continue
		}
		val := a.Value.Value
		switch v := val.(type) {
		case *StringValue:
			strVal := v.StringValue
			return doStringComparison(strVal, rattr, op)
		case *ArrayValue:
			return doArrayComparison(v, rattr, op)
		case *BoolValue:
			boolVal := v.BoolValue
			return doBooleanComparison(boolVal, rattr, op)
		case *BytesValue:
			bytVal := v.BytesValue
			return doBytesComparison(bytVal, rattr, op)
		case *DoubleValue:
			dblVal := v.DoubleValue
			return doDoubleComparison(dblVal, rattr, op)
		case *IntValue:
			intVal := v.IntValue
			return doIntComparison(intVal, rattr, op)
		case *KvListValue:
			kvListVal := v.KvlistValue.Values
			return ops.compareHelper(lattr, rattr, op, kvListVal)
		}
	}
	return false
}

func (ops *Operations) Compare(lattr string, rattr string, op source.Operator) (policy.Criterion[*ResourceLogs], error) {
	//TODO check if rattr is literal or in rule
	f := func(rl *ResourceLogs) bool {
		allAttrs := getAllAttributes(rl)
		return ops.compareHelper(lattr, rattr, op, allAttrs)
	}

	return policy.Criterion[*ResourceLogs]{Pred: f}, nil
}

func (ops *Operations) FoldAny(attr string, list []string, op source.Operator) (policy.Criterion[*ResourceLogs], error) {
	f := func(rl *ResourceLogs) bool {
		allAttrs := getAllAttributes(rl)
		for _, rattr := range list {
			isTrue := ops.compareHelper(attr, rattr, op, allAttrs)
			if isTrue {
				return true
			}
		}
		return false
	}
	return policy.Criterion[*ResourceLogs]{Pred: f}, nil
}

func (ops *Operations) FoldAll(attr string, list []string, op source.Operator) (policy.Criterion[*ResourceLogs], error) {
	f := func(rl *ResourceLogs) bool {
		allAttrs := getAllAttributes(rl)
		for _, rattr := range list {
			isTrue := ops.compareHelper(attr, rattr, op, allAttrs)
			if !isTrue {
				return false
			}
		}
		return true
	}
	return policy.Criterion[*ResourceLogs]{Pred: f}, nil
}

func (op *Operations) RegExp(attr string, re string) (policy.Criterion[*ResourceLogs], error) {
	if regexp, err := regexp.Compile(re); err == nil {
		p := func(r *ResourceLogs) bool {
			allAttrs := getAllAttributes(r)
			for _, a := range allAttrs {
				if a.Key != attr {
					continue
				}
				v := a.Value.Value
				strV, ok := v.(*StringValue)
				if !ok {
					return false
				}
				return regexp.FindString(strV.StringValue) != ""
			}
			return false
		}
		return policy.Criterion[*ResourceLogs]{Pred: p}, nil
	}
	return policy.False[*ResourceLogs](), errors.Errorf("could not compile regular expression %s", re)
}
