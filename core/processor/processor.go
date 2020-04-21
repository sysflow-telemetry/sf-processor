package processor

import (
	"sync"

	"github.ibm.com/sysflow/sf-processor/core/cache"

	hdl "github.com/sysflow-telemetry/sf-apis/go/handlers"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/sf-processor/common/logger"
)

// SysFlowProcessor defines the main processor class.
type SysFlowProcessor struct {
	hdr     *sfgo.SFHeader
	Hdl     hdl.SFHandler
	tables  *cache.SFTables
	OutChan interface{}
}

// NewSysFlowProc creates a new SysFlowProcessor instance.
func NewSysFlowProc(hdl hdl.SFHandler) sp.SFProcessor {
	logger.Trace.Println("Calling NewSysFlowProc")
	p := new(SysFlowProcessor)
	p.Hdl = hdl
	return p
}

// NewSysFlowChan creates a new processor channel instance.
func NewSysFlowChan(size int) interface{} {
	return &sp.SFChannel{In: make(chan *sfgo.SysFlow, size)}
}

// Init initializes the processor with a configuration map.
func (s *SysFlowProcessor) Init(conf map[string]string, tables interface{}) error {
	s.tables = tables.(*cache.SFTables)
	return nil
}

// SetOutChan sets the output channel of the plugin.
func (s *SysFlowProcessor) SetOutChan(ch interface{}) {
	s.OutChan = ch
	s.Hdl.SetOutChan(ch)
}

// Cleanup tears down the plugin resources.
func (s *SysFlowProcessor) Cleanup() {
	s.Hdl.Cleanup()
}

// Process implements the main processor method of the plugin.
func (s *SysFlowProcessor) Process(ch interface{}, wg *sync.WaitGroup) {
	entEnabled := s.Hdl.IsEntityEnabled()
	cha := ch.(*sp.SFChannel)
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
				s.Hdl.HandleHeader(s.hdr)
			}
		case sfgo.SF_CONT:
			cont := sf.Rec.Container
			s.tables.SetCont(cont.Id, cont)
			if entEnabled {
				s.Hdl.HandleContainer(s.hdr, cont)
			}
		case sfgo.SF_PROCESS:
			proc := sf.Rec.Process
			s.tables.SetProc(*proc.Oid, proc)
			if entEnabled {
				cont := s.getContFromProc(proc)
				s.Hdl.HandleProcess(s.hdr, cont, proc)
			}
		case sfgo.SF_FILE:
			file := sf.Rec.File
			s.tables.SetFile(file.Oid, file)
			if entEnabled {
				cont := s.getContFromFile(file)
				s.Hdl.HandleFile(s.hdr, cont, file)
			}
		case sfgo.SF_PROC_EVT:
			pe := sf.Rec.ProcessEvent
			cont, proc := s.getContAndProc(pe.ProcOID)
			s.Hdl.HandleProcEvt(s.hdr, cont, proc, pe)
		case sfgo.SF_NET_FLOW:
			nf := sf.Rec.NetworkFlow
			cont, proc := s.getContAndProc(nf.ProcOID)
			s.Hdl.HandleNetFlow(s.hdr, cont, proc, nf)
		case sfgo.SF_FILE_FLOW:
			ff := sf.Rec.FileFlow
			cont, proc := s.getContAndProc(ff.ProcOID)
			file := s.getFile(ff.FileOID)
			s.Hdl.HandleFileFlow(s.hdr, cont, proc, file, ff)
		case sfgo.SF_FILE_EVT:
			fe := sf.Rec.FileEvent
			cont, proc := s.getContAndProc(fe.ProcOID)
			file := s.getFile(fe.FileOID)
			file2 := s.getOptFile(fe.NewFileOID)
			s.Hdl.HandleFileEvt(s.hdr, cont, proc, file, file2, fe)
		case sfgo.SF_NET_EVT:
		default:
			logger.Warn.Println("Error unsupported SysFlow Type: ", sf.Rec.UnionType)

		}
	}
	s.Cleanup()
}

func (s *SysFlowProcessor) getContFromProc(proc *sfgo.Process) *sfgo.Container {
	if proc.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		if c := s.tables.GetCont(proc.ContainerId.String); c != nil {
			return c
		}
		logger.Warn.Println("No container object for ID: ", proc.ContainerId.String)
	}
	return nil
}

func (s *SysFlowProcessor) getContAndProc(oid *sfgo.OID) (*sfgo.Container, *sfgo.Process) {
	if p := s.tables.GetProc(*oid); p != nil {
		if p.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
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
	if unf.UnionType == sfgo.UnionNullFOIDTypeEnumFOID {
		return s.getFile(unf.FOID)
	}
	return nil
}

func (s *SysFlowProcessor) getContFromFile(file *sfgo.File) *sfgo.Container {
	if file.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		if c := s.tables.GetCont(file.ContainerId.String); c != nil {
			return c
		}
		logger.Warn.Println("Not container object for ID: ", file.ContainerId.String)
	}
	return nil
}
