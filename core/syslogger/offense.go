package syslogger

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.ibm.com/sysflow/sf-processor/core/sfpe/engine"
)

// Offense type
type Offense struct {
	GroupID      string        `json:"groupId"`
	Observations []Observation `json:"observations"`
}

// Policy type
type Policy struct {
	ID       string   `json:"id"`
	Desc     string   `json:"desc"`
	Priority int      `json:"priority"`
	Tags     []string `json:"tags"`
}

// Observation type
type Observation struct {
	Type          string   `json:"type"`
	TS            int64    `json:"ts"`
	EndTS         int64    `json:"endts"`
	PID           int64    `json:"procId"`
	TID           int64    `json:"threadId"`
	PExe          string   `json:"procExe"`
	PArgs         string   `json:"procArgs"`
	PpID          int64    `json:"pProcId"`
	PpExe         string   `json:"pProcExe"`
	PpArgs        string   `json:"pProcArgs"`
	PAExe         []string `json:"procAncExes"`
	PAPID         []string `json:"procAncIds"`
	FilePath      string   `json:"filePath"`
	OpFlags       []string `json:"opFlags"`
	SIP           string   `json:"sIp"`
	SPort         int64    `json:"sPort"`
	DIP           string   `json:"dIp"`
	DPort         int64    `json:"dPort"`
	Proto         string   `json:"proto"`
	ContID        string   `json:"contId"`
	ContName      string   `json:"contName"`
	ContImageID   string   `json:"contImageId"`
	ContImageName string   `json:"contImageName"`
	Policies      []Policy `json:"policies"`
}

// CreateOffenses creates offense instances based on a list of records
func CreateOffenses(occs []*engine.Record) []*Offense {
	var offenses = make([]*Offense, 0)
	var cobs = make(map[string][]Observation)
	for _, o := range extractObservations(occs) {
		if m, ok := cobs[o.ContID]; ok {
			cobs[o.ContID] = append(m, o)
		} else {
			cobs[o.ContID] = append(make([]Observation, 0), o)
		}
	}
	for k, v := range cobs {
		o := &Offense{
			GroupID:      k,
			Observations: v,
		}
		offenses = append(offenses, o)
	}
	return offenses
}

// ToJSONStr returns a JSON string representation of an offense
func (s *Offense) ToJSONStr() string {
	return string(s.ToJSON())
}

// ToJSON returns a JSON bytearray representation of an offense
func (s *Offense) ToJSON() []byte {
	o, _ := json.Marshal(s)
	return o
}

func extractObservations(recs []*engine.Record) []Observation {
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
			Policies:      extractPolicySet(r.Ctx.GetRules()),
		}
		observations = append(observations, o)
	}
	return observations
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
