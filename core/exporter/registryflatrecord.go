//go:build flatrecord
// +build flatrecord

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

// Package exporter implements a module plugin for encoding and exporting telemetry records and events.
package exporter

import (
	"github.com/sysflow-telemetry/sf-processor/core/exporter/encoders"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/transports"
)

// registerCodecs register encoders for exporting processor data.
func (s *Exporter) registerCodecs() {
	(&encoders.JSONEncoder{}).Register(codecs)
	(&encoders.ECSEncoder{}).Register(codecs)
}

// registerExportProtocols register transport protocols for exporting processor data.
func (s *Exporter) registerExportProtocols() {
	(&transports.SyslogProto{}).Register(protocols)
	(&transports.TerminalProto{}).Register(protocols)
	(&transports.TextFileProto{}).Register(protocols)
	(&transports.NullProto{}).Register(protocols)
	(&transports.ElasticProto{}).Register(protocols)
}
