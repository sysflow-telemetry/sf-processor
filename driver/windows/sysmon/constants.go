package sysmon

const (
	cSysmonProcessCreate              = 1
	cSysmonNetworkConnection          = 3
	cSysmonProcessExit                = 5
	cSysmonLoadImage                  = 7
	cSysmonProcessAccess              = 10
	cSysmonSetRegistryValue           = 13
	cSysmonCreateDeleteRegistryObject = 12
	cSysmonFileCreated                = 11
	cSysmonPipeCreated                = 17
	cSysmonPipeConnected              = 18
	cEvtLogProvider                   = "Microsoft-Windows-Sysmon/Operational"

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

	cImageLoaded     = "ImageLoaded"
	cSigned          = "Signed"
	cSignature       = "Signature"
	cSignatureStatus = "SignatureStatus"

	cProtocol            = "Protocol"
	cInitiated           = "Initiated"
	cSourceIsIpv6        = "SourceIsIpv6"
	cSourceIP            = "SourceIp"
	cSourceHostname      = "SourceHostname"
	cSourcePort          = "SourcePort"
	cSourcePortName      = "SourcePortName"
	cDestinationIsIpv6   = "DestinationIsIpv6"
	cDestinationIP       = "DestinationIp"
	cDestinationHostname = "DestinationHostname"
	cDestinationPort     = "DestinationPort"
	cDestinationPortName = "DestinationPortName"

	cTargetFilename  = "TargetFilename"
	cCreationUtcTime = "CreationUtcTime"

	cTargetObject = "TargetObject"
	cDetails      = "Details"

	cEventType = "EventType"

	cTimeFormat = "2006-01-02 15:04:05.000"

	cHashRegex = "^SHA1=([A-Z0-9]+),MD5=([A-Z0-9]+),SHA256=([A-Z0-9]+),IMPHASH=([A-Z0-9]+)$"

	cDeleteValue = "DeleteValue"
	cSetValue    = "SetValue"
)
