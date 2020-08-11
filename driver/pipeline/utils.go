package pipeline

import (
	"strconv"

	"github.ibm.com/sysflow/sf-processor/driver/manifest"
)

const (
	// ChanSize is the default size of the go channels in the pipeline
	ChanSize = 100000
)

// setManifestInfo sets manifest attributes to plugins configuration items.
func setManifestInfo(conf *Config) {
	addGlobalConfigItem(conf, manifest.VersionKey, manifest.Version)
	addGlobalConfigItem(conf, manifest.JSONSchemaVersionKey, manifest.JSONSchemaVersion)
	addGlobalConfigItem(conf, manifest.BuildNumberKey, manifest.BuildNumber)
}

// addGlobalConfigItem adds a config item to all processors in the pipeline.
func addGlobalConfigItem(conf *Config, k string, v interface{}) {
	for _, c := range conf.Pipeline {
		if _, ok := c[ProcConfig]; ok {
			if s, ok := v.(string); ok {
				c[k] = s
			} else if i, ok := v.(int); ok {
				c[k] = strconv.Itoa(i)
			}
		}
	}
}
