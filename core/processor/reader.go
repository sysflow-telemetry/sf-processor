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

// Package processor implements a processor plugin.
package processor

import (
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
)

const (
	readerPluginName string = "sysflowreader"
)

// SysFlowReader defines the main reader class, which process SysFlow records and builds the cache.
// This plugin should typically be first in the pipeline.
type SysFlowReader struct {
	SysFlowProcessor
}

// var sPluginCache plugins.SFPluginCache
// var sPCOnce sync.Once

// NewSysFlowProcessor creates a new SysFlowProcessor instance.
func NewSysFlowReader() plugins.SFProcessor {
	logger.Trace.Println("Calling NewSysFlowReader")
	p := new(SysFlowReader)
	p.buildTables = true
	return p
}

// GetName returns the plugin name.
func (s *SysFlowReader) GetName() string {
	return readerPluginName
}

// // NewSysFlowChan creates a new processor channel instance.
// func NewSysFlowChan(size int) interface{} {
// 	return &plugins.SFChannel{In: make(chan *sfgo.SysFlow, size)}
// }

// Register registers plugin to plugin cache.
func (s *SysFlowReader) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(readerPluginName, NewSysFlowReader)
	pc.AddChannel(channelName, NewSysFlowChan)
	sPCOnce.Do(func() {
		sPluginCache = pc
	})
}

// // Init initializes the processor with a configuration map.
// func (s *SysFlowReader) Init(conf map[string]interface{}) (err error) {
// 	s.tables = cache.GetInstance()
// 	hdlCache := GetHandlerCacheInstance(sPluginCache)
// 	s.hdl, err = hdlCache.GetHandler(conf)
// 	if err != nil {
// 		return errors.Wrap(err, "couldn't obtain the processor handler from cache")
// 	}
// 	if err = s.hdl.Init(conf); err != nil {
// 		return errors.Wrap(err, "couldn't initialize processor handler")
// 	}
// 	return nil
// }

// // SetOutChan sets the output channel of the plugin.
// func (s *SysFlowReader) SetOutChan(ch []interface{}) {
// 	s.hdl.SetOutChan(ch)
// }

// // Process implements the main processor method of the plugin.
// func (s *SysFlowReader) Process(ch interface{}, wg *sync.WaitGroup) {
// 	s.SysFlowProcessor.Process(ch, wg)
// }
