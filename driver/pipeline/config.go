//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
package pipeline

// Config attributes
const (
	ModConfig     string = "mod"
	ProcConfig    string = "processor"
	HdlConfig     string = "handler"
	InChanConfig  string = "in"
	OutChanConfig string = "out"
)

// PluginConfig defines a map for plugin configuration
type PluginConfig map[string]string

// Config defines a pipeline configuration object
type Config struct {
	Pipeline []PluginConfig `json,mapstructures:"pipeline"`
}
