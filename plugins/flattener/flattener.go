package flattener

import (
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/sf-processor/plugins/flattener/types"
)

type Flattener struct {
	Pool *types.FlatRecordPool
}

func (g Flattener) IsEntityEnabled() bool {
	return false
}

func (g Flattener) HandleHeader(hdr *sfgo.SFHeader) error {
	println("HEAD: ", hdr.Version, " ", hdr.Exporter)
	return nil
}

func (g Flattener) HandleContainer(hdr *sfgo.SFHeader, cont *sfgo.Container) error {
	println("CONT: ", cont.Id, " ", cont.Image, " ", cont.Imageid)
	return nil
}

func (g Flattener) HandleProcess(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process) error {
	println("types.PROC: ", proc.Exe, " ", proc.ExeArgs, " ", proc.Oid.Hpid)
	return nil
}

func (g Flattener) HandleFile(hdr *sfgo.SFHeader, cont *sfgo.Container, file *sfgo.File) error {
	println("FILE: ", file.Path, " ", file.Restype, " ", file.Ts)
	return nil
}

func (g Flattener) FillEntities(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, fr *types.FlatRecord) {
	fr.Ints[types.SFHE_VERSION_INT] = hdr.Version
	fr.Strs[types.SFHE_EXPORTER_STR] = hdr.Exporter

	fr.Strs[types.CONT_ID_STR] = cont.Id
	fr.Strs[types.CONT_NAME_STR] = cont.Name
	fr.Strs[types.CONT_IMAGE_STR] = cont.Image
	fr.Strs[types.CONT_IMAGEID_STR] = cont.Imageid
	fr.Ints[types.CONT_TYPE_INT] = int64(cont.Type)
	if cont.Privileged {
		fr.Ints[types.CONT_PRIVILEGED_INT] = 1
	} else {
		fr.Ints[types.CONT_PRIVILEGED_INT] = 0
	}

	fr.Ints[types.PROC_STATE_INT] = int64(proc.State)
	fr.Ints[types.PROC_OID_CREATETS_INT] = int64(proc.Oid.CreateTS)
	fr.Ints[types.PROC_OID_HPID_INT] = int64(proc.Oid.Hpid)

	if proc.Poid.UnionType != sfgo.UnionNullOIDTypeEnumNull {
		fr.Ints[types.PROC_POID_CREATETS_INT] = proc.Poid.OID.CreateTS
		fr.Ints[types.PROC_POID_HPID_INT] = proc.Poid.OID.Hpid
	} else {
		fr.Ints[types.PROC_POID_CREATETS_INT] = -1
		fr.Ints[types.PROC_POID_HPID_INT] = -1
	}
	fr.Ints[types.PROC_TS_INT] = proc.Ts
	fr.Strs[types.PROC_EXE_STR] = proc.Exe
	fr.Strs[types.PROC_EXEARGS_STR] = proc.ExeArgs
	fr.Ints[types.PROC_UID_INT] = int64(proc.Uid)
	fr.Strs[types.PROC_USERNAME_STR] = proc.UserName
	fr.Ints[types.PROC_GID_INT] = int64(proc.Gid)
	fr.Strs[types.PROC_GROUPNAME_STR] = proc.GroupName

	if proc.Tty {
		fr.Ints[types.PROC_TTY_INT] = 1
	} else {
		fr.Ints[types.PROC_TTY_INT] = 0
	}

	if proc.ContainerId.UnionType != sfgo.UnionNullStringTypeEnumNull {
		fr.Strs[types.PROC_CONTAINERID_STRING_STR] = proc.ContainerId.String
	} else {
		fr.Strs[types.PROC_CONTAINERID_STRING_STR] = ""
	}
}

func (g Flattener) HandleNetFlow(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, nf *sfgo.NetworkFlow) error {
	fr := g.Pool.Get()
	g.FillEntities(hdr, cont, proc, fr)
	fr.Ints[types.FL_NETW_TS_INT] = nf.Ts
	fr.Ints[types.FL_NETW_TID_INT] = nf.Tid
	fr.Ints[types.FL_NETW_OPFLAGS_INT] = int64(nf.OpFlags)
	fr.Ints[types.FL_NETW_ENDTS_INT] = nf.EndTs
	fr.Ints[types.FL_NETW_SIP_INT] = int64(nf.Sip)
	fr.Ints[types.FL_NETW_SPORT_INT] = int64(nf.Sport)
	fr.Ints[types.FL_NETW_DIP_INT] = int64(nf.Dip)
	fr.Ints[types.FL_NETW_DPORT_INT] = int64(nf.Dport)
	fr.Ints[types.FL_NETW_PROTO_INT] = int64(nf.Proto)
	fr.Ints[types.FL_NETW_FD_INT] = int64(nf.Fd)
	fr.Ints[types.FL_NETW_NUMRRECVOPS_INT] = nf.NumRRecvOps
	fr.Ints[types.FL_NETW_NUMWSENDOPS_INT] = nf.NumWSendOps
	fr.Ints[types.FL_NETW_NUMRRECVBYTES_INT] = nf.NumRRecvBytes
	fr.Ints[types.FL_NETW_NUMWSENDBYTES_INT] = nf.NumWSendBytes

	return nil
}

func (g Flattener) HandleFileFlow(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, file *sfgo.File, ff *sfgo.FileFlow) error {
	fr := g.Pool.Get()
	fr.Ints[types.FILE_STATE_INT] = int64(file.State)
	fr.Ints[types.FILE_TS_INT] = file.Ts
	fr.Ints[types.FILE_RESTYPE_INT] = int64(file.Restype)
	fr.Strs[types.FILE_PATH_STR] = file.Path
	if file.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		fr.Strs[types.FILE_CONTAINERID_STRING_STR] = file.ContainerId.String
	} else {
		fr.Strs[types.FILE_CONTAINERID_STRING_STR] = ""
	}
	g.FillEntities(hdr, cont, proc, fr)
	fr.Ints[types.FL_FILE_TS_INT] = ff.Ts
	fr.Ints[types.FL_FILE_TID_INT] = ff.Tid
	fr.Ints[types.FL_FILE_OPFLAGS_INT] = int64(ff.OpFlags)
	fr.Ints[types.FL_FILE_OPENFLAGS_INT] = int64(ff.OpenFlags)
	fr.Ints[types.FL_FILE_ENDTS_INT] = ff.EndTs
	fr.Ints[types.FL_FILE_FD_INT] = int64(ff.Fd)
	fr.Ints[types.FL_FILE_NUMRRECVOPS_INT] = ff.NumRRecvOps
	fr.Ints[types.FL_FILE_NUMWSENDOPS_INT] = ff.NumWSendOps
	fr.Ints[types.FL_FILE_NUMRRECVBYTES_INT] = ff.NumRRecvBytes
	fr.Ints[types.FL_FILE_NUMWSENDBYTES_INT] = ff.NumWSendBytes

	return nil
}

func (g Flattener) HandleFileEvt(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, file1 *sfgo.File, file2 *sfgo.File, fe *sfgo.FileEvent) error {
	//ts := utils.GetTimeStrLocal(fe.Ts)
	/*println("FE", proc.Exe, proc.ExeArgs, ts, utils.GetOpFlagsStr(fe.OpFlags), " Tid ", fe.Tid, file1.Path, string(file1.Restype), " File 2 Path: ", func() string {
		if file2 != nil {
			return file2.Path
		} else {
			return ""
		}
	}())*/
	fr := g.Pool.Get()
	fr.Ints[types.FILE_STATE_INT] = int64(file1.State)
	fr.Ints[types.FILE_TS_INT] = file1.Ts
	fr.Ints[types.FILE_RESTYPE_INT] = int64(file1.Restype)
	fr.Strs[types.FILE_PATH_STR] = file1.Path
	if file1.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		fr.Strs[types.FILE_CONTAINERID_STRING_STR] = file1.ContainerId.String
	} else {
		fr.Strs[types.FILE_CONTAINERID_STRING_STR] = ""
	}

	if file2 != nil {
		fr.Ints[types.SEC_FILE_STATE_INT] = int64(file2.State)
		fr.Ints[types.SEC_FILE_TS_INT] = file2.Ts
		fr.Ints[types.SEC_FILE_RESTYPE_INT] = int64(file2.Restype)
		fr.Strs[types.SEC_FILE_PATH_STR] = file2.Path
		if file2.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
			fr.Strs[types.SEC_FILE_CONTAINERID_STRING_STR] = file2.ContainerId.String
		} else {
			fr.Strs[types.SEC_FILE_CONTAINERID_STRING_STR] = ""
		}

	}

	fr.Ints[types.SF_REC_TYPE] = types.FILE_EVT
	g.FillEntities(hdr, cont, proc, fr)
	fr.Ints[types.EV_FILE_TS_INT] = fe.Ts
	fr.Ints[types.EV_FILE_TID_INT] = fe.Tid
	fr.Ints[types.EV_FILE_OPFLAGS_INT] = int64(fe.OpFlags)
	fr.Ints[types.EV_FILE_RET_INT] = int64(fe.Ret)
	return nil
}

func (g Flattener) HandleProcEvt(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, pe *sfgo.ProcessEvent) error {
	//ts := utils.GetTimeStrLocal(pe.Ts)
	//println("PE", proc.Exe, proc.ExeArgs, ts, utils.GetOpFlagsStr(pe.OpFlags), " Tid ", pe.Tid, " Return ", pe.Ret)
	fr := g.Pool.Get()
	fr.Ints[types.SF_REC_TYPE] = types.PROC_EVT
	g.FillEntities(hdr, cont, proc, fr)
	fr.Ints[types.EV_PROC_TS_INT] = pe.Ts
	fr.Ints[types.EV_PROC_TID_INT] = pe.Tid
	fr.Ints[types.EV_PROC_OPFLAGS_INT] = int64(pe.OpFlags)
	fr.Ints[types.EV_PROC_RET_INT] = int64(pe.Ret)
	return nil
}
