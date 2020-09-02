package sysmon

import (
	"fmt"
	"regexp"

	"github.com/elastic/beats/v7/winlogbeat/eventlog"
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
	efr.Sources = make([]flattener.Source, 3)
	efr.Ints = make([][]int64, 3)
	efr.Strs = make([][]string, 3)
	efr.Sources[flattener.SYSFLOW_IDX] = flattener.SYSFLOW_SRC
	efr.Sources[flattener.PROC_IDX] = flattener.PROCESS_SRC
	efr.Sources[flattener.FILE_IDX] = flattener.FILE_SRC

	efr.Ints[flattener.SYSFLOW_IDX] = make([]int64, sfgo.INT_ARRAY_SIZE)
	efr.Strs[flattener.SYSFLOW_IDX] = make([]string, sfgo.STR_ARRAY_SIZE)

	efr.Ints[flattener.PROC_IDX] = make([]int64, flattener.NUM_EXT_PROC_ATTRS_INT)
	efr.Strs[flattener.PROC_IDX] = make([]string, flattener.NUM_EXT_PROC_ATTRS_STR)

	efr.Ints[flattener.FILE_IDX] = make([]int64, flattener.NUM_EXT_FILE_INT)
	efr.Strs[flattener.FILE_IDX] = make([]string, flattener.NUM_EXT_FILE_STR)
	return efr

}

// NewEnSysFlowFlatRecord returns a new sysflow flat record
func NewEnSysFlowFlatRecord() *flattener.EnrichedFlatRecord {
	efr := new(flattener.EnrichedFlatRecord)
	efr.Sources = append(efr.Sources, flattener.SYSFLOW_SRC)
	efr.Sources = append(efr.Sources, flattener.PROCESS_SRC)
	efr.Ints = make([][]int64, 2, 3)
	efr.Strs = make([][]string, 2, 3)
	efr.Ints[flattener.SYSFLOW_IDX] = make([]int64, sfgo.INT_ARRAY_SIZE)
	efr.Strs[flattener.SYSFLOW_IDX] = make([]string, sfgo.STR_ARRAY_SIZE)
	efr.Ints[flattener.PROC_IDX] = make([]int64, flattener.NUM_EXT_PROC_ATTRS_INT)
	efr.Strs[flattener.PROC_IDX] = make([]string, flattener.NUM_EXT_PROC_ATTRS_STR)
	return efr
}

func (s *Converter) fillExtProcess(procObj *ProcessObj, intFields []int64, strFields []string) {
	//Fill SysFlow Fields
	proc := procObj.Process
	if proc != nil {
		intFields[flattener.EVT_TARG_PROC_STATE_INT] = int64(proc.State)
		intFields[flattener.EVT_TARG_PROC_OID_CREATETS_INT] = int64(proc.Oid.CreateTS)
		intFields[flattener.EVT_TARG_PROC_OID_HPID_INT] = int64(proc.Oid.Hpid)
		if proc.Poid != nil && proc.Poid.UnionType == sfgo.UnionNullOIDTypeEnumOID {
			intFields[flattener.EVT_TARG_PROC_POID_CREATETS_INT] = proc.Poid.OID.CreateTS
			intFields[flattener.EVT_TARG_PROC_POID_HPID_INT] = proc.Poid.OID.Hpid
		} else {
			intFields[flattener.EVT_TARG_PROC_POID_CREATETS_INT] = sfgo.Zeros.Int64
			intFields[flattener.EVT_TARG_PROC_POID_HPID_INT] = sfgo.Zeros.Int64
		}
		intFields[flattener.EVT_TARG_PROC_TS_INT] = proc.Ts
		strFields[flattener.EVT_TARG_PROC_EXE_STR] = proc.Exe
		strFields[flattener.EVT_TARG_PROC_EXEARGS_STR] = proc.ExeArgs
		intFields[flattener.EVT_TARG_PROC_UID_INT] = int64(proc.Uid)
		strFields[flattener.EVT_TARG_PROC_USERNAME_STR] = proc.UserName
		intFields[flattener.EVT_TARG_PROC_GID_INT] = int64(proc.Gid)
		strFields[flattener.EVT_TARG_PROC_GROUPNAME_STR] = proc.GroupName
		if proc.Tty {
			intFields[flattener.EVT_TARG_PROC_TTY_INT] = 1
		} else {
			intFields[flattener.EVT_TARG_PROC_TTY_INT] = 0
		}
		if proc.Entry {
			intFields[flattener.EVT_TARG_PROC_ENTRY_INT] = 1
		} else {
			intFields[flattener.EVT_TARG_PROC_ENTRY_INT] = 0
		}
		if proc.ContainerId != nil && proc.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
			strFields[flattener.EVT_TARG_PROC_CONTAINERID_STRING_STR] = proc.ContainerId.String
		} else {
			strFields[flattener.EVT_TARG_PROC_CONTAINERID_STRING_STR] = sfgo.Zeros.String
		}
	} else {
		logger.Warn.Println("Event does not have a related process.  This should not happen.")
		intFields[flattener.EVT_TARG_PROC_STATE_INT] = sfgo.Zeros.Int64
		intFields[flattener.EVT_TARG_PROC_OID_CREATETS_INT] = sfgo.Zeros.Int64
		intFields[flattener.EVT_TARG_PROC_OID_HPID_INT] = sfgo.Zeros.Int64
		intFields[flattener.EVT_TARG_PROC_POID_CREATETS_INT] = sfgo.Zeros.Int64
		intFields[flattener.EVT_TARG_PROC_POID_HPID_INT] = sfgo.Zeros.Int64
		intFields[flattener.EVT_TARG_PROC_TS_INT] = sfgo.Zeros.Int64
		strFields[flattener.EVT_TARG_PROC_EXE_STR] = sfgo.Zeros.String
		strFields[flattener.EVT_TARG_PROC_EXEARGS_STR] = sfgo.Zeros.String
		intFields[flattener.EVT_TARG_PROC_UID_INT] = sfgo.Zeros.Int64
		strFields[flattener.EVT_TARG_PROC_USERNAME_STR] = sfgo.Zeros.String
		intFields[flattener.EVT_TARG_PROC_GID_INT] = sfgo.Zeros.Int64
		strFields[flattener.EVT_TARG_PROC_GROUPNAME_STR] = sfgo.Zeros.String
		intFields[flattener.EVT_TARG_PROC_TTY_INT] = sfgo.Zeros.Int64
		intFields[flattener.EVT_TARG_PROC_ENTRY_INT] = sfgo.Zeros.Int64
		strFields[flattener.EVT_TARG_PROC_CONTAINERID_STRING_STR] = sfgo.Zeros.String
	}

	strFields[flattener.EVT_TARG_PROC_GUID_STR] = procObj.GUID
	strFields[flattener.EVT_TARG_PROC_IMAGE_STR] = procObj.Image
	strFields[flattener.EVT_TARG_PROC_CURR_DIRECTORY_STR] = procObj.CurrentDirectory
	strFields[flattener.EVT_TARG_PROC_LOGON_GUID_STR] = procObj.LogonGUID
	strFields[flattener.EVT_TARG_PROC_LOGON_ID_STR] = procObj.LogonID
	strFields[flattener.EVT_TARG_PROC_TERMINAL_SESSION_ID_STR] = procObj.TerminalSessionID
	strFields[flattener.EVT_TARG_PROC_INTEGRITY_LEVEL_STR] = procObj.Integrity

	//Fill Hashing Fields
	hashes := s.hashParser.FindStringSubmatch(procObj.Hashes)
	if len(hashes) == 5 {
		strFields[flattener.EVT_TARG_PROC_SHA1_HASH_STR] = hashes[flattener.SHA1_HASH_STR+1]
		strFields[flattener.EVT_TARG_PROC_MD5_HASH_STR] = hashes[flattener.MD5_HASH_STR+1]
		strFields[flattener.EVT_TARG_PROC_SHA256_HASH_STR] = hashes[flattener.SHA256_HASH_STR+1]
		strFields[flattener.EVT_TARG_PROC_IMP_HASH_STR] = hashes[flattener.IMP_HASH_STR+1]
	}
	strFields[flattener.EVT_TARG_PROC_SIGNATURE_STR] = procObj.Signature
	strFields[flattener.EVT_TARG_PROC_SIGNATURE_STATUS_STR] = procObj.SignatureStatus
	intFields[flattener.EVT_TARG_PROC_SIGNED_INT] = procObj.Signed

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
	extStrFields := fr.Strs[flattener.PROC_IDX]
	extIntFields := fr.Ints[flattener.PROC_IDX]

	extStrFields[flattener.PROC_GUID_STR] = procObj.GUID
	extStrFields[flattener.PROC_IMAGE_STR] = procObj.Image
	extStrFields[flattener.PROC_CURR_DIRECTORY_STR] = procObj.CurrentDirectory
	extStrFields[flattener.PROC_LOGON_GUID_STR] = procObj.LogonGUID
	extStrFields[flattener.PROC_LOGON_ID_STR] = procObj.LogonID
	extStrFields[flattener.PROC_TERMINAL_SESSION_ID_STR] = procObj.TerminalSessionID
	extStrFields[flattener.PROC_INTEGRITY_LEVEL_STR] = procObj.Integrity

	//Fill Hashing Fields
	hashes := s.hashParser.FindStringSubmatch(procObj.Hashes)
	if len(hashes) == 5 {
		extStrFields[flattener.PROC_SHA1_HASH_STR] = hashes[flattener.SHA1_HASH_STR+1]
		extStrFields[flattener.PROC_MD5_HASH_STR] = hashes[flattener.MD5_HASH_STR+1]
		extStrFields[flattener.PROC_SHA256_HASH_STR] = hashes[flattener.SHA256_HASH_STR+1]
		extStrFields[flattener.PROC_IMP_HASH_STR] = hashes[flattener.IMP_HASH_STR+1]
	}
	extStrFields[flattener.PROC_SIGNATURE_STR] = procObj.Signature
	extStrFields[flattener.PROC_SIGNATURE_STATUS_STR] = procObj.SignatureStatus
	extIntFields[flattener.PROC_SIGNED_INT] = procObj.Signed

}

func (s *Converter) fillProcessEvent(efr *flattener.EnrichedFlatRecord, ts int64, tid int64, opFlags int32, ret int32, intFs []int64, strFs []string) {
	intFields := efr.Ints[flattener.SYSFLOW_IDX]
	intFields[sfgo.SF_REC_TYPE] = sfgo.PROC_EVT
	intFields[sfgo.EV_PROC_TS_INT] = ts
	intFields[sfgo.EV_PROC_TID_INT] = tid
	intFields[sfgo.EV_PROC_OPFLAGS_INT] = int64(opFlags)
	intFields[sfgo.EV_PROC_RET_INT] = int64(ret)

	if intFs != nil && strFs != nil {
		efr.Sources = append(efr.Sources, flattener.TARG_PROC_SRC)
		efr.Ints = append(efr.Ints, intFs)
		efr.Strs = append(efr.Strs, strFs)
	}

}

func (s *Converter) fillHeader(efr *flattener.EnrichedFlatRecord, record eventlog.Record) {
	intFields := efr.Ints[flattener.SYSFLOW_IDX]
	strFields := efr.Strs[flattener.SYSFLOW_IDX]
	intFields[sfgo.SFHE_VERSION_INT] = 0
	strFields[sfgo.SFHE_EXPORTER_STR] = record.Provider.Name
	strFields[sfgo.SFHE_IP_STR] = record.Computer
}

func (s *Converter) fillFile(efr *flattener.EnrichedFlatRecord, ts int64, fileName string, fileType byte, fileHashes string, signed bool, signature string, sigStatus string, details string) {
	intFields := efr.Ints[flattener.SYSFLOW_IDX]
	strFields := efr.Strs[flattener.SYSFLOW_IDX]
	intFields[sfgo.FILE_STATE_INT] = int64(sfgo.SFObjectStateCREATED)
	intFields[sfgo.FILE_TS_INT] = ts
	intFields[sfgo.FILE_RESTYPE_INT] = int64(fileType)
	strFields[sfgo.FILE_PATH_STR] = fileName

	hashStrFields := efr.Strs[flattener.FILE_IDX]
	hashIntFields := efr.Ints[flattener.FILE_IDX]
	//Fill Hashing Fields
	hashes := s.hashParser.FindStringSubmatch(fileHashes)
	if len(hashes) == 5 {
		hashStrFields[flattener.FILE_SHA1_HASH_STR] = hashes[flattener.SHA1_HASH_STR+1]
		hashStrFields[flattener.FILE_MD5_HASH_STR] = hashes[flattener.MD5_HASH_STR+1]
		hashStrFields[flattener.FILE_SHA256_HASH_STR] = hashes[flattener.SHA256_HASH_STR+1]
		hashStrFields[flattener.FILE_IMP_HASH_STR] = hashes[flattener.IMP_HASH_STR+1]
	}
	hashStrFields[flattener.FILE_SIGNATURE_STR] = signature
	hashStrFields[flattener.FILE_SIGNATURE_STATUS_STR] = sigStatus
	hashStrFields[flattener.FILE_DETAILS_STR] = details
	if signed {
		hashIntFields[flattener.FILE_SIGNED_INT] = 1
	} else {
		hashIntFields[flattener.FILE_SIGNED_INT] = 0
	}

}

func (s *Converter) createSFProcEvent(record eventlog.Record, procObj *ProcessObj, ts int64, tid int64, opFlags int32, ret int32, intFs []int64, strFs []string) {
	efr := NewEnSysFlowFlatRecord()
	s.fillHeader(efr, record)
	s.fillProcess(efr, procObj)
	s.fillProcessEvent(efr, ts, tid, opFlags, ret, intFs, strFs)
	s.efrChan <- efr
}

func (s *Converter) fillFileFlow(efr *flattener.EnrichedFlatRecord, ts int64, tid int64, opFlags int32, endTs int64, openFlags int32) {
	intFields := efr.Ints[flattener.SYSFLOW_IDX]
	intFields[sfgo.SF_REC_TYPE] = sfgo.FILE_FLOW
	intFields[sfgo.FL_FILE_TS_INT] = ts
	intFields[sfgo.FL_FILE_TID_INT] = tid
	intFields[sfgo.FL_FILE_OPFLAGS_INT] = int64(opFlags)
	intFields[sfgo.FL_FILE_OPENFLAGS_INT] = int64(openFlags)
	intFields[sfgo.FL_FILE_ENDTS_INT] = endTs
	intFields[sfgo.FL_FILE_FD_INT] = 0
	intFields[sfgo.FL_FILE_NUMRRECVOPS_INT] = 0
	intFields[sfgo.FL_FILE_NUMWSENDOPS_INT] = 0
	intFields[sfgo.FL_FILE_NUMRRECVBYTES_INT] = 0
	intFields[sfgo.FL_FILE_NUMWSENDBYTES_INT] = 0
}

func (s *Converter) createSFFileFlow(record eventlog.Record, procObj *ProcessObj, ts int64, endTs int64, tid int64, opFlags int32, fileName string, openFlags int32,
	signed bool, signature string, sigStatus string, fileType byte, fileHashes string, details string) {
	efr := NewEnrichedFlatRecord()
	s.fillHeader(efr, record)
	if fileType == 'i' && procObj.Process.Exe == fileName {
		fmt.Printf("Assigning signed, signature, and status values to process %s %s %s\n", procObj.Process.Exe, signature, sigStatus)
		procObj.Hashes = fileHashes
		if signed {
			procObj.Signed = 1
		} else {
			procObj.Signed = 0
		}
		procObj.Signature = signature
		procObj.SignatureStatus = sigStatus
	}
	s.fillProcess(efr, procObj)
	s.fillFile(efr, ts, fileName, fileType, fileHashes, signed, signature, sigStatus, details)
	s.fillFileFlow(efr, ts, tid, opFlags, endTs, openFlags)
	s.efrChan <- efr
}

func (s *Converter) fillFileEvt(efr *flattener.EnrichedFlatRecord, ts int64, tid int64, opFlags int32, ret int64) {
	intFields := efr.Ints[flattener.SYSFLOW_IDX]
	intFields[sfgo.SF_REC_TYPE] = sfgo.FILE_EVT
	intFields[sfgo.EV_FILE_TS_INT] = ts
	intFields[sfgo.EV_FILE_TID_INT] = tid
	intFields[sfgo.EV_FILE_OPFLAGS_INT] = int64(opFlags)
	intFields[sfgo.EV_FILE_RET_INT] = ret
}

func (s *Converter) createSFFileEvent(record eventlog.Record, procObj *ProcessObj, ts int64, tid int64, opFlags int32, fileName string,
	signed bool, signature string, sigStatus string, fileType byte, fileHashes string, details string) {
	efr := NewEnrichedFlatRecord()
	s.fillHeader(efr, record)
	s.fillProcess(efr, procObj)
	s.fillFile(efr, ts, fileName, fileType, fileHashes, signed, signature, sigStatus, details)
	s.fillFileEvt(efr, ts, tid, opFlags, 0)
	s.efrChan <- efr
}

func (s *Converter) fillNetFlow(efr *flattener.EnrichedFlatRecord, ts int64, endTs int64, tid int64, opFlags int, sip uint32, sport int64, dip uint32, dport int64, proto int64) {
	intFields := efr.Ints[flattener.SYSFLOW_IDX]
	intFields[sfgo.SF_REC_TYPE] = sfgo.NET_FLOW
	intFields[sfgo.FL_NETW_TS_INT] = ts
	intFields[sfgo.FL_NETW_TID_INT] = tid
	intFields[sfgo.FL_NETW_OPFLAGS_INT] = int64(opFlags)
	intFields[sfgo.FL_NETW_ENDTS_INT] = endTs
	intFields[sfgo.FL_NETW_SIP_INT] = int64(sip)
	intFields[sfgo.FL_NETW_SPORT_INT] = sport
	intFields[sfgo.FL_NETW_DIP_INT] = int64(dip)
	intFields[sfgo.FL_NETW_DPORT_INT] = dport
	intFields[sfgo.FL_NETW_PROTO_INT] = proto
	intFields[sfgo.FL_NETW_FD_INT] = 0
	intFields[sfgo.FL_NETW_NUMRRECVOPS_INT] = 0
	intFields[sfgo.FL_NETW_NUMWSENDOPS_INT] = 0
	intFields[sfgo.FL_NETW_NUMRRECVBYTES_INT] = 0
	intFields[sfgo.FL_NETW_NUMWSENDBYTES_INT] = 0
}

func (s *Converter) createSFNetworkFlow(record eventlog.Record, procObj *ProcessObj, ts int64, endTs int64, tid int64, opFlags int, sip uint32, sport int64, dip uint32, dport int64, proto int64, networkEnrich []string) {
	efr := NewEnSysFlowFlatRecord()
	s.fillHeader(efr, record)
	s.fillProcess(efr, procObj)
	s.fillNetFlow(efr, ts, endTs, tid, opFlags, sip, sport, dip, dport, proto)
	efr.Sources = append(efr.Sources, flattener.NETWORK_SRC)
	efr.Strs = append(efr.Strs, networkEnrich)
	efr.Ints = append(efr.Ints, nil)
	s.efrChan <- efr
}
