package sysmon

import (
	"regexp"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/goutils/logger"
	"github.ibm.com/sysflow/sf-processor/core/flattener"
)

// Converter converts SysmonAttributes into Flattened objects
type Converter struct {
	hashParser *regexp.Regexp
	efrChan    chan *flattener.EnrichedFlatRecord
}

//  NewConverter creates a new sysmon converter object
func NewConverter(efrChan chan *flattener.EnrichedFlatRecord) *Converter {
	return &Converter{
		hashParser: regexp.MustCompile(cHashRegex),
		efrChan:    efrChan,
	}

}

// NewEnrichedFlatRecord returns a new enriched flat record
func NewEnrichedFlatRecord() *flattener.EnrichedFlatRecord {
	efr := new(flattener.EnrichedFlatRecord)
	efr.Sources = make([]int64, 3)
	efr.Ints = make([][]int64, 3)
	efr.Strs = make([][]string, 3)
	efr.Sources[flattener.SYSFLOW_IDX] = flattener.SYSFLOW_IDX
	efr.Sources[flattener.EXT_WIN_IDX] = flattener.EXT_WIN_IDX
	efr.Sources[flattener.HASH_IDX] = flattener.HASH_IDX

	efr.Ints[flattener.SYSFLOW_IDX] = make([]int64, sfgo.INT_ARRAY_SIZE)
	efr.Strs[flattener.SYSFLOW_IDX] = make([]string, sfgo.STR_ARRAY_SIZE)

	efr.Ints[flattener.EXT_WIN_IDX] = nil
	efr.Strs[flattener.EXT_WIN_IDX] = make([]string, flattener.NUM_EXT_PROC_ATTRS)

	efr.Ints[flattener.HASH_IDX] = nil
	efr.Strs[flattener.HASH_IDX] = make([]string, flattener.PROC_IMP_HASH)
	return efr

}

func (s *Converter) fillProcess(fr *flattener.EnrichedFlatRecord, procObj *ProcessObj) {
	//Fill SysFlow Fields
	intFields := fr.Ints[flattener.SYSFLOW_IDX]
	strFields := fr.Strs[flattener.SYSFLOW_IDX]
	proc := procObj.Process
	if proc != nil {
		intFields[sfgo.PROC_STATE_INT] = int64(proc.State)
		intFields[sfgo.PROC_OID_CREATETS_INT] = int64(proc.Oid.CreateTS)
		intFields[sfgo.PROC_OID_HPID_INT] = int64(proc.Oid.Hpid)
		if proc.Poid != nil && proc.Poid.UnionType == sfgo.UnionNullOIDTypeEnumOID {
			intFields[sfgo.PROC_POID_CREATETS_INT] = proc.Poid.OID.CreateTS
			intFields[sfgo.PROC_POID_HPID_INT] = proc.Poid.OID.Hpid
		} else {
			intFields[sfgo.PROC_POID_CREATETS_INT] = sfgo.Zeros.Int64
			intFields[sfgo.PROC_POID_HPID_INT] = sfgo.Zeros.Int64
		}
		intFields[sfgo.PROC_TS_INT] = proc.Ts
		strFields[sfgo.PROC_EXE_STR] = proc.Exe
		strFields[sfgo.PROC_EXEARGS_STR] = proc.ExeArgs
		intFields[sfgo.PROC_UID_INT] = int64(proc.Uid)
		strFields[sfgo.PROC_USERNAME_STR] = proc.UserName
		intFields[sfgo.PROC_GID_INT] = int64(proc.Gid)
		strFields[sfgo.PROC_GROUPNAME_STR] = proc.GroupName
		if proc.Tty {
			intFields[sfgo.PROC_TTY_INT] = 1
		} else {
			intFields[sfgo.PROC_TTY_INT] = 0
		}
		if proc.Entry {
			intFields[sfgo.PROC_ENTRY_INT] = 1
		} else {
			intFields[sfgo.PROC_ENTRY_INT] = 0
		}
		if proc.ContainerId != nil && proc.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
			strFields[sfgo.PROC_CONTAINERID_STRING_STR] = proc.ContainerId.String
		} else {
			strFields[sfgo.PROC_CONTAINERID_STRING_STR] = sfgo.Zeros.String
		}
	} else {
		logger.Warn.Println("Event does not have a related process.  This should not happen.")
		intFields[sfgo.PROC_STATE_INT] = sfgo.Zeros.Int64
		intFields[sfgo.PROC_OID_CREATETS_INT] = sfgo.Zeros.Int64
		intFields[sfgo.PROC_OID_HPID_INT] = sfgo.Zeros.Int64
		intFields[sfgo.PROC_POID_CREATETS_INT] = sfgo.Zeros.Int64
		intFields[sfgo.PROC_POID_HPID_INT] = sfgo.Zeros.Int64
		intFields[sfgo.PROC_TS_INT] = sfgo.Zeros.Int64
		strFields[sfgo.PROC_EXE_STR] = sfgo.Zeros.String
		strFields[sfgo.PROC_EXEARGS_STR] = sfgo.Zeros.String
		intFields[sfgo.PROC_UID_INT] = sfgo.Zeros.Int64
		strFields[sfgo.PROC_USERNAME_STR] = sfgo.Zeros.String
		intFields[sfgo.PROC_GID_INT] = sfgo.Zeros.Int64
		strFields[sfgo.PROC_GROUPNAME_STR] = sfgo.Zeros.String
		intFields[sfgo.PROC_TTY_INT] = sfgo.Zeros.Int64
		intFields[sfgo.PROC_ENTRY_INT] = sfgo.Zeros.Int64
		strFields[sfgo.PROC_CONTAINERID_STRING_STR] = sfgo.Zeros.String
	}

	//Fill Extended Windows Fields
	//extIntFields := fr.Ints[1]
	extStrFields := fr.Strs[flattener.EXT_WIN_IDX]

	extStrFields[flattener.PROC_GUID] = procObj.GUID
	extStrFields[flattener.PROC_IMAGE] = procObj.Image
	extStrFields[flattener.PROC_CURR_DIRECTORY] = procObj.CurrentDirectory
	extStrFields[flattener.PROC_LOGIN_GUID] = procObj.LoginGUID
	extStrFields[flattener.PROC_LOGIN_ID] = procObj.LoginID
	extStrFields[flattener.PROC_TERMINAL_SESSION_ID] = procObj.TerminalSessionID
	extStrFields[flattener.PROC_INTEGRITY_LEVEL] = procObj.Integrity

	//Fill Hashing Fields
	hashStrFields := fr.Strs[flattener.HASH_IDX]
	hashes := s.hashParser.FindStringSubmatch(procObj.Hashes)
	if len(hashes) == 5 {
		hashStrFields[flattener.PROC_SHA1_HASH] = hashes[flattener.PROC_SHA1_HASH+1]
		hashStrFields[flattener.PROC_MD5_HASH] = hashes[flattener.PROC_MD5_HASH+1]
		hashStrFields[flattener.PROC_SHA256_HASH] = hashes[flattener.PROC_SHA256_HASH+1]
		hashStrFields[flattener.PROC_IMP_HASH] = hashes[flattener.PROC_IMP_HASH+1]
	}

}

func (s *Converter) fillProcessEvent(efr *flattener.EnrichedFlatRecord, ts int64, tid int64, opFlags int32, ret int32) {
	intFields := efr.Ints[flattener.SYSFLOW_IDX]
	intFields[sfgo.SF_REC_TYPE] = sfgo.PROC_EVT
	intFields[sfgo.EV_PROC_TS_INT] = ts
	intFields[sfgo.EV_PROC_TID_INT] = tid
	intFields[sfgo.EV_PROC_OPFLAGS_INT] = int64(opFlags)
	intFields[sfgo.EV_PROC_RET_INT] = int64(ret)
}

func (s *Converter) createSFProcEvent(procObj *ProcessObj, ts int64, tid int64, opFlags int32, ret int32) {
	efr := NewEnrichedFlatRecord()
	s.fillProcess(efr, procObj)
	s.fillProcessEvent(efr, ts, tid, opFlags, ret)
	s.efrChan <- efr
}

/*
func (g *Flattener) fillEntities(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, file *sfgo.File, fr *sfgo.FlatRecord) {
	if hdr != nil {
		fr.Ints[sfgo.SFHE_VERSION_INT] = hdr.Version
		fr.Strs[sfgo.SFHE_EXPORTER_STR] = hdr.Exporter
		fr.Strs[sfgo.SFHE_IP_STR] = hdr.Ip
	} else {
		logger.Warn.Println("Event does not have a related header.  This should not happen.")
		fr.Ints[sfgo.SFHE_VERSION_INT] = sfgo.Zeros.Int64
		fr.Strs[sfgo.SFHE_EXPORTER_STR] = sfgo.Zeros.String
		fr.Strs[sfgo.SFHE_IP_STR] = sfgo.Zeros.String
	}
	if cont != nil {
		fr.Strs[sfgo.CONT_ID_STR] = cont.Id
		fr.Strs[sfgo.CONT_NAME_STR] = cont.Name
		fr.Strs[sfgo.CONT_IMAGE_STR] = cont.Image
		fr.Strs[sfgo.CONT_IMAGEID_STR] = cont.Imageid
		fr.Ints[sfgo.CONT_TYPE_INT] = int64(cont.Type)
		if cont.Privileged {
			fr.Ints[sfgo.CONT_PRIVILEGED_INT] = 1
		} else {
			fr.Ints[sfgo.CONT_PRIVILEGED_INT] = 0
		}
	} else {
		fr.Strs[sfgo.CONT_ID_STR] = sfgo.Zeros.String
		fr.Strs[sfgo.CONT_NAME_STR] = sfgo.Zeros.String
		fr.Strs[sfgo.CONT_IMAGE_STR] = sfgo.Zeros.String
		fr.Strs[sfgo.CONT_IMAGEID_STR] = sfgo.Zeros.String
		fr.Ints[sfgo.CONT_TYPE_INT] = sfgo.Zeros.Int64
		fr.Ints[sfgo.CONT_PRIVILEGED_INT] = sfgo.Zeros.Int64

	}
	if proc != nil {
		fr.Ints[sfgo.PROC_STATE_INT] = int64(proc.State)
		fr.Ints[sfgo.PROC_OID_CREATETS_INT] = int64(proc.Oid.CreateTS)
		fr.Ints[sfgo.PROC_OID_HPID_INT] = int64(proc.Oid.Hpid)
		if proc.Poid != nil && proc.Poid.UnionType == sfgo.UnionNullOIDTypeEnumOID {
			fr.Ints[sfgo.PROC_POID_CREATETS_INT] = proc.Poid.OID.CreateTS
			fr.Ints[sfgo.PROC_POID_HPID_INT] = proc.Poid.OID.Hpid
		} else {
			fr.Ints[sfgo.PROC_POID_CREATETS_INT] = sfgo.Zeros.Int64
			fr.Ints[sfgo.PROC_POID_HPID_INT] = sfgo.Zeros.Int64
		}
		fr.Ints[sfgo.PROC_TS_INT] = proc.Ts
		fr.Strs[sfgo.PROC_EXE_STR] = proc.Exe
		fr.Strs[sfgo.PROC_EXEARGS_STR] = proc.ExeArgs
		fr.Ints[sfgo.PROC_UID_INT] = int64(proc.Uid)
		fr.Strs[sfgo.PROC_USERNAME_STR] = proc.UserName
		fr.Ints[sfgo.PROC_GID_INT] = int64(proc.Gid)
		fr.Strs[sfgo.PROC_GROUPNAME_STR] = proc.GroupName
		if proc.Tty {
			fr.Ints[sfgo.PROC_TTY_INT] = 1
		} else {
			fr.Ints[sfgo.PROC_TTY_INT] = 0
		}
		if proc.Entry {
			fr.Ints[sfgo.PROC_ENTRY_INT] = 1
		} else {
			fr.Ints[sfgo.PROC_ENTRY_INT] = 0
		}
		if proc.ContainerId != nil && proc.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
			fr.Strs[sfgo.PROC_CONTAINERID_STRING_STR] = proc.ContainerId.String
		} else {
			fr.Strs[sfgo.PROC_CONTAINERID_STRING_STR] = sfgo.Zeros.String
		}
	} else {
		logger.Warn.Println("Event does not have a related process.  This should not happen.")
		fr.Ints[sfgo.PROC_STATE_INT] = sfgo.Zeros.Int64
		fr.Ints[sfgo.PROC_OID_CREATETS_INT] = sfgo.Zeros.Int64
		fr.Ints[sfgo.PROC_OID_HPID_INT] = sfgo.Zeros.Int64
		fr.Ints[sfgo.PROC_POID_CREATETS_INT] = sfgo.Zeros.Int64
		fr.Ints[sfgo.PROC_POID_HPID_INT] = sfgo.Zeros.Int64
		fr.Ints[sfgo.PROC_TS_INT] = sfgo.Zeros.Int64
		fr.Strs[sfgo.PROC_EXE_STR] = sfgo.Zeros.String
		fr.Strs[sfgo.PROC_EXEARGS_STR] = sfgo.Zeros.String
		fr.Ints[sfgo.PROC_UID_INT] = sfgo.Zeros.Int64
		fr.Strs[sfgo.PROC_USERNAME_STR] = sfgo.Zeros.String
		fr.Ints[sfgo.PROC_GID_INT] = sfgo.Zeros.Int64
		fr.Strs[sfgo.PROC_GROUPNAME_STR] = sfgo.Zeros.String
		fr.Ints[sfgo.PROC_TTY_INT] = sfgo.Zeros.Int64
		fr.Ints[sfgo.PROC_ENTRY_INT] = sfgo.Zeros.Int64
		fr.Strs[sfgo.PROC_CONTAINERID_STRING_STR] = sfgo.Zeros.String
	}
	if file != nil {
		fr.Ints[sfgo.FILE_STATE_INT] = int64(file.State)
		fr.Ints[sfgo.FILE_TS_INT] = file.Ts
		fr.Ints[sfgo.FILE_RESTYPE_INT] = int64(file.Restype)
		fr.Strs[sfgo.FILE_PATH_STR] = file.Path
		if file.ContainerId != nil && file.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
			fr.Strs[sfgo.FILE_CONTAINERID_STRING_STR] = file.ContainerId.String
		} else {
			fr.Strs[sfgo.FILE_CONTAINERID_STRING_STR] = sfgo.Zeros.String
		}
	} else {
		fr.Ints[sfgo.FILE_STATE_INT] = sfgo.Zeros.Int64
		fr.Ints[sfgo.FILE_TS_INT] = sfgo.Zeros.Int64
		fr.Ints[sfgo.FILE_RESTYPE_INT] = sfgo.Zeros.Int64
		fr.Strs[sfgo.FILE_PATH_STR] = sfgo.Zeros.String
		fr.Strs[sfgo.FILE_CONTAINERID_STRING_STR] = sfgo.Zeros.String
	}
}*/
