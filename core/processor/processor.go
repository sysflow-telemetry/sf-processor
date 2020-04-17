package processor

import (
	"sync"

	hdl "github.com/sysflow-telemetry/sf-apis/go/handlers"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/sf-processor/common/logger"
)

// SysFlowProcessor defines the main processor class.
type SysFlowProcessor struct {
	hdr       *sfgo.SFHeader
	Hdl       hdl.SFHandler
	contTable map[string]*sfgo.Container
	procTable map[sfgo.OID]*sfgo.Process
	fileTable map[sfgo.FOID]*sfgo.File
	OutChan   interface{}
}

// NewSysFlowProc creates a new SysFlowProcessor instance.
func NewSysFlowProc(hdl hdl.SFHandler) sp.SFProcessor {
	logger.Trace.Println("Calling NewSysFlowProc")
	p := new(SysFlowProcessor)
	p.Hdl = hdl
	p.contTable = make(map[string]*sfgo.Container)
	p.procTable = make(map[sfgo.OID]*sfgo.Process)
	p.fileTable = make(map[sfgo.FOID]*sfgo.File)
	return p
}

// NewSysFlowChan creates a new processor channel instance.
func NewSysFlowChan(size int) interface{} {
	return &sp.SFChannel{In: make(chan *sfgo.SysFlow, size)}
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
			s.contTable = make(map[string]*sfgo.Container)
			s.procTable = make(map[sfgo.OID]*sfgo.Process)
			s.fileTable = make(map[sfgo.FOID]*sfgo.File)
			if entEnabled {
				s.Hdl.HandleHeader(s.hdr)
			}
		case sfgo.SF_CONT:
			cont := sf.Rec.Container
			s.contTable[cont.Id] = cont
			if entEnabled {
				s.Hdl.HandleContainer(s.hdr, cont)
			}
		case sfgo.SF_PROCESS:
			proc := sf.Rec.Process
			s.procTable[*proc.Oid] = proc
			if entEnabled {
				cont := s.getContFromProc(proc)
				s.Hdl.HandleProcess(s.hdr, cont, proc)
			}
		case sfgo.SF_FILE:
			file := sf.Rec.File
			s.fileTable[file.Oid] = file
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
		if c, ok := s.contTable[proc.ContainerId.String]; ok {
			return c
		}
		logger.Warn.Println("No container object for ID: ", proc.ContainerId.String)
	}
	return nil
}

func (s *SysFlowProcessor) getContAndProc(oid *sfgo.OID) (*sfgo.Container, *sfgo.Process) {
	p, ok := s.procTable[*oid]
	if !ok {
		logger.Warn.Println("No process object for PID, Create Time ", oid.Hpid, oid.CreateTS)
		return nil, nil
	}
	if p.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		if c, ok := s.contTable[p.ContainerId.String]; ok {
			return c, p
		}
		logger.Warn.Println("No container object for ID: ", p.ContainerId.String)
	}
	return nil, p
}

func (s *SysFlowProcessor) getFile(foid sfgo.FOID) *sfgo.File {
	if f, ok := s.fileTable[foid]; ok {
		return f
	}
	logger.Error.Println("No file object for FOID")
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
		if c, ok := s.contTable[file.ContainerId.String]; ok {
			return c
		} else {
			logger.Warn.Println("Not container object for ID: ", file.ContainerId.String)
		}
	}
	return nil
}
