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

// ECS_AGENT_TYPE denotes the ECS agent type.
const ECS_AGENT_TYPE = "SysFlow"

// ECS attributes used in JSONData.
const (
	ECS_CONTAINER_ID      = "id"
	ECS_CONTAINER_NAME    = "name"
	ECS_CONTAINER_RUNTIME = "runtime"
	ECS_CONTAINER_PRIV    = "sf_privileged"

	ECS_IMAGE      = "image"
	ECS_IMAGE_ID   = "id"
	ECS_IMAGE_NAME = "name"

	ECS_HOST_ID = "id"
        ECS_HOST_IP = "ip"

	ECS_EVENT_KIND     = "kind"
	ECS_EVENT_CATEGORY = "category"
	ECS_EVENT_TYPE     = "type"
	ECS_EVENT_ACTION   = "action"
	ECS_EVENT_ORIGINAL = "original"
	ECS_EVENT_START    = "start"
	ECS_EVENT_END      = "end"
	ECS_EVENT_DURATION = "duration"
	ECS_EVENT_SFTYPE   = "sf_type"
	ECS_EVENT_SFRET    = "sf_ret"
	ECS_EVENT_REASON   = "reason"
	ECS_EVENT_SEVERITY = "severity"

	ECS_FILE_DIR    = "directory"
	ECS_FILE_NAME   = "name"
	ECS_FILE_PATH   = "path"
	ECS_FILE_TARGET = "target_path"
	ECS_FILE_TYPE   = "type"

	ECS_GROUP      = "group"
	ECS_GROUP_ID   = "id"
	ECS_GROUP_NAME = "name"

	// used in proc and file fields
	ECS_HASH        = "hash"
	ECS_HASH_MD5    = "md5"
	ECS_HASH_SHA1   = "sha1"
	ECS_HASH_SHA256 = "sha256"

	ECS_NET_BYTES = "bytes"
	ECS_NET_CID   = "community_id"
	ECS_NET_IANA  = "iana_number"
	ECS_NET_PROTO = "protocol"

	// used in source and destination fields
	ECS_ENDPOINT_ADDR    = "address"
	ECS_ENDPOINT_BYTES   = "bytes"
	ECS_ENDPOINT_IP      = "ip"
	ECS_ENDPOINT_PACKETS = "packets"
	ECS_ENDPOINT_PORT    = "port"

	ECS_ORCHESTRATOR_NAMESPACE = "namespace"
	ECS_ORCHESTRATOR_RESOURCE  = "resource"
	ECS_RESOURCE_NAME          = "name"
	ECS_RESOURCE_TYPE          = "type"
	ECS_ORCHESTRATOR_TYPE      = "type"

	ECS_POD_TS           = "ts"
	ECS_POD_ID           = "id"
	ECS_POD_NAME         = "name"
	ECS_POD_NAMESPACE    = "namespace"
	ECS_POD_NODENAME     = "nodename"
	ECS_POD_HOSTIP       = "hostip"
	ECS_POD_INTERNALIP   = "internalip"
	ECS_POD_RESTARTCOUNT = "restartcnt"

	ECS_PROC_ARGS_COUNT = "args_count"
	ECS_PROC_ARGS       = "args"
	ECS_PROC_CMDLINE    = "command_line"
	ECS_PROC_EXE        = "executable"
	ECS_PROC_NAME       = "name"
	ECS_PROC_PARENT     = "parent"
	ECS_PROC_PID        = "pid"
	ECS_PROC_THREAD     = "thread"
	ECS_PROC_TID        = "id"
	ECS_PROC_START      = "start"

	ECS_SF_FA_RBYTES = "bytes_read"
	ECS_SF_FA_ROPS   = "read_ops"
	ECS_SF_FA_WBYTES = "bytes_written"
	ECS_SF_FA_WOPS   = "write_ops"

	ECS_SERVICE_ID          = "id"
	ECS_SERVICE_NAME        = "name"
	ECS_SERVICE_NAMESPACE   = "namespace"
	ECS_SERVICE_CLUSTERIP   = "clusterip"
	ECS_SERVICE_PORT        = "port"
	ECS_SERVICE_TARGETPORT  = "targetport"
	ECS_SERVICE_NODEPORT    = "nodeport"
	ECS_SERVICE_PROTO       = "proto"

	ECS_USER_ID   = "id"
	ECS_USER_NAME = "name"

	ECS_THREAT_FRAMEWORK    = "framework"
	ECS_THREAT_TECHNIQUE_ID = "id"

	ECS_TAGS = "tags"
)

// ECS kind values.
const (
	ECS_KIND_ALERT = "alert"
	ECS_KIND_EVENT = "event"
)

// ECS category values.
const (
	ECS_CAT_DIR     = "directory"
	ECS_CAT_FILE    = "file"
	ECS_CAT_NETWORK = "network"
	ECS_CAT_PROCESS = "process"
	ECS_CAT_ORCH    = "orchestration"
)

// ECS type values.
const (
	ECS_TYPE_ACCESS     = "access"
	ECS_TYPE_CHANGE     = "change"
	ECS_TYPE_CONNECTION = "connection"
	ECS_TYPE_CREATE     = "creation"
	ECS_TYPE_DELETE     = "deletion"
	ECS_TYPE_START      = "start"
	ECS_TYPE_EXIT       = "exit"
	ECS_TYPE_TSTART     = "thread-start"
	ECS_TYPE_TEXIT      = "thread-exit"
	ECS_TYPE_K8S        = "k8s"
)

// ECS action suffixes that differ from ECS types.
// Action values are typically <catogory>-<type> or <catogory>-<action>
const (
	ECS_ACTION_READ    = "read"
	ECS_ACTION_WRITE   = "write"
	ECS_ACTION_CREATE  = "create"
	ECS_ACTION_DELETE  = "delete"
	ECS_ACTION_LINK    = "link"
	ECS_ACTION_RENAME  = "rename"
	ECS_ACTION_TRAFFIC = "connection-traffic"
)
