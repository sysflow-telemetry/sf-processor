package pipeline

type Plugin map[string]string

type PluginConfig struct {
	Pipeline []Plugin `json,mapstructures:"pipeline"`
}
