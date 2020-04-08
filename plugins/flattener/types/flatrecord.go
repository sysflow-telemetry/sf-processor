package types

import "sync"

type Attribute uint32

const (
	ARRAY_INT Attribute = 0
	ARRAY_STR Attribute = 0

	SF_REC_TYPE Attribute = ARRAY_INT
	//Type name:  SFHeader
	SFHE_VERSION_INT  Attribute = SF_REC_TYPE + 1
	SFHE_EXPORTER_STR Attribute = ARRAY_STR

	//Type name:  Container
	CONT_ID_STR         Attribute = SFHE_EXPORTER_STR + 1
	CONT_NAME_STR       Attribute = CONT_ID_STR + 1
	CONT_IMAGE_STR      Attribute = CONT_NAME_STR + 1
	CONT_IMAGEID_STR    Attribute = CONT_IMAGE_STR + 1
	CONT_TYPE_INT       Attribute = SFHE_VERSION_INT + 1
	CONT_PRIVILEGED_INT Attribute = CONT_TYPE_INT + 1

	//Type name:  Process
	PROC_STATE_INT              Attribute = CONT_PRIVILEGED_INT + 1
	PROC_OID_CREATETS_INT       Attribute = PROC_STATE_INT + 1
	PROC_OID_HPID_INT           Attribute = PROC_OID_CREATETS_INT + 1
	PROC_POID_CREATETS_INT      Attribute = PROC_OID_HPID_INT + 1
	PROC_POID_HPID_INT          Attribute = PROC_POID_CREATETS_INT + 1
	PROC_TS_INT                 Attribute = PROC_POID_HPID_INT + 1
	PROC_EXE_STR                Attribute = CONT_IMAGEID_STR + 1
	PROC_EXEARGS_STR            Attribute = PROC_EXE_STR + 1
	PROC_UID_INT                Attribute = PROC_TS_INT + 1
	PROC_USERNAME_STR           Attribute = PROC_EXEARGS_STR + 1
	PROC_GID_INT                Attribute = PROC_UID_INT + 1
	PROC_GROUPNAME_STR          Attribute = PROC_USERNAME_STR + 1
	PROC_TTY_INT                Attribute = PROC_GID_INT + 1
	PROC_CONTAINERID_STRING_STR Attribute = PROC_GROUPNAME_STR + 1

	//Type name:  File
	FILE_STATE_INT              Attribute = PROC_TTY_INT + 1
	FILE_TS_INT                 Attribute = FILE_STATE_INT + 1
	FILE_RESTYPE_INT            Attribute = FILE_TS_INT + 1
	FILE_PATH_STR               Attribute = PROC_CONTAINERID_STRING_STR + 1
	FILE_CONTAINERID_STRING_STR Attribute = FILE_PATH_STR + 1

	//Type name:  FileEvent
	EV_FILE_TS_INT      Attribute = FILE_RESTYPE_INT + 1
	EV_FILE_TID_INT     Attribute = EV_FILE_TS_INT + 1
	EV_FILE_OPFLAGS_INT Attribute = EV_FILE_TID_INT + 1
	EV_FILE_RET_INT     Attribute = EV_FILE_OPFLAGS_INT + 1

	//Type name:  File  number 2
	SEC_FILE_STATE_INT              Attribute = EV_FILE_RET_INT + 1
	SEC_FILE_TS_INT                 Attribute = SEC_FILE_STATE_INT + 1
	SEC_FILE_RESTYPE_INT            Attribute = SEC_FILE_TS_INT + 1
	SEC_FILE_PATH_STR               Attribute = FILE_CONTAINERID_STRING_STR + 1
	SEC_FILE_CONTAINERID_STRING_STR Attribute = SEC_FILE_PATH_STR + 1

	//Type name:  FileFlow
	FL_FILE_TS_INT            Attribute = FILE_RESTYPE_INT + 1
	FL_FILE_TID_INT           Attribute = FL_FILE_TS_INT + 1
	FL_FILE_OPFLAGS_INT       Attribute = FL_FILE_TID_INT + 1
	FL_FILE_OPENFLAGS_INT     Attribute = FL_FILE_OPFLAGS_INT + 1
	FL_FILE_ENDTS_INT         Attribute = FL_FILE_OPENFLAGS_INT + 1
	FL_FILE_FD_INT            Attribute = FL_FILE_ENDTS_INT + 1
	FL_FILE_NUMRRECVOPS_INT   Attribute = FL_FILE_FD_INT + 1
	FL_FILE_NUMWSENDOPS_INT   Attribute = FL_FILE_NUMRRECVOPS_INT + 1
	FL_FILE_NUMRRECVBYTES_INT Attribute = FL_FILE_NUMWSENDOPS_INT + 1
	FL_FILE_NUMWSENDBYTES_INT Attribute = FL_FILE_NUMRRECVBYTES_INT + 1

	//Type name:  NetworkFlow
	FL_NETW_TS_INT            Attribute = PROC_TTY_INT + 1
	FL_NETW_TID_INT           Attribute = FL_NETW_TS_INT + 1
	FL_NETW_OPFLAGS_INT       Attribute = FL_NETW_TID_INT + 1
	FL_NETW_ENDTS_INT         Attribute = FL_NETW_OPFLAGS_INT + 1
	FL_NETW_SIP_INT           Attribute = FL_NETW_ENDTS_INT + 1
	FL_NETW_SPORT_INT         Attribute = FL_NETW_SIP_INT + 1
	FL_NETW_DIP_INT           Attribute = FL_NETW_SPORT_INT + 1
	FL_NETW_DPORT_INT         Attribute = FL_NETW_DIP_INT + 1
	FL_NETW_PROTO_INT         Attribute = FL_NETW_DPORT_INT + 1
	FL_NETW_FD_INT            Attribute = FL_NETW_PROTO_INT + 1
	FL_NETW_NUMRRECVOPS_INT   Attribute = FL_NETW_FD_INT + 1
	FL_NETW_NUMWSENDOPS_INT   Attribute = FL_NETW_NUMRRECVOPS_INT + 1
	FL_NETW_NUMRRECVBYTES_INT Attribute = FL_NETW_NUMWSENDOPS_INT + 1
	FL_NETW_NUMWSENDBYTES_INT Attribute = FL_NETW_NUMRRECVBYTES_INT + 1

	//Type name:  ProcessEvent
	EV_PROC_TS_INT      Attribute = PROC_TTY_INT + 1
	EV_PROC_TID_INT     Attribute = EV_PROC_TS_INT + 1
	EV_PROC_OPFLAGS_INT Attribute = EV_PROC_TID_INT + 1
	EV_PROC_RET_INT     Attribute = EV_PROC_OPFLAGS_INT + 1

	INT_ARRAY_SIZE Attribute = 27
	STR_ARRAY_SIZE Attribute = 14

	HEADER    int64 = 0
	CONT      int64 = 1
	PROC      int64 = 2
	FILE      int64 = 3
	PROC_EVT  int64 = 4
	NET_FLOW  int64 = 5
	FILE_FLOW int64 = 6
	FILE_EVT  int64 = 7
)

// FlatRecord represents a flattened SysFlow record
type FlatRecord struct {
	Ints [INT_ARRAY_SIZE]int64
	Strs [STR_ARRAY_SIZE]string
}

type FlatRecordPool struct {
	pool *sync.Pool
}

func NewFlatRecordPool() *FlatRecordPool {
	p := new(FlatRecordPool)
	p.pool = &sync.Pool{
		New: func() interface{} {
			return new(FlatRecord)
		},
	}
	return p
}

func (f *FlatRecordPool) Get() *FlatRecord {
	return f.pool.Get().(*FlatRecord)
}

func (f *FlatRecordPool) Put(fr *FlatRecord) {
	f.pool.Put(fr)
}
