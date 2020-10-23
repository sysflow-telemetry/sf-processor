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
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/sf-processor/core/policyengine/engine"
)

// SysFlow record components
const (
	PROC      = "proc"
	PPROC     = "pproc"
	NET       = "net"
	FILEF     = "file"
	FLOW      = "flow"
	CONTAINER = "container"
	NODE      = "node"

	VERSION_STR = "{\"version\":"
	BEGIN_STATE = iota
	PROC_STATE
	PPROC_STATE
	NET_STATE
	FILE_STATE
	FLOW_STATE
	CONT_STATE
	NODE_STATE
)

type JSONExporter struct {
	fieldCache [][]string
	proto      ExportProtocol
	config     Config
	buf        *bytes.Buffer
}

func NewJSONExporter(p ExportProtocol, config Config) *JSONExporter {
	t := &JSONExporter{}
	t.fieldCache = make([][]string, len(engine.Fields))
	for ind, k := range engine.Fields {
		t.fieldCache[ind] = strings.Split(k, ".")
	}
	t.proto = p
	t.config = config
	buffer := make([]byte, 0, 2048)
	t.buf = bytes.NewBuffer(buffer)
	return t
}

func (t *JSONExporter) exportOffense(recs []*engine.Record, groupID string, contID string) error {
	t.buf.WriteString("{\"groupId\":\"")
	t.buf.WriteString(groupID)
	if contID != sfgo.Zeros.String {
		t.buf.WriteByte('/')
		t.buf.WriteString(contID)
	}
	t.buf.WriteString("\",")
	t.buf.WriteString("\"observations\":[")
	for _, rec := range recs {
		t.encodeTelemetry(rec)
		t.buf.WriteByte(',')
	}
	t.buf.WriteString("]}")
	return t.proto.Export(t.buf)

}

func (t *JSONExporter) ExportOffenses(recs []*engine.Record) error {
	if len(recs) == 1 {
		groupID := engine.Mapper.MapStr(engine.SF_NODE_ID)(recs[0])
		contID := engine.Mapper.MapStr(engine.SF_CONTAINER_ID)(recs[0])
		t.buf.Reset()
		err := t.exportOffense(recs, groupID, contID)
		return err
	} else {
		var cobs = make(map[string][]*engine.Record)
		for _, rec := range recs {
			groupID := engine.Mapper.MapStr(engine.SF_NODE_ID)(rec)
			contID := engine.Mapper.MapStr(engine.SF_CONTAINER_ID)(rec)
			if contID != sfgo.Zeros.String {
				groupID = fmt.Sprintf("%s/%s", groupID, contID)
			}
			if m, ok := cobs[contID]; ok {
				cobs[groupID] = append(m, rec)
			} else {
				cobs[groupID] = append(make([]*engine.Record, 0), rec)
			}
		}
		for k, v := range cobs {
			t.buf.Reset()
			err := t.exportOffense(v, k, sfgo.Zeros.String)
			if err != nil {
				return err
			}

		}

	}
	return nil

}

func (t *JSONExporter) ExportTelemetryRecords(recs []*engine.Record) error {
	for _, rec := range recs {
		t.buf.Reset()
		t.encodeTelemetry(rec)
		err := t.proto.Export(t.buf)
		if err != nil {
			return err
		}
	}
	return nil

}

func (t *JSONExporter) writeAttribute(k string, fieldIdx int, fieldId int, rec *engine.Record) {
	t.buf.WriteByte('"')
	t.buf.WriteString(t.fieldCache[fieldIdx][fieldId])
	t.buf.WriteString("\":")
	t.buf.WriteString(engine.Mapper.MapBuffer(k, t.buf)(rec))
}

func (t *JSONExporter) writeSectionBegin(section string) {
	t.buf.WriteByte('"')
	t.buf.WriteString(section)
	t.buf.WriteString("\":{")
}

func (t *JSONExporter) encodeTelemetry(rec *engine.Record) {
	t.buf.WriteString(VERSION_STR)
	t.buf.WriteString(t.config.JSONSchemaVersion)
	t.buf.WriteByte(',')
	state := BEGIN_STATE
	pprocID := engine.Mapper.MapInt(engine.SF_PPROC_PID)(rec)
	sftype := engine.Mapper.MapStr(engine.SF_TYPE)(rec)
	pprocExists := !reflect.ValueOf(pprocID).IsZero()
	ct := engine.Mapper.MapStr(engine.SF_CONTAINER_ID)(rec)
	ctExists := !reflect.ValueOf(ct).IsZero()
	existed := true
	/* //Need to add flat support
		if config.Flat {
	                r.FlatRecord = new(FlatRecord)
	                r.FlatRecord.Data = make(map[string]interface{})
	                for _, k := range engine.Fields {
	                        r.Data[k] = engine.Mapper.Mappers[k](rec)
	                }*/

	for ind, k := range engine.Fields {
		numFields := len(t.fieldCache[ind])
		if numFields == 2 {
			t.writeAttribute(k, ind, 1, rec)
			t.buf.WriteByte(',')
		} else if numFields == 3 {
			switch t.fieldCache[ind][1] {
			case PROC:
				if state != PROC_STATE {
					if state != BEGIN_STATE && existed {
						t.buf.WriteString("},")
					}
					existed = true
					t.writeSectionBegin(PROC)
					t.writeAttribute(k, ind, 2, rec)
					state = PROC_STATE
				} else {
					t.buf.WriteByte(',')
					t.writeAttribute(k, ind, 2, rec)
				}
			case PPROC:
				if state != PPROC_STATE {
					if state != BEGIN_STATE && existed {
						t.buf.WriteString("},")
					}
					if pprocExists {
						existed = true
						t.writeSectionBegin(PPROC)
						t.writeAttribute(k, ind, 2, rec)
					} else {
						existed = false
					}
					state = PPROC_STATE
				} else if pprocExists {
					t.buf.WriteByte(',')
					t.writeAttribute(k, ind, 2, rec)
				}
			case NET:
				if state != NET_STATE {
					if state != BEGIN_STATE && existed {
						t.buf.WriteString("},")
					}
					if sftype == engine.TyNF {
						t.writeSectionBegin(NET)
						t.writeAttribute(k, ind, 2, rec)
						existed = true
					} else {
						existed = false
					}
					state = NET_STATE
				} else if sftype == engine.TyNF {
					t.buf.WriteByte(',')
					t.writeAttribute(k, ind, 2, rec)
				}
			case FILEF:
				if state != FILE_STATE {
					if state != BEGIN_STATE && existed {
						t.buf.WriteString("},")
					}
					if sftype == engine.TyFF || sftype == engine.TyFE {
						t.writeSectionBegin(FILEF)
						t.writeAttribute(k, ind, 2, rec)
						existed = true
					} else {
						existed = false
					}
					state = FILE_STATE
				} else if sftype == engine.TyFF || sftype == engine.TyFE {
					t.buf.WriteByte(',')
					t.writeAttribute(k, ind, 2, rec)
				}
			case FLOW:
				if state != FLOW_STATE {
					if state != BEGIN_STATE && existed {
						t.buf.WriteString("},")
					}
					if sftype == engine.TyFF || sftype == engine.TyNF {
						t.writeSectionBegin(FLOW)
						t.writeAttribute(k, ind, 2, rec)
						existed = true
					} else {
						existed = false
					}
					state = FLOW_STATE
				} else if sftype == engine.TyFF || sftype == engine.TyNF {
					t.buf.WriteByte(',')
					t.writeAttribute(k, ind, 2, rec)
				}
			case CONTAINER:
				if state != CONT_STATE {
					if state != BEGIN_STATE && existed {
						t.buf.WriteString("},")
					}
					if ctExists {
						t.writeSectionBegin(CONTAINER)
						t.writeAttribute(k, ind, 2, rec)
						existed = true
					} else {
						existed = false
					}
					state = CONT_STATE
				}
				if ctExists {
					t.buf.WriteByte(',')
					t.writeAttribute(k, ind, 2, rec)
				}
			case NODE:
				if state != NODE_STATE {
					if state != BEGIN_STATE && existed {
						t.buf.WriteString("},")
					}
					existed = true
					t.writeSectionBegin(NODE)
					t.writeAttribute(k, ind, 2, rec)
					state = NODE_STATE
				}
				t.buf.WriteByte(',')
				t.writeAttribute(k, ind, 2, rec)
			}
		}

	}
	t.buf.WriteString("},")
	/* // Need to add hash support
	hashset := rec.Ctx.GetHashes()
	if !reflect.ValueOf(hashset.MD5).IsZero() {
		r.Hashes = &hashset
	} */
	rules := rec.Ctx.GetRules()
	if len(rules) > 0 {
		t.buf.WriteString("\"policies\":[")
		for _, r := range rules {
			t.buf.WriteString("{\"id\":\"")
			t.buf.WriteString(r.Name)
			t.buf.WriteString("\",\"desc\":\"")
			t.buf.WriteString(r.Desc)
			t.buf.WriteString("\",\"priority\":")
			t.buf.WriteString(strconv.Itoa(int(r.Priority)))
			if len(r.Tags) > 0 {
				t.buf.WriteString(",\"tags\":[")
				for _, tag := range r.Tags {
					switch tag.(type) {
					case []string:
						tags := tag.([]string)
						for _, s := range tags {
							t.buf.WriteByte('"')
							t.buf.WriteString(s)
							t.buf.WriteByte('"')
							t.buf.WriteByte(',')
						}
					default:
						t.buf.WriteByte('"')
						t.buf.WriteString(fmt.Sprintf("%v", tag))
						t.buf.WriteByte('"')
						t.buf.WriteByte(',')
					}
				}
				t.buf.WriteByte(']')
			}
			t.buf.WriteString("},")

		}
		t.buf.WriteByte(']')
	}
	t.buf.WriteByte('}')

}
