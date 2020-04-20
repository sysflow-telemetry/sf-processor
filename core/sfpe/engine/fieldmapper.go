package engine

import (
	"strconv"
	"strings"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-apis/go/utils"
	"github.ibm.com/sysflow/sf-processor/common/logger"
)

// FieldMap is a functional type denoting a SysFlow attribute mapper.
type FieldMap func(r sfgo.FlatRecord) interface{}

// IntFieldMap is a functional type denoting a numerical attribute mapper.
type IntFieldMap func(r sfgo.FlatRecord) int64

// StrFieldMap is a functional type denoting a string attribute mapper.
type StrFieldMap func(r sfgo.FlatRecord) string

// FieldMapper is an adapter for SysFlow attribute mappers.
type FieldMapper struct {
	Mappers map[string]FieldMap
}

// Map retrieves a field map based on a SysFlow attribute.
func (m FieldMapper) Map(attr string) FieldMap {
	if mapper, ok := m.Mappers[attr]; ok {
		return mapper
	}
	return func(r sfgo.FlatRecord) interface{} { return attr }
}

// MapInt retrieves a numerical field map based on a SysFlow attribute.
func (m FieldMapper) MapInt(attr string) IntFieldMap {
	return func(r sfgo.FlatRecord) int64 {
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
	return func(r sfgo.FlatRecord) string {
		if v, ok := m.Map(attr)(r).(string); ok {
			return v
		} else if v, ok := m.Map(attr)(r).(int64); ok {
			return strconv.FormatInt(v, 10)
		}
		return sfgo.Zeros.String
	}
}

// Mapper defines a global attribute mapper instance.
var Mapper = FieldMapper{
	map[string]FieldMap{
		"sf.type":                 mapRecType(),
		"sf.opflags":              mapOpFlags(),
		"sf.ret":                  mapInt(sfgo.RET_INT),
		"sf.ts":                   mapInt(sfgo.TS_INT),
		"sf.endts":                mapEndTs(),
		"sf.proc.pid":             mapInt(sfgo.PROC_OID_HPID_INT),
		"sf.proc.name":            mapName(sfgo.PROC_EXE_STR),
		"sf.proc.exe":             mapStr(sfgo.PROC_EXE_STR),
		"sf.proc.args":            mapStr(sfgo.PROC_EXEARGS_STR),
		"sf.proc.uid":             mapInt(sfgo.PROC_UID_INT),
		"sf.proc.user":            mapStr(sfgo.PROC_USERNAME_STR),
		"sf.proc.tid":             mapInt(sfgo.TID_INT),
		"sf.proc.gid":             mapInt(sfgo.PROC_GID_INT),
		"sf.proc.group":           mapStr(sfgo.PROC_GROUPNAME_STR),
		"sf.proc.createts":        mapInt(sfgo.PROC_OID_CREATETS_INT),
		"sf.proc.duration":        mapDuration(sfgo.PROC_OID_CREATETS_INT),
		"sf.proc.tty":             mapInt(sfgo.PROC_TTY_INT),
		"sf.proc.cmdline":         mapJoin(sfgo.PROC_EXE_STR, sfgo.PROC_EXEARGS_STR),
		"sf.proc.aname":           mapNa("sf.proc.aname"), // TBD
		"sf.proc.apid":            mapNa("sf.proc.apid"),  // TBD
		"sf.pproc.pid":            mapInt(sfgo.PROC_POID_HPID_INT),
		"sf.pproc.name":           mapNa("sf.pproc.name"),  // TBD
		"sf.pproc.exe":            mapNa("sf.pproc.exe"),   // TBD
		"sf.pproc.args":           mapNa("sf.pproc.args"),  // TBD
		"sf.pproc.uid":            mapNa("sf.pproc.uid"),   // TBD
		"sf.pproc.user":           mapNa("sf.pproc.user"),  // TBD
		"sf.pproc.gid":            mapNa("sf.pproc.gid"),   // TBD
		"sf.pproc.group":          mapNa("sf.pproc.group"), // TBD
		"sf.pproc.createts":       mapInt(sfgo.PROC_POID_CREATETS_INT),
		"sf.pproc.duration":       mapDuration(sfgo.PROC_POID_CREATETS_INT),
		"sf.pproc.tty":            mapNa("sf.pproc.tty"),     // TBD
		"sf.pproc.cmdline":        mapNa("sf.pproc.cmdline"), // TBD
		"sf.file.name":            mapName(sfgo.FILE_PATH_STR),
		"sf.file.path":            mapStr(sfgo.FILE_PATH_STR),
		"sf.file.directory":       mapDir(sfgo.FILE_PATH_STR),
		"sf.file.newname":         mapName(sfgo.SEC_FILE_PATH_STR),
		"sf.file.newpath":         mapStr(sfgo.SEC_FILE_PATH_STR),
		"sf.file.newdirectory":    mapDir(sfgo.SEC_FILE_PATH_STR),
		"sf.file.type":            mapFileType(sfgo.FILE_RESTYPE_INT),
		"sf.file.is_open_write":   mapIsOpenWrite(sfgo.FL_FILE_OPENFLAGS_INT),
		"sf.file.is_open_read":    mapIsOpenRead(sfgo.FL_FILE_OPENFLAGS_INT),
		"sf.file.fd":              mapInt(sfgo.FL_FILE_FD_INT),
		"sf.file.openflags":       mapOpenFlags(sfgo.FL_FILE_OPENFLAGS_INT),
		"sf.net.proto":            mapInt(sfgo.FL_NETW_PROTO_INT),
		"sf.net.protoname":        mapProto(sfgo.FL_NETW_PROTO_INT),
		"sf.net.sport":            mapPort(sfgo.FL_NETW_SPORT_INT),
		"sf.net.dport":            mapPort(sfgo.FL_NETW_DPORT_INT),
		"sf.net.port":             mapPort(sfgo.FL_NETW_SPORT_INT, sfgo.FL_NETW_DPORT_INT),
		"sf.net.sip":              mapIP(sfgo.FL_NETW_SIP_INT),
		"sf.net.dip":              mapIP(sfgo.FL_NETW_DIP_INT),
		"sf.net.ip":               mapIP(sfgo.FL_NETW_SIP_INT, sfgo.FL_NETW_DIP_INT),
		"sf.flow.rbytes":          mapSum(sfgo.FL_FILE_NUMRRECVBYTES_INT, sfgo.FL_NETW_NUMRRECVBYTES_INT),
		"sf.flow.rops":            mapSum(sfgo.FL_FILE_NUMRRECVOPS_INT, sfgo.FL_NETW_NUMRRECVOPS_INT),
		"sf.flow.wbytes":          mapSum(sfgo.FL_FILE_NUMWSENDBYTES_INT, sfgo.FL_NETW_NUMWSENDBYTES_INT),
		"sf.flow.wops":            mapSum(sfgo.FL_FILE_NUMWSENDOPS_INT, sfgo.FL_NETW_NUMWSENDOPS_INT),
		"sf.container.id":         mapStr(sfgo.CONT_ID_STR),
		"sf.container.name":       mapStr(sfgo.CONT_NAME_STR),
		"sf.container.imageid":    mapStr(sfgo.CONT_IMAGEID_STR),
		"sf.container.image":      mapStr(sfgo.CONT_IMAGE_STR),
		"sf.container.type":       mapContType(sfgo.CONT_TYPE_INT),
		"sf.container.privileged": mapInt(sfgo.CONT_PRIVILEGED_INT),
	},
}

func mapStr(attr sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} { return r.Strs[attr] }
}

func mapInt(attr sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} { return r.Ints[attr] }
}

func mapSum(attrs ...sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		var sum int64 = 0
		for _, attr := range attrs {
			sum += r.Ints[attr]
		}
		return sum
	}
}

func mapJoin(attrs ...sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		var join string = r.Strs[attrs[0]]
		for _, attr := range attrs[1:] {
			join += SPACE + r.Strs[attr]
		}
		return join
	}
}

func mapRecType() FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		switch r.Ints[sfgo.SF_REC_TYPE] {
		case sfgo.PROC:
			return "P"
		case sfgo.FILE:
			return "F"
		case sfgo.CONT:
			return "C"
		case sfgo.PROC_EVT:
			return "PE"
		case sfgo.FILE_EVT:
			return "FE"
		case sfgo.FILE_FLOW:
			return "FF"
		case sfgo.NET_FLOW:
			return "NF"
		case sfgo.HEADER:
			return "H"
		default:
			return ""
		}
	}
}

func mapOpFlags() FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		opflags := r.Ints[sfgo.EV_PROC_OPFLAGS_INT]
		return strings.Join(utils.GetOpFlags(int32(opflags)), LISTSEP)
	}
}

func mapEndTs() FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		switch r.Ints[sfgo.SF_REC_TYPE] {
		case sfgo.FILE_FLOW:
			return r.Ints[sfgo.FL_FILE_ENDTS_INT]
		case sfgo.NET_FLOW:
			return r.Ints[sfgo.FL_NETW_ENDTS_INT]
		default:
			return sfgo.Zeros.Int64
		}
	}
}

func mapDuration(attr sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		return time.Now().Unix() - r.Ints[attr]
	}
}

func mapName(attr sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		return r.Strs[attr]
	}
}

func mapDir(attr sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		return r.Strs[attr]
	}
}

func mapFileType(attr sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		return r.Ints[attr]
	}
}

func mapIsOpenWrite(attr sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		return r.Ints[attr]
	}
}

func mapOpenFlags(attrs ...sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		return r.Ints[attrs[0]]
	}
}

func mapIsOpenRead(attr sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		return r.Ints[attr]
	}
}

func mapProto(attr sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		return r.Ints[attr]
	}
}

func mapPort(attrs ...sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		return r.Ints[attrs[0]]
	}
}

func mapIP(attrs ...sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		return r.Ints[attrs[0]]
	}
}

func mapContType(attr sfgo.Attribute) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		return r.Ints[attr]
	}
}

func mapNa(attr string) FieldMap {
	return func(r sfgo.FlatRecord) interface{} {
		logger.Warn.Println("Attribute not supported ", attr)
		return sfgo.Zeros.String
	}
}
