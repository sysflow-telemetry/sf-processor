package commons

// Configuration keys.
const (
	PathConfigKey string = "file.path"
)

// FileConfig holds file output specific configuration.
type FileConfig struct {
	Path string
}

// CreateSyslogConfig creates a new config object from config dictionary.
func CreateFileConfig(bc Config, conf map[string]interface{}) (c FileConfig, err error) {
	// default values
	c = FileConfig{Path: "./export.out"}

	// parse config map
	if v, ok := conf[PathConfigKey].(string); ok {
		c.Path = v
	}
	return
}
