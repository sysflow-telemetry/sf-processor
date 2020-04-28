package engine

// Configuration keys.
const (
	PoliciesConfigKey   string = "policies"
	ContRuntimeType     string = "runtime"
	ContRuntimeEndpoint string = "runtimeEndpoint"
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
