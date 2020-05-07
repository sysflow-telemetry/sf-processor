package exporter

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.ibm.com/sysflow/sf-processor/core/sfpe/engine"
)

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
	*FlatRecord `json:",omitempty"`
	*DataRecord `json:",omitempty"`
}

// FlatRecord type
type FlatRecord struct {
	Data map[string]interface{} `json:"data"`
}

// DataRecord type (warning: make sure field names have only first letter capitalized)
type DataRecord struct {
	Type       string `json:"type"`
	Opflags    string `json:"opflags"`
	Ret        int64  `json:"ret"`
	Ts         int64  `json:"ts"`
	Endts      int64  `json:"endts,omitempty"`
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

// CreateTelemetryRecord creates offense instances based on a list of records
func CreateTelemetryRecord(occs []*engine.Record, config Config) []Event {
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
			value := v(rec)
			if len(kc) == 2 {
				switch value.(type) {
				case string:
					reflect.ValueOf(r.DataRecord).Elem().FieldByName(strings.Title(kc[1])).SetString(value.(string))
				case int64:
					reflect.ValueOf(r.DataRecord).Elem().FieldByName(strings.Title(kc[1])).SetInt(value.(int64))
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
	return r
}

// Observation type
type Observation struct {
	Type          string         `json:"type"`
	TS            int64          `json:"ts"`
	EndTS         int64          `json:"endts"`
	PID           int64          `json:"procId"`
	TID           int64          `json:"threadId"`
	PExe          string         `json:"procExe"`
	PArgs         string         `json:"procArgs"`
	PpID          int64          `json:"pProcId"`
	PpExe         string         `json:"pProcExe"`
	PpArgs        string         `json:"pProcArgs"`
	PAExe         []string       `json:"procAncExes"`
	PAPID         []string       `json:"procAncIds"`
	FilePath      string         `json:"filePath"`
	OpFlags       []string       `json:"opFlags"`
	SIP           string         `json:"sIp"`
	SPort         int64          `json:"sPort"`
	DIP           string         `json:"dIp"`
	DPort         int64          `json:"dPort"`
	Proto         string         `json:"proto"`
	ContID        string         `json:"contId"`
	ContName      string         `json:"contName"`
	ContImageID   string         `json:"contImageId"`
	ContImageName string         `json:"contImageName"`
	Hashes        engine.HashSet `json:"hashes"`
	Policies      []Policy       `json:"policies"`
}

// CreateObservations creates offense instances based on a list of records
func CreateObservations(occs []*engine.Record) []Event {
	var observations = make([]Event, 0)
	for _, o := range ExtractObservations(occs) {
		observations = append(observations, o)
	}
	return observations
}

// ToJSONStr returns a JSON string representation of an observation
func (s Observation) ToJSONStr() string {
	return string(s.ToJSON())
}

// ToJSON returns a JSON bytearray representation of an observation
func (s Observation) ToJSON() []byte {
	o, _ := json.Marshal(s)
	return o
}

// ExtractObservations produces a list of observations based on an occ set.
func ExtractObservations(recs []*engine.Record) []Observation {
	var observations = make([]Observation, 0)
	for _, r := range recs {
		o := Observation{
			Type:          engine.Mapper.MapStr("sf.type")(r),
			TS:            engine.Mapper.MapInt("sf.ts")(r),
			PID:           engine.Mapper.MapInt("sf.proc.pid")(r),
			TID:           engine.Mapper.MapInt("sf.proc.tid")(r),
			PExe:          engine.Mapper.MapStr("sf.proc.exe")(r),
			PArgs:         engine.Mapper.MapStr("sf.proc.args")(r),
			PpID:          engine.Mapper.MapInt("sf.pproc.pid")(r),
			PpExe:         engine.Mapper.MapStr("sf.pproc.exe")(r),
			PpArgs:        engine.Mapper.MapStr("sf.pproc.args")(r),
			PAExe:         strings.Split(engine.Mapper.MapStr("sf.proc.aexe")(r), engine.LISTSEP),
			PAPID:         strings.Split(engine.Mapper.MapStr("sf.proc.apid")(r), engine.LISTSEP),
			FilePath:      engine.Mapper.MapStr("sf.file.path")(r),
			SIP:           engine.Mapper.MapStr("sf.net.sip")(r),
			SPort:         engine.Mapper.MapInt("sf.net.sport")(r),
			DIP:           engine.Mapper.MapStr("sf.net.dip")(r),
			DPort:         engine.Mapper.MapInt("sf.net.dport")(r),
			Proto:         engine.Mapper.MapStr("sf.net.proto")(r),
			OpFlags:       strings.Split(engine.Mapper.MapStr("sf.opflags")(r), engine.LISTSEP),
			ContID:        engine.Mapper.MapStr("sf.container.id")(r),
			ContName:      engine.Mapper.MapStr("sf.container.name")(r),
			ContImageID:   engine.Mapper.MapStr("sf.container.imageid")(r),
			ContImageName: engine.Mapper.MapStr("sf.container.image")(r),
			Hashes:        r.Ctx.GetHashes(),
			Policies:      extractPolicySet(r.Ctx.GetRules()),
		}
		observations = append(observations, o)
	}
	return observations
}
