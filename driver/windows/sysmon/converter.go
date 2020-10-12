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
package sysmon

import (
	"fmt"
	"regexp"

	"github.com/elastic/beats/v7/winlogbeat/eventlog"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/goutils/logger"
)

// Converter converts SysmonAttributes into Flattened objects
type Converter struct {
	hashParser *regexp.Regexp
	frChan     chan *sfgo.FlatRecord
}

// NewConverter creates a new sysmon converter object
func NewConverter(frChan chan *sfgo.FlatRecord) *Converter {
	return &Converter{
		hashParser: regexp.MustCompile(cHashRegex),
		frChan:     frChan,
	}

}

// NewEnrichedFlatRecord returns a new enriched flat record
func NewEnrichedFlatRecord() *sfgo.FlatRecord {
	efr := new(sfgo.FlatRecord)
	efr.Sources = make([]sfgo.Source, 3)
	efr.Ints = make([][]int64, 3)
	efr.Strs = make([][]string, 3)
	efr.Sources[sfgo.SYSFLOW_IDX] = sfgo.SYSFLOW_SRC
	efr.Sources[sfgo.PROC_IDX] = sfgo.PROCESS_SRC
	efr.Sources[sfgo.FILE_IDX] = sfgo.FILE_SRC

	efr.Ints[sfgo.SYSFLOW_IDX] = make([]int64, sfgo.INT_ARRAY_SIZE)
	efr.Strs[sfgo.SYSFLOW_IDX] = make([]string, sfgo.STR_ARRAY_SIZE)

	efr.Ints[sfgo.PROC_IDX] = make([]int64, sfgo.NUM_EXT_PROC_ATTRS_INT)
	efr.Strs[sfgo.PROC_IDX] = make([]string, sfgo.NUM_EXT_PROC_ATTRS_STR)

	efr.Ints[sfgo.FILE_IDX] = make([]int64, sfgo.NUM_EXT_FILE_INT)
	efr.Strs[sfgo.FILE_IDX] = make([]string, sfgo.NUM_EXT_FILE_STR)
	return efr

}

// NewEnSysFlowFlatRecord returns a new sysflow flat record
func NewEnSysFlowFlatRecord() *sfgo.FlatRecord {
	efr := new(sfgo.FlatRecord)
	efr.Sources = append(efr.Sources, sfgo.SYSFLOW_SRC)
	efr.Sources = append(efr.Sources, sfgo.PROCESS_SRC)
	efr.Ints = make([][]int64, 2, 3)
	efr.Strs = make([][]string, 2, 3)
	efr.Ints[sfgo.SYSFLOW_IDX] = make([]int64, sfgo.INT_ARRAY_SIZE)
	efr.Strs[sfgo.SYSFLOW_IDX] = make([]string, sfgo.STR_ARRAY_SIZE)
	efr.Ints[sfgo.PROC_IDX] = make([]int64, sfgo.NUM_EXT_PROC_ATTRS_INT)
	efr.Strs[sfgo.PROC_IDX] = make([]string, sfgo.NUM_EXT_PROC_ATTRS_STR)
	return efr
}

func (s *Converter) fillExtProcess(procObj *ProcessObj, intFields []int64, strFields []string) {
	//Fill SysFlow Fields
	proc := procObj.Process
	if proc != nil {
		intFields[sfgo.EVT_TARG_PROC_STATE_INT] = int64(proc.State)
		intFields[sfgo.EVT_TARG_PROC_OID_CREATETS_INT] = int64(proc.Oid.CreateTS)
		intFields[sfgo.EVT_TARG_PROC_OID_HPID_INT] = int64(proc.Oid.Hpid)
		if proc.Poid != nil && proc.Poid.UnionType == sfgo.UnionNullOIDTypeEnumOID {
			intFields[sfgo.EVT_TARG_PROC_POID_CREATETS_INT] = proc.Poid.OID.CreateTS
			intFields[sfgo.EVT_TARG_PROC_POID_HPID_INT] = proc.Poid.OID.Hpid
		} else {
			intFields[sfgo.EVT_TARG_PROC_POID_CREATETS_INT] = sfgo.Zeros.Int64
			intFields[sfgo.EVT_TARG_PROC_POID_HPID_INT] = sfgo.Zeros.Int64
		}
		intFields[sfgo.EVT_TARG_PROC_TS_INT] = proc.Ts
		strFields[sfgo.EVT_TARG_PROC_EXE_STR] = proc.Exe
		strFields[sfgo.EVT_TARG_PROC_EXEARGS_STR] = proc.ExeArgs
		intFields[sfgo.EVT_TARG_PROC_UID_INT] = int64(proc.Uid)
		strFields[sfgo.EVT_TARG_PROC_USERNAME_STR] = proc.UserName
		intFields[sfgo.EVT_TARG_PROC_GID_INT] = int64(proc.Gid)
		strFields[sfgo.EVT_TARG_PROC_GROUPNAME_STR] = proc.GroupName
		if proc.Tty {
			intFields[sfgo.EVT_TARG_PROC_TTY_INT] = 1
		} else {
			intFields[sfgo.EVT_TARG_PROC_TTY_INT] = 0
		}
		if proc.Entry {
			intFields[sfgo.EVT_TARG_PROC_ENTRY_INT] = 1
		} else {
			intFields[sfgo.EVT_TARG_PROC_ENTRY_INT] = 0
		}
		if proc.ContainerId != nil && proc.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
			strFields[sfgo.EVT_TARG_PROC_CONTAINERID_STRING_STR] = proc.ContainerId.String
		} else {
			strFields[sfgo.EVT_TARG_PROC_CONTAINERID_STRING_STR] = sfgo.Zeros.String
		}
	} else {
		logger.Warn.Println("Event does not have a related process.  This should not happen.")
		intFields[sfgo.EVT_TARG_PROC_STATE_INT] = sfgo.Zeros.Int64
		intFields[sfgo.EVT_TARG_PROC_OID_CREATETS_INT] = sfgo.Zeros.Int64
		intFields[sfgo.EVT_TARG_PROC_OID_HPID_INT] = sfgo.Zeros.Int64
		intFields[sfgo.EVT_TARG_PROC_POID_CREATETS_INT] = sfgo.Zeros.Int64
		intFields[sfgo.EVT_TARG_PROC_POID_HPID_INT] = sfgo.Zeros.Int64
		intFields[sfgo.EVT_TARG_PROC_TS_INT] = sfgo.Zeros.Int64
		strFields[sfgo.EVT_TARG_PROC_EXE_STR] = sfgo.Zeros.String
		strFields[sfgo.EVT_TARG_PROC_EXEARGS_STR] = sfgo.Zeros.String
		intFields[sfgo.EVT_TARG_PROC_UID_INT] = sfgo.Zeros.Int64
		strFields[sfgo.EVT_TARG_PROC_USERNAME_STR] = sfgo.Zeros.String
		intFields[sfgo.EVT_TARG_PROC_GID_INT] = sfgo.Zeros.Int64
		strFields[sfgo.EVT_TARG_PROC_GROUPNAME_STR] = sfgo.Zeros.String
		intFields[sfgo.EVT_TARG_PROC_TTY_INT] = sfgo.Zeros.Int64
		intFields[sfgo.EVT_TARG_PROC_ENTRY_INT] = sfgo.Zeros.Int64
		strFields[sfgo.EVT_TARG_PROC_CONTAINERID_STRING_STR] = sfgo.Zeros.String
	}

	strFields[sfgo.EVT_TARG_PROC_GUID_STR] = procObj.GUID
	strFields[sfgo.EVT_TARG_PROC_IMAGE_STR] = procObj.Image
	strFields[sfgo.EVT_TARG_PROC_CURR_DIRECTORY_STR] = procObj.CurrentDirectory
	strFields[sfgo.EVT_TARG_PROC_LOGON_GUID_STR] = procObj.LogonGUID
	strFields[sfgo.EVT_TARG_PROC_LOGON_ID_STR] = procObj.LogonID
	strFields[sfgo.EVT_TARG_PROC_TERMINAL_SESSION_ID_STR] = procObj.TerminalSessionID
	strFields[sfgo.EVT_TARG_PROC_INTEGRITY_LEVEL_STR] = procObj.Integrity

	//Fill Hashing Fields
	hashes := s.hashParser.FindStringSubmatch(procObj.Hashes)
	if len(hashes) == 5 {
		strFields[sfgo.EVT_TARG_PROC_SHA1_HASH_STR] = hashes[sfgo.SHA1_HASH_STR+1]
		strFields[sfgo.EVT_TARG_PROC_MD5_HASH_STR] = hashes[sfgo.MD5_HASH_STR+1]
		strFields[sfgo.EVT_TARG_PROC_SHA256_HASH_STR] = hashes[sfgo.SHA256_HASH_STR+1]
		strFields[sfgo.EVT_TARG_PROC_IMP_HASH_STR] = hashes[sfgo.IMP_HASH_STR+1]
	}
	strFields[sfgo.EVT_TARG_PROC_SIGNATURE_STR] = procObj.Signature
	strFields[sfgo.EVT_TARG_PROC_SIGNATURE_STATUS_STR] = procObj.SignatureStatus
	intFields[sfgo.EVT_TARG_PROC_SIGNED_INT] = procObj.Signed

}

func (s *Converter) fillProcess(fr *sfgo.FlatRecord, procObj *ProcessObj) {
	//Fill SysFlow Fields
	intFields := fr.Ints[sfgo.SYSFLOW_IDX]
	strFields := fr.Strs[sfgo.SYSFLOW_IDX]
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
	extStrFields := fr.Strs[sfgo.PROC_IDX]
	extIntFields := fr.Ints[sfgo.PROC_IDX]

	extStrFields[sfgo.PROC_GUID_STR] = procObj.GUID
	extStrFields[sfgo.PROC_IMAGE_STR] = procObj.Image
	extStrFields[sfgo.PROC_CURR_DIRECTORY_STR] = procObj.CurrentDirectory
	extStrFields[sfgo.PROC_LOGON_GUID_STR] = procObj.LogonGUID
	extStrFields[sfgo.PROC_LOGON_ID_STR] = procObj.LogonID
	extStrFields[sfgo.PROC_TERMINAL_SESSION_ID_STR] = procObj.TerminalSessionID
	extStrFields[sfgo.PROC_INTEGRITY_LEVEL_STR] = procObj.Integrity

	//Fill Hashing Fields
	hashes := s.hashParser.FindStringSubmatch(procObj.Hashes)
	if len(hashes) == 5 {
		extStrFields[sfgo.PROC_SHA1_HASH_STR] = hashes[sfgo.SHA1_HASH_STR+1]
		extStrFields[sfgo.PROC_MD5_HASH_STR] = hashes[sfgo.MD5_HASH_STR+1]
		extStrFields[sfgo.PROC_SHA256_HASH_STR] = hashes[sfgo.SHA256_HASH_STR+1]
		extStrFields[sfgo.PROC_IMP_HASH_STR] = hashes[sfgo.IMP_HASH_STR+1]
	}
	extStrFields[sfgo.PROC_SIGNATURE_STR] = procObj.Signature
	extStrFields[sfgo.PROC_SIGNATURE_STATUS_STR] = procObj.SignatureStatus
	extIntFields[sfgo.PROC_SIGNED_INT] = procObj.Signed

}

func (s *Converter) fillProcessEvent(efr *sfgo.FlatRecord, ts int64, tid int64, opFlags int32, ret int32, intFs []int64, strFs []string) {
	intFields := efr.Ints[sfgo.SYSFLOW_IDX]
	intFields[sfgo.SF_REC_TYPE] = sfgo.PROC_EVT
	intFields[sfgo.EV_PROC_TS_INT] = ts
	intFields[sfgo.EV_PROC_TID_INT] = tid
	intFields[sfgo.EV_PROC_OPFLAGS_INT] = int64(opFlags)
	intFields[sfgo.EV_PROC_RET_INT] = int64(ret)

	if intFs != nil && strFs != nil {
		efr.Sources = append(efr.Sources, sfgo.TARG_PROC_SRC)
		efr.Ints = append(efr.Ints, intFs)
		efr.Strs = append(efr.Strs, strFs)
	}

}

func (s *Converter) fillHeader(efr *sfgo.FlatRecord, record eventlog.Record) {
	intFields := efr.Ints[sfgo.SYSFLOW_IDX]
	strFields := efr.Strs[sfgo.SYSFLOW_IDX]
	intFields[sfgo.SFHE_VERSION_INT] = 0
	strFields[sfgo.SFHE_EXPORTER_STR] = record.Provider.Name
	strFields[sfgo.SFHE_IP_STR] = record.Computer
}

func (s *Converter) fillFile(efr *sfgo.FlatRecord, ts int64, fileName string, fileType byte, fileHashes string, signed bool, signature string, sigStatus string, details string) {
	intFields := efr.Ints[sfgo.SYSFLOW_IDX]
	strFields := efr.Strs[sfgo.SYSFLOW_IDX]
	intFields[sfgo.FILE_STATE_INT] = int64(sfgo.SFObjectStateCREATED)
	intFields[sfgo.FILE_TS_INT] = ts
	intFields[sfgo.FILE_RESTYPE_INT] = int64(fileType)
	strFields[sfgo.FILE_PATH_STR] = fileName

	hashStrFields := efr.Strs[sfgo.FILE_IDX]
	hashIntFields := efr.Ints[sfgo.FILE_IDX]
	//Fill Hashing Fields
	hashes := s.hashParser.FindStringSubmatch(fileHashes)
	if len(hashes) == 5 {
		hashStrFields[sfgo.FILE_SHA1_HASH_STR] = hashes[sfgo.SHA1_HASH_STR+1]
		hashStrFields[sfgo.FILE_MD5_HASH_STR] = hashes[sfgo.MD5_HASH_STR+1]
		hashStrFields[sfgo.FILE_SHA256_HASH_STR] = hashes[sfgo.SHA256_HASH_STR+1]
		hashStrFields[sfgo.FILE_IMP_HASH_STR] = hashes[sfgo.IMP_HASH_STR+1]
	}
	hashStrFields[sfgo.FILE_SIGNATURE_STR] = signature
	hashStrFields[sfgo.FILE_SIGNATURE_STATUS_STR] = sigStatus
	hashStrFields[sfgo.FILE_DETAILS_STR] = details
	if signed {
		hashIntFields[sfgo.FILE_SIGNED_INT] = 1
	} else {
		hashIntFields[sfgo.FILE_SIGNED_INT] = 0
	}

}

func (s *Converter) createSFProcEvent(record eventlog.Record, procObj *ProcessObj, ts int64, tid int64, opFlags int32, ret int32, intFs []int64, strFs []string) {
	efr := NewEnSysFlowFlatRecord()
	s.fillHeader(efr, record)
	s.fillProcess(efr, procObj)
	s.fillProcessEvent(efr, ts, tid, opFlags, ret, intFs, strFs)
	s.frChan <- efr
}

func (s *Converter) fillFileFlow(efr *sfgo.FlatRecord, ts int64, tid int64, opFlags int32, endTs int64, openFlags int32) {
	intFields := efr.Ints[sfgo.SYSFLOW_IDX]
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
	s.frChan <- efr
}

func (s *Converter) fillFileEvt(efr *sfgo.FlatRecord, ts int64, tid int64, opFlags int32, ret int64) {
	intFields := efr.Ints[sfgo.SYSFLOW_IDX]
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
	s.frChan <- efr
}

func (s *Converter) fillNetFlow(efr *sfgo.FlatRecord, ts int64, endTs int64, tid int64, opFlags int, sip uint32, sport int64, dip uint32, dport int64, proto int64) {
	intFields := efr.Ints[sfgo.SYSFLOW_IDX]
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
	efr.Sources = append(efr.Sources, sfgo.NETWORK_SRC)
	efr.Strs = append(efr.Strs, networkEnrich)
	efr.Ints = append(efr.Ints, nil)
	s.frChan <- efr
}
