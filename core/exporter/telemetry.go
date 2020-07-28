//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
package exporter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.ibm.com/sysflow/sf-processor/core/policyengine/engine"
)

const schemaVersion = "0.1"

// SysFlow record components
const (
	proc      = "proc"
	pproc     = "pproc"
	net       = "net"
	file      = "file"
	flow      = "flow"
	container = "container"
	node      = "node"
)

// TelemetryRecord type
type TelemetryRecord struct {
	Version     string `json:"version,omitempty"`
	*FlatRecord `json:",omitempty"`
	*DataRecord `json:",omitempty"`
	Hashes      *engine.HashSet `json:"hashes,omitempty"`
	Policies    []Policy        `json:"policies,omitempty"`
}

// FlatRecord type
type FlatRecord struct {
	Data map[string]interface{} `json:"record"`
}

// DataRecord type (warning: make sure field names have only first letter capitalized)
type DataRecord struct {
	Type       string   `json:"type"`
	Opflags    []string `json:"opflags"`
	Ret        int64    `json:"ret"`
	Ts         int64    `json:"ts"`
	Endts      int64    `json:"endts,omitempty"`
	*ProcData  `json:",omitempty"`
	*PprocData `json:",omitempty"`
	*NetData   `json:",omitempty"`
	*FileData  `json:",omitempty"`
	*FlowData  `json:",omitempty"`
	*ContData  `json:",omitempty"`
	*NodeData  `json:",omitempty"`
}

// ProcData type
type ProcData struct {
	Proc map[string]interface{} `json:"proc"`
}

// PprocData type
type PprocData struct {
	Pproc map[string]interface{} `json:"pproc"`
}

// NetData type
type NetData struct {
	Net map[string]interface{} `json:"net"`
}

// FileData type
type FileData struct {
	File map[string]interface{} `json:"file"`
}

// FlowData type
type FlowData struct {
	Flow map[string]interface{} `json:"flow"`
}

// ContData type
type ContData struct {
	Container map[string]interface{} `json:"container"`
}

// NodeData type
type NodeData struct {
	Node map[string]interface{} `json:"node"`
}

// CreateTelemetryRecords creates offense instances based on a list of records
func CreateTelemetryRecords(occs []*engine.Record, config Config) []Event {
	var recs = make([]Event, 0)
	for _, o := range occs {
		recs = append(recs, extractTelemetryRecord(o, config))
	}
	return recs
}

// ToJSONStr returns a JSON string representation of an observation
func (s TelemetryRecord) ToJSONStr() string {
	return string(s.ToJSON())
}

// ToJSON returns a JSON bytearray representation of an observation
func (s TelemetryRecord) ToJSON() []byte {
	o, _ := json.Marshal(s)
	return o
}

func extractTelemetryRecord(rec *engine.Record, config Config) TelemetryRecord {
	r := TelemetryRecord{}
	r.Version = schemaVersion
	if config.Flat {
		r.FlatRecord = new(FlatRecord)
		r.FlatRecord.Data = make(map[string]interface{})
		for _, k := range engine.Fields {
			r.Data[k] = engine.Mapper.Mappers[k](rec)
		}
	} else {
		r.DataRecord = new(DataRecord)
		pprocID := engine.Mapper.MapInt(engine.SF_PPROC_PID)(rec)
		pprocExists := !reflect.ValueOf(pprocID).IsZero()
		ct := engine.Mapper.MapStr(engine.SF_CONTAINER_ID)(rec)
		ctExists := !reflect.ValueOf(ct).IsZero()
		for _, k := range engine.Fields {
			kc := strings.Split(k, ".")
			value := extractValue(k, engine.Mapper.Mappers[k](rec))
			if len(kc) == 2 {
				switch value.(type) {
				case string:
					reflect.ValueOf(r.DataRecord).Elem().FieldByName(strings.Title(kc[1])).SetString(value.(string))
				case int64:
					reflect.ValueOf(r.DataRecord).Elem().FieldByName(strings.Title(kc[1])).SetInt(value.(int64))
				case []string:
					reflect.ValueOf(r.DataRecord).Elem().FieldByName(strings.Title(kc[1])).Set(reflect.ValueOf(value))
				}
			} else if len(kc) == 3 {
				switch kc[1] {
				case proc:
					if r.ProcData == nil {
						r.ProcData = new(ProcData)
						r.ProcData.Proc = make(map[string]interface{})
					}
					r.Proc[kc[2]] = value
				case pproc:
					if pprocExists {
						if r.PprocData == nil {
							r.PprocData = new(PprocData)
							r.PprocData.Pproc = make(map[string]interface{})
						}
						r.Pproc[kc[2]] = value
					}
				case net:
					if r.Type == engine.TyNF {
						if r.NetData == nil {
							r.NetData = new(NetData)
							r.NetData.Net = make(map[string]interface{})
						}
						r.Net[kc[2]] = value
					}
				case file:
					if r.Type == engine.TyFF || r.Type == engine.TyFE {
						if r.FileData == nil {
							r.FileData = new(FileData)
							r.FileData.File = make(map[string]interface{})
						}
						r.File[kc[2]] = value
					}
				case flow:
					if r.Type == engine.TyFF || r.Type == engine.TyNF {
						if r.FlowData == nil {
							r.FlowData = new(FlowData)
							r.FlowData.Flow = make(map[string]interface{})
						}
						r.Flow[kc[2]] = value
					}
				case container:
					if ctExists {
						if r.ContData == nil {
							r.ContData = new(ContData)
							r.ContData.Container = make(map[string]interface{})
						}
						r.Container[kc[2]] = value
					}
				case node:
					if r.NodeData == nil {
						r.NodeData = new(NodeData)
						r.NodeData.Node = make(map[string]interface{})
					}
					r.Node[kc[2]] = value
				}
			}
		}
	}
	hashset := rec.Ctx.GetHashes()
	if !reflect.ValueOf(hashset.MD5).IsZero() {
		r.Hashes = &hashset
	}
	r.Policies = extractPolicySet(rec.Ctx.GetRules())
	return r
}

func extractPolicySet(rules []engine.Rule) []Policy {
	var pols = make([]Policy, 0)
	for _, r := range rules {
		p := Policy{
			ID:       r.Name,
			Desc:     r.Desc,
			Priority: int(r.Priority),
			Tags:     extracTags(r.Tags),
		}
		pols = append(pols, p)
	}
	return pols
}

func extracTags(tags []engine.EnrichmentTag) []string {
	s := make([]string, 0)
	for _, v := range tags {
		switch v.(type) {
		case []string:
			s = append(s, v.([]string)...)
			break
		default:
			s = append(s, string(fmt.Sprintf("%v", v)))
			break
		}
	}
	return s
}

func extractValue(k string, v interface{}) interface{} {
	switch v.(type) {
	case string:
		if array(k) {
			return strings.Split(v.(string), engine.LISTSEP)
		}
		return v
	default:
		return v
	}
}

func array(k string) bool {
	return k == engine.SF_OPFLAGS || k == engine.SF_PROC_APID || k == engine.SF_PROC_ANAME ||
		k == engine.SF_PROC_AEXE || k == engine.SF_PROC_ACMDLINE || k == engine.SF_FILE_OPENFLAGS ||
		k == engine.SF_NET_IP || k == engine.SF_NET_PORT
}
