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
package main

import (
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

const (
	handlerName string = "print"
)

// Handler exports a symbol for this plugin.
var Handler Printer

// Printer defines the main class for the flatterner plugin.
type Printer struct {
}

// NewPrinter creates a new Printer instance.
func NewPrinter() plugins.SFHandler {
	return new(Printer)
}

// RegisterChannel registers channels to plugin cache.
func (s *Printer) RegisterChannel(pc plugins.SFPluginCache) {
}

// RegisterHandler registers handler to handler cache.
func (s *Printer) RegisterHandler(hc plugins.SFHandlerCache) {
	hc.AddHandler(handlerName, NewPrinter)
}

// Init initializes the handler with a configuration map.
func (s *Printer) Init(conf map[string]interface{}) error {
	return nil
}

// IsEntityEnabled is used to check if the flattener returns entity records.
func (s *Printer) IsEntityEnabled() bool {
	return false
}

// SetOutChan sets the plugin output channel.
func (s *Printer) SetOutChan(chObj []interface{}) {
}

// Cleanup tears down resources.
func (s *Printer) Cleanup() {
	logger.Trace.Println("Calling Cleanup on Printer channel")
}

// HandleHeader processes Header entities.
func (s *Printer) HandleHeader(sf *plugins.CtxSysFlow, hdr *sfgo.SFHeader) error {
	return nil
}

// HandleContainer processes Container entities.
func (s *Printer) HandleContainer(sf *plugins.CtxSysFlow, cont *sfgo.Container) error {
	return nil
}

// HandleProcess processes Process entities.
func (s *Printer) HandleProcess(sf *plugins.CtxSysFlow, proc *sfgo.Process) error {
	return nil
}

// HandleFile processes File entities.
func (s *Printer) HandleFile(sf *plugins.CtxSysFlow, file *sfgo.File) error {
	return nil
}

// HandleNetFlow processes Network Flows.
func (s *Printer) HandleNetFlow(sf *plugins.CtxSysFlow, nf *sfgo.NetworkFlow) error {
	logger.Info.Printf("NetworkFlow %s, %d", sf.Process.Exe, nf.Dport)
	return nil
}

// HandleFileFlow processes File Flows.
func (s *Printer) HandleFileFlow(sf *plugins.CtxSysFlow, ff *sfgo.FileFlow) error {
	logger.Info.Printf("FileFlow %s, %d", sf.Process.Exe, ff.Fd)
	return nil
}

// HandleFileEvt processes File Events.
func (s *Printer) HandleFileEvt(sf *plugins.CtxSysFlow, fe *sfgo.FileEvent) error {
	logger.Info.Printf("FileEvt %s, %d", sf.Process.Exe, fe.Tid)
	return nil
}

// HandleFileEvt processes Net Events.
func (s *Printer) HandleNetEvt(sf *plugins.CtxSysFlow, ne *sfgo.NetworkEvent) error {
	logger.Info.Printf("NetEvt %s, %d", sf.Process.Exe, ne.Tid)
	return nil
}

// HandleProcEvt processes Process Events.
func (s *Printer) HandleProcEvt(sf *plugins.CtxSysFlow, pe *sfgo.ProcessEvent) error {
	logger.Info.Printf("ProcEvt %s, %d", sf.Process.Exe, pe.Tid)
	return nil
}

// HandleProcEvt processes Process Flows.
func (s *Printer) HandleProcFlow(sf *plugins.CtxSysFlow, pf *sfgo.ProcessFlow) error {
	logger.Info.Printf("ProcFlow %s, %v", sf.Process.Exe, pf.ProcOID)
	return nil
}

// This function is not run when module is used as a plugin.
func main() {}
