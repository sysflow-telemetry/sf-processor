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
//
package engine

import "errors"

// Configuration keys.
const (
	PoliciesConfigKey    string = "policies"
	ModeConfigKey        string = "mode"
	VersionKey           string = "version"
	JSONSchemaVersionKey string = "jsonschemaversion"
	BuildNumberKey       string = "buildnumber"
)

// Config defines a configuration object for the engine.
type Config struct {
	PoliciesPath      string
	Mode              Mode
	Version           string
	JSONSchemaVersion string
	BuildNumber       string
}

// CreateConfig creates a new config object from config dictionary.
func CreateConfig(conf map[string]string) (Config, error) {
	var c Config = Config{Mode: AlertMode} // default values
	if v, ok := conf[PoliciesConfigKey]; ok {
		c.PoliciesPath = v
	} else {
		return c, errors.New("Configuration tag 'policies' missing from policy engine plugin settings")
	}
	if v, ok := conf[ModeConfigKey]; ok {
		c.Mode = parseModeConfig(v)
	}
	if v, ok := conf[VersionKey]; ok {
		c.Version = v
	}
	if v, ok := conf[JSONSchemaVersionKey]; ok {
		c.JSONSchemaVersion = v
	}
	if v, ok := conf[BuildNumberKey]; ok {
		c.BuildNumber = v
	}
	return c, nil
}

// Mode type.
type Mode int

// Mode config options.
const (
	AlertMode Mode = iota
	FilterMode
)

func (s Mode) String() string {
	return [...]string{"alert", "filter"}[s]
}

func parseModeConfig(s string) Mode {
	if AlertMode.String() == s {
		return AlertMode
	}
	if FilterMode.String() == s {
		return FilterMode
	}
	return AlertMode
}
