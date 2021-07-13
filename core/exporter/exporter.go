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
	"errors"
	"sync"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/encoders"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/transports"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
)

const (
	pluginName string = "exporter"
)

var codecs = make(map[commons.Format]encoders.EncoderFactory)
var protocols = make(map[commons.Transport]transports.TransportProtocolFactory)

// Exporter defines a telemetry export plugin.
type Exporter struct {
	config    commons.Config
	encoder   encoders.Encoder
	transport transports.TransportProtocol
	recs      []*engine.Record
	counter   int
}

// NewExporter creates a new plugin instance.
func NewExporter() plugins.SFProcessor {
	return &Exporter{}
}

// GetName returns the plugin name.
func (s *Exporter) GetName() string {
	return pluginName
}

// Register registers plugin to plugin cache.
func (s *Exporter) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(pluginName, NewExporter)
}

// registerCodecs register encoders for exporting processor data.
func (s *Exporter) registerCodecs() {
	(&encoders.JSONEncoder{}).Register(codecs)
	(&encoders.ECSEncoder{}).Register(codecs)
	(&encoders.OccurrenceEncoder{}).Register(codecs)
}

// registerExportProtocols register transport protocols for exporting processor data.
func (s *Exporter) registerExportProtocols() {
	(&transports.SyslogProto{}).Register(protocols)
	(&transports.TerminalProto{}).Register(protocols)
	(&transports.TextFileProto{}).Register(protocols)
	(&transports.NullProto{}).Register(protocols)
	(&transports.FindingsAPIProto{}).Register(protocols)
	(&transports.ElasticProto{}).Register(protocols)
}

// Init initializes the plugin with a configuration map and cache.
func (s *Exporter) Init(conf map[string]interface{}) error {
	var err error

	// register encoders
	s.registerCodecs()

	// register export protocols
	s.registerExportProtocols()

	// create and read config object
	s.config, err = commons.CreateConfig(conf)
	if err != nil {
		return err
	}

	// initialize encoder
	if createCodec, ok := codecs[s.config.Format]; ok {
		s.encoder = createCodec(s.config)
	} else {
		return errors.New("Unable to find encoder for " + s.config.Format.String())
	}

	// initiliaze transport protocol
	if createTransport, ok := protocols[s.config.Transport]; ok {
		s.transport = createTransport(s.config)
		err = s.transport.Init()
		if err != nil {
			return err
		}
	} else {
		return errors.New("Unable to find transport protocol for " + s.config.Transport.String())
	}

	return err
}

// Test implements health checks for the plugin.
func (s *Exporter) Test() (bool, error) {
	if t, ok := s.transport.(transports.TestableTransportProtocol); ok {
		return t.Test()
	}
	return true, nil
}

// Process implements the main interface of the plugin.
func (s *Exporter) Process(ch interface{}, wg *sync.WaitGroup) {
	cha := ch.(*engine.RecordChannel)
	record := cha.In
	defer wg.Done()

	maxIdle := 1 * time.Second
	ticker := time.NewTicker(maxIdle)
	defer ticker.Stop()
	lastFlush := time.Now()

	logger.Trace.Printf("Starting exporter in mode %s with channel capacity %d", s.config.Transport.String(), cap(record))

RecLoop:
	for {
		select {
		case fc, ok := <-record:
			if ok {
				s.counter++
				s.recs = append(s.recs, fc)
				if s.counter >= s.config.EventBuffer {
					s.process()
					s.recs = s.recs[:0]
					s.counter = 0
					lastFlush = time.Now()
				}
			} else {
				if s.counter > 0 {
					s.process()
				}
				logger.Trace.Println("Channel closed. Shutting down.")
				break RecLoop
			}
		case <-ticker.C:
			// force flush records after 1sec idle
			if time.Since(lastFlush) > maxIdle && s.counter > 0 {
				s.process()
				s.recs = s.recs[:0]
				s.counter = 0
				lastFlush = time.Now()
			}
		}
	}
}

func (s *Exporter) process() error {
	data, err := s.encoder.Encode(s.recs)
	if err != nil {
		logger.Error.Println(err)
		return err
	}
	if len(data) > 0 {
		err = s.transport.Export(data)
		if err != nil {
			logger.Error.Println(err)
			return err
		}
	}
	return nil
}

// SetOutChan sets the output channel of the plugin.
func (s *Exporter) SetOutChan(ch []interface{}) {}

// Cleanup tears down plugin resources.
func (s *Exporter) Cleanup() {
	logger.Trace.Println("Exiting ", pluginName)
	s.encoder.Cleanup()
	s.transport.Cleanup()
}
