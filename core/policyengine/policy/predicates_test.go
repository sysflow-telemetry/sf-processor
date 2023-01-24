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

// func TestNot(t *testing.T) {
// 	c := False
// 	var r *Record
// 	assert.Equal(t, true, c.Not().Eval(r))
// }

// func TestAnd(t *testing.T) {
// 	c := False
// 	var r *Record
// 	assert.Equal(t, false, c.And(c).Eval(r))
// 	assert.Equal(t, false, c.And(c.Not()).Eval(r))
// 	assert.Equal(t, false, c.Not().And(c).Eval(r))
// 	assert.Equal(t, true, c.Not().And(c.Not()).Eval(r))
// }

// func TestOr(t *testing.T) {
// 	c := False
// 	var r *Record
// 	assert.Equal(t, false, c.Or(c).Eval(r))
// 	assert.Equal(t, true, c.Or(c.Not()).Eval(r))
// 	assert.Equal(t, true, c.Not().Or(c).Eval(r))
// 	assert.Equal(t, true, c.Not().Or(c.Not()).Eval(r))
// }

// func TestAll(t *testing.T) {
// 	var r *Record
// 	assert.Equal(t, true, All([]Criterion{True, True}).Eval(r))
// 	assert.Equal(t, false, All([]Criterion{True, False}).Eval(r))
// 	assert.Equal(t, false, All([]Criterion{False, True}).Eval(r))
// 	assert.Equal(t, false, All([]Criterion{False, False}).Eval(r))
// }

// func TestAny(t *testing.T) {
// 	var r *Record
// 	assert.Equal(t, true, Any([]Criterion{True, True}).Eval(r))
// 	assert.Equal(t, true, Any([]Criterion{True, False}).Eval(r))
// 	assert.Equal(t, true, Any([]Criterion{False, True}).Eval(r))
// 	assert.Equal(t, false, Any([]Criterion{False, False}).Eval(r))
// }

// func TestEq(t *testing.T) {
// 	r := NewRecord(sfgo.FlatRecord{})
// 	assert.Equal(t, false, Eq("0", "1").Eval(r))
// 	assert.Equal(t, true, Eq("0", "0").Eval(r))
// 	assert.Equal(t, true, Eq("sf.proc.gid", "0").Eval(r))
// 	assert.Equal(t, true, Eq("sf.proc.uid", "0").Eval(r))
// 	assert.Equal(t, true, Eq("sf.proc.exe", "").Eval(r))
// 	assert.Equal(t, true, Eq("sf.pproc.gid", "0").Eval(r))
// 	assert.Equal(t, true, Eq("sf.pproc.uid", "0").Eval(r))
// 	assert.Equal(t, true, Eq("sf.pproc.exe", "").Eval(r))
// 	// Privilege escalation condition
// 	assert.Equal(t, false, Eq("sf.pproc.uid", "0").And(Eq("sf.proc.gid", "0")).And(Exists("sf.pproc.gid").Not().Not()).Eval(r))
// }

// func TestExists(t *testing.T) {
// 	r := NewRecord(sfgo.FlatRecord{})
// 	assert.Equal(t, false, Exists("").Eval(r))
// 	assert.Equal(t, true, Exists("0").Eval(r))
// 	assert.Equal(t, false, Exists("sf.proc.gid").Eval(r))
// 	assert.Equal(t, false, Exists("sf.proc.uid").Eval(r))
// 	assert.Equal(t, false, Exists("sf.proc.exe").Eval(r))
// 	assert.Equal(t, false, Exists("sf.pproc.gid").Eval(r))
// 	assert.Equal(t, false, Exists("sf.pproc.uid").Eval(r))
// 	assert.Equal(t, false, Exists("sf.pproc.exe").Eval(r))
// }
