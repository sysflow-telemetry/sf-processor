package pipeline

// PluginConfig defines a map for plugin configuration.
type PluginConfig map[string]string

// Config defines a pipeline configuration object.
type Config struct {
	Pipeline []PluginConfig `json,mapstructures:"pipeline"`
}
