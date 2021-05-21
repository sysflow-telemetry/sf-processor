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
package encoders

import (
	"fmt"
	"hash"
	"strings"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"

	"github.com/cespare/xxhash"
	cqueue "github.com/enriquebris/goconcurrentqueue"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/steakknife/bloomfilter"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/utils"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
)

// Bloom filter settings.
const (
	maxElements = 100000
	probCollide = 0.0000001
)

// EventPool contains an event slice with metadata annotations.
type EventPool struct {
	CID                    string
	Events                 []Event
	Filter                 *bloomfilter.Filter
	RuleTypes              *utils.Set
	TopSeverity            Severity
	LastExportedEventIndex int
}

// NewEventPool creates a new EventPool instace.
func NewEventPool(cid string) (ep *EventPool, err error) {
	bf, err := bloomfilter.NewOptimal(maxElements, probCollide)
	if err != nil {
		return
	}
	return &EventPool{CID: cid, Filter: bf, RuleTypes: utils.NewSet(), TopSeverity: SeverityLow}, nil
}

// State returns a tuple summarizing the state of the event pool.
func (ep *EventPool) State() (int, Severity) {
	return ep.RuleTypes.Len(), ep.TopSeverity
}

// Flush writes off event slice.
func (ep *EventPool) Flush() error {
	//ep = EventPool{Events: make([]Event, 0)}
	return nil
}

// Event is an event associated with an occurrence, used as context for the occurrence.
type Event struct {
	Record      *engine.Record
	Ts          int64  `parquet:"name=ts, type=INT64"`
	Description string `parquet:"name=description, type=BYTE_ARRAY, convertedtype=UTF8"`
	Severity    string `parquet:"name=severity, type=BYTE_ARRAY, convertedtype=UTF8"`
	NodeID      string `parquet:"name=node_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	ContainerID string `parquet:"name=container_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	RecType     string `parquet:"name=record_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	OpFlags     string `parquet:"name=op_flags, type=BYTE_ARRAY, convertedtype=UTF8"`
	PProcCmd    string `parquet:"name=pproc_cmd, type=BYTE_ARRAY, convertedtype=UTF8"`
	ProcCmd     string `parquet:"name=proc_cmd, type=BYTE_ARRAY, convertedtype=UTF8"`
	PProcPID    string `parquet:"name=pproc_pid, type=BYTE_ARRAY, convertedtype=UTF8"`
	ProcPID     string `parquet:"name=proc_pid, type=BYTE_ARRAY, convertedtype=UTF8"`
	Resource    string `parquet:"name=resource, type=BYTE_ARRAY, convertedtype=UTF8"`
	Tags        string `parquet:"name=tags, type=BYTE_ARRAY, convertedtype=UTF8"`
	Trace       string `parquet:"name=trace, type=BYTE_ARRAY, convertedtype=UTF8"`
}

// getExportFileName returns the name of the file where the event should be exported.
func (e Event) getExportFileName() string {
	if e.ContainerID == sfgo.Zeros.String {
		return hostFileName
	}
	return e.ContainerID
}

// Occurrence object for IBM Findings API.
type Occurrence struct {
	ID         string
	ShortDescr string
	LongDescr  string
	Severity   Severity
	Certainty  Certainty
	ResType    string
	ResName    string
	AlertQuery string
	NoteID     string
}

// OccurrenceEncoder is an encoder for IBM Findings' occurrences.
type OccurrenceEncoder struct {
	config      commons.Config
	exportQueue *cqueue.FIFO
}

func NewOccurrenceEncoder(conf commons.Config) Encoder {
	queue := cqueue.NewFIFO()
	queue.Enqueue(cmap.New())
	return &OccurrenceEncoder{config: conf, exportQueue: queue}
}

// Register registers the encoder to the codecs cache.
func (oe *OccurrenceEncoder) Register(codecs map[commons.Format]EncoderFactory) {
	codecs[commons.OccurrenceFormat] = NewOccurrenceEncoder
}

// Encodes a telemetry record into an occurrence representation.
func (oe *OccurrenceEncoder) Encode(r *engine.Record) (data commons.EncodedData, err error) {
	if e, ep, alert := oe.addEvent(r); alert {
		oe.createOccurrence(e, ep)
	}
	return
}

// addEvent adds a record to export queue.
func (oe *OccurrenceEncoder) addEvent(r *engine.Record) (e Event, ep *EventPool, alert bool) {
	// write records to disk
	cid := engine.Mapper.MapStr(engine.SF_CONTAINER_ID)(r)
	ep = oe.getEventPool(cid)

	// record the event pool state prior to adding a new event
	rco, so := ep.State()

	// encode and add event to event pool
	e = oe.encodeEvent(r)
	ep.Events = append(ep.Events, e)
	for _, rule := range r.Ctx.GetRules() {
		ep.RuleTypes.Add(rule.Name)
		ep.TopSeverity = Severity(utils.Max(int(ep.TopSeverity), int(rule.Priority)))
	}

	// check if a semantically equivalent record has been seen before
	h := oe.hash(r)
	if !ep.Filter.Contains(h) {
		ep.Filter.Add(h)
		return e, ep, true
	}

	// check for state changes in the pool after adding the event
	rc, s := ep.State()
	if rco != rc || so != s {
		return e, ep, true
	}

	return e, ep, false
}

// getEventPool retrieves container event pool from cache, or create one if absent.
func (oe *OccurrenceEncoder) getEventPool(cid string) *EventPool {
	head, _ := oe.exportQueue.Get(0)
	m := head.(cmap.ConcurrentMap)
	var ep *EventPool
	if v, ok := m.Get(cid); ok {
		ep = v.(*EventPool)
	} else {
		ep, _ := NewEventPool(cid)
		m.Set(cid, ep)
	}
	return ep
}

// createOccurrence creates a new Occurence object.
func (oe *OccurrenceEncoder) createOccurrence(e Event, ep *EventPool) Occurrence {
	oc := Occurrence{Certainty: CertaintyMedium}
	oc.ID = fmt.Sprintf(noteIDStrFmt, ep.CID, time.Now().UTC().UnixNano()/1000)
	if ep.CID != sfgo.Zeros.String {
		oc.ResName = fmt.Sprintf("%s:%s [%s]", ep.CID, engine.Mapper.MapStr(engine.SF_CONTAINER_NAME)(e.Record), e.NodeID)
		oc.ResType = engine.Mapper.MapStr(engine.SF_CONTAINER_TYPE)(e.Record)
	} else {
		oc.ResName = e.NodeID
		oc.ResType = sfgo.Zeros.String
	}
	rnames, tags, severity := oe.summarizePolicy(e.Record)
	oc.Severity = severity
	polStr := fmt.Sprintf(policiesStrFmt, strings.Join(rnames, listSep))
	tagsStr := fmt.Sprintf(tagsStrFmt, strings.Join(tags, listSep))
	switch e.Record.GetInt(sfgo.SF_REC_TYPE, sfgo.SYSFLOW_SRC) {
	case sfgo.PROC_EVT:
		proc := engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(e.Record)
		pproc := engine.Mapper.MapStr(engine.SF_PPROC_CMDLINE)(e.Record)
		detStr := fmt.Sprintf(peStrFmt, pproc, proc)
		oc.ShortDescr = detStr
		oc.LongDescr = fmt.Sprintf(detailsStrFmt, detStr, polStr, tagsStr)
	case sfgo.FILE_EVT:
		proc := engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(e.Record)
		path := oe.formatResource(e.Record)
		detStr := fmt.Sprintf(feStrFmt, proc, path)
		oc.ShortDescr = detStr
		oc.LongDescr = fmt.Sprintf(detailsStrFmt, detStr, polStr, tagsStr)
	case sfgo.FILE_FLOW:
		proc := engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(e.Record)
		path := oe.formatResource(e.Record)
		detStr := fmt.Sprintf(ffStrFmt, proc, path)
		oc.ShortDescr = detStr
		oc.LongDescr = fmt.Sprintf(detailsStrFmt, detStr, polStr, tagsStr)
	case sfgo.NET_FLOW:
		proc := engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(e.Record)
		conn := oe.formatResource(e.Record)
		detStr := fmt.Sprintf(nfStrFmt, proc, conn)
		oc.ShortDescr = detStr
		oc.LongDescr = fmt.Sprintf(detailsStrFmt, detStr, polStr, tagsStr)
	}
	y, m, d := oe.getTimePartitions(e.Ts)
	oc.AlertQuery = fmt.Sprintf(sqlQueryStrFmt, fmt.Sprintf("%s/%d/%d/%d/%s", e.NodeID, y, m, d, e.getExportFileName()))
	return oc
}

// createDigest creates a new Occurrence object summarizing all findings in the event pool.
func (oe *OccurrenceEncoder) createDigest(ep *EventPool) Occurrence {
	return Occurrence{}
}

// summarizePolicy extracts a summary of rules applied to a record.
func (oe *OccurrenceEncoder) summarizePolicy(r *engine.Record) (rnames []string, tags []string, severity Severity) {
	for _, rule := range r.Ctx.GetRules() {
		rnames = append(rnames, rule.Name)
		severity = Severity(utils.Max(int(severity), int(rule.Priority)))
		for _, tag := range rule.Tags {
			switch tag.(type) {
			case []string:
				tags = append(tags, tag.([]string)...)
			default:
				tags = append(tags, tag.(string))
			}
		}
	}
	return
}

// encodeEvent maps a record into an event that can be associated with an occurrence.
func (oe *OccurrenceEncoder) encodeEvent(r *engine.Record) Event {
	rnames, tags, severity := oe.summarizePolicy(r)
	return Event{Record: r,
		Ts:          engine.Mapper.MapInt(engine.SF_TS)(r),
		Description: strings.Join(rnames, listSep),
		Severity:    severity.String(),
		NodeID:      engine.Mapper.MapStr(engine.SF_NODE_ID)(r),
		ContainerID: engine.Mapper.MapStr(engine.SF_CONTAINER_ID)(r),
		RecType:     engine.Mapper.MapStr(engine.SF_TYPE)(r),
		OpFlags:     engine.Mapper.MapStr(engine.SF_OPFLAGS)(r),
		PProcCmd:    engine.Mapper.MapStr(engine.SF_PPROC_CMDLINE)(r),
		ProcCmd:     engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(r),
		Resource:    oe.formatResource(r),
		Tags:        strings.Join(tags, listSep),
		Trace:       ""}
}

// formatResource formats a file or network resource.
func (oe *OccurrenceEncoder) formatResource(r *engine.Record) (res string) {
	switch r.GetInt(sfgo.SF_REC_TYPE, sfgo.SYSFLOW_SRC) {
	case sfgo.FILE_EVT, sfgo.FILE_FLOW:
		return engine.Mapper.MapStr(engine.SF_FILE_PATH)(r)
	case sfgo.NET_FLOW:
		sip := engine.Mapper.MapStr(engine.SF_NET_SIP)(r)
		sport := engine.Mapper.MapInt(engine.SF_NET_SPORT)(r)
		dip := engine.Mapper.MapStr(engine.SF_NET_DIP)(r)
		dport := engine.Mapper.MapInt(engine.SF_NET_DPORT)(r)
		return fmt.Sprintf(connStrFmt, sip, sport, dip, dport)
	}
	return
}

// getTimePartitions obtains time partitions from timestamp.
func (oe *OccurrenceEncoder) getTimePartitions(ts int64) (year int, month int, day int) {
	timeStamp := time.Unix(ts, 0)
	return timeStamp.Year(), int(timeStamp.Month()), timeStamp.Day()
}

// hash computes a hash value over record attributes denoting the semantics of the record (used in the bloom filter).
func (oe *OccurrenceEncoder) hash(r *engine.Record) hash.Hash64 {
	h := xxhash.New()
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(r)))
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_PROC_UID)(r)))
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_FILE_OID)(r)))
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_OPFLAGS)(r)))
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_PROC_TTY)(r)))
	return h
}
