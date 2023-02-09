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

import (
	"strings"

	"golang.org/x/exp/constraints"
)

// Operator type.
type Operator[T constraints.Ordered] func(T, T) bool

// operators struct.
type Operators struct {
	Eq          Operator[string]
	IEq         Operator[string]
	NEq         Operator[string]
	INEq        Operator[string]
	Contains    Operator[string]
	IContains   Operator[string]
	Startswith  Operator[string]
	IStartswith Operator[string]
	Endswith    Operator[string]
	IEndswith   Operator[string]
	Lt          Operator[int64]
	LEq         Operator[int64]
	Gt          Operator[int64]
	GEq         Operator[int64]
}

// Ops defines boolean comparison operators.
var Ops = Operators{
	Eq:          func(l string, r string) bool { return l == r },
	IEq:         func(l string, r string) bool { return strings.EqualFold(l, r) },
	NEq:         func(l string, r string) bool { return l != r },
	INEq:        func(l string, r string) bool { return !strings.EqualFold(l, r) },
	Contains:    func(l string, r string) bool { return strings.Contains(l, r) },
	IContains:   func(l string, r string) bool { return strings.Contains(strings.ToLower(l), strings.ToLower(r)) },
	Startswith:  func(l string, r string) bool { return strings.HasPrefix(l, r) },
	IStartswith: func(l string, r string) bool { return strings.HasPrefix(strings.ToLower(l), strings.ToLower(r)) },
	Endswith:    func(l string, r string) bool { return strings.HasSuffix(l, r) },
	IEndswith:   func(l string, r string) bool { return strings.HasSuffix(strings.ToLower(l), strings.ToLower(r)) },
	Lt:          func(l int64, r int64) bool { return l < r },
	LEq:         func(l int64, r int64) bool { return l <= r },
	Gt:          func(l int64, r int64) bool { return l > r },
	GEq:         func(l int64, r int64) bool { return l >= r },
}
