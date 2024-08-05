//go:build flatrecord
// +build flatrecord

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
	"encoding/binary"
	"fmt"
	"net"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cespare/xxhash/v2"
	"github.com/satta/gommunityid"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/utils"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/flatrecord"
	"github.com/tidwall/gjson"
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
	Event        JSONData   `json:"event"`
	Host         JSONData   `json:"host"`
	Container    JSONData   `json:"container,omitempty"`
	Orchestrator JSONData   `json:"orchestrator,omitempty"`
	Pod          JSONData   `json:"pod,omitempty"`
	Service      []JSONData `json:"service,omitempty"`
	File         JSONData   `json:"file,omitempty"`
	FileAction   JSONData   `json:"sf_file_action,omitempty"`
	Network      JSONData   `json:"network,omitempty"`
	Source       JSONData   `json:"source,omitempty"`
	Destination  JSONData   `json:"destination,omitempty"`
	Process      JSONData   `json:"process,omitempty"`
	User         JSONData   `json:"user,omitempty"`
	Tags         []string   `json:"tags,omitempty"`
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
func (t *ECSEncoder) Encode(recs []*flatrecord.Record) ([]commons.EncodedData, error) {
	t.batch = t.batch[:0]
	for _, rec := range recs {
		ecs := t.encode(rec)
		t.batch = append(t.batch, ecs)
	}
	return t.batch, nil
}

// Encodes a telemetry record into an ECS representation.
func (t *ECSEncoder) encode(rec *flatrecord.Record) *ECSRecord {
	ecs := &ECSRecord{
		ID:   encodeID(rec),
		Host: encodeHost(rec),
	}
	ecs.Agent.Version = t.config.Version
	ecs.Agent.Type = ECS_AGENT_TYPE
	ecs.Ecs.Version = t.config.EcsVersion
	ecs.Ts = utils.ToIsoTimeStr(flatrecord.Mapper.MapInt(flatrecord.SF_TS)(rec))

	// encode specific record components
	sfType := flatrecord.Mapper.MapStr(flatrecord.SF_TYPE)(rec)
	if sfType != sfgo.TyKEStr {
		ecs.Container = encodeContainer(rec)
		if flatrecord.Mapper.MapStr(flatrecord.SF_POD_ID)(rec) != sfgo.Zeros.String {
			ecs.encodeOrchestrator(rec)
			ecs.encodePod(rec)
		}
		ecs.Process = encodeProcess(rec)
		ecs.User = encodeUser(rec)
	} else {
		ecs.encodeOrchestrator(rec)
	}

	switch sfType {
	case sfgo.TyNFStr:
		ecs.encodeNetworkFlow(rec)
	case sfgo.TyFFStr:
		ecs.encodeFileFlow(rec)
	case sfgo.TyFEStr:
		ecs.encodeFileEvent(rec)
	case sfgo.TyPEStr:
		ecs.encodeProcessEvent(rec)
	case sfgo.TyKEStr:
		ecs.encodeK8sEvent(rec)
	}

	// encode tags and policy information
	tags := rec.Ctx.GetTags()
	rules := rec.Ctx.GetRules()
	if len(rules) > 0 {
		reasons := make([]string, 0)
		priority := int(policy.Low)
		for _, r := range rules {
			reasons = append(reasons, r.Name)
			tags = append(tags, extracTags(r.Tags)...)
			priority = utils.Max(priority, int(r.Priority))
		}
		ecs.Event[ECS_EVENT_REASON] = strings.Join(reasons, ", ")
		ecs.Event[ECS_EVENT_SEVERITY] = priority
	}
	if len(tags) > 0 {
		ecs.Tags = tags
	}

	return ecs
}

var byteInt64 []byte = make([]byte, 8)

// encodeID returns the ECS document identifier.
func encodeID(rec *flatrecord.Record) string {
	h := xxhash.New()
	t := flatrecord.Mapper.MapStr(flatrecord.SF_TYPE)(rec)
	h.Write([]byte(flatrecord.Mapper.MapStr(flatrecord.SF_NODE_ID)(rec)))
	h.Write([]byte(flatrecord.Mapper.MapStr(flatrecord.SF_CONTAINER_ID)(rec)))
	binary.LittleEndian.PutUint64(byteInt64, uint64(rec.GetInt(sfgo.TS_INT, sfgo.SYSFLOW_SRC)))
	h.Write(byteInt64)
	binary.LittleEndian.PutUint64(byteInt64, uint64(rec.GetInt(sfgo.TID_INT, sfgo.SYSFLOW_SRC)))
	h.Write(byteInt64)
	binary.LittleEndian.PutUint64(byteInt64, uint64(rec.GetInt(sfgo.PROC_OID_CREATETS_INT, sfgo.SYSFLOW_SRC)))
	h.Write(byteInt64)
	binary.LittleEndian.PutUint64(byteInt64, uint64(rec.GetInt(sfgo.EV_PROC_OPFLAGS_INT, sfgo.SYSFLOW_SRC)))
	h.Write(byteInt64)
	switch t {
	case sfgo.TyFFStr, sfgo.TyFEStr:
		h.Write([]byte(flatrecord.Mapper.MapStr(flatrecord.SF_FILE_OID)(rec)))
	case sfgo.TyNFStr:
		binary.LittleEndian.PutUint64(byteInt64, uint64(rec.GetInt(sfgo.FL_NETW_SIP_INT, sfgo.SYSFLOW_SRC)))
		h.Write(byteInt64)
		binary.LittleEndian.PutUint64(byteInt64, uint64(rec.GetInt(sfgo.FL_NETW_SPORT_INT, sfgo.SYSFLOW_SRC)))
		h.Write(byteInt64)
		binary.LittleEndian.PutUint64(byteInt64, uint64(rec.GetInt(sfgo.FL_NETW_DIP_INT, sfgo.SYSFLOW_SRC)))
		h.Write(byteInt64)
		binary.LittleEndian.PutUint64(byteInt64, uint64(rec.GetInt(sfgo.FL_NETW_DPORT_INT, sfgo.SYSFLOW_SRC)))
		h.Write(byteInt64)
		binary.LittleEndian.PutUint64(byteInt64, uint64(rec.GetInt(sfgo.FL_NETW_PROTO_INT, sfgo.SYSFLOW_SRC)))
		h.Write(byteInt64)
	case sfgo.TyKEStr:
		binary.LittleEndian.PutUint64(byteInt64, uint64(rec.GetInt(sfgo.K8SE_ACTION_INT, sfgo.SYSFLOW_SRC)))
		h.Write(byteInt64)
		binary.LittleEndian.PutUint64(byteInt64, uint64(rec.GetInt(sfgo.K8SE_KIND_INT, sfgo.SYSFLOW_SRC)))
		h.Write(byteInt64)
		h.Write([]byte(flatrecord.Mapper.MapStr(flatrecord.SF_K8SE_MESSAGE)(rec)))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// encodeNetworkFlow populates the ECS representatiom of a NetworkFlow record.
func (ecs *ECSRecord) encodeNetworkFlow(rec *flatrecord.Record) {
	rbytes := flatrecord.Mapper.MapInt(flatrecord.SF_FLOW_RBYTES)(rec)
	rops := flatrecord.Mapper.MapInt(flatrecord.SF_FLOW_ROPS)(rec)
	wbytes := flatrecord.Mapper.MapInt(flatrecord.SF_FLOW_WBYTES)(rec)
	wops := flatrecord.Mapper.MapInt(flatrecord.SF_FLOW_WOPS)(rec)
	sip := flatrecord.Mapper.MapStr(flatrecord.SF_NET_SIP)(rec)
	dip := flatrecord.Mapper.MapStr(flatrecord.SF_NET_DIP)(rec)
	sport := flatrecord.Mapper.MapInt(flatrecord.SF_NET_SPORT)(rec)
	dport := flatrecord.Mapper.MapInt(flatrecord.SF_NET_DPORT)(rec)
	proto := flatrecord.Mapper.MapInt(flatrecord.SF_NET_PROTO)(rec)

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
func (ecs *ECSRecord) encodeFileFlow(rec *flatrecord.Record) {
	opFlags := rec.GetInt(sfgo.EV_PROC_OPFLAGS_INT, sfgo.SYSFLOW_SRC)
	rbytes := flatrecord.Mapper.MapInt(flatrecord.SF_FLOW_RBYTES)(rec)
	rops := flatrecord.Mapper.MapInt(flatrecord.SF_FLOW_ROPS)(rec)
	wbytes := flatrecord.Mapper.MapInt(flatrecord.SF_FLOW_WBYTES)(rec)
	wops := flatrecord.Mapper.MapInt(flatrecord.SF_FLOW_WOPS)(rec)
	category := ECS_CAT_FILE
	eventType := ECS_TYPE_ACCESS
	action := category + "-" + eventType
	if opFlags&sfgo.OP_READ_RECV == sfgo.OP_READ_RECV && (rbytes > 0 || rops > 0) {
		action = action + "-" + ECS_ACTION_READ
	}
	if opFlags&sfgo.OP_WRITE_SEND == sfgo.OP_WRITE_SEND && (wbytes > 0 || wops > 0) {
		eventType = ECS_TYPE_CHANGE
		action = action + "-" + ECS_ACTION_WRITE
	}
	ecs.Event = encodeEvent(rec, category, eventType, action)
	ecs.File = encodeFile(rec)
	if rbytes > 0 || rops > 0 || wbytes > 0 || wops > 0 {
		ecs.FileAction = JSONData{
			ECS_SF_FA_RBYTES: rbytes,
			ECS_SF_FA_ROPS:   rops,
			ECS_SF_FA_WBYTES: wbytes,
			ECS_SF_FA_WOPS:   wops,
		}
	}
}

// encodeFileEvent populates the ECS representatiom of a FE record
func (ecs *ECSRecord) encodeFileEvent(rec *flatrecord.Record) {
	opFlags := rec.GetInt(sfgo.EV_PROC_OPFLAGS_INT, sfgo.SYSFLOW_SRC)
	targetPath := flatrecord.Mapper.MapStr(flatrecord.SF_FILE_NEWPATH)(rec)
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
func (ecs *ECSRecord) encodeProcessEvent(rec *flatrecord.Record) {
	opFlags := rec.GetInt(sfgo.EV_PROC_OPFLAGS_INT, sfgo.SYSFLOW_SRC)
	pid := flatrecord.Mapper.MapInt(flatrecord.SF_PROC_PID)(rec)
	tid := flatrecord.Mapper.MapInt(flatrecord.SF_PROC_TID)(rec)
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

func k8sActionToEventType(rec *flatrecord.Record) string {
	eventType := ECS_TYPE_INFO
	am := flatrecord.Mapper.Mappers[flatrecord.SF_K8SE_ACTION]
	switch sfgo.K8sAction(rec.Fr.Ints[am.Source][am.FlatIndex]) {
	case sfgo.K8sActionK8S_COMPONENT_ADDED:
		eventType = ECS_TYPE_CREATE
	case sfgo.K8sActionK8S_COMPONENT_DELETED:
		eventType = ECS_TYPE_DELETE
	case sfgo.K8sActionK8S_COMPONENT_MODIFIED:
		eventType = ECS_TYPE_CHANGE
	case sfgo.K8sActionK8S_COMPONENT_ERROR:
		eventType = ECS_TYPE_ERROR
	}
	return eventType
}

// encodeK8sEvent populates the ECS representatiom of a KE record
func (ecs *ECSRecord) encodeK8sEvent(rec *flatrecord.Record) {
	category := ECS_CAT_ORCH
	eventType := k8sActionToEventType(rec)
	action := flatrecord.Mapper.MapStr(flatrecord.SF_K8SE_ACTION)(rec)

	ecs.Event = encodeEvent(rec, category, eventType, action)
	msgStr := flatrecord.Mapper.MapStr(flatrecord.SF_K8SE_MESSAGE)(rec)
	ecs.Event[ECS_EVENT_ORIGINAL] = msgStr

	msg := gjson.Parse(msgStr)
	ecs.Orchestrator = JSONData{
		ECS_ORCHESTRATOR_NAMESPACE: msg.Get("items.0.namespace").String(),
		ECS_ORCHESTRATOR_RESOURCE: JSONData{
			ECS_RESOURCE_TYPE: strings.ToLower(msg.Get("kind").String()),
			ECS_RESOURCE_NAME: msg.Get("items.0.name").String(),
		},
		ECS_ORCHESTRATOR_TYPE: "kubernetes",
	}
}

// encodeOrchestrator creates an ECS orchestrator field.
func (ecs *ECSRecord) encodeOrchestrator(rec *flatrecord.Record) {
	ecs.Orchestrator = JSONData{
		ECS_ORCHESTRATOR_NAMESPACE: flatrecord.Mapper.MapStr(flatrecord.SF_POD_NAMESPACE)(rec),
		ECS_ORCHESTRATOR_RESOURCE: JSONData{
			ECS_RESOURCE_TYPE: "pod",
			ECS_RESOURCE_NAME: flatrecord.Mapper.MapStr(flatrecord.SF_POD_NAME)(rec),
		},
		ECS_ORCHESTRATOR_TYPE: "kubernetes",
	}
}

// encodePod creates a custom ECS pod field.
func (ecs *ECSRecord) encodePod(rec *flatrecord.Record) {
	ecs.Pod = JSONData{
		ECS_POD_TS:           utils.ToIsoTimeStr(flatrecord.Mapper.MapInt(flatrecord.SF_POD_TS)(rec)),
		ECS_POD_ID:           flatrecord.Mapper.MapStr(flatrecord.SF_POD_ID)(rec),
		ECS_POD_NAME:         flatrecord.Mapper.MapStr(flatrecord.SF_POD_NAME)(rec),
		ECS_POD_NODENAME:     flatrecord.Mapper.MapStr(flatrecord.SF_POD_NODENAME)(rec),
		ECS_POD_NAMESPACE:    flatrecord.Mapper.MapStr(flatrecord.SF_POD_NAMESPACE)(rec),
		ECS_POD_HOSTIP:       utils.ToIPStrArray(flatrecord.Mapper.MapIntArray(flatrecord.SF_POD_HOSTIP)(rec)),
		ECS_POD_INTERNALIP:   utils.ToIPStrArray(flatrecord.Mapper.MapIntArray(flatrecord.SF_POD_INTERNALIP)(rec)),
		ECS_POD_RESTARTCOUNT: flatrecord.Mapper.MapInt(flatrecord.SF_POD_RESTARTCOUNT)(rec),
	}

	services := flatrecord.Mapper.MapSvcArray(flatrecord.SF_POD_SERVICES)(rec)
	if services != sfgo.Zeros.Any && len(*services) > 0 {
		ecs.encodeService(services)
	}
}

// encodeServices creates an ECS service field.
func (ecs *ECSRecord) encodeService(svcs *[]*sfgo.Service) {
	ecs.Service = make([]JSONData, len(*svcs))
	for i, svc := range *svcs {
		ecs.Service[i] = JSONData{
			ECS_SERVICE_ID:        svc.Id,
			ECS_SERVICE_NAME:      svc.Name,
			ECS_SERVICE_NAMESPACE: svc.Namespace,
			ECS_SERVICE_CLUSTERIP: utils.ToIPStrArray(&svc.ClusterIP),
			ECS_SERVICE_PORTLIST:  encodePortList(&svc.PortList),
		}
	}
}

// encodePortList creates a ports field for an ECS service field.
func encodePortList(pl *[]*sfgo.Port) []JSONData {
	ports := make([]JSONData, len(*pl))
	for i, p := range *pl {
		ports[i] = JSONData{
			ECS_SERVICE_PORT:       p.Port,
			ECS_SERVICE_TARGETPORT: p.TargetPort,
			ECS_SERVICE_NODEPORT:   p.NodePort,
			ECS_SERVICE_PROTO:      p.Proto,
		}
	}
	return ports
}

// encodeContainer creates an ECS container field.
func encodeContainer(rec *flatrecord.Record) JSONData {
	var container JSONData
	cid := flatrecord.Mapper.MapStr(flatrecord.SF_CONTAINER_ID)(rec)
	if cid != sfgo.Zeros.String {
		container = JSONData{
			ECS_CONTAINER_ID:      cid,
			ECS_CONTAINER_RUNTIME: flatrecord.Mapper.MapStr(flatrecord.SF_CONTAINER_TYPE)(rec),
			ECS_CONTAINER_PRIV:    flatrecord.Mapper.MapInt(flatrecord.SF_CONTAINER_PRIVILEGED)(rec) != 0,
			ECS_CONTAINER_NAME:    flatrecord.Mapper.MapStr(flatrecord.SF_CONTAINER_NAME)(rec),
		}
		imageid := flatrecord.Mapper.MapStr(flatrecord.SF_CONTAINER_IMAGEID)(rec)
		if imageid != sfgo.Zeros.String {
			image := JSONData{
				ECS_IMAGE_ID:   imageid,
				ECS_IMAGE_NAME: flatrecord.Mapper.MapStr(flatrecord.SF_CONTAINER_IMAGE)(rec),
			}
			container[ECS_IMAGE] = image
		}
	}
	return container
}

// encodeHost creates the ECS host field
func encodeHost(rec *flatrecord.Record) JSONData {
	return JSONData{
		ECS_HOST_ID: flatrecord.Mapper.MapStr(flatrecord.SF_NODE_ID)(rec),
		ECS_HOST_IP: flatrecord.Mapper.MapStr(flatrecord.SF_NODE_IP)(rec),
	}
}

// encodeUser creates an ECS user field using user and group of the actual process.
func encodeUser(rec *flatrecord.Record) JSONData {
	gname := flatrecord.Mapper.MapStr(flatrecord.SF_PROC_GROUP)(rec)
	group := JSONData{
		ECS_GROUP_ID: flatrecord.Mapper.MapInt(flatrecord.SF_PROC_GID)(rec),
	}
	if gname != sfgo.Zeros.String {
		group[ECS_GROUP_NAME] = gname
	}
	uname := flatrecord.Mapper.MapStr(flatrecord.SF_PROC_USER)(rec)
	user := JSONData{
		ECS_GROUP:   group,
		ECS_USER_ID: flatrecord.Mapper.MapInt(flatrecord.SF_PROC_UID)(rec),
	}
	if uname != sfgo.Zeros.String {
		user[ECS_USER_NAME] = uname
	}
	return user
}

// encodeProcess creates an ECS process field including the nested parent process.
func encodeProcess(rec *flatrecord.Record) JSONData {
	exe := flatrecord.Mapper.MapStr(flatrecord.SF_PROC_EXE)(rec)
	args_count := 0
	if flatrecord.Mapper.MapStr(flatrecord.SF_PROC_ARGS)(rec) != "" {
		args_count = len(strings.Split(flatrecord.Mapper.MapStr(flatrecord.SF_PROC_ARGS)(rec), " "))
	}
	process := JSONData{
		ECS_PROC_EXE:        exe,
		ECS_PROC_ARGS:       flatrecord.Mapper.MapStr(flatrecord.SF_PROC_ARGS)(rec),
		ECS_PROC_ARGS_COUNT: args_count,
		ECS_PROC_CMDLINE:    flatrecord.Mapper.MapStr(flatrecord.SF_PROC_CMDLINE)(rec),
		ECS_PROC_PID:        flatrecord.Mapper.MapInt(flatrecord.SF_PROC_PID)(rec),
		ECS_PROC_START:      utils.ToIsoTimeStr(flatrecord.Mapper.MapInt(flatrecord.SF_PROC_CREATETS)(rec)),
		ECS_PROC_NAME:       path.Base(exe),
		ECS_PROC_THREAD:     JSONData{ECS_PROC_TID: flatrecord.Mapper.MapInt(flatrecord.SF_PROC_TID)(rec)},
	}
	pexe := flatrecord.Mapper.MapStr(flatrecord.SF_PPROC_EXE)(rec)
	pargs_count := 0
	if flatrecord.Mapper.MapStr(flatrecord.SF_PPROC_ARGS)(rec) != "" {
		pargs_count = len(strings.Split(flatrecord.Mapper.MapStr(flatrecord.SF_PPROC_ARGS)(rec), " "))
	}
	parent := JSONData{
		ECS_PROC_EXE:        pexe,
		ECS_PROC_ARGS:       flatrecord.Mapper.MapStr(flatrecord.SF_PPROC_ARGS)(rec),
		ECS_PROC_ARGS_COUNT: pargs_count,
		ECS_PROC_CMDLINE:    flatrecord.Mapper.MapStr(flatrecord.SF_PPROC_CMDLINE)(rec),
		ECS_PROC_PID:        flatrecord.Mapper.MapInt(flatrecord.SF_PPROC_PID)(rec),
		ECS_PROC_START:      utils.ToIsoTimeStr(flatrecord.Mapper.MapInt(flatrecord.SF_PPROC_CREATETS)(rec)),
		ECS_PROC_NAME:       path.Base(pexe),
	}
	process[ECS_PROC_PARENT] = parent
	return process
}

// encodeEvent creates the central ECS event field and sets the classification attributes
func encodeEvent(rec *flatrecord.Record, category string, eventType string, action string) JSONData {
	start := flatrecord.Mapper.MapInt(flatrecord.SF_TS)(rec)
	end := flatrecord.Mapper.MapInt(flatrecord.SF_ENDTS)(rec)
	if end == sfgo.Zeros.Int64 {
		end = start
	}
	sfType := flatrecord.Mapper.MapStr(flatrecord.SF_TYPE)(rec)
	sfRet := flatrecord.Mapper.MapInt(flatrecord.SF_RET)(rec)

	event := JSONData{
		ECS_EVENT_CATEGORY: category,
		ECS_EVENT_TYPE:     eventType,
		ECS_EVENT_ACTION:   action,
		ECS_EVENT_SFTYPE:   sfType,
		ECS_EVENT_START:    utils.ToIsoTimeStr(start),
		ECS_EVENT_END:      utils.ToIsoTimeStr(end),
		ECS_EVENT_DURATION: end - start,
	}

	if rec.Ctx.IsAlert() {
		event[ECS_EVENT_KIND] = ECS_KIND_ALERT
	} else {
		event[ECS_EVENT_KIND] = ECS_KIND_EVENT
	}

	if sfType == sfgo.TyPEStr || sfType == sfgo.TyFEStr {
		event[ECS_EVENT_SFRET] = sfRet
	}
	return event
}

// encodeFile creates an ECS file field
func encodeFile(rec *flatrecord.Record) JSONData {
	opFlags := rec.GetInt(sfgo.EV_PROC_OPFLAGS_INT, sfgo.SYSFLOW_SRC)
	ft := flatrecord.Mapper.MapStr(flatrecord.SF_FILE_TYPE)(rec)
	fpath := flatrecord.Mapper.MapStr(flatrecord.SF_FILE_PATH)(rec)
	fd := flatrecord.Mapper.MapInt(flatrecord.SF_FILE_FD)(rec)
	pid := flatrecord.Mapper.MapInt(flatrecord.SF_PROC_PID)(rec)

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

func extracTags(tags []policy.EnrichmentTag) []string {
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
