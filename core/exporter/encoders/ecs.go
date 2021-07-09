//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Andreas Schade <san@zurich.ibm.com>
// Frederico Araujo <frederico.araujo@ibm.com>
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
	"fmt"
	"net"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cespare/xxhash"
	"github.com/satta/gommunityid"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/utils"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
)

// JSONData is a map to serialize data to JSON.
type JSONData map[string]interface{}

// ECSRecord is a struct for serializing ECS records.
type ECSRecord struct {
	ID    string `json:"-"`
	Ts    string `json:"@timestamp"`
	Agent struct {
		Type    string `json:"type,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"agent,omitempty"`
	Ecs struct {
		Version string `json:"version,omitempty"`
	} `json:"ecs,omitempty"`
	Event       JSONData `json:"event"`
	Container   JSONData `json:"container"`
	File        JSONData `json:"file,omitempty"`
	FileAction  JSONData `json:"file_action,omitempty"`
	Network     JSONData `json:"network,omitempty"`
	Source      JSONData `json:"source,omitempty"`
	Destination JSONData `json:"destination,omitempty"`
	Process     JSONData `json:"process"`
	User        JSONData `json:"user"`
	Tags        []string `json:"tags,omitempty"`
}

// ECSEncoder implements an ECS encoder for telemetry records.
type ECSEncoder struct {
	config commons.Config
	//jsonencoder JSONEncoder
	batch []commons.EncodedData
}

// NewECSEncoder instantiates an ECS encoder.
func NewECSEncoder(config commons.Config) Encoder {
	return &ECSEncoder{
		config: config,
		batch:  make([]commons.EncodedData, 0, config.EventBuffer)}
}

// Register registers the encoder to the codecs cache.
func (t *ECSEncoder) Register(codecs map[commons.Format]EncoderFactory) {
	codecs[commons.ECSFormat] = NewECSEncoder
}

// Encode encodes telemetry records into an ECS representation.
func (t *ECSEncoder) Encode(recs []*engine.Record) ([]commons.EncodedData, error) {
	t.batch = t.batch[:0]
	for _, rec := range recs {
		ecs := t.encode(rec)
		t.batch = append(t.batch, ecs)
	}
	return t.batch, nil
}

// Encodes a telemetry record into an ECS representation.
func (t *ECSEncoder) encode(rec *engine.Record) *ECSRecord {
	ecs := &ECSRecord{
		ID:        encodeID(rec),
		Container: encodeContainer(rec),
		Process:   encodeProcess(rec),
		User:      encodeUser(rec),
	}
	ecs.Agent.Version = t.config.Version
	ecs.Agent.Type = ECS_AGENT_TYPE
	ecs.Ecs.Version = t.config.EcsVersion
	ecs.Ts = utils.ToIsoTimeStr(engine.Mapper.MapInt(engine.SF_TS)(rec))

	// encode specific record components
	sfType := engine.Mapper.MapStr(engine.SF_TYPE)(rec)
	switch sfType {
	case sfgo.TyNFStr:
		ecs.encodeNetworkFlow(rec)
	case sfgo.TyFFStr:
		ecs.encodeFileFlow(rec)
	case sfgo.TyFEStr:
		ecs.encodeFileEvent(rec)
	case sfgo.TyPEStr:
		ecs.encodeProcessEvent(rec)
	}

	// encode tags and policy information
	rules := rec.Ctx.GetRules()
	if len(rules) > 0 {
		reasons := make([]string, 0)
		tags := make([]string, 0)
		priority := int(engine.Low)
		for _, r := range rules {
			reasons = append(reasons, r.Name)
			tags = append(tags, extracTags(r.Tags)...)
			priority = utils.Max(priority, int(r.Priority))
		}
		ecs.Event[ECS_EVENT_REASON] = strings.Join(reasons, ", ")
		ecs.Event[ECS_EVENT_SEVERITY] = priority
		ecs.Tags = tags
	}
	return ecs
}

// encodeID returns the ECS document identifier.
func encodeID(rec *engine.Record) string {
	h := xxhash.New()
	t := engine.Mapper.MapStr(engine.SF_TYPE)(rec)
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_NODE_ID)(rec)))
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_CONTAINER_ID)(rec)))
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_TS)(rec)))
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_PROC_TID)(rec)))
	h.Write([]byte(engine.Mapper.MapStr(engine.SF_PROC_CREATETS)(rec)))
	h.Write([]byte(t))
	switch t {
	case sfgo.TyFFStr, sfgo.TyFEStr:
		h.Write([]byte(engine.Mapper.MapStr(engine.SF_FILE_OID)(rec)))
	case sfgo.TyNFStr:
		h.Write([]byte(engine.Mapper.MapStr(engine.SF_NET_SIP)(rec)))
		h.Write([]byte(engine.Mapper.MapStr(engine.SF_NET_SPORT)(rec)))
		h.Write([]byte(engine.Mapper.MapStr(engine.SF_NET_DIP)(rec)))
		h.Write([]byte(engine.Mapper.MapStr(engine.SF_NET_DPORT)(rec)))
		h.Write([]byte(engine.Mapper.MapStr(engine.SF_NET_PROTO)(rec)))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// encodeNetworkFlow populates the ECS representatiom of a NetworkFlow record.
func (ecs *ECSRecord) encodeNetworkFlow(rec *engine.Record) {
	rbytes := engine.Mapper.MapInt(engine.SF_FLOW_RBYTES)(rec)
	rops := engine.Mapper.MapInt(engine.SF_FLOW_ROPS)(rec)
	wbytes := engine.Mapper.MapInt(engine.SF_FLOW_WBYTES)(rec)
	wops := engine.Mapper.MapInt(engine.SF_FLOW_WOPS)(rec)
	sip := engine.Mapper.MapStr(engine.SF_NET_SIP)(rec)
	dip := engine.Mapper.MapStr(engine.SF_NET_DIP)(rec)
	sport := engine.Mapper.MapInt(engine.SF_NET_SPORT)(rec)
	dport := engine.Mapper.MapInt(engine.SF_NET_DPORT)(rec)
	proto := engine.Mapper.MapInt(engine.SF_NET_PROTO)(rec)

	cid, _ := gommunityid.GetCommunityIDByVersion(1, 0)
	ft := gommunityid.MakeFlowTuple(net.ParseIP(sip), net.ParseIP(dip), uint16(sport), uint16(dport), uint8(proto))

	// Calculate Base64-encoded value
	ecs.Network = JSONData{
		ECS_NET_BYTES: rbytes + wbytes,
		ECS_NET_CID:   cid.CalcBase64(ft),
		ECS_NET_IANA:  strconv.FormatInt(proto, 10),
		ECS_NET_PROTO: sfgo.GetProto(proto),
	}
	ecs.Source = JSONData{
		ECS_ENDPOINT_IP:      sip,
		ECS_ENDPOINT_PORT:    sport,
		ECS_ENDPOINT_ADDR:    sip,
		ECS_ENDPOINT_BYTES:   wbytes,
		ECS_ENDPOINT_PACKETS: wops,
	}
	ecs.Destination = JSONData{
		ECS_ENDPOINT_IP:      dip,
		ECS_ENDPOINT_PORT:    dport,
		ECS_ENDPOINT_ADDR:    dip,
		ECS_ENDPOINT_BYTES:   rbytes,
		ECS_ENDPOINT_PACKETS: rops,
	}
	ecs.Event = encodeEvent(rec, ECS_CAT_NETWORK, ECS_TYPE_CONNECTION, ECS_CAT_NETWORK+"-"+ECS_ACTION_TRAFFIC)
}

// encodeFileFlow populates the ECS representatiom of a FF record
func (ecs *ECSRecord) encodeFileFlow(rec *engine.Record) {
	opFlags := rec.GetInt(sfgo.EV_PROC_OPFLAGS_INT, sfgo.SYSFLOW_SRC)
	rbytes := engine.Mapper.MapInt(engine.SF_FLOW_RBYTES)(rec)
	rops := engine.Mapper.MapInt(engine.SF_FLOW_ROPS)(rec)
	wbytes := engine.Mapper.MapInt(engine.SF_FLOW_WBYTES)(rec)
	wops := engine.Mapper.MapInt(engine.SF_FLOW_WOPS)(rec)
	category := ECS_CAT_FILE
	eventType := ECS_TYPE_ACCESS
	action := category + "-" + eventType
	if opFlags&sfgo.OP_READ_RECV == sfgo.OP_READ_RECV && rbytes > 0 {
		action = action + "-" + ECS_ACTION_READ
	}
	if opFlags&sfgo.OP_WRITE_SEND == sfgo.OP_WRITE_SEND && wbytes > 0 {
		eventType = ECS_TYPE_CHANGE
		action = action + "-" + ECS_ACTION_WRITE
	}
	ecs.Event = encodeEvent(rec, category, eventType, action)
	ecs.File = encodeFile(rec)
	ecs.FileAction = JSONData{
		ECS_SF_FA_RBYTES: rbytes,
		ECS_SF_FA_ROPS:   rops,
		ECS_SF_FA_WBYTES: wbytes,
		ECS_SF_FA_WOPS:   wops,
	}
}

// encodeFileEvent populates the ECS representatiom of a FE record
func (ecs *ECSRecord) encodeFileEvent(rec *engine.Record) {
	opFlags := rec.GetInt(sfgo.EV_PROC_OPFLAGS_INT, sfgo.SYSFLOW_SRC)
	targetPath := engine.Mapper.MapStr(engine.SF_FILE_NEWPATH)(rec)
	ecs.File = encodeFile(rec)
	category := ECS_CAT_FILE
	eventType := ECS_TYPE_CHANGE
	action := category + "-" + eventType
	if opFlags&sfgo.OP_MKDIR == sfgo.OP_MKDIR {
		category = ECS_CAT_DIR
		eventType = ECS_TYPE_CREATE
		action = category + "-" + ECS_ACTION_CREATE
	} else if opFlags&sfgo.OP_RMDIR == sfgo.OP_RMDIR {
		category = ECS_CAT_DIR
		eventType = ECS_TYPE_DELETE
		action = category + "-" + ECS_ACTION_DELETE
	} else if opFlags&sfgo.OP_UNLINK == sfgo.OP_UNLINK {
		eventType = ECS_TYPE_DELETE
		action = category + "-" + ECS_ACTION_DELETE
	} else if opFlags&sfgo.OP_SYMLINK == sfgo.OP_SYMLINK || opFlags&sfgo.OP_LINK == sfgo.OP_LINK {
		action = category + "-" + ECS_ACTION_LINK
		ecs.File[ECS_FILE_TARGET] = targetPath
	} else if opFlags&sfgo.OP_RENAME == sfgo.OP_RENAME {
		action = category + "-" + ECS_ACTION_RENAME
		ecs.File[ECS_FILE_TARGET] = targetPath
	}
	ecs.Event = encodeEvent(rec, category, eventType, action)
}

// encodeProcessEvent populates the ECS representatiom of a PE record
func (ecs *ECSRecord) encodeProcessEvent(rec *engine.Record) {
	opFlags := rec.GetInt(sfgo.EV_PROC_OPFLAGS_INT, sfgo.SYSFLOW_SRC)
	pid := engine.Mapper.MapInt(engine.SF_PROC_PID)(rec)
	tid := engine.Mapper.MapInt(engine.SF_PROC_TID)(rec)
	category := ECS_CAT_PROCESS
	eventType := ECS_TYPE_START

	if opFlags&sfgo.OP_EXIT == sfgo.OP_EXIT {
		if pid != tid {
			eventType = ECS_TYPE_TEXIT
		} else {
			eventType = ECS_TYPE_EXIT
		}
	} else if opFlags&sfgo.OP_CLONE == sfgo.OP_CLONE || opFlags&sfgo.OP_EXEC == sfgo.OP_EXEC {
		if pid != tid {
			eventType = ECS_TYPE_TSTART
		}
	} else if opFlags&sfgo.OP_SETUID == sfgo.OP_SETUID {
		eventType = ECS_TYPE_CHANGE
	}

	action := category + "-" + eventType
	ecs.Event = encodeEvent(rec, category, eventType, action)
}

// encodeContainer creates an ECS container field.
func encodeContainer(rec *engine.Record) JSONData {
	var container JSONData
	cid := engine.Mapper.MapStr(engine.SF_CONTAINER_ID)(rec)
	if cid != sfgo.Zeros.String {
		container = JSONData{
			ECS_CONTAINER_ID:      cid,
			ECS_CONTAINER_RUNTIME: engine.Mapper.MapStr(engine.SF_CONTAINER_TYPE)(rec),
			ECS_CONTAINER_PRIV:    engine.Mapper.MapInt(engine.SF_CONTAINER_PRIVILEGED)(rec) != 0,
			ECS_CONTAINER_NAME:    engine.Mapper.MapStr(engine.SF_CONTAINER_NAME)(rec),
		}
		imageid := engine.Mapper.MapStr(engine.SF_CONTAINER_IMAGEID)(rec)
		if imageid != sfgo.Zeros.String {
			image := JSONData{
				ECS_IMAGE_ID:   imageid,
				ECS_IMAGE_NAME: engine.Mapper.MapStr(engine.SF_CONTAINER_IMAGE)(rec),
			}
			container[ECS_IMAGE] = image
		}
	}
	return container
}

// encodeUser creates an ECS user field using user and group of the actual process.
func encodeUser(rec *engine.Record) JSONData {
	group := JSONData{
		ECS_GROUP_ID:   engine.Mapper.MapInt(engine.SF_PROC_GID)(rec),
		ECS_GROUP_NAME: engine.Mapper.MapStr(engine.SF_PROC_GROUP)(rec),
	}
	user := JSONData{
		ECS_USER_ID:   engine.Mapper.MapInt(engine.SF_PROC_UID)(rec),
		ECS_USER_NAME: engine.Mapper.MapStr(engine.SF_PROC_USER)(rec),
		ECS_GROUP:     group,
	}
	return user
}

// encodeProcess creates an ECS process field including the nested parent process.
func encodeProcess(rec *engine.Record) JSONData {
	exe := engine.Mapper.MapStr(engine.SF_PROC_EXE)(rec)
	process := JSONData{
		ECS_PROC_EXE:     exe,
		ECS_PROC_ARGS:    engine.Mapper.MapStr(engine.SF_PROC_ARGS)(rec),
		ECS_PROC_CMDLINE: engine.Mapper.MapStr(engine.SF_PROC_CMDLINE)(rec),
		ECS_PROC_PID:     engine.Mapper.MapInt(engine.SF_PROC_PID)(rec),
		ECS_PROC_START:   utils.ToIsoTimeStr(engine.Mapper.MapInt(engine.SF_PROC_CREATETS)(rec)),
		ECS_PROC_NAME:    path.Base(exe),
		ECS_PROC_THREAD:  JSONData{ECS_PROC_TID: engine.Mapper.MapInt(engine.SF_PROC_TID)(rec)},
	}
	pexe := engine.Mapper.MapStr(engine.SF_PPROC_EXE)(rec)
	parent := JSONData{
		ECS_PROC_EXE:     pexe,
		ECS_PROC_ARGS:    engine.Mapper.MapStr(engine.SF_PPROC_ARGS)(rec),
		ECS_PROC_CMDLINE: engine.Mapper.MapStr(engine.SF_PPROC_CMDLINE)(rec),
		ECS_PROC_PID:     engine.Mapper.MapInt(engine.SF_PPROC_PID)(rec),
		ECS_PROC_START:   utils.ToIsoTimeStr(engine.Mapper.MapInt(engine.SF_PPROC_CREATETS)(rec)),
		ECS_PROC_NAME:    path.Base(pexe),
	}
	process[ECS_PROC_PARENT] = parent
	return process
}

// encodeEvent creates the central ECS event field and sets the classification attributes
func encodeEvent(rec *engine.Record, category string, eventType string, action string) JSONData {
	start := engine.Mapper.MapInt(engine.SF_TS)(rec)
	end := engine.Mapper.MapInt(engine.SF_ENDTS)(rec)
	if end == sfgo.Zeros.Int64 {
		end = start
	}
	sfType := engine.Mapper.MapStr(engine.SF_TYPE)(rec)
	sfRet := engine.Mapper.MapInt(engine.SF_RET)(rec)

	// TODO: use JSONEncoder if we want the original
	//orig, _ := json.Marshal(rec)
	event := JSONData{
		ECS_EVENT_KIND:     ECS_KIND_EVENT,
		ECS_EVENT_CATEGORY: category,
		ECS_EVENT_TYPE:     eventType,
		ECS_EVENT_ACTION:   action,
		//ECS_EVENT_ORIGINAL: string(orig),
		ECS_EVENT_SFTYPE:   sfType,
		ECS_EVENT_START:    utils.ToIsoTimeStr(start),
		ECS_EVENT_END:      utils.ToIsoTimeStr(end),
		ECS_EVENT_DURATION: end - start,
	}
	if sfType == sfgo.TyPEStr || sfType == sfgo.TyFEStr {
		event[ECS_EVENT_SFRET] = sfRet
	}
	return event
}

// encodeFile creates an ECS file field
func encodeFile(rec *engine.Record) JSONData {
	opFlags := rec.GetInt(sfgo.EV_PROC_OPFLAGS_INT, sfgo.SYSFLOW_SRC)
	ft := engine.Mapper.MapStr(engine.SF_FILE_TYPE)(rec)
	fpath := engine.Mapper.MapStr(engine.SF_FILE_PATH)(rec)
	fd := engine.Mapper.MapInt(engine.SF_FILE_FD)(rec)
	pid := engine.Mapper.MapInt(engine.SF_PROC_PID)(rec)

	fileType := encodeFileType(ft)
	if opFlags&sfgo.OP_SYMLINK == sfgo.OP_SYMLINK {
		fileType = "symlink"
	}
	file := JSONData{ECS_FILE_TYPE: fileType}

	var name string
	if fpath != sfgo.Zeros.String {
		name = path.Base(fpath)
	} else {
		fpath = fmt.Sprintf("/proc/%d/fd/%d", pid, fd)
		name = strconv.FormatInt(fd, 10)
	}

	if fileType == "dir" {
		file[ECS_FILE_DIR] = fpath
	} else {
		file[ECS_FILE_NAME] = name
		file[ECS_FILE_DIR] = filepath.Dir(fpath)
		if fpath != name {
			file[ECS_FILE_PATH] = fpath
		}
	}
	return file
}

func encodeFileType(ft string) string {
	var fileType string
	switch ft {
	case "f":
		fileType = "file"
	case "d":
		fileType = "dir"
	case "u":
		fileType = "socket"
	case "p":
		fileType = "pipe"
	case "?":
		fallthrough
	default:
		fileType = "unknown"
	}
	return fileType
}

func extracTags(tags []engine.EnrichmentTag) []string {
	s := make([]string, 0)
	for _, v := range tags {
		switch v := v.(type) {
		case []string:
			s = append(s, v...)
		default:
			s = append(s, string(fmt.Sprintf("%v", v)))
		}
	}
	return s
}

// Cleanup cleans up resources.
func (t *ECSEncoder) Cleanup() {}
