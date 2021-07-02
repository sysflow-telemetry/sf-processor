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

// Package commons defines common facilities for exporters.
package commons

import (
	"strconv"
)

// Configuration keys.
const (
	ProtoConfigKey     string = "syslog.proto"
	TagConfigKey       string = "syslog.tag"
	LogSourceConfigKey string = "syslog.source"
	HostConfigKey      string = "syslog.host"
	PortConfigKey      string = "syslog.port"
)

// SyslogConfig holds rsyslog specific configuration.
type SyslogConfig struct {
	Proto     Proto
	Tag       string
	LogSource string
	Host      string
	Port      int
}

// CreateSyslogConfig creates a new config object from config dictionary.
func CreateSyslogConfig(bc Config, conf map[string]interface{}) (c SyslogConfig, err error) {
	// default values
	c = SyslogConfig{
		Host: "localhost",
		Port: 514,
		Tag:  "sysflow"}

	// parse config map
	if v, ok := conf[ProtoConfigKey].(string); ok {
		c.Proto = parseProtoConfig(v)
	}
	if v, ok := conf[TagConfigKey].(string); ok {
		c.Tag = v
	}
	if v, ok := conf[LogSourceConfigKey].(string); ok {
		c.LogSource = v
	}
	if v, ok := conf[HostConfigKey].(string); ok {
		c.Host = v
	}
	if v, ok := conf[PortConfigKey].(string); ok {
		c.Port, err = strconv.Atoi(v)
		if err != nil {
			return c, err
		}
	}
	return
}
