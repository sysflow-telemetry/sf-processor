package pipeline

// Plugin defines a map for plugin configuration.
type Plugin map[string]string

// PluginConfig defines a plugin configuration object.
type PluginConfig struct {
	Pipeline []Plugin `json,mapstructures:"pipeline"`
}
