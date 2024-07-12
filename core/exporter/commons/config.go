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

	"github.com/sysflow-telemetry/sf-apis/go/secrets"
)

// Configuration keys.
const (
	TransportConfigKey     string = "export"
	FormatConfigKey        string = "format"
	VaultEnabledConfigKey  string = "vault.secrets"
	VaultPathConfigKey     string = "vault.path"
	VaultEncodingConfigKey string = "vault.encoding"
	EventBufferConfigKey   string = "buffer"
	VersionKey             string = "version"
	JSONSchemaVersionKey   string = "jsonschemaversion"
	EcsVersionKey          string = "ecsversion"
	BuildNumberKey         string = "buildnumber"
	ClusterIDKey           string = "cluster.id"
)

// Config defines a configuration object for the exporter.
type Config struct {
	Transport         Transport
	Format            Format
	EventBuffer       int
	VaultEnabled      bool
	VaultMountPath    string
	VaultEncoding     VaultEncoding
	secrets           *secrets.Secrets
	Version           string
	JSONSchemaVersion string
	EcsVersion        string
	BuildNumber       string
	ClusterID         string
	FileConfig
	SyslogConfig
	ESConfig
	KafkaConfig
}

// CreateConfig creates a new config object from config dictionary.
func CreateConfig(conf map[string]interface{}) (c Config, err error) {
	c = Config{}

	// wrapper for reading from secrets vault
	if v, ok := conf[VaultEnabledConfigKey].(string); ok && v == "true" {
		c.VaultEnabled = true
		if e, ok := conf[VaultEncodingConfigKey].(string); ok {
			c.VaultEncoding = parseVaultEncodingConfig(e)
		}
		var s *secrets.Secrets
		if p, ok := conf[VaultPathConfigKey].(string); ok {
			s, err = secrets.NewSecretsWithCustomPath(p)
		} else {
			s, err = secrets.NewSecrets()
		}
		if err != nil {
			return
		}
		c.secrets = s
	}

	// parse config map
	if v, ok := conf[TransportConfigKey].(string); ok {
		c.Transport = parseTransportConfig(v)
	}
	if v, ok := conf[FormatConfigKey].(string); ok {
		c.Format = parseFormatConfig(v)
	}
	if v, ok := conf[EventBufferConfigKey].(string); ok {
		c.EventBuffer, err = strconv.Atoi(v)
		if err != nil {
			return c, err
		}
	}
	if v, ok := conf[VersionKey].(string); ok {
		c.Version = v
	}
	if v, ok := conf[JSONSchemaVersionKey].(string); ok {
		c.JSONSchemaVersion = v
	}
	if v, ok := conf[EcsVersionKey].(string); ok {
		c.EcsVersion = v
	}
	if v, ok := conf[BuildNumberKey].(string); ok {
		c.BuildNumber = v
	}
	if v, ok := conf[ClusterIDKey].(string); ok {
		c.ClusterID = v
	}

	// parse specialized configs
	c.FileConfig, err = CreateFileConfig(c, conf)
	if err != nil {
		return
	}
	c.SyslogConfig, err = CreateSyslogConfig(c, conf)
	if err != nil {
		return
	}
	c.ESConfig, err = CreateElasticConfig(c, conf)
	if err != nil {
		return
	}
	c.KafkaConfig, err = CreateKafkaConfig(c, conf)

	return
}

// Transport type.
type Transport int

// Transport config options.
const (
	StdOutTransport Transport = iota
	FileTransport
	SyslogTransport
	ESTransport
	KafkaTransport
	NullTransport
)

func (s Transport) String() string {
	return [...]string{"terminal", "file", "syslog", "es", "kafka", "null"}[s]
}

func parseTransportConfig(s string) Transport {
	if FileTransport.String() == s {
		return FileTransport
	}
	if SyslogTransport.String() == s {
		return SyslogTransport
	}
	if ESTransport.String() == s {
		return ESTransport
	}
	if KafkaTransport.String() == s {
		return KafkaTransport
	}
	if NullTransport.String() == s {
		return NullTransport
	}
	return StdOutTransport
}

// Format type.
type Format int

// Format config options.
const (
	JSONFormat Format = iota // JSON schema
	ECSFormat                // Elastic Common Schema
	OtelFormat               // Open Telemetry schema
)

func (s Format) String() string {
	return [...]string{"json", "ecs", "otel"}[s]
}

func parseFormatConfig(s string) Format {
	switch s {
	case JSONFormat.String():
		return JSONFormat
	case ECSFormat.String():
		return ECSFormat
	case OtelFormat.String():
		return OtelFormat
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

// VaultEncoding type.
type VaultEncoding int

// VaultEncoding config options.
const (
	NoneVaultEncoding VaultEncoding = iota
	Base64VaultEncoding
)

func (s VaultEncoding) String() string {
	return [...]string{"none", "base64"}[s]
}

func parseVaultEncodingConfig(s string) VaultEncoding {
	if NoneVaultEncoding.String() == s {
		return NoneVaultEncoding
	}
	if Base64VaultEncoding.String() == s {
		return Base64VaultEncoding
	}
	return NoneVaultEncoding
}

// GetSecret obtains the secret for a key.
func (c Config) GetSecret(key string) (string, error) {
	return [...]func(string) (string, error){c.secrets.Get, c.secrets.GetDecoded}[c.VaultEncoding](key)
}
