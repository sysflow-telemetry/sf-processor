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
	"unicode/utf8"

	"github.com/mailru/easyjson/jwriter"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/utils"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
)

// JSONEncoder is a JSON encoder.
type JSONEncoder struct {
	config     commons.Config
	fieldCache []*engine.FieldValue
	writer     *jwriter.Writer
	buf        []byte
	batch      []commons.EncodedData
}

// NewJSONEncoder instantiates a JSON encoder.
func NewJSONEncoder(config commons.Config) Encoder {
	return &JSONEncoder{
		fieldCache: engine.FieldValues,
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
func (t *JSONEncoder) Encode(recs []*engine.Record) (data []commons.EncodedData, err error) {
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
func (t *JSONEncoder) encode(rec *engine.Record) (commons.EncodedData, error) {
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
			case engine.SectPProc:
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
			case engine.SectNet:
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
			case engine.SectFile:
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
			case engine.SectFlow:
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
			case engine.SectCont:
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
			case engine.SectNode:
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
			case engine.SectMeta:
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

	// Encode hash values
	hp := rec.Ctx.GetHash(engine.HASH_PROC)
	hf := rec.Ctx.GetHash(engine.HASH_FILE)
	if hp != nil || hf != nil {
		t.writeSectionBegin(EXTENSIONS)
		if hp != nil {
			t.writeSectionBegin(PROC)
			t.writeHash(hp)
			if hf != nil {
				t.writer.RawByte(COMMA)
			}
		}
		if hf != nil {
			t.writeSectionBegin(FILEF)
			t.writeHash(hf)
		}
		t.writer.RawByte(END_CURLY)
	}

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

func (t *JSONEncoder) writeAttribute(fv *engine.FieldValue, fieldID int, rec *engine.Record) {
	t.writer.RawByte(DOUBLE_QUOTE)
	t.writer.RawString(fv.FieldSects[fieldID])
	t.writer.RawString(QUOTE_COLON)
	MapJSON(fv, t.writer, rec)
}

func (t *JSONEncoder) writeSectionBegin(section string) {
	t.writer.RawByte(DOUBLE_QUOTE)
	t.writer.RawString(section)
	t.writer.RawString(QUOTE_COLON_CURLY)
}

func (t *JSONEncoder) writeHash(h *engine.HashSet) {
	existed := false
	if h.Md5 != sfgo.Zeros.String {
		t.writer.RawString(MD5)
		t.writer.RawString(h.Md5)
		t.writer.RawByte(DOUBLE_QUOTE)
		existed = true
	}
	if h.Sha1 != sfgo.Zeros.String {
		if existed {
			t.writer.RawByte(COMMA)
		}
		t.writer.RawString(SHA1)
		t.writer.RawString(h.Sha1)
		t.writer.RawByte(DOUBLE_QUOTE)
		existed = true
	}
	if h.Sha256 != sfgo.Zeros.String {
		if existed {
			t.writer.RawByte(COMMA)
		}
		t.writer.RawString(SHA256)
		t.writer.RawString(h.Sha256)
		t.writer.RawByte(DOUBLE_QUOTE)
	}
}


func mapOpFlags(fv *engine.FieldValue, writer *jwriter.Writer, r *engine.Record) {
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
func mapIPs(fv *engine.FieldValue, writer *jwriter.Writer, r *engine.Record) {
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

func mapOpenFlags(fv *engine.FieldValue, writer *jwriter.Writer, r *engine.Record) {
	flags := sfgo.GetOpenFlags(r.GetInt(fv.Entry.FlatIndex, fv.Entry.Source))
	mapStrArray(writer, flags)
}

func mapPorts(fv *engine.FieldValue, writer *jwriter.Writer, r *engine.Record) {
	srcPort := r.GetInt(sfgo.FL_NETW_SPORT_INT, fv.Entry.Source)
	dstPort := r.GetInt(sfgo.FL_NETW_DPORT_INT, fv.Entry.Source)
	writer.RawByte(BEGIN_SQUARE)
	writer.Int64(srcPort)
	writer.RawByte(COMMA)
	writer.Int64(dstPort)
	writer.RawByte(END_SQUARE)
}

// MapJSON writes a SysFlow attribute to a JSON stream.
func MapJSON(fv *engine.FieldValue, writer *jwriter.Writer, r *engine.Record) {
	switch fv.Entry.FlatIndex {
	case engine.A_IDS, engine.PARENT_IDS:
		oid := sfgo.OID{CreateTS: r.GetInt(sfgo.PROC_OID_CREATETS_INT, fv.Entry.Source), Hpid: r.GetInt(sfgo.PROC_OID_HPID_INT, fv.Entry.Source)}
		setCachedValueToJSON(r, oid, fv.Entry.AuxAttr, writer)
		return
	}
	switch fv.Entry.Type {
	case engine.MapStrVal:
		v := r.GetStr(fv.Entry.FlatIndex, fv.Entry.Source)
		writer.String(utils.TrimBoundingQuotes(v))
	case engine.MapIntVal:
		writer.Int64(r.GetInt(fv.Entry.FlatIndex, fv.Entry.Source))
	case engine.MapBoolVal:
		writer.Bool(r.GetInt(fv.Entry.FlatIndex, fv.Entry.Source) == 1)
	case engine.MapSpecialStr:
		v := fv.Entry.Map(r).(string)
		writer.String(utils.TrimBoundingQuotes(v))
	case engine.MapSpecialInt:
		writer.Int64(fv.Entry.Map(r).(int64))
	case engine.MapSpecialBool:
		writer.Bool(fv.Entry.Map(r).(bool))
	case engine.MapArrayStr, engine.MapArrayInt:
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
			}
		}
		v := fv.Entry.Map(r).(string)
		writer.RawByte(BEGIN_SQUARE)
		writer.String(v)
		writer.RawByte(END_SQUARE)
	}
}

// setCachedValueToJSON sets the value of attr from cache for process ID to a JSON writer.
func setCachedValueToJSON(r *engine.Record, ID sfgo.OID, attr engine.RecAttribute, writer *jwriter.Writer) {
	if ptree := r.Fr.Ptree; ptree != nil {
		switch attr {
		case engine.PProcName:
			if len(ptree) > 1 {
				writer.String(utils.TrimBoundingQuotes(filepath.Base(ptree[1].Exe)))
			} else {
				writer.String(EMPTY_STRING)
			}
		case engine.PProcExe:
			if len(ptree) > 1 {
				writer.String(utils.TrimBoundingQuotes(ptree[1].Exe))
			} else {
				writer.String(EMPTY_STRING)
			}
		case engine.PProcArgs:
			if len(ptree) > 1 {
				writer.String(utils.TrimBoundingQuotes(ptree[1].ExeArgs))
			} else {
				writer.String(EMPTY_STRING)
			}
		case engine.PProcUID:
			if len(ptree) > 1 {
				writer.Int64(int64(ptree[1].Uid))
			} else {
				writer.Int64(sfgo.Zeros.Int64)
			}
		case engine.PProcUser:
			if len(ptree) > 1 {
				writer.String(utils.TrimBoundingQuotes(ptree[1].UserName))
			} else {
				writer.String(EMPTY_STRING)
			}
		case engine.PProcGID:
			if len(ptree) > 1 {
				writer.Int64(int64(ptree[1].Gid))
			} else {
				writer.Int64(sfgo.Zeros.Int64)
			}
		case engine.PProcGroup:
			if len(ptree) > 1 {
				writer.String(utils.TrimBoundingQuotes(ptree[1].GroupName))
			} else {
				writer.String(EMPTY_STRING)
			}
		case engine.PProcTTY:
			if len(ptree) > 1 {
				writer.Bool(ptree[1].Tty)
			} else {
				writer.Bool(false)
			}
		case engine.PProcEntry:
			if len(ptree) > 1 {
				writer.Bool(ptree[1].Entry)
			} else {
				writer.Bool(false)
			}
		case engine.PProcCmdLine:
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
		case engine.ProcAName:
			l := len(ptree)
			writer.RawByte(BEGIN_SQUARE)
			for i, p := range ptree {
				writer.String(utils.TrimBoundingQuotes(filepath.Base(p.Exe)))
				if i < (l - 1) {
					writer.RawByte(COMMA)
				}
			}
			writer.RawByte(END_SQUARE)
		case engine.ProcAExe:
			l := len(ptree)
			writer.RawByte(BEGIN_SQUARE)
			for i, p := range ptree {
				writer.String(utils.TrimBoundingQuotes(p.Exe))
				if i < (l - 1) {
					writer.RawByte(COMMA)
				}
			}
			writer.RawByte(END_SQUARE)
		case engine.ProcACmdLine:
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
		case engine.ProcAPID:
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
