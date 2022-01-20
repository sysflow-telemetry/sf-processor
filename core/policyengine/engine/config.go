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
	"strconv"
	"time"
)

// Configuration keys.
const (
	PoliciesConfigKey    string = "policies"
	ModeConfigKey        string = "mode"
	VersionKey           string = "version"
	JSONSchemaVersionKey string = "jsonschemaversion"
	BuildNumberKey       string = "buildnumber"
	MonitorKey           string = "monitor"
	MonitorIntervalKey   string = "monitorinterval"
	ConcurrencyKey       string = "concurrency"
	ActionDirKey         string = "action_dir"
)

// Config defines a configuration object for the engine.
type Config struct {
	PoliciesPath      string
	Mode              Mode
	Version           string
	JSONSchemaVersion string
	BuildNumber       string
	Monitor           MonitorType
	MonitorInterval   time.Duration
	Concurrency       int
	ActionDir         string
}

// CreateConfig creates a new config object from config dictionary.
func CreateConfig(conf map[string]interface{}) (Config, error) {
	var c Config = Config{Mode: EnrichMode, Concurrency: 5, Monitor: NoneType, MonitorInterval: 30 * time.Second} // default values
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
	if v, ok := conf[MonitorKey].(string); ok {
		c.Monitor = parseMonitorType(v)
	}
	if v, ok := conf[MonitorIntervalKey].(string); ok {
		var duration int
		duration, err = strconv.Atoi(v)
		if err == nil {
			c.MonitorInterval = time.Duration(duration) * time.Second
		}
	}
	if v, ok := conf[ConcurrencyKey].(string); ok {
		c.Concurrency, err = strconv.Atoi(v)
	}
	if v, ok := conf[ActionDirKey].(string); ok {
		c.ActionDir = v
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
	return [...]string{"enrich", "alert"}[s]
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

func parseMonitorType(s string) MonitorType {
	if NoneType.String() == s {
		return NoneType
	}
	if LocalType.String() == s {
		return LocalType
	}
	return NoneType
}
