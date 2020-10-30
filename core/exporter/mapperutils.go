package exporter

import (
	"github.com/mailru/easyjson/jwriter"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/sf-processor/core/policyengine/engine"
	"path/filepath"
	"reflect"
	"unsafe"
)

// MapBuffer retrieves a field map based on a SysFlow attribute storing it in a bytes buffer.
/*func MapBuffer(fv *FieldValue, buf *bytes.Buffer) engine.VoidFieldMap {
	return func(r *engine.Record) {
		switch fv.Entry.Id {
		case A_IDS, PARENT_IDS:
			oid := sfgo.OID{CreateTS: r.GetInt(sfgo.PROC_OID_CREATETS_INT, fv.Entry.Source), Hpid: r.GetInt(sfgo.PROC_OID_HPID_INT, fv.Entry.Source)}
			r.SetCachedValueBuffer(oid, fv.Entry.AuxAttr, buf)
			return
		}

		switch fv.Entry.Type {
		case engine.MapStrVal:
			v := r.GetStr(fv.Entry.Id, fv.Entry.Source)
			l := len(v)
			if l > 0 && (v[0] == '"' || v[0] == '\'') {
				buf.WriteString(v)
			} else {
				buf.WriteByte('"')
				buf.WriteString(v)
				buf.WriteByte('"')
			}
		case engine.MapIntVal:
			buf.WriteString(strconv.FormatInt(r.GetInt(fv.Entry.Id, fv.Entry.Source), 10))
		case engine.MapSpecialStr:
			v := fv.Entry.Map(r).(string)
			l := len(v)
			if l > 0 && (v[0] == '"' || v[0] == '\'') {
				buf.WriteString(v)
			} else {
				buf.WriteByte('"')
				buf.WriteString(v)
				buf.WriteByte('"')
			}
		case engine.MapSpecialInt:
			buf.WriteString(strconv.FormatInt(fv.Entry.Map(r).(int64), 10))
		case engine.MapSpecialBool:
			buf.WriteString(strconv.FormatBool(fv.Entry.Map(r).(bool)))
		case engine.MapArrayStr, MapArrayInt:
			v := fv.Entry.Map(r).(string)
			buf.WriteString("[")
			buf.WriteString(v)
			buf.WriteString("]")
		}
	}
}*/

func mapOpFlags(fv *engine.FieldValue, writer *jwriter.Writer, r *engine.Record) {
	opflags := r.GetInt(fv.Entry.Id, fv.Entry.Source)
	rtype := engine.GetRecType(r, fv.Entry.Source)
	flags := sfgo.GetOpFlags(int32(opflags), rtype)
	mapStrArray(writer, flags)
}

func mapStrArray(writer *jwriter.Writer, ss []string) {
	l := len(ss)
	writer.RawByte(BEGIN_SQUARE)
	for idx, s := range ss {
		writer.RawByte(DOUBLE_QUOTE)
		writer.RawString(s)
		writer.RawByte(DOUBLE_QUOTE)
		if idx < (l - 1) {
			writer.RawByte(COMMA)
		}
	}
	writer.RawByte(END_SQUARE)

}

func mapIPStr(ip int64, w *jwriter.Writer) {
	w.Int64(ip >> 0 & 0xFF)
	w.RawByte(PERIOD)
	w.Int64(ip >> 8 & 0xFF)
	w.RawByte(PERIOD)
	w.Int64(ip >> 16 & 0xFF)
	w.RawByte(PERIOD)
	w.Int64(ip >> 24 & 0xFF)
}
func mapIPs(fv *engine.FieldValue, writer *jwriter.Writer, r *engine.Record) {
	srcIP := r.GetInt(sfgo.FL_NETW_SIP_INT, fv.Entry.Source)
	dstIP := r.GetInt(sfgo.FL_NETW_DIP_INT, fv.Entry.Source)
	writer.RawByte(BEGIN_SQUARE)
	mapIPStr(srcIP, writer)
	writer.RawByte(COMMA)
	mapIPStr(dstIP, writer)
	writer.RawByte(END_SQUARE)
}

func mapOpenFlags(fv *engine.FieldValue, writer *jwriter.Writer, r *engine.Record) {
	flags := sfgo.GetOpenFlags(r.GetInt(fv.Entry.Id, fv.Entry.Source))
	mapStrArray(writer, flags)
}

func mapPorts(fv *engine.FieldValue, writer *jwriter.Writer, r *engine.Record) {
	srcPort := r.GetInt(sfgo.FL_NETW_SPORT_INT, fv.Entry.Source)
	dstPort := r.GetInt(sfgo.FL_NETW_DPORT_INT, fv.Entry.Source)
	writer.RawByte(BEGIN_SQUARE)
	writer.Int64(srcPort)
	writer.RawByte(COMMA)
	writer.Int64(dstPort)
	writer.RawByte(END_SQUARE)
}

// MapStr retrieves a string field map based on a SysFlow attribute.
func MapJSON(fv *engine.FieldValue, writer *jwriter.Writer, r *engine.Record) {
	switch fv.Entry.Id {
	case engine.A_IDS, engine.PARENT_IDS:
		oid := sfgo.OID{CreateTS: r.GetInt(sfgo.PROC_OID_CREATETS_INT, fv.Entry.Source), Hpid: r.GetInt(sfgo.PROC_OID_HPID_INT, fv.Entry.Source)}
		SetCachedValueJSON(r, oid, fv.Entry.AuxAttr, writer)
		return
	}

	switch fv.Entry.Type {
	case engine.MapStrVal:
		v := r.GetStr(fv.Entry.Id, fv.Entry.Source)
		l := len(v)
		if l > 0 && (v[0] == '"' || v[0] == '\'') {
			boundingQuotes := trimBoundingQuotes(v)
			writer.String(boundingQuotes)
		} else {
			writer.String(v)
		}
	case engine.MapIntVal:
		writer.Int64(r.GetInt(fv.Entry.Id, fv.Entry.Source))
	case engine.MapSpecialStr:
		v := fv.Entry.Map(r).(string)
		l := len(v)
		if l > 0 && (v[0] == '"' || v[0] == '\'') {
			boundingQuotes := trimBoundingQuotes(v)
			writer.String(boundingQuotes)
		} else {
			writer.String(v)
		}
	case engine.MapSpecialInt:
		writer.Int64(fv.Entry.Map(r).(int64))
	case engine.MapSpecialBool:
		writer.Bool(fv.Entry.Map(r).(bool))
	case engine.MapArrayStr, engine.MapArrayInt:
		if fv.Entry.Source == sfgo.SYSFLOW_SRC {
			switch fv.Entry.Id {
			case sfgo.EV_PROC_OPFLAGS_INT:
				recType := r.GetInt(sfgo.SF_REC_TYPE, fv.Entry.Source)
				if recType == sfgo.NET_FLOW {
					mapIPs(fv, writer, r)
					return
				}
				mapOpFlags(fv, writer, r)
				return
			case sfgo.FL_FILE_OPENFLAGS_INT:
				mapOpenFlags(fv, writer, r)
				return
			case sfgo.FL_NETW_SPORT_INT:
				mapPorts(fv, writer, r)
				return
			}
		}

		v := fv.Entry.Map(r).(string)
		writer.RawByte('[')
		writer.String(v)
		writer.RawByte(']')
	}
}

func trimBoundingQuotes(s string) string {
	if len(s) > 0 && (s[0] == '"' || s[0] == '\'') {
		s = s[1:]
	}
	if len(s) > 0 && (s[len(s)-1] == '"' || s[len(s)-1] == '\'') {
		s = s[:len(s)-1]
	}
	return s
}

func CheckForQuotes(v string, writer *jwriter.Writer) {
	l := len(v)
	if l > 0 && (v[0] == '"' || v[0] == '\'') {
		boundingQuotes := trimBoundingQuotes(v)
		writer.String(boundingQuotes)
	} else {
		writer.String(v)
	}
}

func UnsafeBytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

// SetCachedValueJSON sets the value of attr from cache for process ID to a JSON writer.
func SetCachedValueJSON(r *engine.Record, ID sfgo.OID, attr engine.RecAttribute, writer *jwriter.Writer) {
	if ptree := r.MemoizePtree(ID); ptree != nil {
		switch attr {
		case engine.PProcName:
			if len(ptree) > 1 {
				CheckForQuotes(filepath.Base(ptree[1].Exe), writer)
			}
			break
		case engine.PProcExe:
			if len(ptree) > 1 {
				CheckForQuotes(ptree[1].Exe, writer)
			}
			break
		case engine.PProcArgs:
			if len(ptree) > 1 {
				CheckForQuotes(ptree[1].ExeArgs, writer)
			}
			break
		case engine.PProcUID:
			if len(ptree) > 1 {
				writer.Int64(int64(ptree[1].Uid))
			}
			break
		case engine.PProcUser:
			if len(ptree) > 1 {
				CheckForQuotes(ptree[1].UserName, writer)
			}
			break
		case engine.PProcGID:
			if len(ptree) > 1 {
				writer.Int64(int64(ptree[1].Gid))
			}
			break
		case engine.PProcGroup:
			if len(ptree) > 1 {
				CheckForQuotes(ptree[1].GroupName, writer)
			}
			break
		case engine.PProcTTY:
			if len(ptree) > 1 {
				if ptree[1].Tty {
					writer.RawByte('1')
				} else {
					writer.RawByte('0')
				}
			}
			break
		case engine.PProcEntry:
			if len(ptree) > 1 {
				if ptree[1].Entry {
					writer.RawByte('1')
				} else {
					writer.RawByte('0')
				}
			}
			break
		case engine.PProcCmdLine:
			if len(ptree) > 1 {
				exe := trimBoundingQuotes(ptree[1].Exe)
				exeArgs := trimBoundingQuotes(ptree[1].ExeArgs)
				writer.RawByte('"')
				StringNoQuotes(exe, writer)
				if len(exeArgs) > 0 {
					writer.RawByte(' ')
					StringNoQuotes(exeArgs, writer)
				}
				writer.RawByte('"')
			}
			break
		case engine.ProcAName:
			//var s []string
			l := len(ptree)
			writer.RawByte('[')
			for i, p := range ptree {
				CheckForQuotes(filepath.Base(p.Exe), writer)
				if i < (l - 1) {
					writer.RawByte(',')
				}
			}
			writer.RawByte(']')
		case engine.ProcAExe:
			l := len(ptree)
			writer.RawByte('[')
			for i, p := range ptree {
				CheckForQuotes(p.Exe, writer)
				if i < (l - 1) {
					writer.RawByte(',')
				}
			}
			writer.RawByte(']')
		case engine.ProcACmdLine:
			l := len(ptree)
			writer.RawByte('[')
			for i, p := range ptree {
				exe := trimBoundingQuotes(p.Exe)
				exeArgs := trimBoundingQuotes(p.ExeArgs)
				writer.RawByte('"')
				StringNoQuotes(exe, writer)
				if len(exeArgs) > 0 {
					writer.RawByte(' ')
					StringNoQuotes(exeArgs, writer)
				}
				writer.RawByte('"')
				if i < (l - 1) {
					writer.RawByte(',')
				}
			}
			writer.RawByte(']')
		case engine.ProcAPID:
			l := len(ptree)
			writer.RawByte('[')
			for i, p := range ptree {
				writer.Int64(p.Oid.Hpid)
				if i < (l - 1) {
					writer.RawByte(',')
				}
			}
			writer.RawByte(']')
		}
	}
}
