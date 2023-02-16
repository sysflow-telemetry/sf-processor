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

// Package source implements a backend for policy compilers.
package source

import (
	"errors"
	"strings"

	"golang.org/x/exp/constraints"
)

// Operator enum type.
type Operator int32

// Operator enums.
const (
	Eq Operator = iota
	IEq
	Contains
	IContains
	Startswith
	IStartswith
	Endswith
	IEndswith
	Lt
	LEq
	Gt
	GEq
)

func (s Operator) String() string {
	return [...]string{"Eq", "IEq", "Contains", "IContains", "Startswith", "IStartswith", "Endswith", "IEndswith", "Lt", "LEq", "Gt", "GEq"}[s]
}

// Operator function type.
type OpFunc[T constraints.Ordered | ~bool] func(T, T) bool

// Operator functions over strings.
type StrOps struct{}

func (StrOps) OpFunc(op Operator) (OpFunc[string], error) {
	switch op {
	case Eq:
		return func(l string, r string) bool { return l == r }, nil
	case IEq:
		return func(l string, r string) bool { return strings.EqualFold(l, r) }, nil
	case Contains:
		return func(l string, r string) bool { return strings.Contains(l, r) }, nil
	case IContains:
		return func(l string, r string) bool { return strings.Contains(strings.ToLower(l), strings.ToLower(r)) }, nil
	case Startswith:
		return func(l string, r string) bool { return strings.HasPrefix(l, r) }, nil
	case IStartswith:
		return func(l string, r string) bool { return strings.HasPrefix(strings.ToLower(l), strings.ToLower(r)) }, nil
	case Endswith:
		return func(l string, r string) bool { return strings.HasSuffix(l, r) }, nil
	case IEndswith:
		return func(l string, r string) bool { return strings.HasSuffix(strings.ToLower(l), strings.ToLower(r)) }, nil
	}
	return nil, errors.New("not a string operator")
}

// Operator functions over booleans.
type BoolOps struct{}

func (op BoolOps) Eq(l bool, r bool) bool { return l == r }

// Operator functions over integers.
type IntOps[T constraints.Integer] struct{}

func (IntOps[T]) OpFunc(op Operator) (OpFunc[T], error) {
	switch op {
	case Eq:
		return func(l T, r T) bool { return l == r }, nil
	case Lt:
		return func(l T, r T) bool { return l < r }, nil
	case LEq:
		return func(l T, r T) bool { return l <= r }, nil
	case Gt:
		return func(l T, r T) bool { return l > r }, nil
	case GEq:
		return func(l T, r T) bool { return l >= r }, nil
	}
	return nil, errors.New("not an integer operator")
}
