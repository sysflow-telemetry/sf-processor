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
package statistic

import (
	"math"
	"sync"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.ibm.com/sysflow/goutils/logger"
	"github.ibm.com/sysflow/sf-processor/core/policyengine/engine"
)

const (
	pluginName = "statistic"
	channelName = "statchan"
)

type statistics struct {
	Total uint64
}

// StatisticExporter defines a driver for the statistic exporter plugin.
type StatisticExporter struct {
	config      Config
	outRecordCh chan<- *engine.Record
	stat        statistics
	period      time.Duration
}

// NewStatisticExporter constructs a new Statistic Exporter plugin.
func NewStatisticExporter() plugins.SFProcessor {
	return new(StatisticExporter)
}

// GetName returns the plugin name.
func (_ *StatisticExporter) GetName() string {
	return pluginName
}

// NewEventChan creates a new event record channel instance.
func NewStatisticChan(size int) interface{} {
	return &engine.RecordChannel{In: make(chan *engine.Record, size)}
}


// Register registers plugin to plugin cache.
func (_ *StatisticExporter) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(pluginName, NewStatisticExporter)
	pc.AddChannel(channelName, NewStatisticChan)
}

// Init initializes the plugin.
func (s *StatisticExporter) Init(conf map[string]string) error {
	if conf, err := CreateConfig(conf); err != nil {
		return err
	} else {
		s.config = *conf
		s.period = conf.Period
	}
	return nil
}

// Process implements the main loop of the plugin.
func (s *StatisticExporter) Process(inputCh interface{}, wg *sync.WaitGroup) {
	defer s.Cleanup()
	defer wg.Done()
	switch ch := inputCh.(type) {
	case *engine.RecordChannel:
		logger.Trace.Println("Statistic exporter receive from record channel")
		s.processRecordChan(ch.In)
	default:
		logger.Error.Fatalf("Invalid input channel type for statistic exporter: %T", inputCh)
	}
	logger.Trace.Println("Exiting statistic exporter")
}

// SetOutChan sets the output channel of the plugin.
func (s *StatisticExporter) SetOutChan(outputCh interface{}) {
	switch ch := outputCh.(type) {
	case *engine.RecordChannel:
		logger.Trace.Println("Statistic exporter output to record channel")
		s.outRecordCh = ch.In
	default:
		logger.Error.Fatalf("Invalid output channel type for statistic exporter: %T", outputCh)
	}
}

// Cleanup clean up the plugin resources.
func (s *StatisticExporter) Cleanup() {}

func (s *StatisticExporter) processRecordChan(inputCh <-chan *engine.Record) {
	ticker := time.NewTicker(s.period)
	defer ticker.Stop()
	for {
		select {
		case rec, ok := <-inputCh:
			if ok {
				if s.stat.Total < math.MaxUint64 {
					s.stat.Total++
				}
				s.outRecordCh <- rec
			} else {
				s.flushStatistics()
				return
			}
		case <-ticker.C:
			s.flushStatistics()
		}
	}
}

func (s *StatisticExporter) flushStatistics() {
	overflowSign := ""
	if s.stat.Total == math.MaxUint64 {
		overflowSign = ">="
	}
	logger.Info.Printf("total output in past %v: %s%d", s.period, overflowSign, s.stat.Total)
	s.stat = statistics{}
}
