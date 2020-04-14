package engine

import (
	"strconv"
	"strings"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/utils"
	"github.ibm.com/sysflow/sf-processor/common/logger"
	"github.ibm.com/sysflow/sf-processor/plugins/flattener/types"
)

// FieldMap is a functional type denoting a SysFlow attribute mapper.
type FieldMap func(r types.FlatRecord) interface{}

// IntFieldMap is a functional type denoting a numerical attribute mapper.
type IntFieldMap func(r types.FlatRecord) int64

// StrFieldMap is a functional type denoting a string attribute mapper.
type StrFieldMap func(r types.FlatRecord) string

// Types is used to obtain zero values for supported types.
type Types struct {
	Int64  int64
	String string
}

// Zeros is a zero-initialized struct used to obtain zero values for supported types.
var Zeros = Types{}

// FieldMapper is an adapter for SysFlow attribute mappers.
type FieldMapper struct {
	Mappers map[string]FieldMap
}

// Map retrieves a field map based on a SysFlow attribute.
func (m FieldMapper) Map(attr string) FieldMap {
	if mapper, ok := m.Mappers[attr]; ok {
		return mapper
	}
	return func(r types.FlatRecord) interface{} { return attr }
}

// MapInt retrieves a numerical field map based on a SysFlow attribute.
func (m FieldMapper) MapInt(attr string) IntFieldMap {
	return func(r types.FlatRecord) int64 {
		if v, ok := m.Map(attr)(r).(int64); ok {
			return v
		} else if v, err := strconv.ParseInt(attr, 10, 64); err == nil {
			return v
		}
		return Zeros.Int64
	}
}

// MapStr retrieves a string field map based on a SysFlow attribute.
func (m FieldMapper) MapStr(attr string) StrFieldMap {
	return func(r types.FlatRecord) string {
		if v, ok := m.Map(attr)(r).(string); ok {
			return v
		} else if v, ok := m.Map(attr)(r).(int64); ok {
			return strconv.FormatInt(v, 10)
		}
		return Zeros.String
	}
}

// Mapper defines a global attribute mapper instance.
var Mapper = FieldMapper{
	map[string]FieldMap{
		"sf.type":                 mapRecType(),
		"sf.opflags":              mapOpFlags(),
		"sf.ret":                  mapInt(types.EV_PROC_RET_INT), // normalize
		"sf.ts":                   mapInt(types.EV_PROC_TS_INT),  // normalize
		"sf.endts":                mapEndTs(),
		"sf.proc.pid":             mapInt(types.PROC_OID_HPID_INT),
		"sf.proc.name":            mapStr(types.PROC_EXE_STR),
		"sf.proc.exe":             mapStr(types.PROC_EXE_STR),
		"sf.proc.args":            mapStr(types.PROC_EXEARGS_STR),
		"sf.proc.uid":             mapInt(types.PROC_UID_INT),
		"sf.proc.user":            mapStr(types.PROC_USERNAME_STR),
		"sf.proc.tid":             mapInt(types.EV_PROC_TID_INT), // normalize
		"sf.proc.gid":             mapInt(types.PROC_GID_INT),
		"sf.proc.group":           mapStr(types.PROC_GROUPNAME_STR),
		"sf.proc.createts":        mapInt(types.PROC_OID_CREATETS_INT),
		"sf.proc.duration":        mapDuration(types.PROC_OID_CREATETS_INT),
		"sf.proc.tty":             mapInt(types.PROC_TTY_INT),
		"sf.proc.cmdline":         mapJoin(types.PROC_EXE_STR, types.PROC_EXEARGS_STR),
		"sf.proc.aname":           mapNa("sf.proc.aname"), // TBD
		"sf.proc.apid":            mapNa("sf.proc.apid"),  // TBD
		"sf.pproc.pid":            mapInt(types.PROC_POID_HPID_INT),
		"sf.pproc.name":           mapNa("sf.pproc.name"),  // TBD
		"sf.pproc.exe":            mapNa("sf.pproc.exe"),   // TBD
		"sf.pproc.args":           mapNa("sf.pproc.args"),  // TBD
		"sf.pproc.uid":            mapNa("sf.pproc.uid"),   // TBD
		"sf.pproc.user":           mapNa("sf.pproc.user"),  // TBD
		"sf.pproc.gid":            mapNa("sf.pproc.gid"),   // TBD
		"sf.pproc.group":          mapNa("sf.pproc.group"), // TBD
		"sf.pproc.createts":       mapInt(types.PROC_POID_CREATETS_INT),
		"sf.pproc.duration":       mapDuration(types.PROC_POID_CREATETS_INT),
		"sf.pproc.tty":            mapNa("sf.pproc.tty"),     // TBD
		"sf.pproc.cmdline":        mapNa("sf.pproc.cmdline"), // TBD
		"sf.file.name":            mapName(types.FILE_PATH_STR),
		"sf.file.path":            mapStr(types.FILE_PATH_STR),
		"sf.file.directory":       mapDir(types.FILE_PATH_STR),
		"sf.file.newname":         mapName(types.SEC_FILE_PATH_STR),
		"sf.file.newpath":         mapStr(types.SEC_FILE_PATH_STR),
		"sf.file.newdirectory":    mapDir(types.SEC_FILE_PATH_STR),
		"sf.file.type":            mapFileType(types.FILE_RESTYPE_INT),
		"sf.file.is_open_write":   mapIsOpenWrite(types.FL_FILE_OPENFLAGS_INT),
		"sf.file.is_open_read":    mapIsOpenRead(types.FL_FILE_OPENFLAGS_INT),
		"sf.file.fd":              mapInt(types.FL_FILE_FD_INT),
		"sf.file.openflags":       mapOpenFlags(types.FL_FILE_OPENFLAGS_INT),
		"sf.net.proto":            mapInt(types.FL_NETW_PROTO_INT),
		"sf.net.protoname":        mapProto(types.FL_NETW_PROTO_INT),
		"sf.net.sport":            mapPort(types.FL_NETW_SPORT_INT),
		"sf.net.dport":            mapPort(types.FL_NETW_DPORT_INT),
		"sf.net.port":             mapPort(types.FL_NETW_SPORT_INT, types.FL_NETW_DPORT_INT),
		"sf.net.sip":              mapIP(types.FL_NETW_SIP_INT),
		"sf.net.dip":              mapIP(types.FL_NETW_DIP_INT),
		"sf.net.ip":               mapIP(types.FL_NETW_SIP_INT, types.FL_NETW_DIP_INT),
		"sf.flow.rbytes":          mapSum(types.FL_FILE_NUMRRECVBYTES_INT, types.FL_NETW_NUMRRECVBYTES_INT),
		"sf.flow.rops":            mapSum(types.FL_FILE_NUMRRECVOPS_INT, types.FL_NETW_NUMRRECVOPS_INT),
		"sf.flow.wbytes":          mapSum(types.FL_FILE_NUMWSENDBYTES_INT, types.FL_NETW_NUMWSENDBYTES_INT),
		"sf.flow.wops":            mapSum(types.FL_FILE_NUMWSENDOPS_INT, types.FL_NETW_NUMWSENDOPS_INT),
		"sf.container.id":         mapStr(types.CONT_ID_STR),
		"sf.container.name":       mapStr(types.CONT_NAME_STR),
		"sf.container.imageid":    mapStr(types.CONT_IMAGEID_STR),
		"sf.container.image":      mapStr(types.CONT_IMAGE_STR),
		"sf.container.type":       mapContType(types.CONT_TYPE_INT),
		"sf.container.privileged": mapInt(types.CONT_PRIVILEGED_INT),
	},
}

func mapStr(attr types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} { return r.Strs[attr] }
}

func mapInt(attr types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} { return r.Ints[attr] }
}

func mapSum(attrs ...types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		var sum int64 = 0
		for _, attr := range attrs {
			sum += r.Ints[attr]
		}
		return sum
	}
}

func mapJoin(attrs ...types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		var join string = r.Strs[attrs[0]]
		for _, attr := range attrs[1:] {
			join += SPACE + r.Strs[attr]
		}
		return join
	}
}

func mapRecType() FieldMap {
	return func(r types.FlatRecord) interface{} {
		switch r.Ints[types.SF_REC_TYPE] {
		case types.PROC:
			return "P"
		case types.FILE:
			return "F"
		case types.CONT:
			return "C"
		case types.PROC_EVT:
			return "PE"
		case types.FILE_EVT:
			return "FE"
		case types.FILE_FLOW:
			return "FF"
		case types.NET_FLOW:
			return "NF"
		case types.HEADER:
			return "H"
		default:
			return ""
		}
	}
}

func mapOpFlags() FieldMap {
	return func(r types.FlatRecord) interface{} {
		opflags := r.Ints[types.EV_PROC_OPFLAGS_INT]
		return strings.Join(utils.GetOpFlags(int32(opflags)), LISTSEP)
	}
}

func mapEndTs() FieldMap {
	return func(r types.FlatRecord) interface{} {
		switch r.Ints[types.SF_REC_TYPE] {
		case types.FILE_FLOW:
			return r.Ints[types.FL_FILE_ENDTS_INT]
		case types.NET_FLOW:
			return r.Ints[types.FL_NETW_ENDTS_INT]
		default:
			return Zeros.Int64
		}
	}
}

func mapDuration(attr types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		return time.Now().Unix() - r.Ints[attr]
	}
}

func mapName(attr types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		return r.Strs[attr]
	}
}

func mapDir(attr types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		return r.Strs[attr]
	}
}

func mapFileType(attr types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		return r.Ints[attr]
	}
}

func mapIsOpenWrite(attr types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		return r.Ints[attr]
	}
}

func mapOpenFlags(attrs ...types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		return r.Ints[attrs[0]]
	}
}

func mapIsOpenRead(attr types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		return r.Ints[attr]
	}
}

func mapProto(attr types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		return r.Ints[attr]
	}
}

func mapPort(attrs ...types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		return r.Ints[attrs[0]]
	}
}

func mapIP(attrs ...types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		return r.Ints[attrs[0]]
	}
}

func mapContType(attr types.Attribute) FieldMap {
	return func(r types.FlatRecord) interface{} {
		return r.Ints[attr]
	}
}

func mapNa(attr string) FieldMap {
	return func(r types.FlatRecord) interface{} {
		logger.Warn.Println("Attribute not supported ", attr)
		return Zeros.String
	}
}
