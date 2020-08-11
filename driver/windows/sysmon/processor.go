package sysmon

import (
	"fmt"
	"strconv"

	"github.com/elastic/beats/v7/winlogbeat/eventlog"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/sf-processor/core/cache"
	"github.ibm.com/sysflow/sf-processor/core/flattener"
)

// SMProcessor is an object for processing sysmon events and
// converting them into sysflow.
type SMProcessor struct {
	efrChan   chan *flattener.EnrichedFlatRecord
	procTable ProcessTable
	tables    *cache.SFTables
	converter *Converter
}

// NewSMProcessor instantiates a new SMProcessor object.
func NewSMProcessor(channel *flattener.EFRChannel) *SMProcessor {
	return &SMProcessor{
		efrChan:   channel.In,
		procTable: make(ProcessTable),
		tables:    cache.GetInstance(),
		converter: NewConverter(channel.In),
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
		s.converter.createSFProcEvent(val, ts,
			val.Process.Oid.Hpid, sfgo.OP_EXIT, 0)
	} else {
		fmt.Printf("Uh oh! Process not in process table %s %d\n", image, processID)
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
			procObj.LoginGUID = pairs.Value
		case cLogonID:
			procObj.LoginID = pairs.Value
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
	cmd, args := GetExeAndArgs(procObj.CommandLine)
	procObj.Process.Exe = cmd
	procObj.Process.ExeArgs = args
	procObj.Process.State = sfgo.SFObjectStateCREATED
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
		s.converter.createSFProcEvent(ppObj, record.TimeCreated.SystemTime.UnixNano(),
			ppObj.Process.Oid.Hpid, sfgo.OP_CLONE, int32(procObj.Process.Oid.Hpid))
		procExe := procObj.Process.Exe
		procExeArgs := procObj.Process.ExeArgs
		procObj.Process.Exe = ppObj.Process.Exe
		procObj.Process.ExeArgs = ppObj.Process.ExeArgs
		procObj.Process.Poid = createPOID(ppObj.Process.Oid)
		s.tables.SetProc(*procObj.Process.Oid, procObj.Process)
		s.converter.createSFProcEvent(procObj, record.TimeCreated.SystemTime.UnixNano(),
			procObj.Process.Oid.Hpid, sfgo.OP_CLONE, 0)
		procObj.Process.Exe = procExe
		procObj.Process.ExeArgs = procExeArgs
		if procObj.Process.Exe != ppObj.Process.Exe || procObj.Process.ExeArgs != ppObj.Process.ExeArgs {
			procObj.Process.State = sfgo.SFObjectStateMODIFIED
			s.tables.SetProc(*procObj.Process.Oid, procObj.Process)
			s.converter.createSFProcEvent(procObj, record.TimeCreated.SystemTime.UnixNano(),
				procObj.Process.Oid.Hpid, sfgo.OP_EXEC, 0)
		}
	} else {
		s.converter.createSFProcEvent(procObj, record.TimeCreated.SystemTime.UnixNano(),
			procObj.Process.Oid.Hpid, sfgo.OP_EXEC, 0)
	}
	s.procTable[procObj.GUID] = procObj
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
		default:
			s.printRecord(record)
		}
	}
	//event := record.XML
}
