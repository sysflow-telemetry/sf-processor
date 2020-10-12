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
	"strconv"

	"github.com/elastic/beats/v7/winlogbeat/eventlog"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/goutils/logger"
	"github.ibm.com/sysflow/sf-processor/core/cache"
	"github.ibm.com/sysflow/sf-processor/core/flattener"
)

// SMProcessor is an object for processing sysmon events and
// converting them into sysflow.
type SMProcessor struct {
	frChan    chan *sfgo.FlatRecord
	procTable ProcessTable
	tables    *cache.SFTables
	converter *Converter
	protoMap  map[string]int64
}

// NewSMProcessor instantiates a new SMProcessor object.
func NewSMProcessor(channel *flattener.FlatChannel) *SMProcessor {
	protoMap := map[string]int64{"tcp": 6, "udp": 17}
	return &SMProcessor{
		frChan:    channel.In,
		procTable: make(ProcessTable),
		tables:    cache.GetInstance(),
		converter: NewConverter(channel.In),
		protoMap:  protoMap,
	}
}

// GetProvider returns the name of the sysmon provider as a string
func (s *SMProcessor) GetProvider() string {
	return cEvtLogProvider
}

func (s *SMProcessor) createParentProcess(proc *ProcessObj) *ProcessObj {
	ppObj := NewProcessObj()
	ppObj.GUID = proc.ParentProcessGUID
	if n, err := strconv.ParseInt(proc.ParentProcessID, 10, 64); err == nil {
		ppObj.Process.Oid.Hpid = n
	}
	ppObj.Process.Ts = proc.Process.Ts
	ppObj.Process.State = sfgo.SFObjectStateCREATED
	ppObj.Image = proc.ParentProcessImage
	ppObj.CommandLine = proc.ParentCommandLine
	ppObj.Process.Tty = false
	ppObj.Process.Entry = (ppObj.Process.Oid.Hpid == 1)
	cmd, args := GetExeAndArgs(ppObj.CommandLine)
	ppObj.Process.Exe = cmd
	ppObj.Process.ExeArgs = args
	ppObj.Written = false
	s.procTable[ppObj.GUID] = ppObj
	return ppObj
}

/*EventID  {%!s(uint16=0) %!s(uint32=5)} Provider: {Microsoft-Windows-Sysmon {5770385f-c22a-43e0-bf4c-06f5698ffbd9} } Record ID: 384822  Computer windy2.sl.cloud9.ibm.com User SID Identifier[S-1-5-18] Name[SYSTEM] Domain[NT AUTHORITY] Type[Well Known Group] Time {2020-08-05 01:37:52.3868518 +0000 UTC}
EventData Type sys.EventData
{[{RuleName -} {UtcTime 2020-08-05 01:37:52.386} {ProcessGuid {8ce7f76f-0d70-5f2a-29a6-000000000f00}} {ProcessId 3540} {Image C:\Users\terylt_ibm.com\go\bin\go-outline.exe}]}*/
func (s *SMProcessor) processExited(record eventlog.Record) {
	var procGUID string
	var ts int64
	var image string
	var processID int64
	for _, pairs := range record.EventData.Pairs {
		switch pairs.Key {
		case cUtcTime:
			//fmt.Printf("UTC Time type: %T\n", pairs.Value)
			ts = GetTimestamp(pairs.Value)
		case cProcessGUID:
			procGUID = pairs.Value
		case cImage:
			image = pairs.Value
		case cProcessID:
			if n, err := strconv.ParseInt(pairs.Value, 10, 64); err == nil {
				processID = n
			}

		}
	}

	if val, ok := s.procTable[procGUID]; ok {
		s.tables.SetProc(*val.Process.Oid, val.Process)
		s.converter.createSFProcEvent(record, val, ts,
			val.Process.Oid.Hpid, sfgo.OP_EXIT, 0, nil, nil)
	} else {
		logger.Trace.Printf("Uh oh! Process not in process table for exit process %s %d", image, processID)
	}
}

/*
EventID  {%!s(uint16=0) %!s(uint32=1)} Provider: {Microsoft-Windows-Sysmon {5770385f-c22a-43e0-bf4c-06f5698ffbd9} } Record ID: 10505  Computer windy2.sl.cloud9.ibm.com User SID Identifier[S-1-5-18] Name[SYSTEM] Domain[NT AUTHORITY] Type[Well Known Group] Time {2020-07-30 19:41:14.2458567 +0000 UTC}
{[{RuleName -} {UtcTime 2020-07-30 19:41:14.236} {ProcessGuid {8ce7f76f-225a-5f23-320c-000000000e00}} {ProcessId 5140} {Image C:\Program Files\Git\usr\bin\bash.exe} {FileVersion -} {Description -} {Product -} {Company -} {OriginalFileName -} {CommandLine "C:\Program Files\Git\bin\..\usr\bin\bash.exe"} {CurrentDirectory C:\Users\tery
lt_ibm.com\go\src\github.ibm.com\sysflow\sf-apis\} {User AD-RES\terylt_ibm.com} {LogonGuid {8ce7f76f-16d6-5f23-c786-e90000000000}} {LogonId 0xe986c7} {TerminalSessionId 0} {IntegrityLevel High} {Hashes SHA1=363150831615BCE57EC9585223A17D771E8697EF,MD5=32275787C7C51D2310B8FE2FACF2A935,SHA256=744343E01351BA92E365B7E24EEDD4ED18ED3EBE26
E68C69D9B5E324FE64A1B5,IMPHASH=7358EF16984261EC8925E382CDDC1FB6} {ParentProcessGuid {8ce7f76f-225a-5f23-310c-000000000e00}} {ParentProcessId 5056} {ParentImage C:\Program Files\Git\usr\bin\bash.exe} {ParentCommandLine "C:\Program Files\Git\bin\..\usr\bin\bash.exe"}]}*/
func (s *SMProcessor) processCreated(record eventlog.Record) {
	procObj := NewProcessObj()
	for _, pairs := range record.EventData.Pairs {
		switch pairs.Key {
		case cUtcTime:
			//fmt.Printf("UTC Time type: %T\n", pairs.Value)
			procObj.Process.Oid.CreateTS = GetTimestamp(pairs.Value)
		case cProcessGUID:
			//fmt.Printf("ProcessGuid type: %T\n", pairs.Value)
			procObj.GUID = pairs.Value
		case cProcessID:
			if n, err := strconv.ParseInt(pairs.Value, 10, 64); err == nil {
				procObj.Process.Oid.Hpid = n
			}
		case cUser:
			procObj.Process.UserName = pairs.Value
		case cImage:
			procObj.Image = pairs.Value
		case cCurrentDirectory:
			procObj.CurrentDirectory = pairs.Value
		case cLogonGUID:
			procObj.LogonGUID = pairs.Value
		case cLogonID:
			procObj.LogonID = pairs.Value
		case cCommandLine:
			procObj.CommandLine = pairs.Value
		case cTerminalSessionID:
			procObj.TerminalSessionID = pairs.Value
		case cIntegrityLevel:
			procObj.Integrity = pairs.Value
		case cHashes:
			procObj.Hashes = pairs.Value
		case cParentProcessGUID:
			procObj.ParentProcessGUID = pairs.Value
		case cParentProcessID:
			procObj.ParentProcessID = pairs.Value
		case cParentImage:
			procObj.ParentProcessImage = pairs.Value
		case cParentCommandLine:
			procObj.ParentCommandLine = pairs.Value
		}

	}
	procObj.Process.Ts = record.TimeCreated.SystemTime.UnixNano()
	procObj.Process.Tty = false
	procObj.Process.Entry = (procObj.Process.Oid.Hpid == 1)
	_, args := GetExeAndArgs(procObj.CommandLine)
	procObj.Process.Exe = procObj.Image
	procObj.Process.ExeArgs = args
	procObj.Process.State = sfgo.SFObjectStateCREATED
	procObj.Signed = -1
	var ppObj *ProcessObj
	if len(procObj.ParentProcessGUID) > 0 {
		if val, ok := s.procTable[procObj.ParentProcessGUID]; ok {
			ppObj = val
		} else {
			ppObj = s.createParentProcess(procObj)
		}
		s.tables.SetProc(*ppObj.Process.Oid, ppObj.Process)
	}
	/*if ppObj != nil && !ppObj.Written {
		s.sysFlowChan <- createSFProcess(ppObj.Process)
		ppObj.Written = true
	}*/
	if ppObj != nil {
		s.converter.createSFProcEvent(record, ppObj, record.TimeCreated.SystemTime.UnixNano(),
			ppObj.Process.Oid.Hpid, sfgo.OP_CLONE, int32(procObj.Process.Oid.Hpid), nil, nil)
		procExe := procObj.Process.Exe
		procExeArgs := procObj.Process.ExeArgs
		procObj.Process.Exe = ppObj.Process.Exe
		procObj.Process.ExeArgs = ppObj.Process.ExeArgs
		procObj.Process.Poid = createPOID(ppObj.Process.Oid)
		s.tables.SetProc(*procObj.Process.Oid, procObj.Process)
		s.converter.createSFProcEvent(record, procObj, record.TimeCreated.SystemTime.UnixNano(),
			procObj.Process.Oid.Hpid, sfgo.OP_CLONE, 0, nil, nil)
		procObj.Process.Exe = procExe
		procObj.Process.ExeArgs = procExeArgs
		if procObj.Process.Exe != ppObj.Process.Exe || procObj.Process.ExeArgs != ppObj.Process.ExeArgs {
			procObj.Process.State = sfgo.SFObjectStateMODIFIED
			s.tables.SetProc(*procObj.Process.Oid, procObj.Process)
			s.converter.createSFProcEvent(record, procObj, record.TimeCreated.SystemTime.UnixNano(),
				procObj.Process.Oid.Hpid, sfgo.OP_EXEC, 0, nil, nil)
		}
	} else {
		s.converter.createSFProcEvent(record, procObj, record.TimeCreated.SystemTime.UnixNano(),
			procObj.Process.Oid.Hpid, sfgo.OP_EXEC, 0, nil, nil)
	}
	s.procTable[procObj.GUID] = procObj
}

//{[{RuleName -} {UtcTime 2020-08-05 01:23:09.526} {ProcessGuid {8ce7f76f-56ea-5f23-0100-000000000f00}} {ProcessId 4} {Image System}
//{User NT AUTHORITY\SYSTEM} {Protocol tcp} {Initiated true} {SourceIsIpv6 false} {SourceIp 10.191.226.105} {SourceHostname windy2.sl.cloud9.ibm.com} {SourcePort 63815} {SourcePortName -} {DestinationIsIpv
//6 false} {DestinationIp 10.162.185.211} {DestinationHostname -} {DestinationPort 445} {DestinationPortName microsoft-ds}]}
func (s *SMProcessor) createNetworkConnection(record eventlog.Record) {
	var procGUID string
	var ts int64
	var image string
	var processID int64
	var proto int64
	opFlags := sfgo.OP_CONNECT
	var sourceIP uint32 = 0
	var sourcePort int64
	var destIP uint32 = 0
	var destPort int64
	extNetworkAttrsStr := make([]string, sfgo.NUM_EXT_NET_STR)
	for _, pairs := range record.EventData.Pairs {
		switch pairs.Key {
		case cUtcTime:
			//fmt.Printf("UTC Time type: %T\n", pairs.Value)
			ts = GetTimestamp(pairs.Value)
		case cProcessGUID:
			procGUID = pairs.Value
		case cImage:
			image = pairs.Value
		case cProcessID:
			if n, err := strconv.ParseInt(pairs.Value, 10, 64); err == nil {
				processID = n
			} else {
				logger.Warn.Println("Unable to parse ProcessId sysmon attribute: " + err.Error())
			}
		case cProtocol:
			if prot, ok := s.protoMap[pairs.Value]; ok {
				proto = prot
			} else {
				proto = 0
			}
		case cInitiated:
			if b, err := strconv.ParseBool(pairs.Value); err == nil {
				if !b {
					opFlags = sfgo.OP_ACCEPT
				}
			} else {
				logger.Warn.Println("Unable to parse Initiated sysmon attribute: " + err.Error())
			}
		case cSourceIsIpv6:
			if b, err := strconv.ParseBool(pairs.Value); err == nil {
				if b {
					logger.Warn.Println("Do not currently support IPv6")
					return
				}
			} else {
				logger.Warn.Println("Unable to parse SourceIsIpv6 sysmon attribute: " + err.Error())
			}
		case cSourceIP:
			ip, err := ip2Int(pairs.Value)
			if err == nil {
				sourceIP = ip
			} else {
				logger.Warn.Println("Unable to parse SourceIp sysmon attribute: " + err.Error())
			}
		case cSourceHostname:
			extNetworkAttrsStr[sfgo.NET_SOURCE_HOST_NAME_STR] = pairs.Value
		case cSourcePort:
			if n, err := strconv.ParseInt(pairs.Value, 10, 64); err == nil {
				sourcePort = n
			} else {
				logger.Warn.Println("Unable to parse SourcePort sysmon attribute: " + err.Error())
			}
		case cSourcePortName:
			extNetworkAttrsStr[sfgo.NET_SOURCE_PORT_NAME_STR] = pairs.Value
		case cDestinationIsIpv6:
			if b, err := strconv.ParseBool(pairs.Value); err == nil {
				if b {
					logger.Warn.Println("Do not currently support IPv6")
					return
				}
			} else {
				logger.Warn.Println("Unable to parse SourceIsIpv6 sysmon attribute: " + err.Error())
			}

		case cDestinationIP:
			ip, err := ip2Int(pairs.Value)
			if err == nil {
				destIP = ip
			} else {
				logger.Warn.Println("Unable to parse SourceIp sysmon attribute: " + err.Error())
			}
		case cDestinationHostname:
			extNetworkAttrsStr[sfgo.NET_DEST_HOST_NAME_STR] = pairs.Value
		case cDestinationPort:
			if n, err := strconv.ParseInt(pairs.Value, 10, 64); err == nil {
				destPort = n
			} else {
				logger.Warn.Println("Unable to parse SourcePort sysmon attribute: " + err.Error())
			}
		case cDestinationPortName:
			extNetworkAttrsStr[sfgo.NET_DEST_PORT_NAME_STR] = pairs.Value
		}
	}
	if val, ok := s.procTable[procGUID]; ok {
		s.tables.SetProc(*val.Process.Oid, val.Process)
		s.converter.createSFNetworkFlow(record, val, ts, ts,
			val.Process.Oid.Hpid, opFlags, sourceIP, sourcePort, destIP, destPort, proto, extNetworkAttrsStr)
	} else {
		logger.Trace.Printf("Uh oh! Process not in process table for exit process %s %d", image, processID)
	}

}

func (s *SMProcessor) accessRemoteProcess(record eventlog.Record, evtID int) {
	var sourceProcGUID string
	var targetProcGUID string
	var ts int64
	var sourceImage string
	var targetImage string
	var sourceProcessID int64
	var targetProcessID int64
	var sourceThreadID int64
	intFields := make([]int64, sfgo.NUM_EXT_EVT_INT)
	strFields := make([]string, sfgo.NUM_EXT_EVT_STR)
	var procObj *ProcessObj

	for _, pairs := range record.EventData.Pairs {
		switch pairs.Key {
		case cUtcTime:
			//fmt.Printf("UTC Time type: %T\n", pairs.Value)
			ts = GetTimestamp(pairs.Value)
		case cSourceProcessGUID:
			sourceProcGUID = pairs.Value
		case cSourceImage:
			sourceImage = pairs.Value
		case cSourceProcessID:
			if n, err := strconv.ParseInt(pairs.Value, 10, 64); err == nil {
				sourceProcessID = n
			} else {
				logger.Warn.Println("Unable to parse ProcessId sysmon attribute: " + err.Error())
			}
		case cTargetProcessGUID:
			targetProcGUID = pairs.Value
		case cTargetImage:
			targetImage = pairs.Value
		case cTargetProcessID:
			if n, err := strconv.ParseInt(pairs.Value, 10, 64); err == nil {
				targetProcessID = n
			} else {
				logger.Warn.Println("Unable to parse ProcessId sysmon attribute: " + err.Error())
			}
		case cNewThreadID:
			if n, err := strconv.ParseInt(pairs.Value, 10, 64); err == nil {
				intFields[sfgo.EVT_TARG_PROC_NEW_THREAD_ID_INT] = n
			} else {
				logger.Warn.Println("Unable to parse ProcessId sysmon attribute: " + err.Error())
			}
		case cStartAddress:
			strFields[sfgo.EVT_TARG_PROC_START_ADDR_STR] = pairs.Value
		case cStartModule:
			strFields[sfgo.EVT_TARG_PROC_START_MODULE_STR] = pairs.Value
		case cStartFunction:
			strFields[sfgo.EVT_TARG_PROC_START_FUNCTION_STR] = pairs.Value
		case cSourceThreadID:
			if n, err := strconv.ParseInt(pairs.Value, 10, 64); err == nil {
				sourceThreadID = n
			} else {
				logger.Warn.Println("Unable to parse ProcessId sysmon attribute: " + err.Error())
			}
		case cGrantedAccess:
			strFields[sfgo.EVT_TARG_PROC_GRANT_ACCESS_STR] = pairs.Value
		case cCallTrace:
			strFields[sfgo.EVT_TARG_PROC_CALL_TRACE_STR] = pairs.Value
		}
	}
	if val, ok := s.procTable[sourceProcGUID]; ok {
		s.tables.SetProc(*val.Process.Oid, val.Process)
		procObj = val
	} else {
		logger.Trace.Printf("Uh oh! Process not in process table for access process %s %d", sourceImage, sourceProcessID)
		return
	}
	if val, ok := s.procTable[targetProcGUID]; ok {
		s.tables.SetProc(*val.Process.Oid, val.Process)
		s.converter.fillExtProcess(val, intFields, strFields)
	} else {
		logger.Trace.Printf("Uh oh! Process not in process table for load image %s %d", targetImage, targetProcessID)
	}
	if evtID == cSysmonProcessAccess {
		strFields[sfgo.EVT_TARG_PROC_ACCESS_TYPE_STR] = "AP"
		s.converter.createSFProcEvent(record, procObj, ts,
			sourceThreadID, sfgo.OP_PTRACE, 0, intFields, strFields)
	} else {
		strFields[sfgo.EVT_TARG_PROC_ACCESS_TYPE_STR] = "RT"
		s.converter.createSFProcEvent(record, procObj, ts,
			procObj.Process.Oid.Hpid, sfgo.OP_PTRACE, 0, intFields, strFields)
	}

}

func (s *SMProcessor) loadImage(record eventlog.Record) {
	var procGUID string
	var ts int64
	var image string
	var processID int64
	var imageLoaded string
	var hashes string
	var signed bool = false
	var signature string
	var sigStatus string
	for _, pairs := range record.EventData.Pairs {
		switch pairs.Key {
		case cUtcTime:
			//fmt.Printf("UTC Time type: %T\n", pairs.Value)
			ts = GetTimestamp(pairs.Value)
		case cProcessGUID:
			procGUID = pairs.Value
		case cImage:
			image = pairs.Value
		case cProcessID:
			if n, err := strconv.ParseInt(pairs.Value, 10, 64); err == nil {
				processID = n
			} else {
				logger.Warn.Println("Unable to parse ProcessId sysmon attribute: " + err.Error())
			}
		case cImageLoaded:
			imageLoaded = pairs.Value
		case cHashes:
			hashes = pairs.Value
		case cSigned:
			if b, err := strconv.ParseBool(pairs.Value); err == nil {
				signed = b
			} else {
				logger.Warn.Println("Unable to parse signed sysmon attribute: " + err.Error())
			}
		case cSignature:
			signature = pairs.Value
		case cSignatureStatus:
			sigStatus = pairs.Value
		}

	}

	if val, ok := s.procTable[procGUID]; ok {
		s.tables.SetProc(*val.Process.Oid, val.Process)
		s.converter.createSFFileFlow(record, val, ts, ts,
			val.Process.Oid.Hpid, sfgo.OP_OPEN|sfgo.OP_READ_RECV|sfgo.OP_CLOSE,
			imageLoaded, sfgo.O_RDONLY, signed, signature, sigStatus, 'i', hashes, "")
	} else {
		logger.Trace.Printf("Uh oh! Process not in process table for load image %s %d", image, processID)
	}
}

//{[{RuleName EXE} {UtcTime 2020-08-05 01:20:46.027} {ProcessGuid {8ce7f76f-096d-5f2a-2ea5-000000000f00}} {ProcessId 8836} {Image c:\go\pkg\tool\windows_amd64\link.exe} {TargetFilename C:\Users\terylt_ibm.com\AppData\Local\Temp\go-build916057974\b001\exe\a.out.exe} {CreationUtcTime 2020-08-05 01:20:46.027}]}
func (s *SMProcessor) createFile(record eventlog.Record) {
	var procGUID string
	var ts int64
	//var creationTS int64
	var image string
	var processID int64
	var fileName string
	for _, pairs := range record.EventData.Pairs {
		switch pairs.Key {
		case cUtcTime:
			//fmt.Printf("UTC Time type: %T\n", pairs.Value)
			ts = GetTimestamp(pairs.Value)
		case cProcessGUID:
			procGUID = pairs.Value
		case cImage:
			image = pairs.Value
		case cProcessID:
			if n, err := strconv.ParseInt(pairs.Value, 10, 64); err == nil {
				processID = n
			} else {
				logger.Warn.Println("Unable to parse ProcessId sysmon attribute: " + err.Error())
			}
		case cTargetFilename:
			fileName = pairs.Value
			//case cCreationUtcTime:
			//	creationTS = GetTimestamp(pairs.Value)
		}
	}
	if val, ok := s.procTable[procGUID]; ok {
		s.tables.SetProc(*val.Process.Oid, val.Process)
		s.converter.createSFFileFlow(record, val, ts, ts,
			val.Process.Oid.Hpid, sfgo.OP_OPEN,
			fileName, sfgo.O_CREAT, false, "", "", 'f', "", "")
	} else {
		logger.Trace.Printf("Uh oh! Process not in process table for create file %s %d", image, processID)
	}

}

//[{RuleName T1183,IFEO} {EventType DeleteValue} {UtcTime 2020-08-05 02:53:04.922} {ProcessGuid {8ce7f76f-56f2-5f23-2300-000000000f00}} {ProcessId 1852} {Image C:\Windows\system32\svchost.exe} {TargetObject HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\LSASS.exe\AuditLevel}]}
//{[{RuleName T1031,T1050} {EventType SetValue} {UtcTime 2020-08-05 01:54:56.049} {ProcessGuid {8ce7f76f-56ea-5f23-0100-000000000f00}} {ProcessId 4} {Image System} {TargetObject HKLM\SYSTEM\CrowdStrike\{36903b4a-6f88-46c6-a6f6-3a0de10f42b9}\{0000000e-0000-0000-0000-000000000000}\{00000000-000001a9}\start} {Details Binary Data}]}
func (s *SMProcessor) modifyRegistryValue(record eventlog.Record) {
	var procGUID string
	var ts int64
	//var creationTS int64
	var image string
	var processID int64
	var fileName string
	var eventType string
	var details string
	for _, pairs := range record.EventData.Pairs {
		switch pairs.Key {
		case cUtcTime:
			//fmt.Printf("UTC Time type: %T\n", pairs.Value)
			ts = GetTimestamp(pairs.Value)
		case cProcessGUID:
			procGUID = pairs.Value
		case cImage:
			image = pairs.Value
		case cProcessID:
			if n, err := strconv.ParseInt(pairs.Value, 10, 64); err == nil {
				processID = n
			} else {
				logger.Warn.Println("Unable to parse ProcessId sysmon attribute: " + err.Error())
			}
		case cEventType:
			eventType = pairs.Value
		case cTargetObject:
			fileName = pairs.Value
		case cDetails:
			details = pairs.Value
		}
	}
	if val, ok := s.procTable[procGUID]; ok {
		s.tables.SetProc(*val.Process.Oid, val.Process)
		switch eventType {
		case cSetValue:
			s.converter.createSFFileFlow(record, val, ts, ts,
				val.Process.Oid.Hpid, sfgo.OP_OPEN|sfgo.OP_WRITE_SEND|sfgo.OP_CLOSE,
				fileName, sfgo.O_WRONLY, false, "", "", 'r', "", details)
		case cDeleteValue:
			s.converter.createSFFileEvent(record, val, ts,
				val.Process.Oid.Hpid, sfgo.OP_UNLINK,
				fileName, false, "", "", 'r', "", details)

		default:
			logger.Trace.Println("Registry Event Type not supported: " + eventType)
		}
	} else {
		logger.Trace.Printf("Uh oh! Process not in process table for modify registry %s %d", image, processID)
	}
}

func (s *SMProcessor) printRecord(record eventlog.Record) {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("EventID  %s Provider: %s Record ID: %d  Computer %s User %s Time %s\n", record.EventIdentifier, record.Provider, record.RecordID, record.Computer, record.User, record.TimeCreated)
	fmt.Printf("EventData Type %T\n", record.EventData)
	fmt.Println(record.EventData)
	fmt.Println("--------------------------------------------------")
}

// Process analyzes a set of sysmon event logs and turns them into
// SysFlow records.
func (s *SMProcessor) Process(records []eventlog.Record) {

	for _, record := range records {
		switch record.EventIdentifier.ID {
		case cSysmonProcessCreate:
			//s.printRecord(record)
			s.processCreated(record)
		case cSysmonProcessExit:
			//s.printRecord(record)
			s.processExited(record)
		case cSysmonLoadImage:
			s.loadImage(record)
		case cSysmonNetworkConnection:
			s.createNetworkConnection(record)
		case cSysmonFileCreated:
			s.createFile(record)
		case cSysmonSetRegistryValue:
			s.modifyRegistryValue(record)
		case cSysmonCreateDeleteRegistryObject:
			s.modifyRegistryValue(record)
		case cSysmonProcessAccess:
			s.accessRemoteProcess(record, cSysmonProcessAccess)
		case cSysmonCreateRemoteThread:
			s.accessRemoteProcess(record, cSysmonCreateRemoteThread)
		case cSysmonPipeCreated:
		case cSysmonPipeConnected:
		default:
			s.printRecord(record)
		}
	}
	//event := record.XML
}
