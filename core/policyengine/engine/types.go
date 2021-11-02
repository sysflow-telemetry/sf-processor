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

// Package engine implements a rules engine for telemetry records.
package engine

import (
	"crypto"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

// EnrichmentTag denotes the type for enrichment tags.
type EnrichmentTag interface{}

// Priority denotes the type for rule priority.
type Priority int

// Priority enumeration.
const (
	Low Priority = iota
	Medium
	High
)

// String returns the string representation of a priority instance.
func (p Priority) String() string {
	return [...]string{"low", "medium", "high"}[p]
}

// Rule type
type Rule struct {
	Name      string
	Desc      string
	condition Criterion
	Actions   []string
	Tags      []EnrichmentTag
	Priority  Priority
	Prefilter []string
	Enabled   bool
}

func (s Rule) isApplicable(r *Record) bool {
	if len(s.Prefilter) == 0 {
		return true
	}
	rtype := Mapper.MapStr(SF_TYPE)(r)
	for _, pf := range s.Prefilter {
		if rtype == pf {
			return true
		}
	}
	return false
}

// Filter type
type Filter struct {
	Name      string
	condition Criterion
	Enabled   bool
}

// Record type
type Record struct {
	Fr  sfgo.FlatRecord
	Ctx Context
}

// NewRecord creates a new Record isntance.
func NewRecord(fr sfgo.FlatRecord) *Record {
	var r = new(Record)
	r.Fr = fr
	r.Ctx = make(Context, 3)
	return r
}

// RecordChannel type
type RecordChannel struct {
	In chan *Record
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
					return ptree[1].Exe + SPACE + ptree[1].ExeArgs
				}
				return ptree[1].Exe
			}
		case ProcAName:
			var s []string
			for _, p := range ptree {
				s = append(s, filepath.Base(p.Exe))
			}
			return strings.Join(s, LISTSEP)
		case ProcAExe:
			var s []string
			for _, p := range ptree {
				s = append(s, p.Exe)
			}
			return strings.Join(s, LISTSEP)
		case ProcACmdLine:
			var s []string
			for _, p := range ptree {
				if len(p.ExeArgs) > 0 {
					s = append(s, p.Exe+SPACE+p.ExeArgs)
				} else {
					s = append(s, p.Exe)
				}
			}
			return strings.Join(s, LISTSEP)
		case ProcAPID:
			var s []string
			for _, p := range ptree {
				s = append(s, strconv.FormatInt(p.Oid.Hpid, 10))
			}
			return strings.Join(s, LISTSEP)
		}
	}
	return sfgo.Zeros.String
}

// Context denotes the type for contextual information obtained during rule processing.
type Context []interface{}

// ContextKey type
type contextKey int

// ContextKey enum
const (
	ruleCtxKey contextKey = iota
	tagCtxKey
	hashCtxKey
)

// AddRule stores add a rule instance to the set of rules matching a record.
func (s Context) AddRule(r Rule) {
	if s[ruleCtxKey] == nil {
		s[ruleCtxKey] = make([]Rule, 0)
	}
	s[ruleCtxKey] = append(s[ruleCtxKey].([]Rule), r)
}

// GetRules retrieves the list of stored rules associated with a record context.
func (s Context) GetRules() []Rule {
	if s[ruleCtxKey] != nil {
		return s[ruleCtxKey].([]Rule)
	}
	return nil
}

// SetTags stores tags into context object.
func (s Context) SetTags(tags []string) {
	s[tagCtxKey] = tags
}

// Adds tags to context object.
func (s Context) AddTag(tag string) {
	if s[tagCtxKey] == nil {
		s[tagCtxKey] = make([]string, 0)
	}
	s[tagCtxKey] = append(s[tagCtxKey].([]string), tag)
}

// GetTags retrieves hashes from context object.
func (s Context) GetTags() []string {
	if s[tagCtxKey] != nil {
		return s[tagCtxKey].([]string)
	}
	return nil
}

// SetHashes stores hashes into context object.
func (s Context) SetHashes(h []*HashSet) {
	s[hashCtxKey] = h
}

// Adds a hash value to context object.
func (s Context) AddHash(h *HashSet) {
	if s[hashCtxKey] == nil {
		s[hashCtxKey] = make([]*HashSet, 0)
	}
	s[hashCtxKey] = append(s[hashCtxKey].([]*HashSet), h)
}

// GetHashes retrieves hashes from context object.
func (s Context) GetHashes() []*HashSet {
	if s[hashCtxKey] != nil {
		return s[hashCtxKey].([]*HashSet)
	}
	return nil
}

// HashSet type
type HashSet struct {
	Source    sfgo.Source
	Algorithm crypto.Hash
	Value     string
}
