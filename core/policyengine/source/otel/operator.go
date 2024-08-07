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
	"bytes"
	"strconv"
	"strings"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

func doStringComparison(strVal string, rattr string, op source.Operator) bool {
	switch op {
	case source.Eq:
		return strVal == rattr
	case source.IEq:
		return strings.EqualFold(strVal, rattr)
	case source.Contains:
		return strings.Contains(strVal, rattr)
	case source.IContains:
		lLower := strings.ToLower(strVal)
		rLower := strings.ToLower(rattr)
		return strings.Contains(lLower, rLower)
	case source.Startswith:
		return strings.HasPrefix(strVal, rattr)
	case source.IStartswith:
		lLower := strings.ToLower(strVal)
		rLower := strings.ToLower(rattr)
		return strings.HasPrefix(lLower, rLower)
	case source.Endswith:
		return strings.HasSuffix(strVal, rattr)
	case source.IEndswith:
		lLower := strings.ToLower(strVal)
		rLower := strings.ToLower(rattr)
		return strings.HasSuffix(lLower, rLower)
	case source.Lt:
		return strVal < rattr
	case source.LEq:
		return strVal <= rattr
	case source.Gt:
		return strVal > rattr
	case source.GEq:
		return strVal >= rattr
	default:
		logger.Warn.Printf("unrecognized operator: %v", op)
		return false
	}
}

func doArrayComparison(arrVal *ArrayValue, rattr string, op source.Operator) bool {
	vals := arrVal.ArrayValue
	switch op {
	case source.Eq:
		return false
	case source.IEq:
		return false
	case source.Contains:
		if vals == nil {
			//not a valid array
			return false
		}

		for _, i := range vals.Values {
			str, ok := i.Value.(*StringValue)
			if !ok {
				continue
			}
			if doStringComparison(str.StringValue, rattr, op) {
				return true
			}
		}
		return false
	case source.IContains:
		for _, i := range vals.Values {
			str, ok := i.Value.(*StringValue)
			if !ok {
				continue
			}
			if doStringComparison(str.StringValue, rattr, op) {
				return true
			}
		}
		return false
	case source.Startswith:
		firstVal := vals.Values[0]
		str, ok := firstVal.Value.(*StringValue)
		if !ok {
			return false
		}
		return str.StringValue == rattr
	case source.IStartswith:
		firstVal := vals.Values[0]
		str, ok := firstVal.Value.(*StringValue)
		if !ok {
			return false
		}
		return strings.EqualFold(str.StringValue, rattr)
	case source.Endswith:
		l := len(vals.Values)
		lastVal := vals.Values[l-1]
		str, ok := lastVal.Value.(*StringValue)
		if !ok {
			return false
		}
		return str.StringValue == rattr
	case source.IEndswith:
		l := len(vals.Values)
		lastVal := vals.Values[l-1]
		str, ok := lastVal.Value.(*StringValue)
		if !ok {
			return false
		}
		return strings.EqualFold(str.StringValue, rattr)
	case source.Lt:
		return false
	case source.LEq:
		return false
	case source.Gt:
		return false
	case source.GEq:
		return false
	default:
		logger.Warn.Printf("unrecognized operator: %v", op)
		return false
	}
}

func doBooleanComparison(boolValue bool, rattr string, op source.Operator) bool {
	rattrVal := true
	if rattr == "False" {
		rattrVal = false
	}
	switch op {
	case source.Eq:
		return boolValue == rattrVal
	case source.IEq:
		return boolValue == rattrVal
	case source.Contains:
		return false
	case source.IContains:
		return false
	case source.Startswith:
		return false
	case source.IStartswith:
		return false
	case source.Endswith:
		return false
	case source.IEndswith:
		return false
	case source.Lt:
		return false
	case source.LEq:
		return false
	case source.Gt:
		return false
	case source.GEq:
		return false
	default:
		logger.Warn.Printf("unrecognized operator: %v", op)
		return false
	}
}

func doBytesComparison(byts []byte, rattr string, op source.Operator) bool {
	brattr := []byte(rattr)
	switch op {
	case source.Eq:
		return bytes.Equal(byts, brattr)
	case source.IEq:
		return bytes.EqualFold(byts, brattr)
	case source.Contains:
		return bytes.Contains(byts, brattr)
	case source.IContains:
		return false
	case source.Startswith:
		return bytes.HasPrefix(byts, brattr)
	case source.IStartswith:
		return false
	case source.Endswith:
		return bytes.HasSuffix(byts, brattr)
	case source.IEndswith:
		return false
	case source.Lt:
		return false
	case source.LEq:
		return false
	case source.Gt:
		return false
	case source.GEq:
		return false
	default:
		logger.Warn.Printf("unrecognized operator: %v", op)
		return false
	}
}

func doDoubleComparison(dbl float64, rattr string, op source.Operator) bool {
	drattr, err := strconv.ParseFloat(strings.TrimSpace(rattr), 64)
	if err != nil {
		return false
	}
	switch op {
	case source.Eq:
		return dbl == drattr
	case source.IEq:
		return false
	case source.Contains:
		return false
	case source.IContains:
		return false
	case source.Startswith:
		return false
	case source.IStartswith:
		return false
	case source.Endswith:
		return false
	case source.IEndswith:
		return false
	case source.Lt:
		return dbl < drattr
	case source.LEq:
		return dbl <= drattr
	case source.Gt:
		return dbl > drattr
	case source.GEq:
		return dbl >= drattr
	default:
		logger.Warn.Printf("unrecognized operator: %v", op)
		return false
	}
}

func doIntComparison(intVal int64, rattr string, op source.Operator) bool {
	irattr, err := strconv.ParseInt(strings.TrimSpace(rattr), 10, 64)
	if err != nil {
		return false
	}
	switch op {
	case source.Eq:
		return intVal == irattr
	case source.IEq:
		return false
	case source.Contains:
		return false
	case source.IContains:
		return false
	case source.Startswith:
		return false
	case source.IStartswith:
		return false
	case source.Endswith:
		return false
	case source.IEndswith:
		return false
	case source.Lt:
		return intVal < irattr
	case source.LEq:
		return intVal <= irattr
	case source.Gt:
		return intVal > irattr
	case source.GEq:
		return intVal >= irattr
	default:
		logger.Warn.Printf("unrecognized operator: %v", op)
		return false
	}
}
