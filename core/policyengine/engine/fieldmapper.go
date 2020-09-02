//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
package engine

import (
	"crypto/sha256"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-apis/go/utils"
	"github.ibm.com/sysflow/goutils/logger"
	"github.ibm.com/sysflow/sf-processor/core/flattener"
)

// FieldMap is a functional type denoting a SysFlow attribute mapper.
type FieldMap func(r *Record) interface{}

// IntFieldMap is a functional type denoting a numerical attribute mapper.
type IntFieldMap func(r *Record) int64

// StrFieldMap is a functional type denoting a string attribute mapper.
type StrFieldMap func(r *Record) string

// FieldMapper is an adapter for SysFlow attribute mappers.
type FieldMapper struct {
	Mappers map[string]FieldMap
}

// Map retrieves a field map based on a SysFlow attribute.
func (m FieldMapper) Map(attr string) FieldMap {
	if mapper, ok := m.Mappers[attr]; ok {
		return mapper
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
			return m.trimBoundingQuotes(v)
		} else if v, ok := m.Map(attr)(r).(int64); ok {
			return strconv.FormatInt(v, 10)
		} else if v, ok := m.Map(attr)(r).(bool); ok {
			return strconv.FormatBool(v)
		}
		return sfgo.Zeros.String
	}
}

func (m FieldMapper) trimBoundingQuotes(s string) string {
	if len(s) > 0 && (s[0] == '"' || s[0] == '\'') {
		s = s[1:]
	}
	if len(s) > 0 && (s[len(s)-1] == '"' || s[len(s)-1] == '\'') {
		s = s[:len(s)-1]
	}
	return s
}

func getFields() []string {
	keys := make([]string, 0, len(Mapper.Mappers))
	for k := range Mapper.Mappers {
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

// Fields defines a sorted array of all field mapper keys.
var Fields = getFields()

// Mapper defines a global attribute mapper instance.
var Mapper = FieldMapper{
	map[string]FieldMap{
		SF_TYPE:                 mapRecType(flattener.SYSFLOW_SRC),
		SF_OPFLAGS:              mapOpFlags(flattener.SYSFLOW_SRC),
		SF_RET:                  mapRet(flattener.SYSFLOW_SRC),
		SF_TS:                   mapInt(sfgo.TS_INT, flattener.SYSFLOW_SRC),
		SF_ENDTS:                mapEndTs(flattener.SYSFLOW_SRC),
		SF_PROC_OID:             mapOID(flattener.SYSFLOW_SRC, sfgo.PROC_OID_HPID_INT, sfgo.PROC_OID_CREATETS_INT),
		SF_PROC_PID:             mapInt(sfgo.PROC_OID_HPID_INT, flattener.SYSFLOW_SRC),
		SF_PROC_NAME:            mapName(sfgo.PROC_EXE_STR, flattener.SYSFLOW_SRC),
		SF_PROC_EXE:             mapStr(sfgo.PROC_EXE_STR, flattener.SYSFLOW_SRC),
		SF_PROC_ARGS:            mapStr(sfgo.PROC_EXEARGS_STR, flattener.SYSFLOW_SRC),
		SF_PROC_UID:             mapInt(sfgo.PROC_UID_INT, flattener.SYSFLOW_SRC),
		SF_PROC_USER:            mapStr(sfgo.PROC_USERNAME_STR, flattener.SYSFLOW_SRC),
		SF_PROC_TID:             mapInt(sfgo.TID_INT, flattener.SYSFLOW_SRC),
		SF_PROC_GID:             mapInt(sfgo.PROC_GID_INT, flattener.SYSFLOW_SRC),
		SF_PROC_GROUP:           mapStr(sfgo.PROC_GROUPNAME_STR, flattener.SYSFLOW_SRC),
		SF_PROC_CREATETS:        mapInt(sfgo.PROC_OID_CREATETS_INT, flattener.SYSFLOW_SRC),
		SF_PROC_TTY:             mapInt(sfgo.PROC_TTY_INT, flattener.SYSFLOW_SRC),
		SF_PROC_ENTRY:           mapEntry(sfgo.PROC_ENTRY_INT, flattener.SYSFLOW_SRC),
		SF_PROC_CMDLINE:         mapJoin(flattener.SYSFLOW_SRC, sfgo.PROC_EXE_STR, sfgo.PROC_EXEARGS_STR),
		SF_PROC_ANAME:           mapCachedValue(ProcAName, flattener.SYSFLOW_SRC),
		SF_PROC_AEXE:            mapCachedValue(ProcAExe, flattener.SYSFLOW_SRC),
		SF_PROC_ACMDLINE:        mapCachedValue(ProcACmdLine, flattener.SYSFLOW_SRC),
		SF_PROC_APID:            mapCachedValue(ProcAPID, flattener.SYSFLOW_SRC),
		SF_PPROC_OID:            mapOID(flattener.SYSFLOW_SRC, sfgo.PROC_POID_HPID_INT, sfgo.PROC_POID_CREATETS_INT),
		SF_PPROC_PID:            mapInt(sfgo.PROC_POID_HPID_INT, flattener.SYSFLOW_SRC),
		SF_PPROC_NAME:           mapCachedValue(PProcName, flattener.SYSFLOW_SRC),
		SF_PPROC_EXE:            mapCachedValue(PProcExe, flattener.SYSFLOW_SRC),
		SF_PPROC_ARGS:           mapCachedValue(PProcArgs, flattener.SYSFLOW_SRC),
		SF_PPROC_UID:            mapCachedValue(PProcUID, flattener.SYSFLOW_SRC),
		SF_PPROC_USER:           mapCachedValue(PProcUser, flattener.SYSFLOW_SRC),
		SF_PPROC_GID:            mapCachedValue(PProcGID, flattener.SYSFLOW_SRC),
		SF_PPROC_GROUP:          mapCachedValue(PProcGroup, flattener.SYSFLOW_SRC),
		SF_PPROC_CREATETS:       mapInt(sfgo.PROC_POID_CREATETS_INT, flattener.SYSFLOW_SRC),
		SF_PPROC_TTY:            mapCachedValue(PProcTTY, flattener.SYSFLOW_SRC),
		SF_PPROC_ENTRY:          mapCachedValue(PProcEntry, flattener.SYSFLOW_SRC),
		SF_PPROC_CMDLINE:        mapCachedValue(PProcCmdLine, flattener.SYSFLOW_SRC),
		SF_FILE_NAME:            mapName(sfgo.FILE_PATH_STR, flattener.SYSFLOW_SRC),
		SF_FILE_PATH:            mapStr(sfgo.FILE_PATH_STR, flattener.SYSFLOW_SRC),
		SF_FILE_CANONICALPATH:   mapLinkPath(sfgo.FILE_PATH_STR, flattener.SYSFLOW_SRC),
		SF_FILE_DIRECTORY:       mapDir(sfgo.FILE_PATH_STR, flattener.SYSFLOW_SRC),
		SF_FILE_NEWNAME:         mapName(sfgo.SEC_FILE_PATH_STR, flattener.SYSFLOW_SRC),
		SF_FILE_NEWPATH:         mapStr(sfgo.SEC_FILE_PATH_STR, flattener.SYSFLOW_SRC),
		SF_FILE_NEWDIRECTORY:    mapDir(sfgo.SEC_FILE_PATH_STR, flattener.SYSFLOW_SRC),
		SF_FILE_TYPE:            mapFileType(sfgo.FILE_RESTYPE_INT, flattener.SYSFLOW_SRC),
		SF_FILE_IS_OPEN_WRITE:   mapIsOpenWrite(sfgo.FL_FILE_OPENFLAGS_INT, flattener.SYSFLOW_SRC),
		SF_FILE_IS_OPEN_READ:    mapIsOpenRead(sfgo.FL_FILE_OPENFLAGS_INT, flattener.SYSFLOW_SRC),
		SF_FILE_FD:              mapInt(sfgo.FL_FILE_FD_INT, flattener.SYSFLOW_SRC),
		SF_FILE_OPENFLAGS:       mapOpenFlags(sfgo.FL_FILE_OPENFLAGS_INT, flattener.SYSFLOW_SRC),
		SF_NET_PROTO:            mapInt(sfgo.FL_NETW_PROTO_INT, flattener.SYSFLOW_SRC),
		SF_NET_PROTONAME:        mapProto(sfgo.FL_NETW_PROTO_INT, flattener.SYSFLOW_SRC),
		SF_NET_SPORT:            mapInt(sfgo.FL_NETW_SPORT_INT, flattener.SYSFLOW_SRC),
		SF_NET_DPORT:            mapInt(sfgo.FL_NETW_DPORT_INT, flattener.SYSFLOW_SRC),
		SF_NET_PORT:             mapPort(flattener.SYSFLOW_SRC, sfgo.FL_NETW_SPORT_INT, sfgo.FL_NETW_DPORT_INT),
		SF_NET_SIP:              mapIP(flattener.SYSFLOW_SRC, sfgo.FL_NETW_SIP_INT),
		SF_NET_DIP:              mapIP(flattener.SYSFLOW_SRC, sfgo.FL_NETW_DIP_INT),
		SF_NET_IP:               mapIP(flattener.SYSFLOW_SRC, sfgo.FL_NETW_SIP_INT, sfgo.FL_NETW_DIP_INT),
		SF_FLOW_RBYTES:          mapSum(flattener.SYSFLOW_SRC, sfgo.FL_FILE_NUMRRECVBYTES_INT, sfgo.FL_NETW_NUMRRECVBYTES_INT),
		SF_FLOW_ROPS:            mapSum(flattener.SYSFLOW_SRC, sfgo.FL_FILE_NUMRRECVOPS_INT, sfgo.FL_NETW_NUMRRECVOPS_INT),
		SF_FLOW_WBYTES:          mapSum(flattener.SYSFLOW_SRC, sfgo.FL_FILE_NUMWSENDBYTES_INT, sfgo.FL_NETW_NUMWSENDBYTES_INT),
		SF_FLOW_WOPS:            mapSum(flattener.SYSFLOW_SRC, sfgo.FL_FILE_NUMWSENDOPS_INT, sfgo.FL_NETW_NUMWSENDOPS_INT),
		SF_CONTAINER_ID:         mapStr(sfgo.CONT_ID_STR, flattener.SYSFLOW_SRC),
		SF_CONTAINER_NAME:       mapStr(sfgo.CONT_NAME_STR, flattener.SYSFLOW_SRC),
		SF_CONTAINER_IMAGEID:    mapStr(sfgo.CONT_IMAGEID_STR, flattener.SYSFLOW_SRC),
		SF_CONTAINER_IMAGE:      mapStr(sfgo.CONT_IMAGE_STR, flattener.SYSFLOW_SRC),
		SF_CONTAINER_TYPE:       mapContType(sfgo.CONT_TYPE_INT, flattener.SYSFLOW_SRC),
		SF_CONTAINER_PRIVILEGED: mapInt(sfgo.CONT_PRIVILEGED_INT, flattener.SYSFLOW_SRC),
		SF_NODE_ID:              mapStr(sfgo.SFHE_EXPORTER_STR, flattener.SYSFLOW_SRC),
		SF_NODE_IP:              mapStr(sfgo.SFHE_IP_STR, flattener.SYSFLOW_SRC),
		SF_SCHEMA_VERSION:       mapInt(sfgo.SFHE_VERSION_INT, flattener.SYSFLOW_SRC),

		//Ext processes
		EXT_PROC_GUID_STR:                mapStr(flattener.PROC_GUID_STR, flattener.PROCESS_SRC),
		EXT_PROC_IMAGE_STR:               mapStr(flattener.PROC_IMAGE_STR, flattener.PROCESS_SRC),
		EXT_PROC_CURR_DIRECTORY_STR:      mapDir(flattener.PROC_CURR_DIRECTORY_STR, flattener.PROCESS_SRC),
		EXT_PROC_LOGON_GUID_STR:          mapStr(flattener.PROC_LOGON_GUID_STR, flattener.PROCESS_SRC),
		EXT_PROC_LOGON_ID_STR:            mapStr(flattener.PROC_LOGON_ID_STR, flattener.PROCESS_SRC),
		EXT_PROC_TERMINAL_SESSION_ID_STR: mapStr(flattener.PROC_TERMINAL_SESSION_ID_STR, flattener.PROCESS_SRC),
		EXT_PROC_INTEGRITY_LEVEL_STR:     mapStr(flattener.PROC_INTEGRITY_LEVEL_STR, flattener.PROCESS_SRC),
		EXT_PROC_SIGNATURE_STR:           mapStr(flattener.PROC_SIGNATURE_STR, flattener.PROCESS_SRC),
		EXT_PROC_SIGNATURE_STATUS_STR:    mapStr(flattener.PROC_SIGNATURE_STATUS_STR, flattener.PROCESS_SRC),
		EXT_PROC_SHA1_HASH_STR:           mapStr(flattener.PROC_SHA1_HASH_STR, flattener.PROCESS_SRC),
		EXT_PROC_MD5_HASH_STR:            mapStr(flattener.PROC_MD5_HASH_STR, flattener.PROCESS_SRC),
		EXT_PROC_SHA256_HASH_STR:         mapStr(flattener.PROC_SHA256_HASH_STR, flattener.PROCESS_SRC),
		EXT_PROC_IMP_HASH_STR:            mapStr(flattener.PROC_IMP_HASH_STR, flattener.PROCESS_SRC),
		EXT_PROC_SIGNED_INT:              mapInt(flattener.PROC_SIGNED_INT, flattener.PROCESS_SRC),

		//Ext files
		EXT_FILE_SIGNATURE_STR:        mapStr(flattener.FILE_SIGNATURE_STR, flattener.FILE_SRC),
		EXT_FILE_SIGNATURE_STATUS_STR: mapStr(flattener.FILE_SIGNATURE_STATUS_STR, flattener.FILE_SRC),
		EXT_FILE_SHA1_HASH_STR:        mapStr(flattener.FILE_SHA1_HASH_STR, flattener.FILE_SRC),
		EXT_FILE_MD5_HASH_STR:         mapStr(flattener.FILE_MD5_HASH_STR, flattener.FILE_SRC),
		EXT_FILE_SHA256_HASH_STR:      mapStr(flattener.FILE_SHA256_HASH_STR, flattener.FILE_SRC),
		EXT_FILE_IMP_HASH_STR:         mapStr(flattener.FILE_IMP_HASH_STR, flattener.FILE_SRC),
		EXT_FILE_SIGNED_INT:           mapInt(flattener.FILE_SIGNED_INT, flattener.FILE_SRC),

		//Ext network
		EXT_NET_SOURCE_HOST_NAME_STR: mapStr(flattener.NET_SOURCE_HOST_NAME_STR, flattener.NETWORK_SRC),
		EXT_NET_SOURCE_PORT_NAME_STR: mapStr(flattener.NET_SOURCE_PORT_NAME_STR, flattener.NETWORK_SRC),
		EXT_NET_DEST_HOST_NAME_STR:   mapStr(flattener.NET_DEST_HOST_NAME_STR, flattener.NETWORK_SRC),
		EXT_NET_DEST_PORT_NAME_STR:   mapStr(flattener.NET_DEST_PORT_NAME_STR, flattener.NETWORK_SRC),

		//Ext target proc

		EXT_TARG_PROC_OID_CREATETS_INT:       mapInt(flattener.EVT_TARG_PROC_OID_CREATETS_INT, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_OID_HPID_INT:           mapInt(flattener.EVT_TARG_PROC_OID_HPID_INT, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_TS_INT:                 mapInt(flattener.EVT_TARG_PROC_TS_INT, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_POID_CREATETS_INT:      mapInt(flattener.EVT_TARG_PROC_POID_CREATETS_INT, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_POID_HPID_INT:          mapInt(flattener.EVT_TARG_PROC_POID_HPID_INT, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_EXE_STR:                mapStr(flattener.EVT_TARG_PROC_EXE_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_EXEARGS_STR:            mapStr(flattener.EVT_TARG_PROC_EXEARGS_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_UID_INT:                mapInt(flattener.EVT_TARG_PROC_UID_INT, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_GID_INT:                mapInt(flattener.EVT_TARG_PROC_GID_INT, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_USERNAME_STR:           mapStr(flattener.EVT_TARG_PROC_USERNAME_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_GROUPNAME_STR:          mapStr(flattener.EVT_TARG_PROC_GROUPNAME_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_TTY_INT:                mapInt(flattener.EVT_TARG_PROC_TTY_INT, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_CONTAINERID_STRING_STR: mapStr(flattener.EVT_TARG_PROC_CONTAINERID_STRING_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_ENTRY_INT:              mapEntry(flattener.EVT_TARG_PROC_ENTRY_INT, flattener.TARG_PROC_SRC),

		EXT_TARG_PROC_GUID_STR:                mapStr(flattener.EVT_TARG_PROC_GUID_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_IMAGE_STR:               mapStr(flattener.EVT_TARG_PROC_IMAGE_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_CURR_DIRECTORY_STR:      mapDir(flattener.EVT_TARG_PROC_CURR_DIRECTORY_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_LOGON_GUID_STR:          mapStr(flattener.EVT_TARG_PROC_LOGON_GUID_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_LOGON_ID_STR:            mapStr(flattener.EVT_TARG_PROC_LOGON_ID_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_TERMINAL_SESSION_ID_STR: mapStr(flattener.EVT_TARG_PROC_TERMINAL_SESSION_ID_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_INTEGRITY_LEVEL_STR:     mapStr(flattener.EVT_TARG_PROC_INTEGRITY_LEVEL_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_SIGNATURE_STR:           mapStr(flattener.EVT_TARG_PROC_SIGNATURE_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_SIGNATURE_STATUS_STR:    mapStr(flattener.EVT_TARG_PROC_SIGNATURE_STATUS_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_SHA1_HASH_STR:           mapStr(flattener.EVT_TARG_PROC_SHA1_HASH_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_MD5_HASH_STR:            mapStr(flattener.EVT_TARG_PROC_MD5_HASH_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_SHA256_HASH_STR:         mapStr(flattener.EVT_TARG_PROC_SHA256_HASH_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_IMP_HASH_STR:            mapStr(flattener.EVT_TARG_PROC_IMP_HASH_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_SIGNED_INT:              mapInt(flattener.EVT_TARG_PROC_SIGNED_INT, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_START_ADDR_STR:          mapStr(flattener.EVT_TARG_PROC_START_ADDR_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_START_MODULE_STR:        mapStr(flattener.EVT_TARG_PROC_START_MODULE_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_START_FUNCTION_STR:      mapStr(flattener.EVT_TARG_PROC_START_FUNCTION_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_GRANT_ACCESS_STR:        mapStr(flattener.EVT_TARG_PROC_GRANT_ACCESS_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_CALL_TRACE_STR:          mapStr(flattener.EVT_TARG_PROC_CALL_TRACE_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_ACCESS_TYPE_STR:         mapStr(flattener.EVT_TARG_PROC_ACCESS_TYPE_STR, flattener.TARG_PROC_SRC),
		EXT_TARG_PROC_NEW_THREAD_ID_INT:       mapInt(flattener.EVT_TARG_PROC_NEW_THREAD_ID_INT, flattener.TARG_PROC_SRC),
	},
}

func mapStr(attr sfgo.Attribute, src flattener.Source) FieldMap {
	return func(r *Record) interface{} { return r.GetStr(attr, src) }
}

func mapInt(attr sfgo.Attribute, src flattener.Source) FieldMap {
	return func(r *Record) interface{} { return r.GetInt(attr, src) }
}

func mapSum(src flattener.Source, attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var sum int64 = 0
		for _, attr := range attrs {
			sum += r.GetInt(attr, src)
		}
		return sum
	}
}

func mapJoin(src flattener.Source, attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var join string = r.GetStr(attrs[0], src)
		for _, attr := range attrs[1:] {
			join += SPACE + r.GetStr(attr, src)
		}
		return join
	}
}

func mapRecType(src flattener.Source) FieldMap {
	return func(r *Record) interface{} {
		switch r.GetInt(sfgo.SF_REC_TYPE, src) {
		case sfgo.PROC:
			return TyP
		case sfgo.FILE:
			return TyF
		case sfgo.CONT:
			return TyC
		case sfgo.PROC_EVT:
			return TyPE
		case sfgo.FILE_EVT:
			return TyFE
		case sfgo.FILE_FLOW:
			return TyFF
		case sfgo.NET_FLOW:
			return TyNF
		case sfgo.HEADER:
			return TyH
		default:
			return TyUnknow
		}
	}
}

func mapOpFlags(src flattener.Source) FieldMap {
	return func(r *Record) interface{} {
		opflags := r.GetInt(sfgo.EV_PROC_OPFLAGS_INT, src)
		rtype := mapRecType(src)(r).(string)
		return strings.Join(utils.GetOpFlags(int32(opflags), rtype), LISTSEP)
	}
}

func mapRet(src flattener.Source) FieldMap {
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

func mapEndTs(src flattener.Source) FieldMap {
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

func mapEntry(attr sfgo.Attribute, src flattener.Source) FieldMap {
	return func(r *Record) interface{} {
		if r.GetInt(attr, src) == 1 {
			return true
		}
		return false
	}
}

func mapName(attr sfgo.Attribute, src flattener.Source) FieldMap {
	return func(r *Record) interface{} {
		return filepath.Base(r.GetStr(attr, src))
	}
}

func mapDir(attr sfgo.Attribute, src flattener.Source) FieldMap {
	return func(r *Record) interface{} {
		return filepath.Dir(r.GetStr(attr, src))
	}
}

func mapLinkPath(attr sfgo.Attribute, src flattener.Source) FieldMap {
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

func mapFileType(attr sfgo.Attribute, src flattener.Source) FieldMap {
	return func(r *Record) interface{} {
		return utils.GetFileType(r.GetInt(attr, src))
	}
}

func mapIsOpenWrite(attr sfgo.Attribute, src flattener.Source) FieldMap {
	return func(r *Record) interface{} {
		if utils.IsOpenWrite(r.GetInt(attr, src)) {
			return true
		}
		return false
	}
}

func mapIsOpenRead(attr sfgo.Attribute, src flattener.Source) FieldMap {
	return func(r *Record) interface{} {
		if utils.IsOpenRead(r.GetInt(attr, src)) {
			return true
		}
		return false
	}
}

func mapOpenFlags(attr sfgo.Attribute, src flattener.Source) FieldMap {
	return func(r *Record) interface{} {
		return strings.Join(utils.GetOpenFlags(r.GetInt(attr, src)), LISTSEP)
	}
}

func mapProto(attr sfgo.Attribute, src flattener.Source) FieldMap {
	return func(r *Record) interface{} {
		return r.GetInt(attr, src)
	}
}

func mapPort(src flattener.Source, attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var ports = make([]string, 0)
		for _, attr := range attrs {
			ports = append(ports, strconv.FormatInt(r.GetInt(attr, src), 10))
		}
		// logger.Info.Println(ports)
		return strings.Join(ports, LISTSEP)
	}
}

func mapIP(src flattener.Source, attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var ips = make([]string, 0)
		for _, attr := range attrs {
			ips = append(ips, utils.GetIPStr(int32(r.GetInt(attr, src))))
		}
		// logger.Info.Println(ips)
		return strings.Join(ips, LISTSEP)
	}
}

func mapContType(attr sfgo.Attribute, src flattener.Source) FieldMap {
	return func(r *Record) interface{} {
		return utils.GetContType(r.GetInt(attr, src))
	}
}

func mapCachedValue(attr RecAttribute, src flattener.Source) FieldMap {
	return func(r *Record) interface{} {
		oid := sfgo.OID{CreateTS: r.GetInt(sfgo.PROC_OID_CREATETS_INT, src), Hpid: r.GetInt(sfgo.PROC_OID_HPID_INT, src)}
		return r.GetCachedValue(oid, attr)
	}
}

func mapOID(src flattener.Source, attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		h := sha256.New()
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
