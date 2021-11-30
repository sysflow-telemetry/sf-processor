//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
// Andreas Schade <san@zurich.ibm.com>
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

// Package engine implements a rules engine for telemetry records.
package engine

import (
	"errors"
	"strconv"
)

// Configuration keys.
const (
	PoliciesConfigKey    string = "policies"
	ModeConfigKey	string = "mode"
	VersionKey	   string = "version"
	JSONSchemaVersionKey string = "jsonschemaversion"
	BuildNumberKey       string = "buildnumber"
	MonitorKey	   string = "monitor"
	ConcurrencyKey       string = "concurrency"
	ActionDirKey	     string = "action_dir"
	ContainerConfigKey   string = "container_config"
	MaxFileSizeKey       string = "max_file_size"
	HashCacheExpKey      string = "hash_cache_expiration"
	HashCachePurgeKey    string = "hash_cache_purge"
)

// Config defines a configuration object for the engine.
type Config struct {
	PoliciesPath      string
	Mode	      Mode
	Version	   string
	JSONSchemaVersion string
	BuildNumber       string
	Monitor	   MonitorType
	Concurrency       int
	ActionDir	 string
	ContainerConfig   string
	MaxFileSize       int64
	HashCacheExp      int
	HashCachePurge    int
}

// Defaults
const (
	// Concurrency of policy engine
	CONCURRENCY   = 5

	// Max hashing file size default is 256 MiB
	MAX_FILE_SIZE = 1 << 28

	// Hash cache entry expiration
	EXPIRATION    = 5

	// Hash cache purge of expired entries
	PURGE	 = 7
)

// CreateConfig creates a new config object from config dictionary.
func CreateConfig(conf map[string]interface{}) (Config, error) {
	var c Config = Config{Mode: AlertMode} // default values
	var err error

	if v, ok := conf[PoliciesConfigKey].(string); ok {
		c.PoliciesPath = v
	}
	if v, ok := conf[ModeConfigKey].(string); ok {
		c.Mode = parseModeConfig(v)
	}
	if v, ok := conf[VersionKey].(string); ok {
		c.Version = v
	}
	if v, ok := conf[JSONSchemaVersionKey].(string); ok {
		c.JSONSchemaVersion = v
	}
	if v, ok := conf[BuildNumberKey].(string); ok {
		c.BuildNumber = v
	}
	if v, ok := conf[ActionDirKey].(string); ok {
		c.ActionDir = v
	}
	if v, ok := conf[ContainerConfigKey].(string); ok {
		c.ContainerConfig = v
	}
	c.Concurrency = CONCURRENCY
	if v, ok := conf[ConcurrencyKey].(string); ok {
		c.Concurrency, err = strconv.Atoi(v)
	}
	c.MaxFileSize = MAX_FILE_SIZE
	if v, ok := conf[MaxFileSizeKey].(string); ok {
		c.MaxFileSize, err = strconv.ParseInt(v, 10, 64)
	}
	c.HashCacheExp = EXPIRATION
	if v, ok := conf[HashCacheExpKey].(string); ok {
		c.HashCacheExp, err = strconv.Atoi(v)
	}
	c.HashCachePurge = PURGE
	if v, ok := conf[HashCachePurgeKey].(string); ok {
		c.HashCachePurge, err = strconv.Atoi(v)
	}
	c.Monitor = NoneType
	if v, ok := conf[MonitorKey].(string); ok {
		if v == "local" {
			c.Monitor = LocalType
		} else if v == "none" {
			c.Monitor = NoneType
		} else {
			err = errors.New("Configuration tag 'monitor' must be set to 'none', 'local'")
		}
	}

	return c, err
}

// Mode type.
type Mode int

// Mode config options.
const (
	EnrichMode Mode = iota
	AlertMode
)

func (s Mode) String() string {
	return [...]string{"enrich", "alert", "filter"}[s]
}

func parseModeConfig(s string) Mode {
	if EnrichMode.String() == s {
		return EnrichMode
	}
	if AlertMode.String() == s {
		return AlertMode
	}
	return EnrichMode
}

// MonitorType defines a policy monitor type.
type MonitorType uint32

// Monitor types.
const (
	NoneType MonitorType = iota
	LocalType
)

func (s MonitorType) String() string {
	return [...]string{"none", "local"}[s]
}
