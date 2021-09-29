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

// Package processor implements a processor plugin.
package processor

import (
	"strings"
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/cache"
)

const (
	readerPluginName  string = "sysflowreader"
	readerChannelName string = "sysflowchan"
)

// SysFlowReader defines the main reader class, which process SysFlow records and builds the cache.
// This plugin should typically be first in the pipeline.
type SysFlowReader struct {
	SysFlowProcessor
	hdr    *sfgo.SFHeader
	tables *cache.SFTables
}

// NewSysFlowProcessor creates a new SysFlowProcessor instance.
func NewSysFlowReader() plugins.SFProcessor {
	logger.Trace.Println("Calling NewSysFlowReader")
	p := new(SysFlowReader)
	return p
}

// GetName returns the plugin name.
func (s *SysFlowReader) GetName() string {
	return readerPluginName
}

// NewSysFlowChan creates a new processor channel instance.
func NewSysFlowChan(size int) interface{} {
	return &plugins.SFChannel{In: make(chan *sfgo.SysFlow, size)}
}

// Register registers plugin to plugin cache.
func (s *SysFlowReader) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(readerPluginName, NewSysFlowReader)
	pc.AddChannel(readerChannelName, NewSysFlowChan)
	sPCOnce.Do(func() {
		sPluginCache = pc
	})
}

// Init initializes the processor with a configuration map.
func (s *SysFlowReader) Init(conf map[string]interface{}) (err error) {
	s.tables = cache.NewSFTables()
	return s.SysFlowProcessor.Init(conf)
}

// Process implements the main processor method of the plugin.
func (s *SysFlowReader) Process(ch interface{}, wg *sync.WaitGroup) {
	entEnabled := s.hdl.IsEntityEnabled()
	cha := ch.(*plugins.SFChannel)
	record := cha.In
	defer wg.Done()
	logger.Trace.Println("Starting SysFlow Reader...")
	for {
		r, ok := <-record
		if !ok {
			logger.Trace.Println("SysFlow Reader channel closed. Shutting down.")
			break
		}
		sf := new(plugins.CtxSysFlow)
		sf.SysFlow = r
		sf.Header = s.hdr
		switch sf.Rec.UnionType {
		case sfgo.SF_HEADER:
			s.hdr = sf.Rec.SFHeader
			s.tables.Reset()
			if entEnabled {
				s.hdl.HandleHeader(sf, s.hdr)
			}
		case sfgo.SF_CONT:
			cont := sf.Rec.Container
			s.tables.SetCont(cont.Id, cont)
			if entEnabled {
				s.hdl.HandleContainer(sf, cont)
			}
		case sfgo.SF_PROCESS:
			proc := sf.Rec.Process
			proc.Exe = strings.TrimSpace(proc.Exe)
			proc.ExeArgs = strings.TrimSpace(proc.ExeArgs)
			s.tables.SetProc(*proc.Oid, proc)
			if entEnabled {
				sf.Process = proc
				sf.PTree = s.tables.GetPtree(*proc.Oid)
				sf.Container = s.getContFromProc(proc)
				s.hdl.HandleProcess(sf, proc)
			}
		case sfgo.SF_FILE:
			sf.File = sf.Rec.File
			s.tables.SetFile(sf.File.Oid, sf.File)
			if entEnabled {
				sf.Container = s.getContFromFile(sf.File)
				s.hdl.HandleFile(sf, sf.File)
			}
		case sfgo.SF_PROC_EVT:
			pe := sf.Rec.ProcessEvent
			sf.Container, sf.Process, sf.PTree = s.getContAndProc(pe.ProcOID)
			s.hdl.HandleProcEvt(sf, pe)
		case sfgo.SF_NET_FLOW:
			nf := sf.Rec.NetworkFlow
			sf.Container, sf.Process, sf.PTree = s.getContAndProc(nf.ProcOID)
			s.hdl.HandleNetFlow(sf, nf)
		case sfgo.SF_FILE_FLOW:
			ff := sf.Rec.FileFlow
			sf.Container, sf.Process, sf.PTree = s.getContAndProc(ff.ProcOID)
			sf.File = s.getFile(ff.FileOID)
			s.hdl.HandleFileFlow(sf, ff)
		case sfgo.SF_FILE_EVT:
			fe := sf.Rec.FileEvent
			sf.Container, sf.Process, sf.PTree = s.getContAndProc(fe.ProcOID)
			sf.File = s.getFile(fe.FileOID)
			sf.NewFile = s.getOptFile(fe.NewFileOID)
			s.hdl.HandleFileEvt(sf, fe)
		case sfgo.SF_PROC_FLOW:
		case sfgo.SF_NET_EVT:
		default:
			logger.Warn.Println("Error unsupported SysFlow Type: ", sf.Rec.UnionType)
		}
	}
}

// Cleanup tears down the plugin resources.
func (s *SysFlowReader) Cleanup() {
	logger.Trace.Println("Exiting ", readerPluginName)
	s.hdl.Cleanup()
}

func (s *SysFlowReader) getContFromProc(proc *sfgo.Process) *sfgo.Container {
	if proc.ContainerId != nil && proc.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		if c := s.tables.GetCont(proc.ContainerId.String); c != nil {
			return c
		}
		logger.Warn.Println("No container object for ID: ", proc.ContainerId.String)
	}
	return nil
}

func (s *SysFlowReader) getContAndProc(oid *sfgo.OID) (*sfgo.Container, *sfgo.Process, []*sfgo.Process) {
	if p := s.tables.GetProc(*oid); p != nil {
		ptree := s.tables.GetPtree(*oid)
		if p.ContainerId != nil && p.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
			if c := s.tables.GetCont(p.ContainerId.String); c != nil {
				return c, p, ptree
			}
			logger.Warn.Println("No container object for ID: ", p.ContainerId.String)
		}
		return nil, p, ptree
	}
	logger.Error.Println("No process object for ID: ", *oid)
	return nil, nil, nil
}

func (s *SysFlowReader) getFile(foid sfgo.FOID) *sfgo.File {
	if f := s.tables.GetFile(foid); f != nil {
		return f
	}
	logger.Error.Println("No file object for FOID: ", foid)
	return nil
}

func (s *SysFlowReader) getOptFile(unf *sfgo.UnionNullFOID) *sfgo.File {
	if unf != nil && unf.UnionType == sfgo.UnionNullFOIDTypeEnumFOID {
		return s.getFile(unf.FOID)
	}
	return nil
}

func (s *SysFlowReader) getContFromFile(file *sfgo.File) *sfgo.Container {
	if file != nil && file.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		if c := s.tables.GetCont(file.ContainerId.String); c != nil {
			return c
		}
		logger.Warn.Println("Not container object for ID: ", file.ContainerId.String)
	}
	return nil
}
