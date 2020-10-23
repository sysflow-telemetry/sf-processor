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
	"errors"
	"sync"

	syslog "github.com/RackSec/srslog"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.ibm.com/sysflow/sf-processor/core/policyengine/engine"
)

const (
	pluginName string = "exporter"
)

// Exporter defines a syslogger plugin.
type Exporter struct {
	recs             []*engine.Record
	counter          int
	sysl             *syslog.Writer
	config           Config
	exporter         *JSONExporter
	exportProto      ExportProtocol
	exportProtoCache map[string]interface{}
}

// NewExporter creates a new plugin instance.
func NewExporter() plugins.SFProcessor {
	e := &Exporter{exportProtoCache: make(map[string]interface{})}
	return e
}

// GetName returns the plugin name.
func (s *Exporter) GetName() string {
	return pluginName
}

// Register registers plugin to plugin cache.
func (s *Exporter) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(pluginName, NewExporter)
}

// AddExportProtocol registers an export protocol object with the Exporter
func (s *Exporter) AddExportProtocol(protoName string, ep interface{}) {
	s.exportProtoCache[protoName] = ep
}

func (s *Exporter) initProtos() {
	(&SyslogProto{}).Register(s)
	(&TerminalProto{}).Register(s)
	(&TextFileProto{}).Register(s)
	(&NullProto{}).Register(s)

}

// Init initializes the plugin with a configuration map and cache.
func (s *Exporter) Init(conf map[string]string) error {
	var err error
	s.config = CreateConfig(conf)
	s.initProtos()
	if val, ok := s.exportProtoCache[s.config.Export.String()]; ok {
		funct := val.(func() ExportProtocol)
		s.exportProto = funct()
		err = s.exportProto.Init(conf)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Unable to find export protocol: " + s.config.Export.String())
	}
	s.exporter = NewJSONExporter(s.exportProto, s.config)

	return err
}

// Process implements the main interface of the plugin.
func (s *Exporter) Process(ch interface{}, wg *sync.WaitGroup) {
	cha := ch.(*engine.RecordChannel)
	record := cha.In
	defer wg.Done()
	logger.Trace.Printf("Starting Exporter in mode %s with channel capacity %d", s.config.Export.String(), cap(record))
	for {
		fc, ok := <-record
		if !ok {
			s.process()
			logger.Trace.Println("Channel closed. Shutting down.")
			break
		}
		s.counter++
		s.recs = append(s.recs, fc)
		if s.counter > s.config.EventBuffer {
			s.process()
			s.recs = s.recs[:0] // make([]*engine.Record, 0)
			s.counter = 0
		}
	}
}

func (s *Exporter) process() {
	if s.config.ExpType == AlertType {
		err := s.exporter.ExportOffenses(s.recs)
		if err != nil {
			logger.Error.Println("Error exporting events: " + err.Error())
		}
	} else {
		err := s.exporter.ExportTelemetryRecords(s.recs)
		if err != nil {
			logger.Error.Println("Error exporting events: " + err.Error())
		}
	}
}

// SetOutChan sets the output channel of the plugin.
func (s *Exporter) SetOutChan(ch interface{}) {}

// Cleanup tears down plugin resources.
func (s *Exporter) Cleanup() {
	logger.Trace.Println("Exiting ", pluginName)
	s.exportProto.Cleanup()
}

// This function is not run when module is used as a plugin.
func main() {}
