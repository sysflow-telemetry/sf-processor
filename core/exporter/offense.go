package exporter

import (
	"encoding/json"
	"fmt"

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

// CreateOffenses creates offense instances based on a list of records
func CreateOffenses(occs []*engine.Record) []Event {
	var offenses = make([]Event, 0)
	var cobs = make(map[string][]Observation)
	for _, o := range ExtractObservations(occs) {
		if m, ok := cobs[o.ContID]; ok {
			cobs[o.ContID] = append(m, o)
		} else {
			cobs[o.ContID] = append(make([]Observation, 0), o)
		}
	}
	for k, v := range cobs {
		o := Offense{
			GroupID:      k,
			Observations: v,
		}
		offenses = append(offenses, o)
	}
	return offenses
}

// ToJSONStr returns a JSON string representation of an offense
func (s Offense) ToJSONStr() string {
	return string(s.ToJSON())
}

// ToJSON returns a JSON bytearray representation of an offense
func (s Offense) ToJSON() []byte {
	o, _ := json.Marshal(s)
	return o
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
