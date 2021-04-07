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
package exporter

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
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
		groupID := engine.Mapper.MapStr(engine.SF_NODE_ID)(recs[i])
		contID := engine.Mapper.MapStr(engine.SF_CONTAINER_ID)(recs[i])
		if contID != sfgo.Zeros.String {
			groupID = fmt.Sprintf("%s/%s", groupID, contID)
		}
		if m, ok := cobs[contID]; ok {
			cobs[groupID] = append(m, o)
		} else {
			cobs[groupID] = append(make([]TelemetryRecord, 0), o)
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
	var b bytes.Buffer
	b.WriteString("{\"groupId\":")
	m, _ := json.Marshal(s.GroupID)
	b.Write(m)
	b.WriteString(",\"observations\":[")
	for i, tr := range s.Observations {
		b.Write(tr.ToJSON())
		if i + 1 < len(s.Observations) {
			b.WriteString(",")
		}
	}
	b.WriteString("]}")
	return b.Bytes()

// Old code here:
//	o, _ := json.Marshal(s)
//	return o
}

func (s Offense) ID() string {
	var b bytes.Buffer
	for _, tr := range s.Observations {
		b.WriteString(tr.ID())
	}
	return Sha256Hex(b.Bytes())
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
