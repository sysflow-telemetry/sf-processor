package flattener

import (
	hdl "github.com/sysflow-telemetry/sf-apis/go/handlers"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/sf-processor/common/logger"
)

// Flattener defines the main class for the flatterner plugin.
type Flattener struct {
	Pool     *sfgo.FlatRecordPool
	flatChan *hdl.FlatChannel
	ch       chan *sfgo.FlatRecord
}

// NewFlattener creates a new Flattener instance.
func NewFlattener() hdl.SFHandler {
	return new(Flattener)
}

// NewFlattenerChan creates a new channel with given capacity.
func NewFlattenerChan(size int) interface{} {
	ch := hdl.FlatChannel{In: make(chan *sfgo.FlatRecord, size)}
	return &ch
}

// IsEntityEnabled is used to check if the flattener returns entity records.
func (g *Flattener) IsEntityEnabled() bool {
	return false
}

// SetOutChan sets the plugin output channel.
func (g *Flattener) SetOutChan(chObj interface{}) {
	g.flatChan = chObj.(*hdl.FlatChannel)
	g.ch = g.flatChan.In
}

// Cleanup tears down resources.
func (g *Flattener) Cleanup() {
	logger.Trace.Println("Calling Cleanup on Flattener channel")
	close(g.ch)
}

// HandleHeader processes Header entities.
func (g *Flattener) HandleHeader(hdr *sfgo.SFHeader) error {
	return nil
}

// HandleContainer processes Container entities.
func (g *Flattener) HandleContainer(hdr *sfgo.SFHeader, cont *sfgo.Container) error {
	return nil
}

// HandleProcess processes Process entities.
func (g *Flattener) HandleProcess(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process) error {
	return nil
}

// HandleFile processes File entities.
func (g *Flattener) HandleFile(hdr *sfgo.SFHeader, cont *sfgo.Container, file *sfgo.File) error {
	return nil
}

// HandleNetFlow processes Network Flows.
func (g *Flattener) HandleNetFlow(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, nf *sfgo.NetworkFlow) error {
	//fr := g.Pool.Get()
	fr := new(sfgo.FlatRecord)
	g.fillEntities(hdr, cont, proc, nil, fr)
	fr.Ints[sfgo.FL_NETW_TS_INT] = nf.Ts
	fr.Ints[sfgo.FL_NETW_TID_INT] = nf.Tid
	fr.Ints[sfgo.FL_NETW_OPFLAGS_INT] = int64(nf.OpFlags)
	fr.Ints[sfgo.FL_NETW_ENDTS_INT] = nf.EndTs
	fr.Ints[sfgo.FL_NETW_SIP_INT] = int64(nf.Sip)
	fr.Ints[sfgo.FL_NETW_SPORT_INT] = int64(nf.Sport)
	fr.Ints[sfgo.FL_NETW_DIP_INT] = int64(nf.Dip)
	fr.Ints[sfgo.FL_NETW_DPORT_INT] = int64(nf.Dport)
	fr.Ints[sfgo.FL_NETW_PROTO_INT] = int64(nf.Proto)
	fr.Ints[sfgo.FL_NETW_FD_INT] = int64(nf.Fd)
	fr.Ints[sfgo.FL_NETW_NUMRRECVOPS_INT] = nf.NumRRecvOps
	fr.Ints[sfgo.FL_NETW_NUMWSENDOPS_INT] = nf.NumWSendOps
	fr.Ints[sfgo.FL_NETW_NUMRRECVBYTES_INT] = nf.NumRRecvBytes
	fr.Ints[sfgo.FL_NETW_NUMWSENDBYTES_INT] = nf.NumWSendBytes
	g.ch <- fr
	return nil
}

// HandleFileFlow processes File Flows.
func (g *Flattener) HandleFileFlow(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, file *sfgo.File, ff *sfgo.FileFlow) error {
	//fr := g.Pool.Get()
	fr := new(sfgo.FlatRecord)
	g.fillEntities(hdr, cont, proc, file, fr)
	fr.Ints[sfgo.FL_FILE_TS_INT] = ff.Ts
	fr.Ints[sfgo.FL_FILE_TID_INT] = ff.Tid
	fr.Ints[sfgo.FL_FILE_OPFLAGS_INT] = int64(ff.OpFlags)
	fr.Ints[sfgo.FL_FILE_OPENFLAGS_INT] = int64(ff.OpenFlags)
	fr.Ints[sfgo.FL_FILE_ENDTS_INT] = ff.EndTs
	fr.Ints[sfgo.FL_FILE_FD_INT] = int64(ff.Fd)
	fr.Ints[sfgo.FL_FILE_NUMRRECVOPS_INT] = ff.NumRRecvOps
	fr.Ints[sfgo.FL_FILE_NUMWSENDOPS_INT] = ff.NumWSendOps
	fr.Ints[sfgo.FL_FILE_NUMRRECVBYTES_INT] = ff.NumRRecvBytes
	fr.Ints[sfgo.FL_FILE_NUMWSENDBYTES_INT] = ff.NumWSendBytes
	g.ch <- fr
	return nil
}

// HandleFileEvt processes File Events.
func (g *Flattener) HandleFileEvt(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, file1 *sfgo.File, file2 *sfgo.File, fe *sfgo.FileEvent) error {
	//fr := g.Pool.Get()
	fr := new(sfgo.FlatRecord)
	if file2 != nil {
		fr.Ints[sfgo.SEC_FILE_STATE_INT] = int64(file2.State)
		fr.Ints[sfgo.SEC_FILE_TS_INT] = file2.Ts
		fr.Ints[sfgo.SEC_FILE_RESTYPE_INT] = int64(file2.Restype)
		fr.Strs[sfgo.SEC_FILE_PATH_STR] = file2.Path
		if file2.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
			fr.Strs[sfgo.SEC_FILE_CONTAINERID_STRING_STR] = file2.ContainerId.String
		} else {
			fr.Strs[sfgo.SEC_FILE_CONTAINERID_STRING_STR] = sfgo.EMPTY_STR
		}
	}
	fr.Ints[sfgo.SF_REC_TYPE] = sfgo.FILE_EVT
	g.fillEntities(hdr, cont, proc, file1, fr)
	fr.Ints[sfgo.EV_FILE_TS_INT] = fe.Ts
	fr.Ints[sfgo.EV_FILE_TID_INT] = fe.Tid
	fr.Ints[sfgo.EV_FILE_OPFLAGS_INT] = int64(fe.OpFlags)
	fr.Ints[sfgo.EV_FILE_RET_INT] = int64(fe.Ret)
	g.ch <- fr
	return nil
}

// HandleProcEvt processes Process Events.
func (g *Flattener) HandleProcEvt(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, pe *sfgo.ProcessEvent) error {
	//fr := g.Pool.Get()
	fr := new(sfgo.FlatRecord)
	fr.Ints[sfgo.SF_REC_TYPE] = sfgo.PROC_EVT
	g.fillEntities(hdr, cont, proc, nil, fr)
	fr.Ints[sfgo.EV_PROC_TS_INT] = pe.Ts
	fr.Ints[sfgo.EV_PROC_TID_INT] = pe.Tid
	fr.Ints[sfgo.EV_PROC_OPFLAGS_INT] = int64(pe.OpFlags)
	fr.Ints[sfgo.EV_PROC_RET_INT] = int64(pe.Ret)
	g.ch <- fr
	return nil
}

func (g *Flattener) fillEntities(hdr *sfgo.SFHeader, cont *sfgo.Container, proc *sfgo.Process, file *sfgo.File, fr *sfgo.FlatRecord) {
	if hdr != nil {
		fr.Ints[sfgo.SFHE_VERSION_INT] = hdr.Version
		fr.Strs[sfgo.SFHE_EXPORTER_STR] = hdr.Exporter
	} else {
		logger.Warn.Println("Event does not have a related header.  This should not happen.")
		fr.Ints[sfgo.SFHE_VERSION_INT] = sfgo.EMPTY_INT
		fr.Strs[sfgo.SFHE_EXPORTER_STR] = sfgo.EMPTY_STR
	}
	if cont != nil {
		fr.Strs[sfgo.CONT_ID_STR] = cont.Id
		fr.Strs[sfgo.CONT_NAME_STR] = cont.Name
		fr.Strs[sfgo.CONT_IMAGE_STR] = cont.Image
		fr.Strs[sfgo.CONT_IMAGEID_STR] = cont.Imageid
		fr.Ints[sfgo.CONT_TYPE_INT] = int64(cont.Type)

		if cont.Privileged {
			fr.Ints[sfgo.CONT_PRIVILEGED_INT] = 1
		} else {
			fr.Ints[sfgo.CONT_PRIVILEGED_INT] = 0
		}
	} else {
		fr.Strs[sfgo.CONT_ID_STR] = sfgo.EMPTY_STR
		fr.Strs[sfgo.CONT_NAME_STR] = sfgo.EMPTY_STR
		fr.Strs[sfgo.CONT_IMAGE_STR] = sfgo.EMPTY_STR
		fr.Strs[sfgo.CONT_IMAGEID_STR] = sfgo.EMPTY_STR
		fr.Ints[sfgo.CONT_TYPE_INT] = sfgo.EMPTY_INT
		fr.Ints[sfgo.CONT_PRIVILEGED_INT] = sfgo.EMPTY_INT

	}
	if proc != nil {
		fr.Ints[sfgo.PROC_STATE_INT] = int64(proc.State)
		fr.Ints[sfgo.PROC_OID_CREATETS_INT] = int64(proc.Oid.CreateTS)
		fr.Ints[sfgo.PROC_OID_HPID_INT] = int64(proc.Oid.Hpid)

		if proc.Poid.UnionType != sfgo.UnionNullOIDTypeEnumNull {
			fr.Ints[sfgo.PROC_POID_CREATETS_INT] = proc.Poid.OID.CreateTS
			fr.Ints[sfgo.PROC_POID_HPID_INT] = proc.Poid.OID.Hpid
		} else {
			fr.Ints[sfgo.PROC_POID_CREATETS_INT] = sfgo.EMPTY_INT
			fr.Ints[sfgo.PROC_POID_HPID_INT] = sfgo.EMPTY_INT
		}
		fr.Ints[sfgo.PROC_TS_INT] = proc.Ts
		fr.Strs[sfgo.PROC_EXE_STR] = proc.Exe
		fr.Strs[sfgo.PROC_EXEARGS_STR] = proc.ExeArgs
		fr.Ints[sfgo.PROC_UID_INT] = int64(proc.Uid)
		fr.Strs[sfgo.PROC_USERNAME_STR] = proc.UserName
		fr.Ints[sfgo.PROC_GID_INT] = int64(proc.Gid)
		fr.Strs[sfgo.PROC_GROUPNAME_STR] = proc.GroupName

		if proc.Tty {
			fr.Ints[sfgo.PROC_TTY_INT] = 1
		} else {
			fr.Ints[sfgo.PROC_TTY_INT] = 0
		}

		if proc.ContainerId.UnionType != sfgo.UnionNullStringTypeEnumNull {
			fr.Strs[sfgo.PROC_CONTAINERID_STRING_STR] = proc.ContainerId.String
		} else {
			fr.Strs[sfgo.PROC_CONTAINERID_STRING_STR] = sfgo.EMPTY_STR
		}
	} else {
		logger.Warn.Println("Event does not have a related process.  This should not happen.")
		fr.Ints[sfgo.PROC_STATE_INT] = sfgo.EMPTY_INT
		fr.Ints[sfgo.PROC_OID_CREATETS_INT] = sfgo.EMPTY_INT
		fr.Ints[sfgo.PROC_OID_HPID_INT] = sfgo.EMPTY_INT
		fr.Ints[sfgo.PROC_POID_CREATETS_INT] = sfgo.EMPTY_INT
		fr.Ints[sfgo.PROC_POID_HPID_INT] = sfgo.EMPTY_INT
		fr.Ints[sfgo.PROC_TS_INT] = sfgo.EMPTY_INT
		fr.Strs[sfgo.PROC_EXE_STR] = sfgo.EMPTY_STR
		fr.Strs[sfgo.PROC_EXEARGS_STR] = sfgo.EMPTY_STR
		fr.Ints[sfgo.PROC_UID_INT] = sfgo.EMPTY_INT
		fr.Strs[sfgo.PROC_USERNAME_STR] = sfgo.EMPTY_STR
		fr.Ints[sfgo.PROC_GID_INT] = sfgo.EMPTY_INT
		fr.Strs[sfgo.PROC_GROUPNAME_STR] = sfgo.EMPTY_STR
		fr.Ints[sfgo.PROC_TTY_INT] = sfgo.EMPTY_INT
		fr.Strs[sfgo.PROC_CONTAINERID_STRING_STR] = sfgo.EMPTY_STR
	}
	if file != nil {
		fr.Ints[sfgo.FILE_STATE_INT] = int64(file.State)
		fr.Ints[sfgo.FILE_TS_INT] = file.Ts
		fr.Ints[sfgo.FILE_RESTYPE_INT] = int64(file.Restype)
		fr.Strs[sfgo.FILE_PATH_STR] = file.Path
		if file.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
			fr.Strs[sfgo.FILE_CONTAINERID_STRING_STR] = file.ContainerId.String
		} else {
			fr.Strs[sfgo.FILE_CONTAINERID_STRING_STR] = sfgo.EMPTY_STR
		}
	} else {
		fr.Ints[sfgo.FILE_STATE_INT] = sfgo.EMPTY_INT
		fr.Ints[sfgo.FILE_TS_INT] = sfgo.EMPTY_INT
		fr.Ints[sfgo.FILE_RESTYPE_INT] = sfgo.EMPTY_INT
		fr.Strs[sfgo.FILE_PATH_STR] = sfgo.EMPTY_STR
		fr.Strs[sfgo.FILE_CONTAINERID_STRING_STR] = sfgo.EMPTY_STR
	}
}
