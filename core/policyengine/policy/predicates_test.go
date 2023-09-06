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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNot(t *testing.T) {
	c := False[any]()
	assert.Equal(t, true, c.Not().Eval(nil))
}

func TestAnd(t *testing.T) {
	c := False[any]()
	assert.Equal(t, false, c.And(c).Eval(nil))
	assert.Equal(t, false, c.And(c.Not()).Eval(nil))
	assert.Equal(t, false, c.Not().And(c).Eval(nil))
	assert.Equal(t, true, c.Not().And(c.Not()).Eval(nil))
}

func TestOr(t *testing.T) {
	c := False[any]()
	assert.Equal(t, false, c.Or(c).Eval(nil))
	assert.Equal(t, true, c.Or(c.Not()).Eval(nil))
	assert.Equal(t, true, c.Not().Or(c).Eval(nil))
	assert.Equal(t, true, c.Not().Or(c.Not()).Eval(nil))
}

func TestAll(t *testing.T) {
	assert.Equal(t, true, All([]Criterion[any]{True[any](), True[any]()}).Eval(nil))
	assert.Equal(t, false, All([]Criterion[any]{True[any](), False[any]()}).Eval(nil))
	assert.Equal(t, false, All([]Criterion[any]{False[any](), True[any]()}).Eval(nil))
	assert.Equal(t, false, All([]Criterion[any]{False[any](), False[any]()}).Eval(nil))
}

func TestAny(t *testing.T) {
	assert.Equal(t, true, Any([]Criterion[any]{True[any](), True[any]()}).Eval(nil))
	assert.Equal(t, true, Any([]Criterion[any]{True[any](), False[any]()}).Eval(nil))
	assert.Equal(t, true, Any([]Criterion[any]{False[any](), True[any]()}).Eval(nil))
	assert.Equal(t, false, Any([]Criterion[any]{False[any](), False[any]()}).Eval(nil))
}
