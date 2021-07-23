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

// Package engine implements a rules engine for telemetry records.
package engine

import (
	"errors"
)

// Configuration keys.
const (
	PoliciesConfigKey    string = "policies"
	ModeConfigKey        string = "mode"
	VersionKey           string = "version"
	JSONSchemaVersionKey string = "jsonschemaversion"
	BuildNumberKey       string = "buildnumber"
	MonitorKey           string = "monitor"
)

// Config defines a configuration object for the engine.
type Config struct {
	PoliciesPath      string
	Mode              Mode
	Version           string
	JSONSchemaVersion string
	BuildNumber       string
	Monitor           MonitorType
}

// CreateConfig creates a new config object from config dictionary.
func CreateConfig(conf map[string]interface{}) (Config, error) {
	var c Config = Config{Mode: AlertMode} // default values

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
	c.Monitor = NoneType
	if v, ok := conf[MonitorKey].(string); ok {
		if v == "local" {
			c.Monitor = LocalType
		} else if v == "none" {
			c.Monitor = NoneType
		} else {
			return c, errors.New("Configuration tag 'monitor' must be set to 'none', 'local'")
		}
	}
	return c, nil
}

// Mode type.
type Mode int

// Mode config options.
const (
	AlertMode Mode = iota
	FilterMode
	BypassMode
)

func (s Mode) String() string {
	return [...]string{"alert", "filter", "bypass"}[s]
}

func parseModeConfig(s string) Mode {
	if AlertMode.String() == s {
		return AlertMode
	}
	if FilterMode.String() == s {
		return FilterMode
	}
	if BypassMode.String() == s {
		return BypassMode
	}
	return AlertMode
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
