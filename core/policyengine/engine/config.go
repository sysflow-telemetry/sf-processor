//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
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
