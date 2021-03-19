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
package exporter

import (
	"strconv"
	"strings"
	"time"
)

// Configuration keys.
const (
	ExportConfigKey      string = "export"
	ExpTypeConfigKey     string = "type"
	FormatConfigKey      string = "format"
	FlatConfigKey        string = "flat"
	ProtoConfigKey       string = "proto"
	TagConfigKey         string = "tag"
	LogSourceConfigKey   string = "source"
	HostConfigKey        string = "host"
	PortConfigKey        string = "port"
	PathConfigKey        string = "path"
        ESAddressesConfigKey string = "es.addresses"
	ESIndexConfigKey     string = "es.index"
	ESUsernameConfigKey  string = "es.username"
	ESPasswordConfigKey  string = "es.password"
	ESWorkersConfigKey   string = "es.bulk.numWorkers"
	ESFBufferConfigKey   string = "es.bulk.flushBuffer"
	ESFTimeoutConfigKey  string = "es.bulk.flushTimeout"
	EventBufferConfigKey string = "buffer"
	VersionKey           string = "version"
	JSONSchemaVersionKey string = "jsonschemaversion"
	BuildNumberKey       string = "buildnumber"
)

// Config defines a configuration object for the exporter.
type Config struct {
	Export            Export
	ExpType           ExportType
	Format            Format
	Flat              bool
	Proto             Proto
	Tag               string
	LogSource         string
	Host              string
	Port              int
	Path              string
	ESAddresses	  []string
	ESIndex           string
	ESUsername        string
	ESPassword        string
	ESNumWorkers      int
	ESFlushBuffer     int
	ESFlushTimeout    time.Duration
	EventBuffer       int
	Version           string
	JSONSchemaVersion string
	BuildNumber       string
}

// CreateConfig creates a new config object from config dictionary.
func CreateConfig(conf map[string]string) Config {
	// default values
	var c Config = Config{
			Host: "localhost",
			Port: 514,
			Path: "./export.out",
			Tag: "sysflow",
			ESNumWorkers: 0,
			ESFlushBuffer: 5e+6,
			ESFlushTimeout: 30 * time.Second}

	if v, ok := conf[ExportConfigKey]; ok {
		c.Export = parseExportConfig(v)
	}
	if v, ok := conf[ExpTypeConfigKey]; ok {
		c.ExpType = parseExportTypeConfig(v)
	}
	if v, ok := conf[FormatConfigKey]; ok {
		c.Format = parseFormatConfig(v)
	}
	if v, ok := conf[FlatConfigKey]; ok && v == "true" {
		c.Flat = true
	}
	if v, ok := conf[ProtoConfigKey]; ok {
		c.Proto = parseProtoConfig(v)
	}
	if v, ok := conf[TagConfigKey]; ok {
		c.Tag = v
	}
	if v, ok := conf[LogSourceConfigKey]; ok {
		c.LogSource = v
	}
	if v, ok := conf[HostConfigKey]; ok {
		c.Host = v
	}
	if v, ok := conf[PortConfigKey]; ok {
		c.Port, _ = strconv.Atoi(v)
	}
	if v, ok := conf[PathConfigKey]; ok {
		c.Path = v
	}
	if v, ok := conf[ESAddressesConfigKey]; ok {
		c.ESAddresses = strings.Split(v, ",")
	}
	if v, ok := conf[ESIndexConfigKey]; ok {
		c.ESIndex = v
	}
	if v, ok := conf[ESUsernameConfigKey]; ok {
		c.ESUsername = v
	}
	if v, ok := conf[ESPasswordConfigKey]; ok {
		c.ESPassword = v
	}
	if v, ok := conf[ESWorkersConfigKey]; ok {
		c.ESNumWorkers, _ = strconv.Atoi(v)
	}
	if v, ok := conf[ESFBufferConfigKey]; ok {
		c.ESFlushBuffer, _ = strconv.Atoi(v)
	}
	if v, ok := conf[ESFTimeoutConfigKey]; ok {
		c.ESFlushTimeout, _ = time.ParseDuration(v)
	}
	if v, ok := conf[EventBufferConfigKey]; ok {
		c.EventBuffer, _ = strconv.Atoi(v)
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
	return c
}

// Export type.
type Export int

// Export config options.
const (
	StdOutExport Export = iota
	FileExport
	SyslogExport
	ESExport
)

func (s Export) String() string {
	return [...]string{"terminal", "file", "syslog", "es"}[s]
}

func parseExportConfig(s string) Export {
	if FileExport.String() == s {
		return FileExport
	}
	if SyslogExport.String() == s {
		return SyslogExport
	}
	if ESExport.String() == s {
		return ESExport
	}
	return StdOutExport
}

// ExportType type.
type ExportType int

// ExportType config options.
const (
	TelemetryType ExportType = iota
	BatchType
)

func (s ExportType) String() string {
	return [...]string{"telemetry", "batch"}[s]
}

func parseExportTypeConfig(s string) ExportType {
	if BatchType.String() == s {
		return BatchType
	}
	return TelemetryType
}

// Format type.
type Format int

// Format config options.
const (
	JSONFormat Format = iota
	ECSFormat
)

func (s Format) String() string {
	return [...]string{"json", "ecs"}[s]
}

func parseFormatConfig(s string) Format {
	switch s {
	case JSONFormat.String():
		return JSONFormat
	case ECSFormat.String():
		return ECSFormat
	}
	return JSONFormat
}

// Proto denotes protocol type.
type Proto int

// Proto config options.
const (
	TCPProto Proto = iota
	TCPTLSProto
	UDPProto
)

func (s Proto) String() string {
	return [...]string{"tcp", "tls", "udp"}[s]
}

func parseProtoConfig(s string) Proto {
	switch s {
	case TCPProto.String():
		return TCPProto
	case TCPTLSProto.String():
		return TCPTLSProto
	case UDPProto.String():
		return UDPProto
	}
	return TCPProto
}
