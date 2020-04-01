package main

import (
	"log"

	"github.com/sysflow-telemetry/sf-apis/go/handlers"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"sync"
)

type SysFlowProcessor struct {
	hdr       *sfgo.SFHeader
	Hdl       handlers.SFHandler
	contTable map[string]*sfgo.Container
	procTable map[sfgo.OID]*sfgo.Process
	fileTable map[sfgo.FOID]*sfgo.File
}

func NewSysFlowProc(hdl handlers.SFHandler) *SysFlowProcessor {
	p := new(SysFlowProcessor)
	p.Hdl = hdl
	p.contTable = make(map[string]*sfgo.Container)
	p.procTable = make(map[sfgo.OID]*sfgo.Process)
	p.fileTable = make(map[sfgo.FOID]*sfgo.File)
	return p
}

func (s SysFlowProcessor) getContFromProc(proc *sfgo.Process) *sfgo.Container {
	if proc.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		if c, ok := s.contTable[proc.ContainerId.String]; ok {
			return c
		} else {
			log.Println("No container object for ID: ", proc.ContainerId.String)
		}
	}
	return nil
}

func (s SysFlowProcessor) getContAndProc(oid *sfgo.OID) (*sfgo.Container, *sfgo.Process) {
	p, ok := s.procTable[*oid]
	if !ok {
		log.Println("No process object for PID, Create Time ", oid.Hpid, oid.CreateTS)
		return nil, nil
	}
	if p.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		if c, ok := s.contTable[p.ContainerId.String]; ok {
			return c, p
		} else {
			log.Println("No container object for ID: ", p.ContainerId.String)
		}
	}
	return nil, p
}

func (s SysFlowProcessor) getFile(foid sfgo.FOID) *sfgo.File {
	if f, ok := s.fileTable[foid]; ok {
		return f
	}
	log.Println("No file object for FOID")
	return nil
}
func (s SysFlowProcessor) getOptFile(unf *sfgo.UnionNullFOID) *sfgo.File {
	if unf.UnionType == sfgo.UnionNullFOIDTypeEnumFOID {
		return s.getFile(unf.FOID)
	}
	return nil
}

func (s SysFlowProcessor) getContFromFile(file *sfgo.File) *sfgo.Container {
	if file.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		if c, ok := s.contTable[file.ContainerId.String]; ok {
			return c
		} else {
			log.Println("Not container object for ID: ", file.ContainerId.String)
		}
	}
	return nil
}

func (s SysFlowProcessor) process(record <-chan *sfgo.SysFlow, wg *sync.WaitGroup) {
	entEnabled := s.Hdl.IsEntityEnabled()
	defer wg.Done()
	for {
		sf, ok := <-record
		if !ok {
			log.Println("Channel closed. Shutting down.")
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
			log.Println("Error unsupported SysFlow Type: ", sf.Rec.UnionType)

		}

	}
}
