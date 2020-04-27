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
)

// String returns the string representation of an action instance.
func (a Action) String() string {
	return [...]string{"alert", "tag"}[a]
}

// EnrichmentTag denotes the type for enrichment tags.
type EnrichmentTag interface{}

// Context denotes the type for contextual information obtained during rule processing.
type Context map[string]interface{}

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
	actions   []Action
	tags      []EnrichmentTag
	priority  Priority
	ctx       Context
}

// Filter type
type Filter struct {
	name      string
	condition Criterion
}

// Record type
type Record struct {
	Fr    sfgo.FlatRecord
	Ctx   *cache.SFTables
	Ptree map[sfgo.OID][]*sfgo.Process
}

// NewRecord creates a new Record isntance.
func NewRecord(fr sfgo.FlatRecord, ctx *cache.SFTables) Record {
	return Record{fr, ctx, make(map[sfgo.OID][]*sfgo.Process)}
}

// RecAttribute denotes a record attribute enumeration.
type RecAttribute int

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
	PProcCmdLine
	ProcAExe
	ProcAName
	ProcACmdLine
	ProcAPID
)

// GetInt returns an integer value from internal flat record.
func (r Record) GetInt(attr sfgo.Attribute) int64 {
	return r.Fr.Ints[attr]
}

// GetStr returns a string value from internal flat record.
func (r Record) GetStr(attr sfgo.Attribute) string {
	return r.Fr.Strs[attr]
}

// GetProc returns a process object by ID.
func (r Record) GetProc(ID sfgo.OID) *sfgo.Process {
	return r.Ctx.GetProc(ID)
}

func (r Record) getProcProv(ID sfgo.OID) []*sfgo.Process {
	var ptree = make([]*sfgo.Process, 0)
	if p := r.Ctx.GetProc(ID); p != nil && p.Poid.UnionType != sfgo.UnionNullOIDTypeEnumNull {
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

// Occurence type
type Occurence struct {
	Record Record
	Rules  []Rule
}

// OccurenceChannel type
type OccurenceChannel struct {
	In chan *Occurence
}

// NewOccurence constructs a new occurence
func NewOccurence(r Record, rlist []Rule) *Occurence {
	o := new(Occurence)
	o.Record = r
	o.Rules = rlist
	return o
}
