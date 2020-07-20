package engine

import (
	"crypto/sha256"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-apis/go/utils"
	"github.ibm.com/sysflow/goutils/logger"
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
		SF_TYPE:                 mapRecType(),
		SF_OPFLAGS:              mapOpFlags(),
		SF_RET:                  mapRet(),
		SF_TS:                   mapInt(sfgo.TS_INT),
		SF_ENDTS:                mapEndTs(),
		SF_PROC_OID:             mapOID(sfgo.PROC_OID_HPID_INT, sfgo.PROC_OID_CREATETS_INT),
		SF_PROC_PID:             mapInt(sfgo.PROC_OID_HPID_INT),
		SF_PROC_NAME:            mapName(sfgo.PROC_EXE_STR),
		SF_PROC_EXE:             mapStr(sfgo.PROC_EXE_STR),
		SF_PROC_ARGS:            mapStr(sfgo.PROC_EXEARGS_STR),
		SF_PROC_UID:             mapInt(sfgo.PROC_UID_INT),
		SF_PROC_USER:            mapStr(sfgo.PROC_USERNAME_STR),
		SF_PROC_TID:             mapInt(sfgo.TID_INT),
		SF_PROC_GID:             mapInt(sfgo.PROC_GID_INT),
		SF_PROC_GROUP:           mapStr(sfgo.PROC_GROUPNAME_STR),
		SF_PROC_CREATETS:        mapInt(sfgo.PROC_OID_CREATETS_INT),
		SF_PROC_DURATION:        mapDuration(sfgo.PROC_OID_CREATETS_INT),
		SF_PROC_TTY:             mapInt(sfgo.PROC_TTY_INT),
		SF_PROC_CMDLINE:         mapJoin(sfgo.PROC_EXE_STR, sfgo.PROC_EXEARGS_STR),
		SF_PROC_ANAME:           mapCachedValue(ProcAName),
		SF_PROC_AEXE:            mapCachedValue(ProcAExe),
		SF_PROC_ACMDLINE:        mapCachedValue(ProcACmdLine),
		SF_PROC_APID:            mapCachedValue(ProcAPID),
		SF_PPROC_OID:            mapOID(sfgo.PROC_POID_HPID_INT, sfgo.PROC_POID_CREATETS_INT),
		SF_PPROC_PID:            mapInt(sfgo.PROC_POID_HPID_INT),
		SF_PPROC_NAME:           mapCachedValue(PProcName),
		SF_PPROC_EXE:            mapCachedValue(PProcExe),
		SF_PPROC_ARGS:           mapCachedValue(PProcArgs),
		SF_PPROC_UID:            mapCachedValue(PProcUID),
		SF_PPROC_USER:           mapCachedValue(PProcUser),
		SF_PPROC_GID:            mapCachedValue(PProcGID),
		SF_PPROC_GROUP:          mapCachedValue(PProcGroup),
		SF_PPROC_CREATETS:       mapInt(sfgo.PROC_POID_CREATETS_INT),
		SF_PPROC_DURATION:       mapDuration(sfgo.PROC_POID_CREATETS_INT),
		SF_PPROC_TTY:            mapCachedValue(PProcTTY),
		SF_PPROC_CMDLINE:        mapCachedValue(PProcCmdLine),
		SF_FILE_NAME:            mapName(sfgo.FILE_PATH_STR),
		SF_FILE_PATH:            mapStr(sfgo.FILE_PATH_STR),
		SF_FILE_CANONICALPATH:   mapLinkPath(sfgo.FILE_PATH_STR),
		SF_FILE_DIRECTORY:       mapDir(sfgo.FILE_PATH_STR),
		SF_FILE_NEWNAME:         mapName(sfgo.SEC_FILE_PATH_STR),
		SF_FILE_NEWPATH:         mapStr(sfgo.SEC_FILE_PATH_STR),
		SF_FILE_NEWDIRECTORY:    mapDir(sfgo.SEC_FILE_PATH_STR),
		SF_FILE_TYPE:            mapFileType(sfgo.FILE_RESTYPE_INT),
		SF_FILE_IS_OPEN_WRITE:   mapIsOpenWrite(sfgo.FL_FILE_OPENFLAGS_INT),
		SF_FILE_IS_OPEN_READ:    mapIsOpenRead(sfgo.FL_FILE_OPENFLAGS_INT),
		SF_FILE_FD:              mapInt(sfgo.FL_FILE_FD_INT),
		SF_FILE_OPENFLAGS:       mapOpenFlags(sfgo.FL_FILE_OPENFLAGS_INT),
		SF_NET_PROTO:            mapInt(sfgo.FL_NETW_PROTO_INT),
		SF_NET_PROTONAME:        mapProto(sfgo.FL_NETW_PROTO_INT),
		SF_NET_SPORT:            mapInt(sfgo.FL_NETW_SPORT_INT),
		SF_NET_DPORT:            mapInt(sfgo.FL_NETW_DPORT_INT),
		SF_NET_PORT:             mapPort(sfgo.FL_NETW_SPORT_INT, sfgo.FL_NETW_DPORT_INT),
		SF_NET_SIP:              mapIP(sfgo.FL_NETW_SIP_INT),
		SF_NET_DIP:              mapIP(sfgo.FL_NETW_DIP_INT),
		SF_NET_IP:               mapIP(sfgo.FL_NETW_SIP_INT, sfgo.FL_NETW_DIP_INT),
		SF_FLOW_RBYTES:          mapSum(sfgo.FL_FILE_NUMRRECVBYTES_INT, sfgo.FL_NETW_NUMRRECVBYTES_INT),
		SF_FLOW_ROPS:            mapSum(sfgo.FL_FILE_NUMRRECVOPS_INT, sfgo.FL_NETW_NUMRRECVOPS_INT),
		SF_FLOW_WBYTES:          mapSum(sfgo.FL_FILE_NUMWSENDBYTES_INT, sfgo.FL_NETW_NUMWSENDBYTES_INT),
		SF_FLOW_WOPS:            mapSum(sfgo.FL_FILE_NUMWSENDOPS_INT, sfgo.FL_NETW_NUMWSENDOPS_INT),
		SF_CONTAINER_ID:         mapStr(sfgo.CONT_ID_STR),
		SF_CONTAINER_NAME:       mapStr(sfgo.CONT_NAME_STR),
		SF_CONTAINER_IMAGEID:    mapStr(sfgo.CONT_IMAGEID_STR),
		SF_CONTAINER_IMAGE:      mapStr(sfgo.CONT_IMAGE_STR),
		SF_CONTAINER_TYPE:       mapContType(sfgo.CONT_TYPE_INT),
		SF_CONTAINER_PRIVILEGED: mapInt(sfgo.CONT_PRIVILEGED_INT),
		SF_NODE_ID:              mapStr(sfgo.SFHE_EXPORTER_STR),
	},
}

func mapStr(attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} { return r.GetStr(attr) }
}

func mapInt(attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} { return r.GetInt(attr) }
}

func mapSum(attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var sum int64 = 0
		for _, attr := range attrs {
			sum += r.GetInt(attr)
		}
		return sum
	}
}

func mapJoin(attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var join string = r.GetStr(attrs[0])
		for _, attr := range attrs[1:] {
			join += SPACE + r.GetStr(attr)
		}
		return join
	}
}

func mapRecType() FieldMap {
	return func(r *Record) interface{} {
		switch r.GetInt(sfgo.SF_REC_TYPE) {
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

func mapOpFlags() FieldMap {
	return func(r *Record) interface{} {
		opflags := r.GetInt(sfgo.EV_PROC_OPFLAGS_INT)
		rtype := mapRecType()(r).(string)
		return strings.Join(utils.GetOpFlags(int32(opflags), rtype), LISTSEP)
	}
}

func mapRet() FieldMap {
	return func(r *Record) interface{} {
		switch r.GetInt(sfgo.SF_REC_TYPE) {
		case sfgo.PROC_EVT:
			fallthrough
		case sfgo.FILE_EVT:
			return r.GetInt(sfgo.RET_INT)
		default:
			return sfgo.Zeros.Int64
		}
	}
}

func mapEndTs() FieldMap {
	return func(r *Record) interface{} {
		switch r.GetInt(sfgo.SF_REC_TYPE) {
		case sfgo.FILE_FLOW:
			return r.GetInt(sfgo.FL_FILE_ENDTS_INT)
		case sfgo.NET_FLOW:
			return r.GetInt(sfgo.FL_NETW_ENDTS_INT)
		default:
			return sfgo.Zeros.Int64
		}
	}
}

func mapDuration(attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return time.Now().UnixNano() - r.GetInt(attr)
	}
}

func mapName(attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return filepath.Base(r.GetStr(attr))
	}
}

func mapDir(attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return filepath.Dir(r.GetStr(attr))
	}
}

func mapLinkPath(attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		orig := r.GetStr(attr)
		// Possible format: aabbccddeeff0011->aabbccddeeff0011 /path/to/target.file
		var src, dst uint64
		var targetPath string
		if _, err := fmt.Sscanf(orig, "%x->%x %s", &src, &dst, &targetPath); nil == err {
			return targetPath
		}
		return orig
	}
}

func mapFileType(attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return utils.GetFileType(r.GetInt(attr))
	}
}

func mapIsOpenWrite(attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		if utils.IsOpenWrite(r.GetInt(attr)) {
			return true
		}
		return false
	}
}

func mapIsOpenRead(attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		if utils.IsOpenRead(r.GetInt(attr)) {
			return true
		}
		return false
	}
}

func mapOpenFlags(attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return strings.Join(utils.GetOpenFlags(r.GetInt(attr)), LISTSEP)
	}
}

func mapProto(attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return r.GetInt(attr)
	}
}

func mapPort(attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var ports = make([]string, 0)
		for _, attr := range attrs {
			ports = append(ports, strconv.FormatInt(r.GetInt(attr), 10))
		}
		// logger.Info.Println(ports)
		return strings.Join(ports, LISTSEP)
	}
}

func mapIP(attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var ips = make([]string, 0)
		for _, attr := range attrs {
			ips = append(ips, utils.GetIPStr(int32(r.GetInt(attr))))
		}
		// logger.Info.Println(ips)
		return strings.Join(ips, LISTSEP)
	}
}

func mapContType(attr sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		return utils.GetContType(r.GetInt(attr))
	}
}

func mapCachedValue(attr RecAttribute) FieldMap {
	return func(r *Record) interface{} {
		oid := sfgo.OID{CreateTS: r.GetInt(sfgo.PROC_OID_CREATETS_INT), Hpid: r.GetInt(sfgo.PROC_OID_HPID_INT)}
		return r.GetCachedValue(oid, attr)
	}
}

func mapOID(attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		h := sha256.New()
		for _, attr := range attrs {
			h.Write([]byte(fmt.Sprintf("%v", r.GetInt(attr))))
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
