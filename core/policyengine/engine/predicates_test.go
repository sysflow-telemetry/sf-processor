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
//
package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNot(t *testing.T) {
	c := False
	var r *Record
	assert.Equal(t, true, c.Not().Eval(r))
}

func TestAnd(t *testing.T) {
	c := False
	var r *Record
	assert.Equal(t, false, c.And(c).Eval(r))
	assert.Equal(t, false, c.And(c.Not()).Eval(r))
	assert.Equal(t, false, c.Not().And(c).Eval(r))
	assert.Equal(t, true, c.Not().And(c.Not()).Eval(r))
}

func TestOr(t *testing.T) {
	c := False
	var r *Record
	assert.Equal(t, false, c.Or(c).Eval(r))
	assert.Equal(t, true, c.Or(c.Not()).Eval(r))
	assert.Equal(t, true, c.Not().Or(c).Eval(r))
	assert.Equal(t, true, c.Not().Or(c.Not()).Eval(r))
}

func TestAll(t *testing.T) {
	var r *Record
	assert.Equal(t, true, All([]Criterion{True, True}).Eval(r))
	assert.Equal(t, false, All([]Criterion{True, False}).Eval(r))
	assert.Equal(t, false, All([]Criterion{False, True}).Eval(r))
	assert.Equal(t, false, All([]Criterion{False, False}).Eval(r))
}

func TestAny(t *testing.T) {
	var r *Record
	assert.Equal(t, true, Any([]Criterion{True, True}).Eval(r))
	assert.Equal(t, true, Any([]Criterion{True, False}).Eval(r))
	assert.Equal(t, true, Any([]Criterion{False, True}).Eval(r))
	assert.Equal(t, false, Any([]Criterion{False, False}).Eval(r))
}
