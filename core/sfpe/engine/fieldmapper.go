package engine

import (
	"crypto/sha256"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-apis/go/utils"
	"github.ibm.com/sysflow/sf-processor/common/logger"
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

// Mapper defines a global attribute mapper instance.
var Mapper = FieldMapper{
	map[string]FieldMap{
		"sf.type":                 mapRecType(),
		"sf.opflags":              mapOpFlags(),
		"sf.ret":                  mapInt(sfgo.RET_INT),
		"sf.ts":                   mapInt(sfgo.TS_INT),
		"sf.endts":                mapEndTs(),
		"sf.proc.oid":             mapOID(sfgo.PROC_OID_HPID_INT, sfgo.PROC_OID_CREATETS_INT),
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
		"sf.proc.aname":           mapCachedValue(ProcAName),
		"sf.proc.aexe":            mapCachedValue(ProcAExe),
		"sf.proc.acmdline":        mapCachedValue(ProcACmdLine),
		"sf.proc.apid":            mapCachedValue(ProcAPID),
		"sf.pproc.oid":            mapOID(sfgo.PROC_POID_HPID_INT, sfgo.PROC_POID_CREATETS_INT),
		"sf.pproc.pid":            mapInt(sfgo.PROC_POID_HPID_INT),
		"sf.pproc.name":           mapCachedValue(PProcName),
		"sf.pproc.exe":            mapCachedValue(PProcExe),
		"sf.pproc.args":           mapCachedValue(PProcArgs),
		"sf.pproc.uid":            mapCachedValue(PProcUID),
		"sf.pproc.user":           mapCachedValue(PProcUser),
		"sf.pproc.gid":            mapCachedValue(PProcGID),
		"sf.pproc.group":          mapCachedValue(PProcGroup),
		"sf.pproc.createts":       mapInt(sfgo.PROC_POID_CREATETS_INT),
		"sf.pproc.duration":       mapDuration(sfgo.PROC_POID_CREATETS_INT),
		"sf.pproc.tty":            mapCachedValue(PProcTTY),
		"sf.pproc.cmdline":        mapCachedValue(PProcCmdLine),
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
		"sf.net.sport":            mapInt(sfgo.FL_NETW_SPORT_INT),
		"sf.net.dport":            mapInt(sfgo.FL_NETW_DPORT_INT),
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
	return func(r *Record) interface{} {
		opflags := r.GetInt(sfgo.EV_PROC_OPFLAGS_INT)
		return strings.Join(utils.GetOpFlags(int32(opflags)), LISTSEP)
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
		return time.Now().Unix() - r.GetInt(attr)
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
		return strings.Join(ports, LISTSEP)
	}
}

func mapIP(attrs ...sfgo.Attribute) FieldMap {
	return func(r *Record) interface{} {
		var ips = make([]string, 0)
		for _, attr := range attrs {
			ips = append(ips, utils.GetIPStr(int32(r.GetInt(attr))))
		}
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
