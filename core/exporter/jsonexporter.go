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
	"fmt"
	"reflect"

	"github.com/mailru/easyjson/jwriter"
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

	BEGIN_STATE = iota
	PROC_STATE
	PPROC_STATE
	NET_STATE
	FILE_STATE
	FLOW_STATE
	CONT_STATE
	NODE_STATE

	BUFFER_SIZE = 10240
)

// JSONExporter  implements a JSON serializer and exporter
type JSONExporter struct {
	fieldCache []*engine.FieldValue
	proto      ExportProtocol
	config     Config
	buf        []byte
	writer     *jwriter.Writer
}

// NewJSONExporter instantiates a JSON exporter
func NewJSONExporter(p ExportProtocol, config Config) *JSONExporter {
	t := &JSONExporter{}
	t.fieldCache = engine.FieldValues
	t.proto = p
	t.config = config
	t.writer = &jwriter.Writer{}
	t.buf = make([]byte, 0, BUFFER_SIZE)
	t.writer.Buffer.Buf = t.buf
	return t
}

const (
	VERSION_STR        = "{\"version\":"
	GROUP_ID           = "{\"groupId\":\""
	FORWARD_SLASH      = '/'
	QUOTE_COMMA        = "\","
	OBSERVATIONS       = "\"observations\":["
	COMMA              = ','
	END_SQ_SQUIGGLE    = "]}"
	DOUBLE_QUOTE       = '"'
	QUOTE_COLON        = "\":"
	QUOTE_COLON_OSUIG  = "\":{"
	END_SQUIGGLE_COMMA = "},"
	END_SQUIGGLE       = '}'
	END_SQUARE         = ']'
	BEGIN_SQUARE       = '['
	POLICIES           = ",\"policies\":["
	ID_TAG             = "{\"id\":"
	DESC               = ",\"desc\":"
	PRIORITY           = ",\"priority\":"
	TAGS               = ",\"tags\":["
	PERIOD             = '.'
)

func (t *JSONExporter) exportOffense(recs []*engine.Record, groupID string, contID string) error {
	t.writer.RawString(GROUP_ID)
	t.writer.RawString(groupID)
	if contID != sfgo.Zeros.String {
		t.writer.RawByte(FORWARD_SLASH)
		t.writer.RawString(contID)
	}
	t.writer.RawString(QUOTE_COMMA)
	t.writer.RawString(OBSERVATIONS)
	numRecs := len(recs)
	for idx, rec := range recs {
		t.encodeTelemetry(rec)
		if idx < (numRecs - 1) {
			t.writer.RawByte(COMMA)
		}
	}
	t.writer.RawString(END_SQ_SQUIGGLE)
	if t.writer.Size() <= BUFFER_SIZE {
		return t.proto.Export(t.writer.Buffer.Buf)
	} else {
		b, err := t.writer.BuildBytes()
		if err != nil {
			return err
		}
		return t.proto.Export(b)
	}

}

// ExportOffenses exports a set of  offesnes as JSON objects
func (t *JSONExporter) ExportOffenses(recs []*engine.Record) error {
	if len(recs) == 1 {
		groupID := engine.Mapper.MapStr(engine.SF_NODE_ID)(recs[0])
		contID := engine.Mapper.MapStr(engine.SF_CONTAINER_ID)(recs[0])
		t.buf = t.buf[:0]
		t.writer.Buffer.Buf = t.buf
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
			t.buf = t.buf[:0]
			t.writer.Buffer.Buf = t.buf
			err := t.exportOffense(v, k, sfgo.Zeros.String)
			if err != nil {
				return err
			}

		}

	}
	return nil

}

// ExportTelemetryRecords exports a set of telemetry records as JSON objects.
func (t *JSONExporter) ExportTelemetryRecords(recs []*engine.Record) error {
	var b []byte
	var err error
	for _, rec := range recs {
		t.buf = t.buf[:0]
		t.writer.Buffer.Buf = t.buf
		t.encodeTelemetry(rec)

		if t.writer.Size() <= BUFFER_SIZE {
			b = t.writer.Buffer.Buf
		} else {
			b, err = t.writer.BuildBytes()
			if err != nil {
				return err
			}
		}
		err = t.proto.Export(b)
		if err != nil {
			return err
		}
	}
	return nil

}

func (t *JSONExporter) writeAttribute(fv *engine.FieldValue, fieldId int, rec *engine.Record) {
	t.writer.RawByte(DOUBLE_QUOTE)
	t.writer.RawString(fv.FieldSects[fieldId])
	t.writer.RawString(QUOTE_COLON)
	MapJSON(fv, t.writer, rec)
}

func (t *JSONExporter) writeSectionBegin(section string) {
	t.writer.RawByte(DOUBLE_QUOTE)
	t.writer.RawString(section)
	t.writer.RawString(QUOTE_COLON_OSUIG)
}

func (t *JSONExporter) encodeTelemetry(rec *engine.Record) {
	t.writer.RawString(VERSION_STR)
	t.writer.RawString(t.config.JSONSchemaVersion)
	t.writer.RawByte(COMMA)
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

	for _, fv := range t.fieldCache {
		numFields := len(fv.FieldSects)
		if numFields == 2 {
			t.writeAttribute(fv, 1, rec)
			t.writer.RawByte(COMMA)
		} else if numFields == 3 {
			switch fv.Entry.Section {
			case engine.SectProc:
				if state != PROC_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_SQUIGGLE_COMMA)
					}
					existed = true
					t.writeSectionBegin(PROC)
					t.writeAttribute(fv, 2, rec)
					state = PROC_STATE
				} else {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			case engine.SectPProc:
				if state != PPROC_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_SQUIGGLE_COMMA)
					}
					if pprocExists {
						existed = true
						t.writeSectionBegin(PPROC)
						t.writeAttribute(fv, 2, rec)
					} else {
						existed = false
					}
					state = PPROC_STATE
				} else if pprocExists {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			case engine.SectNet:
				if state != NET_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_SQUIGGLE_COMMA)
					}
					if sftype == engine.TyNF {
						t.writeSectionBegin(NET)
						t.writeAttribute(fv, 2, rec)
						existed = true
					} else {
						existed = false
					}
					state = NET_STATE
				} else if sftype == engine.TyNF {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			case engine.SectFile:
				if state != FILE_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_SQUIGGLE_COMMA)
					}
					if sftype == engine.TyFF || sftype == engine.TyFE {
						t.writeSectionBegin(FILEF)
						t.writeAttribute(fv, 2, rec)
						existed = true
					} else {
						existed = false
					}
					state = FILE_STATE
				} else if sftype == engine.TyFF || sftype == engine.TyFE {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			case engine.SectFlow:
				if state != FLOW_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_SQUIGGLE_COMMA)
					}
					if sftype == engine.TyFF || sftype == engine.TyNF {
						t.writeSectionBegin(FLOW)
						t.writeAttribute(fv, 2, rec)
						existed = true
					} else {
						existed = false
					}
					state = FLOW_STATE
				} else if sftype == engine.TyFF || sftype == engine.TyNF {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			case engine.SectCont:
				if state != CONT_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_SQUIGGLE_COMMA)
					}
					if ctExists {
						t.writeSectionBegin(CONTAINER)
						t.writeAttribute(fv, 2, rec)
						existed = true
					} else {
						existed = false
					}
					state = CONT_STATE
				} else if ctExists {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			case engine.SectNode:
				if state != NODE_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_SQUIGGLE_COMMA)
					}
					existed = true
					t.writeSectionBegin(NODE)
					t.writeAttribute(fv, 2, rec)
					state = NODE_STATE
				} else {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			}
		}

	}
	t.writer.RawByte(END_SQUIGGLE)
	/* // Need to add hash support
	hashset := rec.Ctx.GetHashes()
	if !reflect.ValueOf(hashset.MD5).IsZero() {
		r.Hashes = &hashset
	} */
	rules := rec.Ctx.GetRules()
	numRules := len(rules)
	if numRules > 0 {
		t.writer.RawString(POLICIES)

		for id, r := range rules {
			t.writer.RawString(ID_TAG)
			t.writer.String(r.Name)
			t.writer.RawString(DESC)
			t.writer.String(r.Desc)
			t.writer.RawString(PRIORITY)
			t.writer.Int64(int64(r.Priority))
			numTags := len(r.Tags)
			currentTag := 0
			if numTags > 0 {
				t.writer.RawString(TAGS)
				for _, tag := range r.Tags {
					switch tag.(type) {
					case []string:
						tags := tag.([]string)
						numTags := numTags + len(tags) - 1
						for _, s := range tags {
							t.writer.String(s)
							if currentTag < (numTags - 1) {
								t.writer.RawByte(COMMA)
							}
							currentTag += 1
						}
					default:
						//t.writer.RawByte(DOUBLE_QUOTE)
						t.writer.String(tag.(string)) //fmt.Sprintf("%v", tag))
						//t.writer.RawByte(DOUBLE_QUOTE)
						if currentTag < (numTags - 1) {
							t.writer.RawByte(COMMA)
						}
						currentTag += 1
					}
				}
				t.writer.RawByte(END_SQUARE)
			}
			t.writer.RawByte(END_SQUIGGLE)
			if id < (numRules - 1) {
				t.writer.RawByte(COMMA)
			}

		}
		t.writer.RawByte(END_SQUARE)
	}
	t.writer.RawByte(END_SQUIGGLE)

}
