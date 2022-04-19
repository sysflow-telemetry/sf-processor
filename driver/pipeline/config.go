//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package pipeline implements a pluggable data processing pipeline infrastructure.
package pipeline

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/driver/manifest"
)

// Global config variables
const (
	ClusterIDEnvKey string = "CLUSTER_ID"
)

// Config attributes
const (
	ModConfig     string = "mod"
	ProcConfig    string = "processor"
	HdlConfig     string = "handler"
	InChanConfig  string = "in"
	OutChanConfig string = "out"
	DrivConfig    string = "driver"
)

// Driver constants/defaults
const (
	SockFile            = "/var/run/sysflow.sock"
	BuffSize            = 16384
	OOBuffSize          = 1024
	DriverDir           = "../resources/drivers"
	PluginDir           = "../resources/plugins"
	ChanSize            = 100000
	HealthChecksTimeout = 10 * time.Second
)

// PluginConfig defines a map for plugin configuration
type PluginConfig map[string]interface{}

// DriverConfig defines a map for plugin configuration
type DriverConfig map[string]interface{}

// Config defines a pipeline configuration object
type Config struct {
	Drivers  []DriverConfig `json,mapstructures:"drivers"`
	Pipeline []PluginConfig `json,mapstructures:"pipeline"`
}

// setManifestInfo sets manifest attributes to plugins configuration items.
func setManifestInfo(conf *Config) {
	addGlobalConfigItem(conf, manifest.VersionKey, manifest.Version)                     //nolint:typecheck
	addGlobalConfigItem(conf, manifest.JSONSchemaVersionKey, manifest.JSONSchemaVersion) //nolint:typecheck
	addGlobalConfigItem(conf, manifest.EcsVersionKey, manifest.EcsVersion)               //nolint:typecheck
	addGlobalConfigItem(conf, manifest.BuildNumberKey, manifest.BuildNumber)             //nolint:typecheck
	addGlobalConfigItem(conf, commons.ClusterIDKey, getEnv(ClusterIDEnvKey))
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

// getEnv retrieves the environment varible for a key.
func getEnv(k string) string {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == k && len(pair) == 2 {
			return pair[1]
		}
	}
	return sfgo.Zeros.String
}
