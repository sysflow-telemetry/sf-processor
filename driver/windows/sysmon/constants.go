package sysmon

const (
	cSysmonProcessCreate = 1
	cSysmonProcessExit   = 5
	cEvtLogProvider      = "Microsoft-Windows-Sysmon/Operational"

	cUtcTime           = "UtcTime"
	cProcessGUID       = "ProcessGuid"
	cProcessID         = "ProcessId"
	cUser              = "User"
	cImage             = "Image"
	cCommandLine       = "CommandLine"
	cCurrentDirectory  = "CurrentDirectory"
	cLogonGUID         = "LogonGuid"
	cLogonID           = "LogonId"
	cTerminalSessionID = "TerminalSessionId"
	cIntegrityLevel    = "IntegrityLevel"
	cHashes            = "Hashes"
	cParentProcessGUID = "ParentProcessGuid"
	cParentProcessID   = "ParentProcessId"
	cParentImage       = "ParentImage"
	cParentCommandLine = "ParentCommandLine"

	cTimeFormat = "2006-01-02 15:04:05.000"

	cHashRegex = "^SHA1=([A-Z0-9]+),MD5=([A-Z0-9]+),SHA256=([A-Z0-9]+),IMPHASH=([A-Z0-9]+)$"
)
