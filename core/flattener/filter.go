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

// Package flattener flattens input telemetry in a flattened representation.
package flattener

import (
	"container/list"
	"encoding/binary"
	"time"

	"github.com/cespare/xxhash/v2"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

var byteInt64 []byte = make([]byte, 8)

// Filter is a time decaying filter with a TTL per entry.
type Filter struct {
	m   map[uint64]int64
	q   *list.List
	ttl time.Duration
}

// Entry encodes a hash value with the time it was first added to the filter.
type Entry struct {
	h         uint64
	firstSeen time.Time
}

// NewFilter creates a new time decaying filter that evicts entries that have been seen longer than t duration.
func NewFilter(t time.Duration) *Filter {
	return &Filter{m: make(map[uint64]int64), q: list.New(), ttl: t}
}

// Test tests if hash h has been seen since maximum ttl.
func (f *Filter) Test(h uint64) bool {
	f.evictAgedEntries()
	_, ok := f.m[h]
	return ok
}

// TestAndAdd tests if hash h has been seen since maximum ttl and adds or increments the element in the filter cache.
func (f *Filter) TestAndAdd(h uint64) bool {
	f.evictAgedEntries()
	_, ok := f.m[h]
	f.Add(h)
	return ok
}

// Contains returns how many times hash h has been seen during its ttl time.
func (f *Filter) Count(h uint64) int64 {
	f.evictAgedEntries()
	if count, ok := f.m[h]; ok {
		return count
	}
	return 0
}

// Add adds hash h to the filter.
func (f *Filter) Add(h uint64) {
	if v, ok := f.m[h]; !ok {
		f.m[h] = 1
		f.q.PushBack(Entry{h: h, firstSeen: time.Now()})
	} else {
		f.m[h] = v + 1
	}
}

func (f *Filter) evictAgedEntries() {
	for f.q.Len() > 0 {
		e := f.q.Front()
		entry := e.Value.(Entry)
		if time.Since(entry.firstSeen) < f.ttl {
			break
		}
		f.q.Remove(e)
		delete(f.m, entry.h)
	}
}

// semanticHash computes a hash value over record attributes denoting the semantics of the record (used in the time decay filter).
func semanticHash(fr *sfgo.FlatRecord) uint64 {
	h := xxhash.New()
	h.Write([]byte(fr.Strs[sfgo.SYSFLOW_SRC][sfgo.PROC_EXE_STR]))
	h.Write([]byte(fr.Strs[sfgo.SYSFLOW_SRC][sfgo.PROC_EXEARGS_STR]))
	binary.LittleEndian.PutUint64(byteInt64, uint64(fr.Ints[sfgo.SYSFLOW_SRC][sfgo.PROC_UID_INT]))
	h.Write(byteInt64)
	binary.LittleEndian.PutUint64(byteInt64, uint64(fr.Ints[sfgo.SYSFLOW_SRC][sfgo.PROC_GID_INT]))
	h.Write(byteInt64)
	binary.LittleEndian.PutUint64(byteInt64, uint64(fr.Ints[sfgo.SYSFLOW_SRC][sfgo.OPFLAGS_INT]))
	h.Write(byteInt64)
	binary.LittleEndian.PutUint64(byteInt64, uint64(fr.Ints[sfgo.SYSFLOW_SRC][sfgo.PROC_TTY_INT]))
	h.Write(byteInt64)
	sfType := fr.Ints[sfgo.SYSFLOW_IDX][sfgo.SF_REC_TYPE]
	if sfType == sfgo.NET_FLOW {
		binary.LittleEndian.PutUint64(byteInt64, uint64(fr.Ints[sfgo.SYSFLOW_SRC][sfgo.FL_NETW_SIP_INT]))
		h.Write(byteInt64)
		binary.LittleEndian.PutUint64(byteInt64, uint64(fr.Ints[sfgo.SYSFLOW_SRC][sfgo.FL_NETW_DIP_INT]))
		h.Write(byteInt64)
		binary.LittleEndian.PutUint64(byteInt64, uint64(fr.Ints[sfgo.SYSFLOW_SRC][sfgo.FL_NETW_DPORT_INT]))
		h.Write(byteInt64)
		binary.LittleEndian.PutUint64(byteInt64, uint64(fr.Ints[sfgo.SYSFLOW_SRC][sfgo.FL_NETW_PROTO_INT]))
		h.Write(byteInt64)
	}
	if sfType == sfgo.FILE_FLOW || sfType == sfgo.FILE_EVT {
		h.Write([]byte(fr.Strs[sfgo.SYSFLOW_SRC][sfgo.FILE_PATH_STR]))
	}
	return h.Sum64()
}
