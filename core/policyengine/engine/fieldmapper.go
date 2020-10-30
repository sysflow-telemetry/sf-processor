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
package engine

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/cespare/xxhash"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

// FieldMap is a functional type denoting a SysFlow attribute mapper.
type FieldMap func(r *Record) interface{}

// integer representation of a mapping type.
type MappingType uint8

const (
	MapIntVal      MappingType = 0
	MapStrVal      MappingType = 1
	MapBoolVal     MappingType = 2
	MapArrayStr    MappingType = 3
	MapArrayInt    MappingType = 4
	MapSpecialInt  MappingType = 5
	MapSpecialStr  MappingType = 6
	MapSpecialBool MappingType = 7
)

type SectionType uint8

const (
	SectNone  SectionType = 0
	SectProc  SectionType = 1
	SectPProc SectionType = 2
	SectFile  SectionType = 3
	SectNet   SectionType = 4
	SectFlow  SectionType = 5
	SectCont  SectionType = 6
	SectNode  SectionType = 7
)

const (
	A_IDS      sfgo.Attribute = (2 << 30) - 1
	PARENT_IDS sfgo.Attribute = (2 << 30) - 2
)

// FieldEntry is an object that stores metadata for each field in the exported map.
type FieldEntry struct {
	Map     FieldMap
	Id      sfgo.Attribute
	Type    MappingType
	Source  sfgo.Source
	Section SectionType
	AuxAttr RecAttribute
}

// IntFieldMap is a functional type denoting a numerical attribute mapper.
type IntFieldMap func(r *Record) int64

// StrFieldMap is a functional type denoting a string attribute mapper.
type StrFieldMap func(r *Record) string

// VoidFieldMap is a functional type denoting a void attribute mapper.
type VoidFieldMap func(r *Record)

// FieldMapper is an adapter for SysFlow attribute mappers.
type FieldMapper struct {
	Mappers map[string]*FieldEntry
}

// Map retrieves a field map based on a SysFlow attribute.
func (m FieldMapper) Map(attr string) FieldMap {
	if mapper, ok := m.Mappers[attr]; ok {
		return mapper.Map
	}
	return func(r *Record) interface{} { return attr }
}

// MapInt retrieves a numerical field map based on a SysFlow attribute.
func (m FieldMapper) MapInt(attr string) IntFieldMap {
	return func(r *Record) int64 {
		if v, ok := m.Map(attr)(r).(int64); ok {
			return v
		} else if v, err := strconv.ParseInt(attr, 10, 64); err == nil {
			return v
		}
		return sfgo.Zeros.Int64
	}
}

// MapStr retrieves a string field map based on a SysFlow attribute.
func (m FieldMapper) MapStr(attr string) StrFieldMap {
	return func(r *Record) string {
		if v, ok := m.Map(attr)(r).(string); ok {
			return trimBoundingQuotes(v)
		} else if v, ok := m.Map(attr)(r).(int64); ok {
			return strconv.FormatInt(v, 10)
		} else if v, ok := m.Map(attr)(r).(bool); ok {
			return strconv.FormatBool(v)
		}
		return sfgo.Zeros.String
	}
}

// Fields defines a sorted array of all exported field mapper keys.
var Fields = getFields()

// Mapper defines a global attribute mapper instance.
var Mapper = FieldMapper{getMappers()}

// getFields returns a sorted array of all exported field mapper keys.
func getFields() []string {
	mappers := getExportedMappers()
	keys := make([]string, 0, len(mappers))
	for k := range mappers {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i int, j int) bool {
		ki := len(strings.Split(keys[i], "."))
		kj := len(strings.Split(keys[j], "."))
		if ki == kj {
			return strings.Compare(keys[i], keys[j]) < 0
		}
		return ki < kj
	})
	return keys
}

type FieldValue struct {
	FieldName  string
	FieldSects []string
	Entry      *FieldEntry
}

var FieldValues = getFieldsAndValues()

// getFields returns a sorted array of all exported field mapper keys.
func getFieldsAndValues() []*FieldValue {
	mappers := getExportedMappers()
	fields := make([]*FieldValue, 0, len(mappers))
	for k, v := range mappers {
		field := &FieldValue{FieldName: k,
			FieldSects: strings.Split(k, "."),
			Entry:      v}

		fields = append(fields, field)
	}
	sort.SliceStable(fields, func(i int, j int) bool {
		ki := len(fields[i].FieldSects)
		kj := len(fields[j].FieldSects)
		if ki == kj {
			return strings.Compare(fields[i].FieldName, fields[j].FieldName) < 0
		}
		return ki < kj
	})
	return fields
}

func getMappers() map[string]*FieldEntry {
	mappers := getExportedMappers()
	for k, v := range getNonExportedMappers() {
		if _, ok := mappers[k]; !ok {
			mappers[k] = v
		} else if ok {
			logger.Warn.Println("Duplicate mapper key: ", k)
		}
	}
	return mappers
}

// getExportedMappers defines all mappers for exported attributes.
func getExportedMappers() map[string]*FieldEntry {
	return map[string]*FieldEntry{
		// SysFlow
		SF_TYPE:                  &FieldEntry{Map: mapRecType(sfgo.SYSFLOW_SRC), Id: sfgo.SF_REC_TYPE, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC},
		SF_OPFLAGS:               &FieldEntry{Map: mapOpFlags(sfgo.SYSFLOW_SRC), Id: sfgo.EV_PROC_OPFLAGS_INT, Type: MapArrayStr, Source: sfgo.SYSFLOW_SRC},
		SF_RET:                   &FieldEntry{Map: mapRet(sfgo.SYSFLOW_SRC), Id: sfgo.SF_REC_TYPE, Type: MapSpecialInt, Source: sfgo.SYSFLOW_SRC},
		SF_TS:                    &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.TS_INT), Id: sfgo.TS_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC},
		SF_ENDTS:                 &FieldEntry{Map: mapEndTs(sfgo.SYSFLOW_SRC), Id: sfgo.FL_FILE_ENDTS_INT, Type: MapSpecialInt, Source: sfgo.SYSFLOW_SRC},
		SF_PROC_OID:              &FieldEntry{Map: mapOID(sfgo.SYSFLOW_SRC, sfgo.PROC_OID_HPID_INT, sfgo.PROC_OID_CREATETS_INT), Id: sfgo.PROC_OID_HPID_INT, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_PID:              &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_OID_HPID_INT), Id: sfgo.PROC_OID_HPID_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_NAME:             &FieldEntry{Map: mapName(sfgo.SYSFLOW_SRC, sfgo.PROC_EXE_STR), Id: sfgo.PROC_EXE_STR, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_EXE:              &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.PROC_EXE_STR), Id: sfgo.PROC_EXE_STR, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_ARGS:             &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.PROC_EXEARGS_STR), Id: sfgo.PROC_EXEARGS_STR, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_UID:              &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_UID_INT), Id: sfgo.PROC_UID_INT, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_USER:             &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.PROC_USERNAME_STR), Id: sfgo.PROC_USERNAME_STR, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_TID:              &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.TID_INT), Id: sfgo.TID_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_GID:              &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_GID_INT), Id: sfgo.PROC_GID_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_GROUP:            &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.PROC_GROUPNAME_STR), Id: sfgo.PROC_GROUPNAME_STR, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_CREATETS:         &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_OID_CREATETS_INT), Id: sfgo.PROC_OID_CREATETS_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_TTY:              &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_TTY_INT), Id: sfgo.PROC_TTY_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_ENTRY:            &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_ENTRY_INT), Id: sfgo.PROC_ENTRY_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_CMDLINE:          &FieldEntry{Map: mapJoin(sfgo.SYSFLOW_SRC, sfgo.PROC_EXE_STR, sfgo.PROC_EXEARGS_STR), Id: sfgo.PROC_EXE_STR, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectProc},
		SF_PROC_ANAME:            &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, ProcAName), Id: A_IDS, Type: MapArrayStr, Source: sfgo.SYSFLOW_SRC, Section: SectProc, AuxAttr: ProcAName},
		SF_PROC_AEXE:             &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, ProcAExe), Id: A_IDS, Type: MapArrayStr, Source: sfgo.SYSFLOW_SRC, Section: SectProc, AuxAttr: ProcAExe},
		SF_PROC_ACMDLINE:         &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, ProcACmdLine), Id: A_IDS, Type: MapArrayStr, Source: sfgo.SYSFLOW_SRC, Section: SectProc, AuxAttr: ProcACmdLine},
		SF_PROC_APID:             &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, ProcAPID), Id: A_IDS, Type: MapArrayInt, Source: sfgo.SYSFLOW_SRC, Section: SectProc, AuxAttr: ProcAPID},
		SF_PPROC_OID:             &FieldEntry{Map: mapOID(sfgo.SYSFLOW_SRC, sfgo.PROC_POID_HPID_INT, sfgo.PROC_POID_CREATETS_INT), Id: sfgo.PROC_POID_HPID_INT, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectPProc},
		SF_PPROC_PID:             &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_POID_HPID_INT), Id: sfgo.PROC_POID_HPID_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectPProc},
		SF_PPROC_NAME:            &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcName), Id: PARENT_IDS, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectPProc, AuxAttr: PProcName},
		SF_PPROC_EXE:             &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcExe), Id: PARENT_IDS, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectPProc, AuxAttr: PProcExe},
		SF_PPROC_ARGS:            &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcArgs), Id: PARENT_IDS, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectPProc, AuxAttr: PProcArgs},
		SF_PPROC_UID:             &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcUID), Id: PARENT_IDS, Type: MapSpecialInt, Source: sfgo.SYSFLOW_SRC, Section: SectPProc, AuxAttr: PProcUID},
		SF_PPROC_USER:            &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcUser), Id: PARENT_IDS, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectPProc, AuxAttr: PProcUser},
		SF_PPROC_GID:             &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcGID), Id: PARENT_IDS, Type: MapSpecialInt, Source: sfgo.SYSFLOW_SRC, Section: SectPProc, AuxAttr: PProcGID},
		SF_PPROC_GROUP:           &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcGroup), Id: PARENT_IDS, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectPProc, AuxAttr: PProcGroup},
		SF_PPROC_CREATETS:        &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_POID_CREATETS_INT), Id: sfgo.PROC_POID_CREATETS_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectPProc},
		SF_PPROC_TTY:             &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcTTY), Id: PARENT_IDS, Type: MapSpecialInt, Source: sfgo.SYSFLOW_SRC, Section: SectPProc, AuxAttr: PProcTTY},
		SF_PPROC_ENTRY:           &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcEntry), Id: PARENT_IDS, Type: MapSpecialInt, Source: sfgo.SYSFLOW_SRC, Section: SectPProc, AuxAttr: PProcEntry},
		SF_PPROC_CMDLINE:         &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcCmdLine), Id: PARENT_IDS, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectPProc, AuxAttr: PProcCmdLine},
		SF_FILE_NAME:             &FieldEntry{Map: mapName(sfgo.SYSFLOW_SRC, sfgo.FILE_PATH_STR), Id: sfgo.FILE_PATH_STR, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_PATH:             &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.FILE_PATH_STR), Id: sfgo.FILE_PATH_STR, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_CANONICALPATH:    &FieldEntry{Map: mapLinkPath(sfgo.SYSFLOW_SRC, sfgo.FILE_PATH_STR), Id: sfgo.FILE_PATH_STR, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_OID:              &FieldEntry{Map: mapOID(sfgo.SYSFLOW_SRC, sfgo.FILE_PATH_STR), Id: sfgo.FILE_PATH_STR, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_DIRECTORY:        &FieldEntry{Map: mapDir(sfgo.SYSFLOW_SRC, sfgo.FILE_PATH_STR), Id: sfgo.FILE_PATH_STR, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_NEWNAME:          &FieldEntry{Map: mapName(sfgo.SYSFLOW_SRC, sfgo.SEC_FILE_PATH_STR), Id: sfgo.SEC_FILE_PATH_STR, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_NEWPATH:          &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.SEC_FILE_PATH_STR), Id: sfgo.SEC_FILE_PATH_STR, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_NEWCANONICALPATH: &FieldEntry{Map: mapLinkPath(sfgo.SYSFLOW_SRC, sfgo.SEC_FILE_PATH_STR), Id: sfgo.SEC_FILE_PATH_STR, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_NEWOID:           &FieldEntry{Map: mapOID(sfgo.SYSFLOW_SRC, sfgo.SEC_FILE_PATH_STR), Id: sfgo.SEC_FILE_PATH_STR, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_NEWDIRECTORY:     &FieldEntry{Map: mapDir(sfgo.SYSFLOW_SRC, sfgo.SEC_FILE_PATH_STR), Id: sfgo.SEC_FILE_PATH_STR, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_TYPE:             &FieldEntry{Map: mapFileType(sfgo.SYSFLOW_SRC, sfgo.FILE_RESTYPE_INT), Id: sfgo.FILE_RESTYPE_INT, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_IS_OPEN_WRITE:    &FieldEntry{Map: mapIsOpenWrite(sfgo.SYSFLOW_SRC, sfgo.FL_FILE_OPENFLAGS_INT), Id: sfgo.FL_FILE_OPENFLAGS_INT, Type: MapSpecialBool, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_IS_OPEN_READ:     &FieldEntry{Map: mapIsOpenRead(sfgo.SYSFLOW_SRC, sfgo.FL_FILE_OPENFLAGS_INT), Id: sfgo.FL_FILE_OPENFLAGS_INT, Type: MapSpecialBool, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_FD:               &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.FL_FILE_FD_INT), Id: sfgo.FL_FILE_FD_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_FILE_OPENFLAGS:        &FieldEntry{Map: mapOpenFlags(sfgo.SYSFLOW_SRC, sfgo.FL_FILE_OPENFLAGS_INT), Id: sfgo.FL_FILE_OPENFLAGS_INT, Type: MapArrayStr, Source: sfgo.SYSFLOW_SRC, Section: SectFile},
		SF_NET_PROTO:             &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_PROTO_INT), Id: sfgo.FL_NETW_PROTO_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectNet},
		//SF_NET_PROTONAME:         &FieldEntry{Map: mapProto(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_PROTO_INT), Id: sfgo.FL_NETW_PROTO_INT, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectNet},
		SF_NET_SPORT:            &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_SPORT_INT), Id: sfgo.FL_NETW_SPORT_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectNet},
		SF_NET_DPORT:            &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_DPORT_INT), Id: sfgo.FL_NETW_DPORT_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectNet},
		SF_NET_PORT:             &FieldEntry{Map: mapPort(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_SPORT_INT, sfgo.FL_NETW_DPORT_INT), Id: sfgo.FL_NETW_SPORT_INT, Type: MapArrayStr, Source: sfgo.SYSFLOW_SRC, Section: SectNet},
		SF_NET_SIP:              &FieldEntry{Map: mapIP(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_SIP_INT), Id: sfgo.FL_NETW_SIP_INT, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectNet},
		SF_NET_DIP:              &FieldEntry{Map: mapIP(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_DIP_INT), Id: sfgo.FL_NETW_DIP_INT, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectNet},
		SF_NET_IP:               &FieldEntry{Map: mapIP(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_SIP_INT, sfgo.FL_NETW_DIP_INT), Id: sfgo.FL_NETW_SIP_INT, Type: MapArrayStr, Source: sfgo.SYSFLOW_SRC, Section: SectNet},
		SF_FLOW_RBYTES:          &FieldEntry{Map: mapSum(sfgo.SYSFLOW_SRC, sfgo.FL_FILE_NUMRRECVBYTES_INT, sfgo.FL_NETW_NUMRRECVBYTES_INT), Id: sfgo.FL_FILE_NUMRRECVBYTES_INT, Type: MapSpecialInt, Source: sfgo.SYSFLOW_SRC, Section: SectFlow},
		SF_FLOW_ROPS:            &FieldEntry{Map: mapSum(sfgo.SYSFLOW_SRC, sfgo.FL_FILE_NUMRRECVOPS_INT, sfgo.FL_NETW_NUMRRECVOPS_INT), Id: sfgo.FL_FILE_NUMRRECVOPS_INT, Type: MapSpecialInt, Source: sfgo.SYSFLOW_SRC, Section: SectFlow},
		SF_FLOW_WBYTES:          &FieldEntry{Map: mapSum(sfgo.SYSFLOW_SRC, sfgo.FL_FILE_NUMWSENDBYTES_INT, sfgo.FL_NETW_NUMWSENDBYTES_INT), Id: sfgo.FL_FILE_NUMWSENDBYTES_INT, Type: MapSpecialInt, Source: sfgo.SYSFLOW_SRC, Section: SectFlow},
		SF_FLOW_WOPS:            &FieldEntry{Map: mapSum(sfgo.SYSFLOW_SRC, sfgo.FL_FILE_NUMWSENDOPS_INT, sfgo.FL_NETW_NUMWSENDOPS_INT), Id: sfgo.FL_FILE_NUMWSENDOPS_INT, Type: MapSpecialInt, Source: sfgo.SYSFLOW_SRC, Section: SectFlow},
		SF_CONTAINER_ID:         &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.CONT_ID_STR), Id: sfgo.CONT_ID_STR, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectCont},
		SF_CONTAINER_NAME:       &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.CONT_NAME_STR), Id: sfgo.CONT_NAME_STR, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectCont},
		SF_CONTAINER_IMAGEID:    &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.CONT_IMAGEID_STR), Id: sfgo.CONT_IMAGEID_STR, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectCont},
		SF_CONTAINER_IMAGE:      &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.CONT_IMAGE_STR), Id: sfgo.CONT_IMAGE_STR, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectCont},
		SF_CONTAINER_TYPE:       &FieldEntry{Map: mapContType(sfgo.SYSFLOW_SRC, sfgo.CONT_TYPE_INT), Id: sfgo.CONT_TYPE_INT, Type: MapSpecialStr, Source: sfgo.SYSFLOW_SRC, Section: SectCont},
		SF_CONTAINER_PRIVILEGED: &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.CONT_PRIVILEGED_INT), Id: sfgo.CONT_PRIVILEGED_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectCont},
		SF_NODE_ID:              &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.SFHE_EXPORTER_STR), Id: sfgo.SFHE_EXPORTER_STR, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectNode},
		SF_NODE_IP:              &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.SFHE_IP_STR), Id: sfgo.SFHE_IP_STR, Type: MapStrVal, Source: sfgo.SYSFLOW_SRC, Section: SectNode},
		SF_SCHEMA_VERSION:       &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.SFHE_VERSION_INT), Id: sfgo.SFHE_VERSION_INT, Type: MapIntVal, Source: sfgo.SYSFLOW_SRC, Section: SectNode},

		/*		//Ext processes
				EXT_PROC_GUID_STR:                mapStr(sfgo.PROCESS_SRC, sfgo.PROC_GUID_STR),
				EXT_PROC_IMAGE_STR:               mapStr(sfgo.PROCESS_SRC, sfgo.PROC_IMAGE_STR),
				EXT_PROC_CURR_DIRECTORY_STR:      mapDir(sfgo.PROCESS_SRC, sfgo.PROC_CURR_DIRECTORY_STR),
				EXT_PROC_LOGON_GUID_STR:          mapStr(sfgo.PROCESS_SRC, sfgo.PROC_LOGON_GUID_STR),
				EXT_PROC_LOGON_ID_STR:            mapStr(sfgo.PROCESS_SRC, sfgo.PROC_LOGON_ID_STR),
				EXT_PROC_TERMINAL_SESSION_ID_STR: mapStr(sfgo.PROCESS_SRC, sfgo.PROC_TERMINAL_SESSION_ID_STR),
				EXT_PROC_INTEGRITY_LEVEL_STR:     mapStr(sfgo.PROCESS_SRC, sfgo.PROC_INTEGRITY_LEVEL_STR),
				EXT_PROC_SIGNATURE_STR:           mapStr(sfgo.PROCESS_SRC, sfgo.PROC_SIGNATURE_STR),
				EXT_PROC_SIGNATURE_STATUS_STR:    mapStr(sfgo.PROCESS_SRC, sfgo.PROC_SIGNATURE_STATUS_STR),
				EXT_PROC_SHA1_HASH_STR:           mapStr(sfgo.PROCESS_SRC, sfgo.PROC_SHA1_HASH_STR),
				EXT_PROC_MD5_HASH_STR:            mapStr(sfgo.PROCESS_SRC, sfgo.PROC_MD5_HASH_STR),
				EXT_PROC_SHA256_HASH_STR:         mapStr(sfgo.PROCESS_SRC, sfgo.PROC_SHA256_HASH_STR),
				EXT_PROC_IMP_HASH_STR:            mapStr(sfgo.PROCESS_SRC, sfgo.PROC_IMP_HASH_STR),
				EXT_PROC_SIGNED_INT:              mapInt(sfgo.PROCESS_SRC, sfgo.PROC_SIGNED_INT),

				//Ext files
				EXT_FILE_SIGNATURE_STR:        mapStr(sfgo.FILE_SRC, sfgo.FILE_SIGNATURE_STR),
				EXT_FILE_SIGNATURE_STATUS_STR: mapStr(sfgo.FILE_SRC, sfgo.FILE_SIGNATURE_STATUS_STR),
				EXT_FILE_SHA1_HASH_STR:        mapStr(sfgo.FILE_SRC, sfgo.FILE_SHA1_HASH_STR),
				EXT_FILE_MD5_HASH_STR:         mapStr(sfgo.FILE_SRC, sfgo.FILE_MD5_HASH_STR),
				EXT_FILE_SHA256_HASH_STR:      mapStr(sfgo.FILE_SRC, sfgo.FILE_SHA256_HASH_STR),
				EXT_FILE_IMP_HASH_STR:         mapStr(sfgo.FILE_SRC, sfgo.FILE_IMP_HASH_STR),
				EXT_FILE_SIGNED_INT:           mapInt(sfgo.FILE_SRC, sfgo.FILE_SIGNED_INT),

				//Ext network
				EXT_NET_SOURCE_HOST_NAME_STR: mapStr(sfgo.NETWORK_SRC, sfgo.NET_SOURCE_HOST_NAME_STR),
				EXT_NET_SOURCE_PORT_NAME_STR: mapStr(sfgo.NETWORK_SRC, sfgo.NET_SOURCE_PORT_NAME_STR),
				EXT_NET_DEST_HOST_NAME_STR:   mapStr(sfgo.NETWORK_SRC, sfgo.NET_DEST_HOST_NAME_STR),
				EXT_NET_DEST_PORT_NAME_STR:   mapStr(sfgo.NETWORK_SRC, sfgo.NET_DEST_PORT_NAME_STR),

				//Ext target proc
				EXT_TARG_PROC_OID_CREATETS_INT:       mapInt(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_OID_CREATETS_INT),
				EXT_TARG_PROC_OID_HPID_INT:           mapInt(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_OID_HPID_INT),
				EXT_TARG_PROC_TS_INT:                 mapInt(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_TS_INT),
				EXT_TARG_PROC_POID_CREATETS_INT:      mapInt(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_POID_CREATETS_INT),
				EXT_TARG_PROC_POID_HPID_INT:          mapInt(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_POID_HPID_INT),
				EXT_TARG_PROC_EXE_STR:                mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_EXE_STR),
				EXT_TARG_PROC_EXEARGS_STR:            mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_EXEARGS_STR),
				EXT_TARG_PROC_UID_INT:                mapInt(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_UID_INT),
				EXT_TARG_PROC_GID_INT:                mapInt(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_GID_INT),
				EXT_TARG_PROC_USERNAME_STR:           mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_USERNAME_STR),
				EXT_TARG_PROC_GROUPNAME_STR:          mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_GROUPNAME_STR),
				EXT_TARG_PROC_TTY_INT:                mapInt(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_TTY_INT),
				EXT_TARG_PROC_CONTAINERID_STRING_STR: mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_CONTAINERID_STRING_STR),
				EXT_TARG_PROC_ENTRY_INT:              mapEntry(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_ENTRY_INT),

				EXT_TARG_PROC_GUID_STR:                mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_GUID_STR),
				EXT_TARG_PROC_IMAGE_STR:               mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_IMAGE_STR),
				EXT_TARG_PROC_CURR_DIRECTORY_STR:      mapDir(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_CURR_DIRECTORY_STR),
				EXT_TARG_PROC_LOGON_GUID_STR:          mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_LOGON_GUID_STR),
				EXT_TARG_PROC_LOGON_ID_STR:            mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_LOGON_ID_STR),
				EXT_TARG_PROC_TERMINAL_SESSION_ID_STR: mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_TERMINAL_SESSION_ID_STR),
				EXT_TARG_PROC_INTEGRITY_LEVEL_STR:     mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_INTEGRITY_LEVEL_STR),
				EXT_TARG_PROC_SIGNATURE_STR:           mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_SIGNATURE_STR),
				EXT_TARG_PROC_SIGNATURE_STATUS_STR:    mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_SIGNATURE_STATUS_STR),
				EXT_TARG_PROC_SHA1_HASH_STR:           mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_SHA1_HASH_STR),
				EXT_TARG_PROC_MD5_HASH_STR:            mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_MD5_HASH_STR),
				EXT_TARG_PROC_SHA256_HASH_STR:         mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_SHA256_HASH_STR),
				EXT_TARG_PROC_IMP_HASH_STR:            mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_IMP_HASH_STR),
				EXT_TARG_PROC_SIGNED_INT:              mapInt(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_SIGNED_INT),
				EXT_TARG_PROC_START_ADDR_STR:          mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_START_ADDR_STR),
				EXT_TARG_PROC_START_MODULE_STR:        mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_START_MODULE_STR),
				EXT_TARG_PROC_START_FUNCTION_STR:      mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_START_FUNCTION_STR),
				EXT_TARG_PROC_GRANT_ACCESS_STR:        mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_GRANT_ACCESS_STR),
				EXT_TARG_PROC_CALL_TRACE_STR:          mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_CALL_TRACE_STR),
				EXT_TARG_PROC_ACCESS_TYPE_STR:         mapStr(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_ACCESS_TYPE_STR),
				EXT_TARG_PROC_NEW_THREAD_ID_INT:       mapInt(sfgo.TARG_PROC_SRC, sfgo.EVT_TARG_PROC_NEW_THREAD_ID_INT),
		*/
	}
}

// getNonExportedMappers defines all mappers for non-exported (query-only) attributes.
func getNonExportedMappers() map[string]*FieldEntry {
	return map[string]*FieldEntry{
		// Falco
		FALCO_EVT_TYPE:          &FieldEntry{Map: mapOpFlags(sfgo.SYSFLOW_SRC)},
		FALCO_EVT_RAW_RES:       &FieldEntry{Map: mapRecType(sfgo.SYSFLOW_SRC)},
		FALCO_EVT_RAW_TIME:      &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.TS_INT)},
		FALCO_EVT_DIR:           &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_OID_HPID_INT)},
		FALCO_EVT_IS_OPEN_READ:  &FieldEntry{Map: mapIsOpenRead(sfgo.SYSFLOW_SRC, sfgo.FL_FILE_OPENFLAGS_INT)},
		FALCO_EVT_IS_OPEN_WRITE: &FieldEntry{Map: mapIsOpenWrite(sfgo.SYSFLOW_SRC, sfgo.FL_FILE_OPENFLAGS_INT)},
		FALCO_EVT_UID:           &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_UID_INT)},
		FALCO_FD_TYPECHAR:       &FieldEntry{Map: mapFileType(sfgo.SYSFLOW_SRC, sfgo.FILE_RESTYPE_INT)},
		FALCO_FD_DIRECTORY:      &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_OID_HPID_INT)},
		FALCO_FD_NAME:           &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_OID_HPID_INT)},
		FALCO_FD_FILENAME:       &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_OID_HPID_INT)},
		FALCO_FD_PROTO:          &FieldEntry{Map: mapDir(sfgo.SYSFLOW_SRC, sfgo.FILE_PATH_STR)},
		FALCO_FD_LPROTO:         &FieldEntry{Map: mapDir(sfgo.SYSFLOW_SRC, sfgo.FILE_PATH_STR)},
		FALCO_FD_L4PROTO:        &FieldEntry{Map: mapName(sfgo.SYSFLOW_SRC, sfgo.FILE_PATH_STR)},
		FALCO_FD_RPROTO:         &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_PROTO_INT)},
		FALCO_FD_SPROTO:         &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_PROTO_INT)},
		FALCO_FD_CPROTO:         &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_PROTO_INT)},
		FALCO_FD_SPORT:          &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_SPORT_INT)},
		FALCO_FD_DPORT:          &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_DPORT_INT)},
		FALCO_FD_SIP:            &FieldEntry{Map: mapIP(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_SIP_INT)},
		FALCO_FD_DIP:            &FieldEntry{Map: mapIP(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_DIP_INT)},
		FALCO_FD_IP:             &FieldEntry{Map: mapIP(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_SIP_INT, sfgo.FL_NETW_DIP_INT)},
		FALCO_FD_PORT:           &FieldEntry{Map: mapPort(sfgo.SYSFLOW_SRC, sfgo.FL_NETW_SPORT_INT, sfgo.FL_NETW_DPORT_INT)},
		FALCO_FD_NUM:            &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.FL_FILE_FD_INT)},
		FALCO_USER_NAME:         &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.PROC_USERNAME_STR)},
		FALCO_PROC_PID:          &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_OID_HPID_INT)},
		FALCO_PROC_TID:          &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.TID_INT)},
		FALCO_PROC_GID:          &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_GID_INT)},
		FALCO_PROC_UID:          &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_UID_INT)},
		FALCO_PROC_GROUP:        &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.PROC_GROUPNAME_STR)},
		FALCO_PROC_TTY:          &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcTTY)},
		FALCO_PROC_USER:         &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.PROC_USERNAME_STR)},
		FALCO_PROC_EXE:          &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.PROC_EXE_STR)},
		FALCO_PROC_NAME:         &FieldEntry{Map: mapName(sfgo.SYSFLOW_SRC, sfgo.PROC_EXE_STR)},
		FALCO_PROC_ARGS:         &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.PROC_EXEARGS_STR)},
		FALCO_PROC_CREATE_TIME:  &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_POID_CREATETS_INT)},
		FALCO_PROC_CMDLINE:      &FieldEntry{Map: mapJoin(sfgo.SYSFLOW_SRC, sfgo.PROC_EXE_STR, sfgo.PROC_EXEARGS_STR)},
		FALCO_PROC_ANAME:        &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, ProcAName)},
		FALCO_PROC_APID:         &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, ProcAPID)},
		FALCO_PROC_PPID:         &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_POID_HPID_INT)},
		FALCO_PROC_PGID:         &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcGID)},
		FALCO_PROC_PUID:         &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcUID)},
		FALCO_PROC_PGROUP:       &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcGroup)},
		FALCO_PROC_PTTY:         &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcTTY)},
		FALCO_PROC_PUSER:        &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcUser)},
		FALCO_PROC_PEXE:         &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcExe)},
		FALCO_PROC_PARGS:        &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcArgs)},
		FALCO_PROC_PCREATE_TIME: &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.PROC_POID_CREATETS_INT)},
		FALCO_PROC_PNAME:        &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcName)},
		FALCO_PROC_PCMDLINE:     &FieldEntry{Map: mapCachedValue(sfgo.SYSFLOW_SRC, PProcCmdLine)},
		FALCO_CONT_ID:           &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.CONT_ID_STR)},
		FALCO_CONT_IMAGE_ID:     &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.CONT_IMAGEID_STR)},
		FALCO_CONT_IMAGE:        &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.CONT_IMAGE_STR)},
		FALCO_CONT_NAME:         &FieldEntry{Map: mapStr(sfgo.SYSFLOW_SRC, sfgo.CONT_NAME_STR)},
		FALCO_CONT_TYPE:         &FieldEntry{Map: mapContType(sfgo.SYSFLOW_SRC, sfgo.CONT_TYPE_INT)},
		FALCO_CONT_PRIVILEGED:   &FieldEntry{Map: mapInt(sfgo.SYSFLOW_SRC, sfgo.CONT_PRIVILEGED_INT)},
	}
}

func mapStr(src sfgo.Source, attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} { return r.GetStr(attr, src) }
}

func mapInt(src sfgo.Source, attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} { return r.GetInt(attr, src) }
}

func mapSum(src sfgo.Source, attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var sum int64 = 0
		for _, attr := range attrs {
			sum += r.GetInt(attr, src)
		}
		return sum
	}
}

func mapJoin(src sfgo.Source, attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var join string = r.GetStr(attrs[0], src)
		for _, attr := range attrs[1:] {
			join += SPACE + r.GetStr(attr, src)
		}
		return join
	}
}

func mapRecType(src sfgo.Source) FieldMap {
	return func(r *Record) interface{} {
		return GetRecType(r, src)
	}
}

func mapOpFlags(src sfgo.Source) FieldMap {
	return func(r *Record) interface{} {
		opflags := r.GetInt(sfgo.EV_PROC_OPFLAGS_INT, src)
		rtype := mapRecType(src)(r).(string)
		return strings.Join(sfgo.GetOpFlags(int32(opflags), rtype), LISTSEP)
	}
}

func mapRet(src sfgo.Source) FieldMap {
	return func(r *Record) interface{} {
		switch r.GetInt(sfgo.SF_REC_TYPE, src) {
		case sfgo.PROC_EVT:
			fallthrough
		case sfgo.FILE_EVT:
			return r.GetInt(sfgo.RET_INT, src)
		default:
			return sfgo.Zeros.Int64
		}
	}
}

func mapEndTs(src sfgo.Source) FieldMap {
	return func(r *Record) interface{} {
		switch r.GetInt(sfgo.SF_REC_TYPE, src) {
		case sfgo.FILE_FLOW:
			return r.GetInt(sfgo.FL_FILE_ENDTS_INT, src)
		case sfgo.NET_FLOW:
			return r.GetInt(sfgo.FL_NETW_ENDTS_INT, src)
		default:
			return sfgo.Zeros.Int64
		}
	}
}

func mapEntry(src sfgo.Source, attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		if r.GetInt(attr, src) == 1 {
			return true
		}
		return false
	}
}

func mapName(src sfgo.Source, attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return filepath.Base(r.GetStr(attr, src))
	}
}

func mapDir(src sfgo.Source, attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return filepath.Dir(r.GetStr(attr, src))
	}
}

func mapLinkPath(src sfgo.Source, attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		orig := r.GetStr(attr, src)
		// Possible format: aabbccddeeff0011->aabbccddeeff0011 /path/to/target.file
		var src, dst uint64
		var targetPath string
		if _, err := fmt.Sscanf(orig, "%x->%x %s", &src, &dst, &targetPath); nil == err {
			return targetPath
		}
		return orig
	}
}

func mapFileType(src sfgo.Source, attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return sfgo.GetFileType(r.GetInt(attr, src))
	}
}

func mapIsOpenWrite(src sfgo.Source, attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		if sfgo.IsOpenWrite(r.GetInt(attr, src)) {
			return true
		}
		return false
	}
}

func mapIsOpenRead(src sfgo.Source, attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		if sfgo.IsOpenRead(r.GetInt(attr, src)) {
			return true
		}
		return false
	}
}

func mapOpenFlags(src sfgo.Source, attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return strings.Join(sfgo.GetOpenFlags(r.GetInt(attr, src)), LISTSEP)
	}
}

func mapProto(src sfgo.Source, attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return sfgo.GetProto(r.GetInt(attr, src))
	}
}

func mapPort(src sfgo.Source, attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var ports = make([]string, 0)
		for _, attr := range attrs {
			ports = append(ports, strconv.FormatInt(r.GetInt(attr, src), 10))
		}
		// logger.Info.Println(ports)
		return strings.Join(ports, LISTSEP)
	}
}

func mapIP(src sfgo.Source, attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var ips = make([]string, 0)
		for _, attr := range attrs {
			ips = append(ips, sfgo.GetIPStr(int32(r.GetInt(attr, src))))
		}
		// logger.Info.Println(ips)
		return strings.Join(ips, LISTSEP)
	}
}

func mapContType(src sfgo.Source, attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return sfgo.GetContType(r.GetInt(attr, src))
	}
}

func mapCachedValue(src sfgo.Source, attr RecAttribute) FieldMap {
	return func(r *Record) interface{} {
		oid := sfgo.OID{CreateTS: r.GetInt(sfgo.PROC_OID_CREATETS_INT, src), Hpid: r.GetInt(sfgo.PROC_OID_HPID_INT, src)}
		return r.GetCachedValue(oid, attr)
	}
}

func mapOID(src sfgo.Source, attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		h := xxhash.New()
		for _, attr := range attrs {
			h.Write([]byte(fmt.Sprintf("%v", r.GetInt(attr, src))))
		}
		return fmt.Sprintf("%x", h.Sum(nil))
	}
}

func mapNa(attr string) FieldMap {
	return func(r *Record) interface{} {
		logger.Warn.Println("Attribute not supported ", attr)
		return sfgo.Zeros.String
	}
}
