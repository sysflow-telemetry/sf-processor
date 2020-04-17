package pipeline

// Plugin denotes the type of a plugin.
type Plugin map[string]string

// PluginConfig defines a plugin configuration object.
type PluginConfig struct {
	Pipeline []Plugin `json,mapstructures:"pipeline"`
}
