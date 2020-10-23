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
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/sf-processor/core/cache"
)

// Action type for enumeration.
type Action int

// Action enumeration.
const (
	Alert Action = iota
	Tag
	Hash
)

// String returns the string representation of an action instance.
func (a Action) String() string {
	return [...]string{"alert", "tag", "hash"}[a]
}

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
	Actions   []Action
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
	Fr    sfgo.FlatRecord
	Cr    *cache.SFTables
	Ptree map[sfgo.OID][]*sfgo.Process
	Ctx   Context
}

// NewRecord creates a new Record isntance.
func NewRecord(fr sfgo.FlatRecord, cr *cache.SFTables) *Record {
	var r = new(Record)
	r.Fr = fr
	r.Cr = cr
	r.Ptree = make(map[sfgo.OID][]*sfgo.Process)
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
			return strings.TrimSpace(r.Fr.Strs[idx][attr])
		}
	}
	return sfgo.Zeros.String
}

// GetProc returns a process object by ID.
func (r Record) GetProc(ID sfgo.OID) *sfgo.Process {
	return r.Cr.GetProc(ID)
}

func (r Record) getProcProv(ID sfgo.OID) []*sfgo.Process {
	var ptree = make([]*sfgo.Process, 0)
	if p := r.Cr.GetProc(ID); p != nil && p.Poid != nil && p.Poid.UnionType == sfgo.UnionNullOIDTypeEnumOID {
		return append(append(ptree, p), r.getProcProv(*p.Poid.OID)...)
	}
	return ptree
}

func (r Record) memoizePtree(ID sfgo.OID) []*sfgo.Process {
	if ptree, ok := r.Ptree[ID]; ok {
		return ptree
	}
	r.Ptree[ID] = r.getProcProv(ID)
	return r.Ptree[ID]
}

// GetCachedValue returns the value of attr from cache for process ID.
func (r Record) GetCachedValue(ID sfgo.OID, attr RecAttribute) interface{} {
	if ptree := r.memoizePtree(ID); ptree != nil {
		switch attr {
		case PProcName:
			if len(ptree) > 1 {
				return filepath.Base(ptree[1].Exe)
			}
			break
		case PProcExe:
			if len(ptree) > 1 {
				return ptree[1].Exe
			}
			break
		case PProcArgs:
			if len(ptree) > 1 {
				return ptree[1].ExeArgs
			}
			break
		case PProcUID:
			if len(ptree) > 1 {
				return ptree[1].Uid
			}
			break
		case PProcUser:
			if len(ptree) > 1 {
				return ptree[1].UserName
			}
			break
		case PProcGID:
			if len(ptree) > 1 {
				return ptree[1].Gid
			}
			break
		case PProcGroup:
			if len(ptree) > 1 {
				return ptree[1].GroupName
			}
			break
		case PProcTTY:
			if len(ptree) > 1 {
				return ptree[1].Tty
			}
			break
		case PProcEntry:
			if len(ptree) > 1 {
				return ptree[1].Entry
			}
			break
		case PProcCmdLine:
			if len(ptree) > 1 {
				return ptree[1].Exe + SPACE + ptree[1].ExeArgs
			}
			break
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
				s = append(s, p.Exe+SPACE+p.ExeArgs)
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

// GetTags retrieves hashes from context object.
func (s Context) GetTags() []string {
	if s[tagCtxKey] != nil {
		return s[tagCtxKey].([]string)
	}
	return nil
}

// SetHashes stores hashes into context object.
func (s Context) SetHashes(h HashSet) {
	s[hashCtxKey] = h
}

// GetHashes retrieves hashes from context object.
func (s Context) GetHashes() HashSet {
	if s[hashCtxKey] != nil {
		return s[hashCtxKey].(HashSet)
	}
	return HashSet{}
}

// HashSet type
type HashSet struct {
	MD5      string
	SHA1     string
	SHA256   string
	Size     int
	UpdateTs int64
}
