package engine

// Configuration keys.
const (
	PoliciesConfigKey            string = "policies"
	ModeConfigKey                string = "mode"
	ContRuntimeTypeConfigKey     string = "runtime"
	ContRuntimeEndpointConfigKey string = "runtimeEndpoint"
	EnrichmentConfigKey          string = "enrichmentconfigpath"
)

// Mode constants.
const (
	AlertMode  string = "alert"
	FilterMode string = "filter"
)

// Runtime constants.
const (
	Docker     string = "docker"
	Crio       string = "crio"
	Containerd string = "containerd"
)

// Parsing constants.
const (
	LISTSEP string = ","
	EMPTY   string = ""
	QUOTE   string = "\""
	SPACE   string = " "
)

// SysFlow object types.
const (
	TyP      string = "P"
	TyF      string = "F"
	TyC      string = "C"
	TyH      string = "H"
	TyPE     string = "PE"
	TyFE     string = "FE"
	TyFF     string = "FF"
	TyNF     string = "NF"
	TyUnknow string = ""
)

// SysFlow attribute names.
const (
	SF_TYPE                 string = "sf.type"
	SF_OPFLAGS              string = "sf.opflags"
	SF_RET                  string = "sf.ret"
	SF_TS                   string = "sf.ts"
	SF_ENDTS                string = "sf.endts"
	SF_PROC_OID             string = "sf.proc.oid"
	SF_PROC_PID             string = "sf.proc.pid"
	SF_PROC_NAME            string = "sf.proc.name"
	SF_PROC_EXE             string = "sf.proc.exe"
	SF_PROC_ARGS            string = "sf.proc.args"
	SF_PROC_UID             string = "sf.proc.uid"
	SF_PROC_USER            string = "sf.proc.user"
	SF_PROC_TID             string = "sf.proc.tid"
	SF_PROC_GID             string = "sf.proc.gid"
	SF_PROC_GROUP           string = "sf.proc.group"
	SF_PROC_CREATETS        string = "sf.proc.createts"
	SF_PROC_DURATION        string = "sf.proc.duration"
	SF_PROC_TTY             string = "sf.proc.tty"
	SF_PROC_CMDLINE         string = "sf.proc.cmdline"
	SF_PROC_ANAME           string = "sf.proc.aname"
	SF_PROC_AEXE            string = "sf.proc.aexe"
	SF_PROC_ACMDLINE        string = "sf.proc.acmdline"
	SF_PROC_APID            string = "sf.proc.apid"
	SF_PPROC_OID            string = "sf.pproc.oid"
	SF_PPROC_PID            string = "sf.pproc.pid"
	SF_PPROC_NAME           string = "sf.pproc.name"
	SF_PPROC_EXE            string = "sf.pproc.exe"
	SF_PPROC_ARGS           string = "sf.pproc.args"
	SF_PPROC_UID            string = "sf.pproc.uid"
	SF_PPROC_USER           string = "sf.pproc.user"
	SF_PPROC_GID            string = "sf.pproc.gid"
	SF_PPROC_GROUP          string = "sf.pproc.group"
	SF_PPROC_CREATETS       string = "sf.pproc.createts"
	SF_PPROC_DURATION       string = "sf.pproc.duration"
	SF_PPROC_TTY            string = "sf.pproc.tty"
	SF_PPROC_CMDLINE        string = "sf.pproc.cmdline"
	SF_FILE_NAME            string = "sf.file.name"
	SF_FILE_PATH            string = "sf.file.path"
	SF_FILE_CANONICALPATH   string = "sf.file.canonicalpath"
	SF_FILE_DIRECTORY       string = "sf.file.directory"
	SF_FILE_NEWNAME         string = "sf.file.newname"
	SF_FILE_NEWPATH         string = "sf.file.newpath"
	SF_FILE_NEWDIRECTORY    string = "sf.file.newdirectory"
	SF_FILE_TYPE            string = "sf.file.type"
	SF_FILE_IS_OPEN_WRITE   string = "sf.file.is_open_write"
	SF_FILE_IS_OPEN_READ    string = "sf.file.is_open_read"
	SF_FILE_FD              string = "sf.file.fd"
	SF_FILE_OPENFLAGS       string = "sf.file.openflags"
	SF_NET_PROTO            string = "sf.net.proto"
	SF_NET_PROTONAME        string = "sf.net.protoname"
	SF_NET_SPORT            string = "sf.net.sport"
	SF_NET_DPORT            string = "sf.net.dport"
	SF_NET_PORT             string = "sf.net.port"
	SF_NET_SIP              string = "sf.net.sip"
	SF_NET_DIP              string = "sf.net.dip"
	SF_NET_IP               string = "sf.net.ip"
	SF_FLOW_RBYTES          string = "sf.flow.rbytes"
	SF_FLOW_ROPS            string = "sf.flow.rops"
	SF_FLOW_WBYTES          string = "sf.flow.wbytes"
	SF_FLOW_WOPS            string = "sf.flow.wops"
	SF_CONTAINER_ID         string = "sf.container.id"
	SF_CONTAINER_NAME       string = "sf.container.name"
	SF_CONTAINER_IMAGEID    string = "sf.container.imageid"
	SF_CONTAINER_IMAGE      string = "sf.container.image"
	SF_CONTAINER_TYPE       string = "sf.container.type"
	SF_CONTAINER_PRIVILEGED string = "sf.container.privileged"
	SF_NODE_ID              string = "sf.node.id"
)
