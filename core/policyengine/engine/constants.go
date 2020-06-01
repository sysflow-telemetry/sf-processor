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
