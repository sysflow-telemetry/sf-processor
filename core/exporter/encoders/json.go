//go:build flatrecord
// +build flatrecord

//
// Copyright (C) 2021 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
// Andreas Schade <san@zurich.ibm.com>
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

// Package encoders implements codecs for exporting records and events in different data formats.
package encoders

import (
	"path/filepath"
	"reflect"
	"strings"
	"unicode/utf8"

	"github.com/mailru/easyjson/jwriter"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/utils"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/flatrecord"
)

// JSONEncoder is a JSON encoder.
type JSONEncoder struct {
	config     commons.Config
	fieldCache []*flatrecord.FieldValue
	writer     *jwriter.Writer
	buf        []byte
	batch      []commons.EncodedData
}

// NewJSONEncoder instantiates a JSON encoder.
func NewJSONEncoder(config commons.Config) Encoder {
	return &JSONEncoder{
		fieldCache: flatrecord.FieldValues,
		config:     config,
		writer:     &jwriter.Writer{},
		buf:        make([]byte, 0, BUFFER_SIZE),
		batch:      make([]commons.EncodedData, 0, config.EventBuffer)}
}

// Register registers the encoder to the codecs cache.
func (t *JSONEncoder) Register(codecs map[commons.Format]EncoderFactory) {
	codecs[commons.JSONFormat] = NewJSONEncoder
}

// Encode encodes telemetry records into a JSON representation.
func (t *JSONEncoder) Encode(recs []*flatrecord.Record) (data []commons.EncodedData, err error) {
	t.batch = t.batch[:0]
	for _, rec := range recs {
		var j commons.EncodedData
		if j, err = t.encode(rec); err != nil {
			return nil, err
		}
		t.batch = append(t.batch, j)
	}
	return t.batch, nil
}

// Encodes a telemetry record into a JSON representation.
func (t *JSONEncoder) encode(rec *flatrecord.Record) (commons.EncodedData, error) {
	t.writer.RawString(VERSION_STR)
	t.writer.RawString(t.config.JSONSchemaVersion)
	t.writer.RawByte(COMMA)
	state := BEGIN_STATE
	sftype := flatrecord.Mapper.MapStr(flatrecord.SF_TYPE)(rec)

	pprocID := flatrecord.Mapper.MapInt(flatrecord.SF_PPROC_PID)(rec)
	pprocExists := !reflect.ValueOf(pprocID).IsZero()
	ct := flatrecord.Mapper.MapStr(flatrecord.SF_CONTAINER_ID)(rec)
	ctExists := !reflect.ValueOf(ct).IsZero()
	pd := flatrecord.Mapper.MapStr(flatrecord.SF_POD_ID)(rec)
	pdExists := !reflect.ValueOf(pd).IsZero()
	existed := true

	for _, fv := range t.fieldCache {
		numFields := len(fv.FieldSects)
		if numFields == 2 {
			t.writeAttribute(fv, 1, rec)
			t.writer.RawByte(COMMA)
		} else if numFields == 3 {
			if sftype == sfgo.TyKEStr {
				switch fv.Entry.Section {
				case flatrecord.SectK8sEvt:
					if state != KE_STATE {
						if state != BEGIN_STATE && existed {
							t.writer.RawString(END_CURLY_COMMA)
						}
						existed = true
						t.writeSectionBegin(KE)
						t.writeAttribute(fv, 2, rec)
						state = KE_STATE
					} else {
						t.writer.RawByte(COMMA)
						t.writeAttribute(fv, 2, rec)
					}
				case flatrecord.SectNode:
					if state != NODE_STATE {
						if state != BEGIN_STATE && existed {
							t.writer.RawString(END_CURLY_COMMA)
						}
						existed = true
						t.writeSectionBegin(NODE)
						t.writeAttribute(fv, 2, rec)
						state = NODE_STATE
					} else {
						t.writer.RawByte(COMMA)
						t.writeAttribute(fv, 2, rec)
					}
				case flatrecord.SectMeta:
					if state != META_STATE {
						if state != BEGIN_STATE && existed {
							t.writer.RawString(END_CURLY_COMMA)
						}
						existed = true
						t.writeSectionBegin(META)
						t.writeAttribute(fv, 2, rec)
						state = META_STATE
					} else {
						t.writer.RawByte(COMMA)
						t.writeAttribute(fv, 2, rec)
					}
				}
				continue
			}

			switch fv.Entry.Section {
			case flatrecord.SectProc:
				if state != PROC_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_CURLY_COMMA)
					}
					existed = true
					t.writeSectionBegin(PROC)
					t.writeAttribute(fv, 2, rec)
					state = PROC_STATE
				} else {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			case flatrecord.SectPProc:
				if state != PPROC_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_CURLY_COMMA)
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
			case flatrecord.SectNet:
				if state != NET_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_CURLY_COMMA)
					}
					if sftype == sfgo.TyNFStr {
						t.writeSectionBegin(NET)
						t.writeAttribute(fv, 2, rec)
						existed = true
					} else {
						existed = false
					}
					state = NET_STATE
				} else if sftype == sfgo.TyNFStr {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			case flatrecord.SectFile:
				if state != FILE_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_CURLY_COMMA)
					}
					if sftype == sfgo.TyFFStr || sftype == sfgo.TyFEStr {
						t.writeSectionBegin(FILEF)
						t.writeAttribute(fv, 2, rec)
						existed = true
					} else {
						existed = false
					}
					state = FILE_STATE
				} else if sftype == sfgo.TyFFStr || sftype == sfgo.TyFEStr {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			case flatrecord.SectFlow:
				if state != FLOW_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_CURLY_COMMA)
					}
					if sftype == sfgo.TyFFStr || sftype == sfgo.TyNFStr {
						t.writeSectionBegin(FLOW)
						t.writeAttribute(fv, 2, rec)
						existed = true
					} else {
						existed = false
					}
					state = FLOW_STATE
				} else if sftype == sfgo.TyFFStr || sftype == sfgo.TyNFStr {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			case flatrecord.SectCont:
				if state != CONT_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_CURLY_COMMA)
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
			case flatrecord.SectPod:
				if state != POD_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_CURLY_COMMA)
					}
					if pdExists {
						t.writeSectionBegin(POD)
						t.writeAttribute(fv, 2, rec)
						existed = true
					} else {
						existed = false
					}
					state = POD_STATE
				} else if pdExists {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			case flatrecord.SectNode:
				if state != NODE_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_CURLY_COMMA)
					}
					existed = true
					t.writeSectionBegin(NODE)
					t.writeAttribute(fv, 2, rec)
					state = NODE_STATE
				} else {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			case flatrecord.SectMeta:
				if state != META_STATE {
					if state != BEGIN_STATE && existed {
						t.writer.RawString(END_CURLY_COMMA)
					}
					existed = true
					t.writeSectionBegin(META)
					t.writeAttribute(fv, 2, rec)
					state = META_STATE
				} else {
					t.writer.RawByte(COMMA)
					t.writeAttribute(fv, 2, rec)
				}
			}
		}
	}
	t.writer.RawByte(END_CURLY)

	// Encode policies
	numRules := len(rec.Ctx.GetRules())
	rtags := make([]string, 0)
	if numRules > 0 {
		t.writer.RawString(POLICIES)
		for num, r := range rec.Ctx.GetRules() {
			t.writer.RawString(ID_TAG)
			t.writer.String(r.Name)
			t.writer.RawString(DESC)
			t.writer.String(r.Desc)
			t.writer.RawString(PRIORITY)
			t.writer.Int64(int64(r.Priority))
			t.writer.RawByte(END_CURLY)
			if num < (numRules - 1) {
				t.writer.RawByte(COMMA)
			}

			for _, tag := range r.Tags {
				switch tag := tag.(type) {
				case []string:
					rtags = append(rtags, tag...)
				default:
					rtags = append(rtags, tag.(string))
				}
			}
		}
		t.writer.RawByte(END_SQUARE)
	}

	// Encode tags as a list of record tag context plus all rule tags
	numTags := len(rtags) + len(rec.Ctx.GetTags())
	if numTags > 0 {
		currentTag := 0
		t.writer.RawString(TAGS)
		for _, tag := range rec.Ctx.GetTags() {
			t.writer.String(tag)
			if currentTag < (numTags - 1) {
				t.writer.RawByte(COMMA)
			}
			currentTag++
		}
		for _, tag := range rtags {
			t.writer.String(tag)
			if currentTag < (numTags - 1) {
				t.writer.RawByte(COMMA)
			}
			currentTag++
		}
		t.writer.RawByte(END_SQUARE)
	}
	t.writer.RawByte(END_CURLY)

	// BuildBytes returns writer data as a single byte slice. It tries to reuse buf.
	//return t.writer.BuildBytes(t.buf)
	return t.writer.BuildBytes()
}

func (t *JSONEncoder) writeAttribute(fv *flatrecord.FieldValue, fieldID int, rec *flatrecord.Record) {
	t.writer.RawByte(DOUBLE_QUOTE)
	name := fv.FieldSects[fieldID]
	if strings.HasSuffix(name, "+") {
		t.writer.RawString(name[:len(name)-1])
	} else {
		t.writer.RawString(name)
	}
	t.writer.RawString(QUOTE_COLON)
	MapJSON(fv, t.writer, rec)
}

func (t *JSONEncoder) writeSectionBegin(section string) {
	t.writer.RawByte(DOUBLE_QUOTE)
	t.writer.RawString(section)
	t.writer.RawString(QUOTE_COLON_CURLY)
}

func mapOpFlags(fv *flatrecord.FieldValue, writer *jwriter.Writer, r *flatrecord.Record) {
	opflags := r.GetInt(fv.Entry.FlatIndex, fv.Entry.Source)
	rtype, _ := sfgo.ParseRecordType(r.GetInt(sfgo.SF_REC_TYPE, fv.Entry.Source))
	flags := sfgo.GetOpFlags(int32(opflags), rtype)
	mapStrArray(writer, flags)
}

func mapStrArray(writer *jwriter.Writer, ss []string) {
	l := len(ss)
	writer.RawByte(BEGIN_SQUARE)
	for idx, s := range ss {
		writer.RawByte(DOUBLE_QUOTE)
		writer.RawString(s)
		writer.RawByte(DOUBLE_QUOTE)
		if idx < (l - 1) {
			writer.RawByte(COMMA)
		}
	}
	writer.RawByte(END_SQUARE)

}

func mapIPStr(ip int64, w *jwriter.Writer) {
	w.Int64(ip >> 0 & 0xFF)
	w.RawByte(PERIOD)
	w.Int64(ip >> 8 & 0xFF)
	w.RawByte(PERIOD)
	w.Int64(ip >> 16 & 0xFF)
	w.RawByte(PERIOD)
	w.Int64(ip >> 24 & 0xFF)
}

func mapIPs(fv *flatrecord.FieldValue, writer *jwriter.Writer, r *flatrecord.Record) {
	srcIP := r.GetInt(sfgo.FL_NETW_SIP_INT, fv.Entry.Source)
	dstIP := r.GetInt(sfgo.FL_NETW_DIP_INT, fv.Entry.Source)
	writer.RawByte(BEGIN_SQUARE)
	writer.RawByte(DOUBLE_QUOTE)
	mapIPStr(srcIP, writer)
	writer.RawByte(DOUBLE_QUOTE)
	writer.RawByte(COMMA)
	writer.RawByte(DOUBLE_QUOTE)
	mapIPStr(dstIP, writer)
	writer.RawByte(DOUBLE_QUOTE)
	writer.RawByte(END_SQUARE)
}

func mapIPArray(ips *[]int64, writer *jwriter.Writer) {
	writer.RawByte(BEGIN_SQUARE)
	for _, ip := range *ips {
		writer.RawByte(DOUBLE_QUOTE)
		mapIPStr(ip, writer)
		writer.RawByte(DOUBLE_QUOTE)
	}
	writer.RawByte(END_SQUARE)
}

func mapOpenFlags(fv *flatrecord.FieldValue, writer *jwriter.Writer, r *flatrecord.Record) {
	flags := sfgo.GetOpenFlags(r.GetInt(fv.Entry.FlatIndex, fv.Entry.Source))
	mapStrArray(writer, flags)
}

func mapPorts(fv *flatrecord.FieldValue, writer *jwriter.Writer, r *flatrecord.Record) {
	srcPort := r.GetInt(sfgo.FL_NETW_SPORT_INT, fv.Entry.Source)
	dstPort := r.GetInt(sfgo.FL_NETW_DPORT_INT, fv.Entry.Source)
	writer.RawByte(BEGIN_SQUARE)
	writer.Int64(srcPort)
	writer.RawByte(COMMA)
	writer.Int64(dstPort)
	writer.RawByte(END_SQUARE)
}

func writeStrField(writer *jwriter.Writer, name string, val string) {
	writer.RawByte(DOUBLE_QUOTE)
	writer.RawString(name)
	writer.RawString(QUOTE_COLON)
	writer.String(val)
}

func writeIntField(writer *jwriter.Writer, name string, val int32) {
	writer.RawByte(DOUBLE_QUOTE)
	writer.RawString(name)
	writer.RawString(QUOTE_COLON)
	writer.Int32(val)
}

func writeIntArrayField(writer *jwriter.Writer, name string, val *[]int64) {
	writer.RawByte(DOUBLE_QUOTE)
	writer.RawString(name)
	writer.RawString(QUOTE_COLON)
	mapIPArray(val, writer)
}

func mapPortList(writer *jwriter.Writer, ports *[]*sfgo.Port) {
	writer.RawByte(DOUBLE_QUOTE)
	writer.RawString("ports")
	writer.RawString(QUOTE_COLON)
	writer.RawByte(BEGIN_SQUARE)
	for i, p := range *ports {
		writer.RawByte(BEGIN_CURLY)
		writeIntField(writer, "port", p.Port)
		writer.RawByte(COMMA)
		writeIntField(writer, "targetport", p.TargetPort)
		writer.RawByte(COMMA)
		writeIntField(writer, "nodeport", p.NodePort)
		writer.RawByte(COMMA)
		writeStrField(writer, "proto", p.Proto)
		if (i + 1) < len(*ports) {
			writer.RawString(END_CURLY_COMMA)
		} else {
			writer.RawByte(END_CURLY)
		}
	}
	writer.RawByte(END_SQUARE)
}

func mapSvcArray(fv *flatrecord.FieldValue, writer *jwriter.Writer, r *flatrecord.Record) {
	writer.RawByte(BEGIN_SQUARE)
	for _, s := range *r.GetSvcArray(fv.Entry.FlatIndex, fv.Entry.Source) {
		writer.RawByte('{')
		writeStrField(writer, "id", s.Id)
		writer.RawByte(COMMA)
		writeStrField(writer, "name", s.Name)
		writer.RawByte(COMMA)
		writeStrField(writer, "namespace", s.Namespace)
		writer.RawByte(COMMA)
		writeIntArrayField(writer, "clusterIP", &s.ClusterIP)
		writer.RawByte(COMMA)
		mapPortList(writer, &s.PortList)
		writer.RawByte(END_CURLY)
	}
	writer.RawByte(END_SQUARE)
}

// MapJSON writes a SysFlow attribute to a JSON stream.
func MapJSON(fv *flatrecord.FieldValue, writer *jwriter.Writer, r *flatrecord.Record) {
	switch fv.Entry.FlatIndex {
	case flatrecord.A_IDS, flatrecord.PARENT_IDS:
		oid := sfgo.OID{CreateTS: r.GetInt(sfgo.PROC_OID_CREATETS_INT, fv.Entry.Source), Hpid: r.GetInt(sfgo.PROC_OID_HPID_INT, fv.Entry.Source)}
		setCachedValueToJSON(r, oid, fv.Entry.AuxAttr, writer)
		return
	}
	switch fv.Entry.Type {
	case flatrecord.MapStrVal:
		v := r.GetStr(fv.Entry.FlatIndex, fv.Entry.Source)
		writer.String(utils.TrimBoundingQuotes(v))
	case flatrecord.MapIntVal:
		writer.Int64(r.GetInt(fv.Entry.FlatIndex, fv.Entry.Source))
	case flatrecord.MapBoolVal:
		writer.Bool(r.GetInt(fv.Entry.FlatIndex, fv.Entry.Source) == 1)
	case flatrecord.MapSpecialStr:
		v := fv.Entry.Map(r).(string)
		writer.String(utils.TrimBoundingQuotes(v))
	case flatrecord.MapSpecialInt:
		writer.Int64(fv.Entry.Map(r).(int64))
	case flatrecord.MapSpecialBool:
		writer.Bool(fv.Entry.Map(r).(bool))
	case flatrecord.MapArrayStr, flatrecord.MapArrayInt:
		if fv.Entry.Source == sfgo.SYSFLOW_SRC {
			switch fv.Entry.FlatIndex {
			case sfgo.EV_PROC_OPFLAGS_INT:
				mapOpFlags(fv, writer, r)
				return
			case sfgo.FL_FILE_OPENFLAGS_INT:
				recType := r.GetInt(sfgo.SF_REC_TYPE, fv.Entry.Source)
				if recType == sfgo.NET_FLOW {
					mapIPs(fv, writer, r)
					return
				}
				mapOpenFlags(fv, writer, r)
				return
			case sfgo.FL_NETW_SPORT_INT:
				mapPorts(fv, writer, r)
				return
			case sfgo.POD_HOSTIP_ANY, sfgo.POD_INTERNALIP_ANY:
				ips := r.GetIntArray(fv.Entry.FlatIndex, fv.Entry.Source)
				mapIPArray(ips, writer)
				return

			}
		}
		v := fv.Entry.Map(r).(string)
		writer.RawByte(BEGIN_SQUARE)
		writer.String(v)
		writer.RawByte(END_SQUARE)
	case flatrecord.MapArraySvc:
		mapSvcArray(fv, writer, r)
	}
}

// setCachedValueToJSON sets the value of attr from cache for process ID to a JSON writer.
func setCachedValueToJSON(r *flatrecord.Record, ID sfgo.OID, attr flatrecord.RecAttribute, writer *jwriter.Writer) {
	if ptree := r.Fr.Ptree; ptree != nil {
		switch attr {
		case flatrecord.PProcName:
			if len(ptree) > 1 {
				writer.String(utils.TrimBoundingQuotes(filepath.Base(ptree[1].Exe)))
			} else {
				writer.String(EMPTY_STRING)
			}
		case flatrecord.PProcExe:
			if len(ptree) > 1 {
				writer.String(utils.TrimBoundingQuotes(ptree[1].Exe))
			} else {
				writer.String(EMPTY_STRING)
			}
		case flatrecord.PProcArgs:
			if len(ptree) > 1 {
				writer.String(utils.TrimBoundingQuotes(ptree[1].ExeArgs))
			} else {
				writer.String(EMPTY_STRING)
			}
		case flatrecord.PProcUID:
			if len(ptree) > 1 {
				writer.Int64(int64(ptree[1].Uid))
			} else {
				writer.Int64(sfgo.Zeros.Int64)
			}
		case flatrecord.PProcUser:
			if len(ptree) > 1 {
				writer.String(utils.TrimBoundingQuotes(ptree[1].UserName))
			} else {
				writer.String(EMPTY_STRING)
			}
		case flatrecord.PProcGID:
			if len(ptree) > 1 {
				writer.Int64(int64(ptree[1].Gid))
			} else {
				writer.Int64(sfgo.Zeros.Int64)
			}
		case flatrecord.PProcGroup:
			if len(ptree) > 1 {
				writer.String(utils.TrimBoundingQuotes(ptree[1].GroupName))
			} else {
				writer.String(EMPTY_STRING)
			}
		case flatrecord.PProcTTY:
			if len(ptree) > 1 {
				writer.Bool(ptree[1].Tty)
			} else {
				writer.Bool(false)
			}
		case flatrecord.PProcEntry:
			if len(ptree) > 1 {
				writer.Bool(ptree[1].Entry)
			} else {
				writer.Bool(false)
			}
		case flatrecord.PProcCmdLine:
			if len(ptree) > 1 {
				exe := utils.TrimBoundingQuotes(ptree[1].Exe)
				exeArgs := utils.TrimBoundingQuotes(ptree[1].ExeArgs)
				writer.RawByte(DOUBLE_QUOTE)
				stringNoQuotes(exe, writer)
				if len(exeArgs) > 0 {
					writer.RawByte(SPACE)
					stringNoQuotes(exeArgs, writer)
				}
				writer.RawByte(DOUBLE_QUOTE)
			} else {
				writer.String(EMPTY_STRING)
			}
		case flatrecord.ProcAName:
			l := len(ptree)
			writer.RawByte(BEGIN_SQUARE)
			for i, p := range ptree {
				writer.String(utils.TrimBoundingQuotes(filepath.Base(p.Exe)))
				if i < (l - 1) {
					writer.RawByte(COMMA)
				}
			}
			writer.RawByte(END_SQUARE)
		case flatrecord.ProcAExe:
			l := len(ptree)
			writer.RawByte(BEGIN_SQUARE)
			for i, p := range ptree {
				writer.String(utils.TrimBoundingQuotes(p.Exe))
				if i < (l - 1) {
					writer.RawByte(COMMA)
				}
			}
			writer.RawByte(END_SQUARE)
		case flatrecord.ProcACmdLine:
			l := len(ptree)
			writer.RawByte(BEGIN_SQUARE)
			for i, p := range ptree {
				exe := utils.TrimBoundingQuotes(p.Exe)
				exeArgs := utils.TrimBoundingQuotes(p.ExeArgs)
				writer.RawByte(DOUBLE_QUOTE)
				stringNoQuotes(exe, writer)
				if len(exeArgs) > 0 {
					writer.RawByte(SPACE)
					stringNoQuotes(exeArgs, writer)
				}
				writer.RawByte(DOUBLE_QUOTE)
				if i < (l - 1) {
					writer.RawByte(COMMA)
				}
			}
			writer.RawByte(END_SQUARE)
		case flatrecord.ProcAPID:
			l := len(ptree)
			writer.RawByte(BEGIN_SQUARE)
			for i, p := range ptree {
				writer.Int64(p.Oid.Hpid)
				if i < (l - 1) {
					writer.RawByte(COMMA)
				}
			}
			writer.RawByte(END_SQUARE)
		}
	}
}

// code taken from github.com/mailru/easyjson/jwriter to support string encoding.
// original version prepends quotes around strings, this doesn't.
func getTable(falseValues ...int) [128]bool {
	table := [128]bool{}
	for i := 0; i < 128; i++ {
		table[i] = true
	}
	for _, v := range falseValues {
		table[v] = false
	}
	return table
}

var (
	htmlEscapeTable   = getTable(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, '"', '&', '<', '>', '\\')
	htmlNoEscapeTable = getTable(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, '"', '\\')
)

// stringNoQuotes writes an escaped string with a JSON writer. Adapted from github.com/mailru/easyjson/jwriter.
func stringNoQuotes(s string, w *jwriter.Writer) {
	p := 0 // last non-escape symbol

	escapeTable := &htmlEscapeTable
	if w.NoEscapeHTML {
		escapeTable = &htmlNoEscapeTable
	}

	for i := 0; i < len(s); {
		c := s[i]

		if c < utf8.RuneSelf {
			if escapeTable[c] {
				// single-width character, no escaping is required
				i++
				continue
			}

			w.Buffer.AppendString(s[p:i])
			switch c {
			case '\t':
				w.Buffer.AppendString(`\t`)
			case '\r':
				w.Buffer.AppendString(`\r`)
			case '\n':
				w.Buffer.AppendString(`\n`)
			case '\\':
				w.Buffer.AppendString(`\\`)
			case '"':
				w.Buffer.AppendString(`\"`)
			default:
				w.Buffer.AppendString(`\u00`)
				w.Buffer.AppendByte(chars[c>>4])
				w.Buffer.AppendByte(chars[c&0xf])
			}

			i++
			p = i
			continue
		}

		// broken utf
		runeValue, runeWidth := utf8.DecodeRuneInString(s[i:])
		if runeValue == utf8.RuneError && runeWidth == 1 {
			w.Buffer.AppendString(s[p:i])
			w.Buffer.AppendString(`\ufffd`)
			i++
			p = i
			continue
		}

		// jsonp stuff - tab separator and line separator
		if runeValue == '\u2028' || runeValue == '\u2029' {
			w.Buffer.AppendString(s[p:i])
			w.Buffer.AppendString(`\u202`)
			w.Buffer.AppendByte(chars[runeValue&0xf])
			i += runeWidth
			p = i
			continue
		}
		i += runeWidth
	}
	w.Buffer.AppendString(s[p:])
}

// Cleanup cleans up resources.
func (t *JSONEncoder) Cleanup() {}
