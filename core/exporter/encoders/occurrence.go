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
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"os"
	"path"
	"strings"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/logger"

	"github.com/cespare/xxhash"
	"github.com/linkedin/goavro"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/steakknife/bloomfilter"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/encoders/avro/occurrence/event"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/utils"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
)

// EventPool contains an event slice with metadata annotations.
type EventPool struct {
	CID           string
	Events        []*Event
	Filter        *bloomfilter.Filter
	RuleTypes     *utils.Set
	TopSeverity   Severity
	LastFlushTime time.Time
	epw           *EventPoolWriter
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

// Aged checks if event pool has aged.
func (ep *EventPool) Aged(maxAge int) bool {
	return time.Since(ep.LastFlushTime).Minutes() > float64(maxAge)
}

// ReachedCapacity indicates whether the pool has reached its configured event capacity.
func (ep *EventPool) ReachedCapacity(capacity int) bool {
	return len(ep.Events) >= capacity
}

// Flush writes off event slice.
func (ep *EventPool) Flush(pathPrefix string, clusterID string) (err error) {
	var events []interface{}
	for _, v := range ep.Events {
		exportPath := fmt.Sprintf("%s/%s", pathPrefix, v.getExportFilePath(clusterID))
		if err = ep.UpdateEventPoolWriter(exportPath, v.Schema()); err != nil {
			return
		}
		var m map[string]interface{}
		logger.Warn.Println(v)
		s, _ := json.Marshal(v.Event)
		json.Unmarshal(s, &m)
		logger.Warn.Println(m)
		events = append(events, m)
	}
	if len(events) > 0 && ep.epw != nil {
		if err = ep.epw.Append(events); err != nil {
			return
		}
	}
	ep.Events = nil
	ep.LastFlushTime = time.Now()
	return
}

// Reset clears event slice and resets sketch counters and filter.
func (ep *EventPool) Reset() (err error) {
	bf, err := bloomfilter.NewOptimal(maxElements, probCollide)
	if err != nil {
		return
	}
	ep.Events = nil
	ep.Filter = bf
	ep.RuleTypes = utils.NewSet()
	ep.TopSeverity = SeverityLow
	ep.LastFlushTime = time.Now()
	return
}

// UpdateEventPoolWriter updates the EventPoolWriter for exportPath.
// It reuses the current EventPoolWriter if already point to the given exportPath.
// Otherwise, it creates a new OCF writer and the export directory structure if not present.
func (ep *EventPool) UpdateEventPoolWriter(exportPath string, schema string) (err error) {
	if ep.epw == nil {
		ep.epw = new(EventPoolWriter)
	}
	if exportPath != ep.epw.currentExportPath {
		dir := path.Dir(exportPath)
		if _, err = os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return
			}
		}
		if err = ep.epw.UpdateOCFWriter(exportPath, schema); err != nil {
			return
		}
		fmt.Println(ep.epw.ocfw)
	}
	// sanity check for cached OCF writer
	if ep.epw.ocfw == nil {
		return errors.New("EventPoolWriter's OCF file writer should not be null")
	}
	return
}

// EventPoolWriter is an EventPool writer.
type EventPoolWriter struct {
	currentExportPath string
	fw                *os.File
	codec             *goavro.Codec
	ocfw              *goavro.OCFWriter
}

// UpdateOCFWriter creates a new OCF writer.
func (epw *EventPoolWriter) UpdateOCFWriter(exportPath string, schema string) (err error) {
	// close the current file writer before creating a new one
	if epw.fw != nil {
		epw.fw.Close()
	}
	epw.currentExportPath = exportPath
	epw.fw, err = os.OpenFile(epw.currentExportPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return
	}
	if epw.codec == nil {
		epw.codec, err = goavro.NewCodec(schema)
		if err != nil {
			logger.Warn.Println(err)
			return
		}
	}
	epw.ocfw, err = goavro.NewOCFWriter(goavro.OCFConfig{
		W:               epw.fw,
		Codec:           epw.codec,
		CompressionName: "snappy",
	})
	logger.Warn.Println(epw.ocfw)
	logger.Warn.Println(err)
	return
}

func (epw *EventPoolWriter) Append(events []interface{}) error {
	if epw.ocfw != nil {
		return epw.ocfw.Append(events)
	}
	return errors.New("Trying to append events using a null OCF file writer reference")
}

func (epw *EventPoolWriter) Cleanup() error {
	return epw.fw.Close()
}

// Event is an event associated with an occurrence, used as context for the occurrence.
type Event struct {
	*event.Event
	Record *engine.Record
}

// getExportFileName returns the name of the file where the event should be exported.
func (e *Event) getExportFileName() string {
	if e.ContainerID == sfgo.Zeros.String {
		return hostFileName
	}
	return e.ContainerID
}

func (e *Event) getExportFilePath(clusterID string) string {
	y, m, d := e.getTimePartitions(e.Ts)
	return fmt.Sprintf("%s/%s/%d/%d/%d/%s.avro", clusterID, e.NodeID, y, m, d, e.getExportFileName())
}

// getTimePartitions obtains time partitions from timestamp.
func (e *Event) getTimePartitions(ts int64) (year int, month int, day int) {
	timeStamp := time.Unix(0, ts)
	return timeStamp.Year(), int(timeStamp.Month()), timeStamp.Day()
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
}

func (occ *Occurrence) NoteID() string {
	if occ.Severity < SeverityHigh {
		return "notification"
	}
	return "actionable-offense"
}

// OccurrenceEncoder is an encoder for IBM Findings' occurrences.
type OccurrenceEncoder struct {
	config      commons.Config
	exportCache cmap.ConcurrentMap
	batch       []commons.EncodedData
}

func NewOccurrenceEncoder(config commons.Config) Encoder {
	return &OccurrenceEncoder{
		config:      config,
		exportCache: cmap.New(),
		batch:       make([]commons.EncodedData, 0, config.EventBuffer)}
}

// Register registers the encoder to the codecs cache.
func (oe *OccurrenceEncoder) Register(codecs map[commons.Format]EncoderFactory) {
	codecs[commons.OccurrenceFormat] = NewOccurrenceEncoder
}

// Encodes a telemetry record into an occurrence representation.
func (oe *OccurrenceEncoder) encode(rec *engine.Record) (data commons.EncodedData, err error) {
	if e, ep, alert := oe.addEvent(rec); alert {
		data = oe.createOccurrence(e, ep)
	}
	return
}

// Encodes telemetry records into an occurrence representation.
func (oe *OccurrenceEncoder) Encode(recs []*engine.Record) ([]commons.EncodedData, error) {
	oe.batch = oe.batch[:0]
	for _, r := range recs {
		if data, _ := oe.encode(r); data != nil {
			oe.batch = append(oe.batch, data)
		}
	}
	return oe.batch, nil
}

// addEvent adds a record to export queue.
func (oe *OccurrenceEncoder) addEvent(r *engine.Record) (e *Event, ep *EventPool, alert bool) {
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
	h := oe.semanticHash(r)
	if !ep.Filter.Contains(h) {
		ep.Filter.Add(h)
		alert = true
	}

	// check for state changes in the pool after adding the event
	rc, s := ep.State()
	if rco != rc || so != s {
		alert = true
	}

	// write events out if
	// (1) an occurrence is generated for the current event, or
	// (2) the event pool has reached its configured capacity, or
	// (3) the event pool has aged.
	full := ep.ReachedCapacity(oe.config.FindingsPoolCapacity)
	aged := ep.Aged(oe.config.FindingsPoolMaxAge)
	if alert || full || aged {
		if err := ep.Flush(oe.config.FindingsPath, oe.config.ClusterID); err != nil {
			logger.Error.Println(err)
		}
		if aged {
			ep.Reset()
		}
	}

	return
}

// getEventPool retrieves container event pool from cache, or create one if absent.
func (oe *OccurrenceEncoder) getEventPool(cid string) *EventPool {
	m := oe.exportCache
	var ep *EventPool
	if v, ok := m.Get(cid); ok {
		ep = v.(*EventPool)
	} else {
		ep, _ = NewEventPool(cid)
		m.Set(cid, ep)
	}
	return ep
}

// createOccurrence creates a new Occurence object.
func (oe *OccurrenceEncoder) createOccurrence(e *Event, ep *EventPool) *Occurrence {
	oc := new(Occurrence)
	oc.Certainty = CertaintyMedium
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
	var detStr string
	switch e.Record.GetInt(sfgo.SF_REC_TYPE, sfgo.SYSFLOW_SRC) {
	case sfgo.PROC_EVT:
		proc := engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(e.Record)
		pproc := engine.Mapper.MapStr(engine.SF_PPROC_CMDLINE)(e.Record)
		detStr = fmt.Sprintf(peStrFmt, pproc, proc)
	case sfgo.FILE_EVT:
		proc := engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(e.Record)
		path := oe.formatResource(e.Record)
		detStr = fmt.Sprintf(feStrFmt, proc, path)
	case sfgo.FILE_FLOW:
		proc := engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(e.Record)
		path := oe.formatResource(e.Record)
		detStr = fmt.Sprintf(ffStrFmt, proc, path)
	case sfgo.NET_FLOW:
		proc := engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(e.Record)
		conn := oe.formatResource(e.Record)
		detStr = fmt.Sprintf(nfStrFmt, proc, conn)
	}
	// sanitizes details string to avoid being flagged by tools like CloudFlare
	encDetStr := strings.ReplaceAll(detStr, "/", fwdSlash)
	oc.ShortDescr = encDetStr
	oc.LongDescr = fmt.Sprintf(detailsStrFmt, encDetStr, polStr, tagsStr)
	oc.AlertQuery = fmt.Sprintf(sqlQueryStrFmt, oe.config.FindingsS3Region, oe.config.FindingsS3Bucket,
		e.getExportFilePath(oe.config.ClusterID), oe.config.FindingsS3Region, oe.config.FindingsS3Bucket)
	return oc
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
func (oe *OccurrenceEncoder) encodeEvent(r *engine.Record) *Event {
	rnames, tags, severity := oe.summarizePolicy(r)
	e := &Event{Record: r, Event: event.NewEvent()}
	e.Ts = engine.Mapper.MapInt(engine.SF_TS)(r)
	e.Description = strings.Join(rnames, listSep)
	e.Severity = severity.String()
	e.NodeID = engine.Mapper.MapStr(engine.SF_NODE_ID)(r)
	e.ContainerID = engine.Mapper.MapStr(engine.SF_CONTAINER_ID)(r)
	e.RecordType = engine.Mapper.MapStr(engine.SF_TYPE)(r)
	e.OpFlags = engine.Mapper.MapStr(engine.SF_OPFLAGS)(r)
	e.PProcCmd = engine.Mapper.MapStr(engine.SF_PPROC_CMDLINE)(r)
	e.PProcPID = engine.Mapper.MapInt(engine.SF_PPROC_PID)(r)
	e.ProcCmd = engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(r)
	e.ProcPID = engine.Mapper.MapInt(engine.SF_PROC_PID)(r)
	e.Resource = oe.formatResource(r)
	e.Tags = strings.Join(tags, listSep)
	e.Trace = engine.Mapper.MapStr(engine.SF_TRACENAME)(r)
	return e
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

// semanticHash computes a hash value over record attributes denoting the semantics of the record (used in the bloom filter).
func (oe *OccurrenceEncoder) semanticHash(r *engine.Record) hash.Hash64 {
	h := xxhash.New()
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(r)))
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_PROC_UID)(r)))
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_FILE_OID)(r)))
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_OPFLAGS)(r)))
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_PROC_TTY)(r)))
	return h
}

// Cleanup cleans up resources.
func (oe *OccurrenceEncoder) Cleanup() {
	for _, v := range oe.exportCache.Items() {
		ep := v.(*EventPool)
		ep.epw.Cleanup()
	}
}
