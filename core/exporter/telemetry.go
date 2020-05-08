package exporter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.ibm.com/sysflow/sf-processor/core/sfpe/engine"
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
		for k, v := range engine.Mapper.Mappers {
			r.Data[k] = v(rec)
		}
	} else {
		r.DataRecord = new(DataRecord)
		pprocID := engine.Mapper.MapInt("sf.pproc.pid")(rec)
		pprocExists := !reflect.ValueOf(pprocID).IsZero()
		ct := engine.Mapper.MapStr("sf.container.id")(rec)
		ctExists := !reflect.ValueOf(ct).IsZero()
		for k, v := range engine.Mapper.Mappers {
			kc := strings.Split(k, ".")
			value := extractValue(k, v(rec))
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
					if r.Type == "NF" {
						if r.NetData == nil {
							r.NetData = new(NetData)
							r.NetData.Net = make(map[string]interface{})
						}
						r.Net[kc[2]] = value
					}
				case file:
					if r.Type == "FF" || r.Type == "FE" {
						if r.FileData == nil {
							r.FileData = new(FileData)
							r.FileData.File = make(map[string]interface{})
						}
						r.File[kc[2]] = value
					}
				case flow:
					if r.Type == "FF" || r.Type == "NF" {
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
	return k == "sf.opflags" || k == "sf.proc.apid" || k == "sf.proc.aname" ||
		k == "sf.proc.aexe" || k == "sf.proc.acmdline" || k == "sf.file.openflags" ||
		k == "sf.net.ip" || k == "sf.net.port"
}
