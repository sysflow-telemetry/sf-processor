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

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/secrets"
)

// Configuration keys.
const (
	ExportConfigKey        string = "export"
	ExpTypeConfigKey       string = "type"
	FormatConfigKey        string = "format"
	FlatConfigKey          string = "flat"
	VaultEnabledConfigKey  string = "vault.secrets"
	VaultPathConfigKey     string = "vault.path"
	ProtoConfigKey         string = "proto"
	TagConfigKey           string = "tag"
	LogSourceConfigKey     string = "source"
	HostConfigKey          string = "host"
	PortConfigKey          string = "port"
	PathConfigKey          string = "path"
	ESAddressesConfigKey   string = "es.addresses"
	ESIndexConfigKey       string = "es.index"
	ESUsernameConfigKey    string = "es.username"
	ESPasswordConfigKey    string = "es.password"
	ESWorkersConfigKey     string = "es.bulk.numWorkers"
	ESFBufferConfigKey     string = "es.bulk.flushBuffer"
	ESFTimeoutConfigKey    string = "es.bulk.flushTimeout"
	SAApiKeyConfigKey      string = "sa.apikey"
	SAUrlConfigKey         string = "sa.url"
	SAAccountIDConfigKey   string = "sa.accountid"
	SAProviderIDConfigKey  string = "sa.provider"
	SANoteIDConfigKey      string = "sa.note"
	SASqlQueryUrlConifgKey string = "sa.sqlqueryurl"
	SASqlQueryCrnConfigKey string = "sa.sqlquerycrn"
	RegionConfigKey        string = "region"
	EventBufferConfigKey   string = "buffer"
	VersionKey             string = "version"
	JSONSchemaVersionKey   string = "jsonschemaversion"
	BuildNumberKey         string = "buildnumber"
)

// Config defines a configuration object for the exporter.
type Config struct {
	Export            Export
	ExpType           ExportType
	Format            Format
	Flat              bool
	VaultEnabled      bool
	VaultMountPath    string
	secrets           *secrets.Secrets
	Proto             Proto
	Tag               string
	LogSource         string
	Host              string
	Port              int
	Path              string
	ESAddresses       []string
	ESIndex           string
	ESUsername        string
	ESPassword        string
	ESNumWorkers      int
	ESFlushBuffer     int
	ESFlushTimeout    time.Duration
	SAApiKey          string
	SAUrl             string
	SAAccountID       string
	SAProviderID      string
	SANoteID          string
	SASqlQueryUrl     string
	SASqlQueryCrn     string
	Region            string
	EventBuffer       int
	Version           string
	JSONSchemaVersion string
	BuildNumber       string
}

// CreateConfig creates a new config object from config dictionary.
func CreateConfig(conf map[string]interface{}) (Config, error) {
	// default values
	var c Config = Config{
		Host:           "localhost",
		Port:           514,
		Path:           "./export.out",
		Tag:            "sysflow",
		ESNumWorkers:   0,
		ESFlushBuffer:  5e+6,
		ESFlushTimeout: 30 * time.Second,
		SAUrl:          "https://us-south.secadvisor.cloud.ibm.com/findings",
		SASqlQueryUrl:  "https://us.sql-query.cloud.ibm.com/sqlquery"}

	// wrapper for reading from secrets vault
	if v, ok := conf[VaultEnabledConfigKey].(string); ok && v == "true" {
		c.VaultEnabled = true
		var s *secrets.Secrets
		var err error
		if p, ok := conf[VaultPathConfigKey].(string); ok {
			s, err = secrets.NewSecretsWithCustomPath(p)
		} else {
			s, err = secrets.NewSecrets()
		}
		if err != nil {
			logger.Error.Printf("Could not read secrets from vault: %v", err)
			return c, err
		}
		c.secrets = s
	}

	if v, ok := conf[ExportConfigKey].(string); ok {
		c.Export = parseExportConfig(v)
	}
	if v, ok := conf[ExpTypeConfigKey].(string); ok {
		c.ExpType = parseExportTypeConfig(v)
	}
	if v, ok := conf[FormatConfigKey].(string); ok {
		c.Format = parseFormatConfig(v)
	}
	if v, ok := conf[FlatConfigKey].(string); ok && v == "true" {
		c.Flat = true
	}
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
		c.Port, _ = strconv.Atoi(v)
	}
	if v, ok := conf[PathConfigKey].(string); ok {
		c.Path = v
	}
	if v, ok := conf[ESAddressesConfigKey].(string); ok {
		c.ESAddresses = strings.Split(v, ",")
	}
	if v, ok := conf[ESIndexConfigKey].(string); ok {
		c.ESIndex = v
	}
	if v, ok := conf[ESUsernameConfigKey].(string); ok {
		c.ESUsername = v
	} else if c.VaultEnabled {
		s, err := c.secrets.GetDecoded(ESUsernameConfigKey)
		if err != nil {
			logger.Error.Printf("Could not read secret %s from vault: %v", ESUsernameConfigKey, err)
		}
		c.ESUsername = string(s)
	}
	if v, ok := conf[ESPasswordConfigKey].(string); ok {
		c.ESPassword = v
	} else if c.VaultEnabled {
		s, err := c.secrets.GetDecoded(ESPasswordConfigKey)
		if err != nil {
			logger.Error.Printf("Could not read secret %s from vault: %v", ESPasswordConfigKey, err)
		}
		c.ESPassword = string(s)
	}
	if v, ok := conf[ESWorkersConfigKey].(string); ok {
		c.ESNumWorkers, _ = strconv.Atoi(v)
	}
	if v, ok := conf[ESFBufferConfigKey].(string); ok {
		c.ESFlushBuffer, _ = strconv.Atoi(v)
	}
	if v, ok := conf[ESFTimeoutConfigKey].(string); ok {
		c.ESFlushTimeout, _ = time.ParseDuration(v)
	}
	if v, ok := conf[SAApiKeyConfigKey].(string); ok {
		c.SAApiKey = v
	} else if c.VaultEnabled {
		s, err := c.secrets.GetDecoded(SAApiKeyConfigKey)
		if err != nil {
			logger.Error.Printf("Could not read secret %s from vault: %v", SAApiKeyConfigKey, err)
		}
		c.SAApiKey = string(s)
	}
	if v, ok := conf[SAAccountIDConfigKey].(string); ok {
		c.SAAccountID = v
	}
	if v, ok := conf[SAUrlConfigKey].(string); ok {
		c.SAUrl = v
	}
	if v, ok := conf[SAProviderIDConfigKey].(string); ok {
		c.SAProviderID = v
	}
	if v, ok := conf[SANoteIDConfigKey].(string); ok {
		c.SANoteID = v
	}
	if v, ok := conf[SASqlQueryUrlConifgKey].(string); ok {
		c.SASqlQueryUrl = v
	}
	if v, ok := conf[SASqlQueryCrnConfigKey].(string); ok {
		c.SASqlQueryCrn = v
	}
	if v, ok := conf[RegionConfigKey].(string); ok {
		c.Region = v
	}
	if v, ok := conf[EventBufferConfigKey].(string); ok {
		c.EventBuffer, _ = strconv.Atoi(v)
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
	return c, nil
}

// Export type.
type Export int

// Export config options.
const (
	StdOutExport Export = iota
	FileExport
	SyslogExport
	ESExport
	SAExport
)

func (s Export) String() string {
	return [...]string{"terminal", "file", "syslog", "es", "sa"}[s]
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
	if SAExport.String() == s {
		return SAExport
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
