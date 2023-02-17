//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
// Andreas Schade <san@zurich.ibm.com>
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

// Package flatrecord implements a flatrecord source for the policy compilers.
package flatrecord

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/common"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
)

// Record type
type Record struct {
	Fr  *sfgo.FlatRecord
	Ctx Context
}

// NewRecord creates a new Record isntance.
func NewRecord(fr *sfgo.FlatRecord) *Record {
	var r = new(Record)
	r.Fr = fr
	r.Ctx = make(Context, 4)
	return r
}

// RecAttribute denotes a record attribute enumeration.
type RecAttribute int8

// List of auxialiary record attributes enumerations.
const (
	PProcName RecAttribute = iota
	PProcExe
	PProcArgs
	PProcUID
	PProcUser
	PProcGID
	PProcGroup
	PProcTTY
	PProcEntry
	PProcCmdLine
	ProcAExe
	ProcAName
	ProcACmdLine
	ProcAPID
)

// GetInt returns an integer value from internal flat record.
func (r Record) GetInt(attr sfgo.Attribute, src sfgo.Source) int64 {
	for idx, s := range r.Fr.Sources {
		if s == src {
			return r.Fr.Ints[idx][attr]
		}
	}
	return sfgo.Zeros.Int64
}

// GetIntArray returns an integer array ptr value from internal flat record.
func (r Record) GetIntArray(attr sfgo.Attribute, src sfgo.Source) *[]int64 {
	for idx, s := range r.Fr.Sources {
		if s == src {
			if v, ok := r.Fr.Anys[idx][attr].(*[]int64); ok {
				return v
			}
			return nil
		}
	}
	return nil
}

// GetSvcArray returns a service array ptr value from internal flat record.
func (r Record) GetSvcArray(attr sfgo.Attribute, src sfgo.Source) *[]*sfgo.Service {
	for idx, s := range r.Fr.Sources {
		if s == src {
			if v, ok := r.Fr.Anys[idx][attr].(*[]*sfgo.Service); ok {
				return v
			}
			return nil
		}
	}
	return nil
}

// GetStr returns a string value from internal flat record.
func (r Record) GetStr(attr sfgo.Attribute, src sfgo.Source) string {
	for idx, s := range r.Fr.Sources {
		if s == src {
			return r.Fr.Strs[idx][attr]
		}
	}
	return sfgo.Zeros.String
}

// GetCachedValue returns the value of attr from cache for process ID.
func (r Record) GetCachedValue(ID sfgo.OID, attr RecAttribute) interface{} {
	if ptree := r.Fr.Ptree; ptree != nil {
		switch attr {
		case PProcName:
			if len(ptree) > 1 {
				return filepath.Base(ptree[1].Exe)
			}
		case PProcExe:
			if len(ptree) > 1 {
				return ptree[1].Exe
			}
		case PProcArgs:
			if len(ptree) > 1 {
				return ptree[1].ExeArgs
			}
		case PProcUID:
			if len(ptree) > 1 {
				return ptree[1].Uid
			}
		case PProcUser:
			if len(ptree) > 1 {
				return ptree[1].UserName
			}
		case PProcGID:
			if len(ptree) > 1 {
				return ptree[1].Gid
			}
		case PProcGroup:
			if len(ptree) > 1 {
				return ptree[1].GroupName
			}
		case PProcTTY:
			if len(ptree) > 1 {
				return ptree[1].Tty
			}
		case PProcEntry:
			if len(ptree) > 1 {
				return ptree[1].Entry
			}
		case PProcCmdLine:
			if len(ptree) > 1 {
				if len(ptree[1].ExeArgs) > 0 {
					return ptree[1].Exe + common.SPACE + ptree[1].ExeArgs
				}
				return ptree[1].Exe
			}
		case ProcAName:
			var s []string
			for _, p := range ptree {
				s = append(s, filepath.Base(p.Exe))
			}
			return strings.Join(s, common.LISTSEP)
		case ProcAExe:
			var s []string
			for _, p := range ptree {
				s = append(s, p.Exe)
			}
			return strings.Join(s, common.LISTSEP)
		case ProcACmdLine:
			var s []string
			for _, p := range ptree {
				if len(p.ExeArgs) > 0 {
					s = append(s, p.Exe+common.SPACE+p.ExeArgs)
				} else {
					s = append(s, p.Exe)
				}
			}
			return strings.Join(s, common.LISTSEP)
		case ProcAPID:
			var s []string
			for _, p := range ptree {
				s = append(s, strconv.FormatInt(p.Oid.Hpid, 10))
			}
			return strings.Join(s, common.LISTSEP)
		}
	}
	switch attr {
	case PProcUID, PProcGID, PProcTTY, PProcEntry:
		return sfgo.Zeros.Int64
	}
	return sfgo.Zeros.String
}

// Context denotes the type for contextual information obtained during rule processing.
type Context []interface{}

// ContextKey type
type contextKey int

// ContextKey enum
const (
	alertCtxKey contextKey = iota
	ruleCtxKey
	tagCtxKey
	hashCtxKey
)

func (s Context) IsAlert() bool {
	if s[alertCtxKey] != nil {
		return s[alertCtxKey].(bool)
	}
	return false
}

func (s Context) SetAlert(isAlert bool) {
	s[alertCtxKey] = isAlert
}

// AddRule adds a rule instance to the set of rules matching a record.
func (s Context) AddRules(rules ...policy.Rule[*Record]) {
	if s[ruleCtxKey] == nil {
		s[ruleCtxKey] = make([]policy.Rule[*Record], 0)
	}
	for _, r := range rules {
		s[ruleCtxKey] = append(s[ruleCtxKey].([]policy.Rule[*Record]), r)
	}
}

// GetRules retrieves the list of stored rules associated with a record context.
func (s Context) GetRules() []policy.Rule[*Record] {
	if s[ruleCtxKey] != nil {
		return s[ruleCtxKey].([]policy.Rule[*Record])
	}
	return nil
}

// SetTags stores tags into context object.
func (s Context) SetTags(tags []string) {
	s[tagCtxKey] = tags
}

// Adds tags to context object.
func (s Context) AddTags(tags ...string) {
	if s[tagCtxKey] == nil {
		s[tagCtxKey] = make([]string, 0)
	}
	for _, tag := range tags {
		s[tagCtxKey] = append(s[tagCtxKey].([]string), tag)
	}
}

// GetTags retrieves hashes from context object.
func (s Context) GetTags() []string {
	if s[tagCtxKey] != nil {
		return s[tagCtxKey].([]string)
	}
	return nil
}

func (s Context) GetHash(ht HashType) *HashSet {
	if s[hashCtxKey] == nil {
		return nil
	}
	hpa := s[hashCtxKey].([]*HashSet)
	return hpa[ht]
}

// Adds a hash set to context object.
func (s Context) SetHashes(ht HashType, hs *HashSet) {
	if s[hashCtxKey] == nil {
		s[hashCtxKey] = make([]*HashSet, 2)
	}
	hpa := s[hashCtxKey].([]*HashSet)

	if hpa[ht] == nil {
		hpa[ht] = hs
	}
}

type HashType uint

const (
	HASH_TYPE_PROC HashType = iota
	HASH_TYPE_FILE
)

type HashSet struct {
	Md5    string `json:"md5,omitempty"`
	Sha1   string `json:"sha1,omitempty"`
	Sha256 string `json:"sha256,omitempty"`
}
