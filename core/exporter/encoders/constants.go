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

// Package encoders implements codecs for exporting records and events in different data formats.
package encoders

// SysFlow record components
const (
	PROC      = "proc"
	PPROC     = "pproc"
	NET       = "net"
	FILEF     = "file"
	FLOW      = "flow"
	CONTAINER = "container"
	POD       = "pod"
	SERVICE   = "service"
	KE        = "k8s"
	NODE      = "node"
	META      = "meta"

	BEGIN_STATE = iota
	PROC_STATE
	PPROC_STATE
	NET_STATE
	FILE_STATE
	FLOW_STATE
	CONT_STATE
	POD_STATE
	SVC_STATE
	KE_STATE
	NODE_STATE
	META_STATE
)

// Export schema shared attribute names.
const (
	VERSION_ATTR      = "version"
	GROUP_ID_ATTR     = "groupId"
	OBSERVATIONS_ATTR = "observations"
	POLICIES_ATTR     = "policies"
	ID_TAG_ATTR       = "id"
	DESC_ATTR         = "desc"
	PRIORITY_ATTR     = "priority"
	TAGS_ATTR         = "tags"
)
