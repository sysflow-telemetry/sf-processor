package main

import (
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-apis/go/utils"
)

type Printer struct {
}

func (g Printer) IsEntityEnabled() bool {
	return false
}

func (g Printer) HandleHeader(hdr *sfgo.SFHeader) error {
	println("HEAD: ", hdr.Version, " ", hdr.Exporter)
	return nil
}

func (g Printer) HandleContainer(hdr *sfgo.SFHeader, cont *sfgo.Container) error {
	println("CONT: ", cont.Id, " ", cont.Image, " ", cont.Imageid)
	return nil
}

func (g Printer) HandleProcess(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process) error {
	println("PROC: ", proc.Exe, " ", proc.ExeArgs, " ", proc.Oid.Hpid)
	return nil
}

func (g Printer) HandleFile(hdr *sfgo.SFHeader, cont *sfgo.Container, file *sfgo.File) error {
	println("FILE: ", file.Path, " ", file.Restype, " ", file.Ts)
	return nil
}

func (g Printer) HandleNetFlow(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, nf *sfgo.NetworkFlow) error {
	nfStr := utils.GetNetworkFlowStr(nf)
	ts := utils.GetTimeStrLocal(nf.Ts)
	println("NF", proc.Exe, proc.ExeArgs, ts, utils.GetOpFlagsStr(nf.OpFlags), nfStr)
	return nil
}

func (g Printer) HandleFileFlow(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, file *sfgo.File, ff *sfgo.FileFlow) error {
	ts := utils.GetTimeStrLocal(ff.Ts)
	println("FF", proc.Exe, proc.ExeArgs, ts, utils.GetOpFlagsStr(ff.OpFlags), " Open Flags ", ff.OpenFlags, " Tid ", ff.Tid, file.Path, string(file.Restype))
	return nil
}

func (g Printer) HandleFileEvt(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, file1 *sfgo.File, file2 *sfgo.File, fe *sfgo.FileEvent) error {
	ts := utils.GetTimeStrLocal(fe.Ts)
	println("FE", proc.Exe, proc.ExeArgs, ts, utils.GetOpFlagsStr(fe.OpFlags), " Tid ", fe.Tid, file1.Path, string(file1.Restype), " File 2 Path: ", func() string {
		if file2 != nil {
			return file2.Path
		} else {
			return ""
		}
	}())
	return nil
}

func (g Printer) HandleProcEvt(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, pe *sfgo.ProcessEvent) error {
	ts := utils.GetTimeStrLocal(pe.Ts)
	println("PE", proc.Exe, proc.ExeArgs, ts, utils.GetOpFlagsStr(pe.OpFlags), " Tid ", pe.Tid, " Return ", pe.Ret)
	return nil
}
