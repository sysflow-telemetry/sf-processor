//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
package processor

import (
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/goutils/logger"
	"github.ibm.com/sysflow/sf-processor/core/cache"
	"github.ibm.com/sysflow/sf-processor/core/flattener"
)

const (
	pluginName  string = "sysflowreader"
	channelName string = "sysflowchan"
)

// SysFlowProcessor defines the main processor class.
type SysFlowProcessor struct {
	hdr    *sfgo.SFHeader
	hdl    plugins.SFHandler
	tables *cache.SFTables
}

// NewSysFlowProcessor creates a new SysFlowProcessor instance.
func NewSysFlowProcessor(hdl plugins.SFHandler) plugins.SFProcessor {
	logger.Trace.Println("Calling NewSysFlowProc")
	p := new(SysFlowProcessor)
	p.hdl = hdl
	return p
}

// GetName returns the plugin name.
func (s *SysFlowProcessor) GetName() string {
	return pluginName
}

// NewSysFlowChan creates a new processor channel instance.
func NewSysFlowChan(size int) interface{} {
	return &plugins.SFChannel{In: make(chan *sfgo.SysFlow, size)}
}

// Register registers plugin to plugin cache.
func (s *SysFlowProcessor) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(pluginName, NewSysFlowProcessor)
	pc.AddChannel(channelName, NewSysFlowChan)
	(&flattener.Flattener{}).Register(pc)
}

// Init initializes the processor with a configuration map.
func (s *SysFlowProcessor) Init(conf map[string]string) error {
	s.tables = cache.GetInstance()
	return nil
}

// SetOutChan sets the output channel of the plugin.
func (s *SysFlowProcessor) SetOutChan(ch interface{}) {
	s.hdl.SetOutChan(ch)
}

// Cleanup tears down the plugin resources.
func (s *SysFlowProcessor) Cleanup() {
	s.hdl.Cleanup()
}

// Process implements the main processor method of the plugin.
func (s *SysFlowProcessor) Process(ch interface{}, wg *sync.WaitGroup) {
	entEnabled := s.hdl.IsEntityEnabled()
	cha := ch.(*plugins.SFChannel)
	record := cha.In
	defer wg.Done()
	logger.Trace.Println("Starting SysFlow processing...")
	for {
		sf, ok := <-record
		if !ok {
			logger.Trace.Println("SysFlow Processor channel closed. Shutting down.")
			break
		}
		switch sf.Rec.UnionType {
		case sfgo.SF_HEADER:
			hdr := sf.Rec.SFHeader
			s.hdr = hdr
			s.tables.Reset()
			if entEnabled {
				s.hdl.HandleHeader(s.hdr)
			}
		case sfgo.SF_CONT:
			cont := sf.Rec.Container
			s.tables.SetCont(cont.Id, cont)
			if entEnabled {
				s.hdl.HandleContainer(s.hdr, cont)
			}
		case sfgo.SF_PROCESS:
			proc := sf.Rec.Process
			s.tables.SetProc(*proc.Oid, proc)
			if entEnabled {
				cont := s.getContFromProc(proc)
				s.hdl.HandleProcess(s.hdr, cont, proc)
			}
		case sfgo.SF_FILE:
			file := sf.Rec.File
			s.tables.SetFile(file.Oid, file)
			if entEnabled {
				cont := s.getContFromFile(file)
				s.hdl.HandleFile(s.hdr, cont, file)
			}
		case sfgo.SF_PROC_EVT:
			pe := sf.Rec.ProcessEvent
			cont, proc := s.getContAndProc(pe.ProcOID)
			s.hdl.HandleProcEvt(s.hdr, cont, proc, pe)
		case sfgo.SF_NET_FLOW:
			nf := sf.Rec.NetworkFlow
			cont, proc := s.getContAndProc(nf.ProcOID)
			s.hdl.HandleNetFlow(s.hdr, cont, proc, nf)
		case sfgo.SF_FILE_FLOW:
			ff := sf.Rec.FileFlow
			cont, proc := s.getContAndProc(ff.ProcOID)
			file := s.getFile(ff.FileOID)
			s.hdl.HandleFileFlow(s.hdr, cont, proc, file, ff)
		case sfgo.SF_FILE_EVT:
			fe := sf.Rec.FileEvent
			cont, proc := s.getContAndProc(fe.ProcOID)
			file := s.getFile(fe.FileOID)
			file2 := s.getOptFile(fe.NewFileOID)
			s.hdl.HandleFileEvt(s.hdr, cont, proc, file, file2, fe)
		case sfgo.SF_NET_EVT:
		default:
			logger.Warn.Println("Error unsupported SysFlow Type: ", sf.Rec.UnionType)

		}
	}
	s.Cleanup()
}

func (s *SysFlowProcessor) getContFromProc(proc *sfgo.Process) *sfgo.Container {
	if proc.ContainerId != nil && proc.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		if c := s.tables.GetCont(proc.ContainerId.String); c != nil {
			return c
		}
		logger.Warn.Println("No container object for ID: ", proc.ContainerId.String)
	}
	return nil
}

func (s *SysFlowProcessor) getContAndProc(oid *sfgo.OID) (*sfgo.Container, *sfgo.Process) {
	if p := s.tables.GetProc(*oid); p != nil {
		if p.ContainerId != nil && p.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
			if c := s.tables.GetCont(p.ContainerId.String); c != nil {
				return c, p
			}
			logger.Warn.Println("No container object for ID: ", p.ContainerId.String)
		}
		return nil, p
	}
	logger.Error.Println("No process object for ID: ", *oid)
	return nil, nil
}

func (s *SysFlowProcessor) getFile(foid sfgo.FOID) *sfgo.File {
	if f := s.tables.GetFile(foid); f != nil {
		return f
	}
	logger.Error.Println("No file object for FOID: ", foid)
	return nil
}

func (s *SysFlowProcessor) getOptFile(unf *sfgo.UnionNullFOID) *sfgo.File {
	if unf != nil && unf.UnionType == sfgo.UnionNullFOIDTypeEnumFOID {
		return s.getFile(unf.FOID)
	}
	return nil
}

func (s *SysFlowProcessor) getContFromFile(file *sfgo.File) *sfgo.Container {
	if file != nil && file.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		if c := s.tables.GetCont(file.ContainerId.String); c != nil {
			return c
		}
		logger.Warn.Println("Not container object for ID: ", file.ContainerId.String)
	}
	return nil
}
