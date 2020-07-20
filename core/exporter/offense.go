package exporter

import (
	"encoding/json"

	"github.ibm.com/sysflow/sf-processor/core/policyengine/engine"
)

// Offense type
type Offense struct {
	GroupID      string            `json:"groupId"`
	Observations []TelemetryRecord `json:"observations"`
}

// Policy type
type Policy struct {
	ID       string   `json:"id"`
	Desc     string   `json:"desc"`
	Priority int      `json:"priority"`
	Tags     []string `json:"tags"`
}

// CreateOffenses creates offense instances based on a list of records
func CreateOffenses(recs []*engine.Record, config Config) []Event {
	var offenses = make([]Event, 0)
	var cobs = make(map[string][]TelemetryRecord)
	for i, o := range extractObservations(recs, config) {
		contID := engine.Mapper.MapStr("sf.container.id")(recs[i])
		if m, ok := cobs[contID]; ok {
			cobs[contID] = append(m, o)
		} else {
			cobs[contID] = append(make([]TelemetryRecord, 0), o)
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

// CreateObservations creates offense instances based on a list of records
func CreateObservations(recs []*engine.Record, config Config) []Event {
	var observations = make([]Event, 0)
	for _, o := range extractObservations(recs, config) {
		observations = append(observations, o)
	}
	return observations
}

func extractObservations(recs []*engine.Record, config Config) []TelemetryRecord {
	var observations = make([]TelemetryRecord, 0)
	for _, r := range recs {
		o := extractTelemetryRecord(r, config)
		observations = append(observations, o)
	}
	return observations
}
