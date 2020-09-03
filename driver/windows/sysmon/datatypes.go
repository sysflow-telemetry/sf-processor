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
