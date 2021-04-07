//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
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
//
package exporter

import (
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
)

type JsonData map[string]interface{}

// struct for serializing ECS
type ECSRecord struct {
	Ts    string `json:"@timestamp"`
	Agent struct {
		Type    string `json:"type,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"agent,omitempty"`
	Container JsonData `json:"container"`
	Ecs       struct {
		Version string `json:"version,omitempty"`
	} `json:"ecs,omitempty"`
	Event       JsonData `json:"event"`
	File        JsonData `json:"file,omitempty"`
	FileAction  JsonData `json:"sf_file_action,omitempty"`
	Network     JsonData `json:"network,omitempty"`
	Source      JsonData `json:"source,omitempty"`
	Destination JsonData `json:"destination,omitempty"`
	Process     JsonData `json:"process"`
	User        JsonData `json:"user"`
}

// extracts the last part of a dot-separated id
func LastPart(id string) string {
	ld := strings.LastIndex(id, ".")
	if ld < 0 {
		return id
	}
	return id[ld+1:]
}

// converts a unix time value in ns to UTC time and returns an RFC3399 string
func ToIsoTimeStr(ts int64) string {
	ts_sec := int64(ts / 1E+9)
	ts_ns := int64(ts % 1E+9)
	t := time.Unix(ts_sec, ts_ns).In(time.UTC)
	return t.Format(time.RFC3339Nano)
}

// creates an ECS container field
func GetContainer(tr *TelemetryRecord) JsonData {
	var cid, ctype, cpriv, cname, cimage, cimageid interface{}

	if tr.FlatRecord != nil {
		data := tr.FlatRecord.Data
		if v, ok := data[engine.SF_CONTAINER_ID]; ok {
			cid = v
		}
		if v, ok := data[engine.SF_CONTAINER_TYPE]; ok {
			ctype = v
		}
		if v, ok := data[engine.SF_CONTAINER_PRIVILEGED]; ok {
			cpriv = v
		}
		if v, ok := data[engine.SF_CONTAINER_NAME]; ok {
			cname = v
		}
		if v, ok := data[engine.SF_CONTAINER_IMAGE]; ok {
			cimage = v
		}
		if v, ok := data[engine.SF_CONTAINER_IMAGEID]; ok {
			cimageid = v
		}
	} else {
		if tr.DataRecord.ContData != nil {
			data := tr.DataRecord.ContData.Container
			cid = data[LastPart(engine.SF_CONTAINER_ID)]
			ctype = data[LastPart(engine.SF_CONTAINER_TYPE)]
			cpriv = data[LastPart(engine.SF_CONTAINER_PRIVILEGED)]
			cname = data[LastPart(engine.SF_CONTAINER_NAME)]
			cimage = data[LastPart(engine.SF_CONTAINER_IMAGE)]
			cimageid = data[LastPart(engine.SF_CONTAINER_IMAGEID)]
		}
	}

	container := JsonData{
		ECS_CONTAINER_ID:      cid,
		ECS_CONTAINER_RUNTIME: ctype,
		ECS_CONTAINER_PRIV:    cpriv != 0,
		ECS_CONTAINER_NAME:    cname,
	}

	image := JsonData{ECS_IMAGE_NAME: cimage}
	if cimageid != "" {
		image[ECS_IMAGE_ID] = cimageid
	}
	container[ECS_IMAGE] = image

	return container
}

// creates an ECS user field using user and group of the actual process
func GetUser(tr *TelemetryRecord) JsonData {
	var uid, uname, gid, gname interface{}

	if tr.FlatRecord != nil {
		data := tr.FlatRecord.Data
		uid = data[engine.SF_PROC_UID]
		uname = data[engine.SF_PROC_USER]
		gid = data[engine.SF_PROC_GID]
		gname = data[engine.SF_PROC_GROUP]
	} else {
		data := tr.DataRecord.ProcData.Proc
		uid = data[LastPart(engine.SF_PROC_UID)]
		uname = data[LastPart(engine.SF_PROC_USER)]
		gid = data[LastPart(engine.SF_PROC_GID)]
		gname = data[LastPart(engine.SF_PROC_GROUP)]
	}

	user := JsonData{ECS_USER_ID: uid}
	if uname != "" {
		user[ECS_USER_NAME] = uname
	}

	if gid != "" {
		group := JsonData{ECS_GROUP_ID: gid}
		if gname != "" {
			group[ECS_GROUP_NAME] = gname
		}
		user[ECS_GROUP] = group
	}

	return user
}

// creates an ECS process field including the nested parent process
func GetProcess(tr *TelemetryRecord) JsonData {
	var args, exe string
	var pid, createts interface{}

	if tr.FlatRecord != nil {
		data := tr.FlatRecord.Data
		args = data[engine.SF_PROC_ARGS].(string)
		exe = data[engine.SF_PROC_EXE].(string)
		pid = data[engine.SF_PROC_PID]
		createts = data[engine.SF_PROC_CREATETS]
	} else {
		data := tr.DataRecord.ProcData.Proc
		args = data[LastPart(engine.SF_PROC_ARGS)].(string)
		exe = data[LastPart(engine.SF_PROC_EXE)].(string)
		pid = data[LastPart(engine.SF_PROC_PID)]
		createts = data[LastPart(engine.SF_PROC_CREATETS)]
	}

	argsCount := len(strings.Fields(args))
	if argsCount == 0 && len(args) != 0 {
		argsCount = 1
	}

	process := JsonData{
		ECS_PROC_ARGS_COUNT: len(strings.Fields(args)),
		ECS_PROC_EXE:        exe,
		ECS_PROC_PID:        pid,
		ECS_PROC_START:      ToIsoTimeStr(createts.(int64)),
	}
	if argsCount != 0 {
		process[ECS_PROC_ARGS] = args
		process[ECS_PROC_CMDLINE] = exe + " " + args
	} else {
		process[ECS_PROC_CMDLINE] = exe
	}
	if exe != "" {
		process[ECS_PROC_NAME] = path.Base(exe)
	} else {
		process[ECS_PROC_NAME] = ""
	}

	var pargs, pexe string
	var ppid, pcreatets interface{}

	if tr.FlatRecord != nil {
		data := tr.FlatRecord.Data
		if v, ok := data[engine.SF_PPROC_ARGS]; ok {
			pargs = v.(string)
		}
		if v, ok := data[engine.SF_PPROC_EXE]; ok {
			pexe = v.(string)
		}
		if v, ok := data[engine.SF_PPROC_PID]; ok {
			ppid = v
		} else {
			return process
		}
		if v, ok := data[engine.SF_PPROC_CREATETS]; ok {
			pcreatets = v
		}
	} else {
		if tr.DataRecord.PprocData != nil {
			data := tr.DataRecord.PprocData.Pproc
			pargs = data[LastPart(engine.SF_PPROC_ARGS)].(string)
			pexe = data[LastPart(engine.SF_PPROC_EXE)].(string)
			ppid = data[LastPart(engine.SF_PPROC_PID)]
			pcreatets = data[LastPart(engine.SF_PPROC_CREATETS)]
		} else {
			return process
		}
	}

	pArgsCount := len(strings.Fields(pargs))
	if pArgsCount == 0 && len(pargs) != 0 {
		pArgsCount = 1
	}

	parent := JsonData{
		ECS_PROC_ARGS_COUNT: pArgsCount,
		ECS_PROC_EXE:        pexe,
		ECS_PROC_PID:        ppid,
		ECS_PROC_START:      ToIsoTimeStr(pcreatets.(int64)),
	}
	if pArgsCount != 0 {
		parent[ECS_PROC_ARGS] = pargs
		parent[ECS_PROC_CMDLINE] = pexe + " " + pargs
	} else {
		parent[ECS_PROC_CMDLINE] = pexe
	}
	if pexe != "" {
		parent[ECS_PROC_NAME] = path.Base(pexe)
	} else {
		parent[ECS_PROC_NAME] = ""
	}

	process[ECS_PROC_PARENT] = parent

	return process
}

// populates the ECS representatiom of a NF record
func (ecs *ECSRecord) HandleNF(tr *TelemetryRecord) {
	var rops, wops, sport, dport, sip, dip, proto interface{}
	var rbytes, wbytes int64

	if tr.FlatRecord != nil {
		data := tr.FlatRecord.Data
		rbytes = data[engine.SF_FLOW_RBYTES].(int64)
		rops = data[engine.SF_FLOW_ROPS]
		wbytes = data[engine.SF_FLOW_WBYTES].(int64)
		wops = data[engine.SF_FLOW_WOPS]
		sip = data[engine.SF_NET_SIP]
		dip = data[engine.SF_NET_DIP]
		sport = data[engine.SF_NET_SPORT]
		dport = data[engine.SF_NET_DPORT]
		proto = data[engine.SF_NET_PROTO]
	} else {
		data := tr.DataRecord
		rbytes = data.FlowData.Flow[LastPart(engine.SF_FLOW_RBYTES)].(int64)
		rops = data.FlowData.Flow[LastPart(engine.SF_FLOW_ROPS)]
		wbytes = data.FlowData.Flow[LastPart(engine.SF_FLOW_WBYTES)].(int64)
		wops = data.FlowData.Flow[LastPart(engine.SF_FLOW_WOPS)]
		sip = data.NetData.Net[LastPart(engine.SF_NET_SIP)]
		dip = data.NetData.Net[LastPart(engine.SF_NET_DIP)]
		sport = data.NetData.Net[LastPart(engine.SF_NET_SPORT)]
		dport = data.NetData.Net[LastPart(engine.SF_NET_DPORT)]
		proto = data.NetData.Net[LastPart(engine.SF_NET_PROTO)]
	}

	(*ecs).Network = JsonData{
		ECS_NET_BYTES: rbytes + wbytes,
		ECS_NET_CID:   fmt.Sprintf("%s:%d-%s:%d", sip, sport, dip, dport),
	}

	if proto != nil {
		(*ecs).Network[ECS_NET_IANA] = strconv.FormatInt(proto.(int64), 10)
		(*ecs).Network[ECS_NET_PROTO] = sfgo.GetProto(proto.(int64))
	}

	(*ecs).Source = JsonData{
		ECS_ENDPOINT_IP:      sip,
		ECS_ENDPOINT_PORT:    sport,
		ECS_ENDPOINT_ADDR:    sip,
		ECS_ENDPOINT_BYTES:   wbytes,
		ECS_ENDPOINT_PACKETS: wops,
	}

	(*ecs).Destination = JsonData{
		ECS_ENDPOINT_IP:      dip,
		ECS_ENDPOINT_PORT:    dport,
		ECS_ENDPOINT_ADDR:    dip,
		ECS_ENDPOINT_BYTES:   rbytes,
		ECS_ENDPOINT_PACKETS: rops,
	}

	(*ecs).Event = GetEvent(tr, ECS_CAT_NETWORK, ECS_TYPE_CONNECTION, ECS_CAT_NETWORK+"-"+ECS_ACTION_TRAFFIC)
}

// creates the central ECS event field and sets the classification attributes
func GetEvent(tr *TelemetryRecord, category string, event_type string, action string) JsonData {
	var start, end, sf_ret int64
	var sf_type string

	if tr.FlatRecord != nil {
		data := tr.FlatRecord.Data
		start = data[engine.SF_TS].(int64)
		end = data[engine.SF_ENDTS].(int64)
		sf_type = data[engine.SF_TYPE].(string)
		sf_ret = data[engine.SF_RET].(int64)
	} else {
		start = tr.DataRecord.Ts
		end = tr.DataRecord.Endts
		sf_type = tr.DataRecord.Type
		sf_ret = tr.DataRecord.Ret
	}
	if end == 0 {
		end = start
	}

	orig, _ := json.Marshal(tr)
	event := JsonData{
		ECS_EVENT_KIND:     ECS_KIND_EVENT,
		ECS_EVENT_CATEGORY: category,
		ECS_EVENT_TYPE:     event_type,
		ECS_EVENT_ACTION:   action,
		ECS_EVENT_ORIGINAL: string(orig),
		ECS_EVENT_SFTYPE:   sf_type,
		ECS_EVENT_START:    ToIsoTimeStr(start),
		ECS_EVENT_END:      ToIsoTimeStr(end),
		ECS_EVENT_DURATION: end - start,
	}
	if sf_type == sfgo.TyPEStr || sf_type == sfgo.TyFEStr {
		event[ECS_EVENT_SFRET] = sf_ret
	}

	return event
}

// creates an ECS file field
func GetFile(tr *TelemetryRecord) JsonData {
	var opFlags string
	var ft string
	var fpath string
	var fd int64
	var pid int64

	if tr.FlatRecord != nil {
		opFlags = tr.FlatRecord.Data[engine.SF_OPFLAGS].(string)
		ft = tr.FlatRecord.Data[engine.SF_FILE_TYPE].(string)
		fpath = tr.FlatRecord.Data[engine.SF_FILE_PATH].(string)
		fd = tr.FlatRecord.Data[engine.SF_FILE_FD].(int64)
		pid = tr.FlatRecord.Data[engine.SF_PROC_PID].(int64)
	} else {
		opFlags = strings.Join(tr.DataRecord.Opflags, engine.LISTSEP)
		ft = tr.DataRecord.FileData.File[LastPart(engine.SF_FILE_TYPE)].(string)
		fpath = tr.DataRecord.FileData.File[LastPart(engine.SF_FILE_PATH)].(string)
		fd = tr.DataRecord.FileData.File[LastPart(engine.SF_FILE_FD)].(int64)
		pid = tr.DataRecord.ProcData.Proc[LastPart(engine.SF_PROC_PID)].(int64)
	}

	file_type := ""
	switch ft {
	case "f":
		file_type = "file"
	case "d":
		file_type = "dir"
	case "u":
		file_type = "socket"
	case "p":
		file_type = "pipe"
	case "?":
		file_type = "unknown"
	}

	if strings.Contains(opFlags, sfgo.OpFlagMkdir) || strings.Contains(opFlags, sfgo.OpFlagRmdir) {
		file_type = "dir"
	} else if strings.Contains(opFlags, sfgo.OpFlagSymlink) {
		file_type = "symlink"
	}

	file := JsonData{ECS_FILE_TYPE: file_type}

	var name string
	if fpath != "" {
		name = path.Base(fpath)
	} else {
		fpath = fmt.Sprintf("/proc/%d/fd/%d", pid, fd)
		name = strconv.FormatInt(fd, 10)
	}

	if file_type == "dir" {
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

// populates the ECS representatiom of a FF record
func (ecs *ECSRecord) HandleFF(tr *TelemetryRecord) {
	var opFlags string
	var rbytes int64
	var rops int64
	var wbytes int64
	var wops int64
	if tr.FlatRecord != nil {
		data := tr.FlatRecord.Data
		opFlags = data[engine.SF_OPFLAGS].(string)
		rbytes = data[engine.SF_FLOW_RBYTES].(int64)
		rops = data[engine.SF_FLOW_ROPS].(int64)
		wbytes = data[engine.SF_FLOW_WBYTES].(int64)
		wops = data[engine.SF_FLOW_WOPS].(int64)
	} else {
		opFlags = strings.Join(tr.DataRecord.Opflags, engine.LISTSEP)
		data := tr.DataRecord.FlowData.Flow
		rbytes = data[LastPart(engine.SF_FLOW_RBYTES)].(int64)
		rops = data[LastPart(engine.SF_FLOW_ROPS)].(int64)
		wbytes = data[LastPart(engine.SF_FLOW_WBYTES)].(int64)
		wops = data[LastPart(engine.SF_FLOW_WOPS)].(int64)
	}

	category := ECS_CAT_FILE
	event_type := ECS_TYPE_ACCESS
	action := category + "-" + event_type

	if strings.Contains(opFlags, sfgo.OpFlagWrite) && wbytes > 0 {
		event_type = ECS_TYPE_CHANGE
		action = category + "-" + ECS_ACTION_WRITE
	}
	if strings.Contains(opFlags, sfgo.OpFlagRead) && rbytes > 0 {
		action = category + "-" + ECS_ACTION_READ
	}

	ecs.Event = GetEvent(tr, category, event_type, action)
	ecs.File = GetFile(tr)
	ecs.FileAction = JsonData{
		ECS_SF_FA_RBYTES: rbytes,
		ECS_SF_FA_ROPS:   rops,
		ECS_SF_FA_WBYTES: wbytes,
		ECS_SF_FA_WOPS:   wops,
	}
}

// populates the ECS representatiom of a FE record
func (ecs *ECSRecord) HandleFE(tr *TelemetryRecord) {
	var opFlags string
	var target_path string
	if tr.FlatRecord != nil {
		opFlags = tr.FlatRecord.Data[engine.SF_OPFLAGS].(string)
		target_path = tr.FlatRecord.Data[engine.SF_FILE_NEWPATH].(string)
	} else {
		opFlags = strings.Join(tr.DataRecord.Opflags, engine.LISTSEP)
		target_path = tr.DataRecord.FileData.File[LastPart(engine.SF_FILE_NEWPATH)].(string)
	}

	ecs.File = GetFile(tr)

	category := ECS_CAT_FILE
	event_type := ECS_TYPE_CHANGE
	action := category + "-" + event_type

	if strings.Contains(opFlags, sfgo.OpFlagMkdir) {
		category = ECS_CAT_DIR
		event_type = ECS_TYPE_CREATE
		action = category + "-" + ECS_ACTION_CREATE
	} else if strings.Contains(opFlags, sfgo.OpFlagRmdir) {
		category = ECS_CAT_DIR
		event_type = ECS_TYPE_DELETE
		action = category + "-" + ECS_ACTION_DELETE
	} else if strings.Contains(opFlags, sfgo.OpFlagUnlink) {
		event_type = ECS_TYPE_DELETE
		action = category + "-" + ECS_ACTION_DELETE
	} else if strings.Contains(opFlags, sfgo.OpFlagSymlink) || strings.Contains(opFlags, sfgo.OpFlagLink) {
		action = category + "-" + ECS_ACTION_LINK
		ecs.File[ECS_FILE_TARGET] = target_path
	} else if strings.Contains(opFlags, sfgo.OpFlagRename) {
		action = category + "-" + ECS_ACTION_RENAME
		ecs.File[ECS_FILE_TARGET] = target_path
	}

	ecs.Event = GetEvent(tr, category, event_type, action)
}

// populates the ECS representatiom of a PE record
func (ecs *ECSRecord) HandlePE(tr *TelemetryRecord) {
	category := ECS_CAT_PROCESS
	event_type := ECS_TYPE_START

	var opFlags string
	var pid int64
	var tid int64
	if tr.FlatRecord != nil {
		opFlags = tr.FlatRecord.Data[engine.SF_OPFLAGS].(string)
		pid = tr.FlatRecord.Data[engine.SF_PROC_PID].(int64)
		tid = tr.FlatRecord.Data[engine.SF_PROC_TID].(int64)
	} else {
		opFlags = strings.Join(tr.DataRecord.Opflags, engine.LISTSEP)
		pid = tr.DataRecord.ProcData.Proc[LastPart(engine.SF_PROC_PID)].(int64)
		pid = tr.DataRecord.ProcData.Proc[LastPart(engine.SF_PROC_TID)].(int64)
	}

	if strings.Contains(opFlags, sfgo.OpFlagExit) {
		if pid != tid {
			event_type = ECS_TYPE_TEND
		} else {
			event_type = ECS_TYPE_END
		}
	} else if strings.Contains(opFlags, sfgo.OpFlagClone) || strings.Contains(opFlags, sfgo.OpFlagExec) {
		if pid != tid {
			event_type = ECS_TYPE_TSTART
		}
	} else if strings.Contains(opFlags, sfgo.OpFlagSetuid) {
		event_type = ECS_TYPE_CHANGE
	}

	action := category + "-" + event_type

	ecs.Event = GetEvent(tr, category, event_type, action)
}

// returns the ECS representation for the given telemetry record
func ToECS(tr TelemetryRecord) *ECSRecord {
	ecs := &ECSRecord{
		Container: GetContainer(&tr),
		Process:   GetProcess(&tr),
		User:      GetUser(&tr),
	}
	ecs.Agent.Version = tr.Version
	ecs.Agent.Type = ECS_AGENT_TYPE
	ecs.Ecs.Version = ECS_VERSION

	var sfType string
	if tr.FlatRecord != nil {
		ecs.Ts = ToIsoTimeStr(tr.FlatRecord.Data[engine.SF_TS].(int64))
		sfType = tr.FlatRecord.Data[engine.SF_TYPE].(string)
	} else {
		ecs.Ts = ToIsoTimeStr(tr.DataRecord.Ts)
		sfType = tr.DataRecord.Type
	}

	switch sfType {
	case sfgo.TyNFStr:
		ecs.HandleNF(&tr)
	case sfgo.TyFFStr:
		ecs.HandleFF(&tr)
	case sfgo.TyFEStr:
		ecs.HandleFE(&tr)
	case sfgo.TyPEStr:
		ecs.HandlePE(&tr)
	}

	// map policy ids to event.reason
	if len(tr.Policies) > 0 {
		reason := make([]string, 0)
		for _, p := range tr.Policies {
			reason = append(reason, p.ID)
		}
		ecs.Event[ECS_EVENT_REASON] = strings.Join(reason, ", ")
	}

	return ecs
}
