package sysmon

import (
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

// ProcessObj contains all information about a Windows process.
type ProcessObj struct {
	Process            *sfgo.Process
	GUID               string
	Image              string
	CurrentDirectory   string
	CommandLine        string
	LogonGUID          string
	LogonID            string
	TerminalSessionID  string
	Integrity          string
	Hashes             string
	ParentProcessGUID  string
	ParentProcessID    string
	ParentProcessImage string
	ParentCommandLine  string
	Signature          string
	SignatureStatus    string
	Signed             int64
	Written            bool
}

// NewProcessObj creates a new ProcessObj
func NewProcessObj() *ProcessObj {
	p := &ProcessObj{
		Process: sfgo.NewProcess(),
		Written: false,
	}
	p.Process.Oid = sfgo.NewOID()
	return p

}

// ProcessTable stores process objects by their GUID
type ProcessTable map[string]*ProcessObj
