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
	"sync"

	"github.com/pkg/errors"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

const (
	pluginName  string = "sysflowprocessor"
	channelName string = "ctxchan"
)

// SysFlowProcessor defines the main processor class.
type SysFlowProcessor struct {
	hdl plugins.SFHandler
}

var sPluginCache plugins.SFPluginCache
var sPCOnce sync.Once

// NewSysFlowProcessor creates a new SysFlowProcessor instance.
func NewSysFlowProcessor() plugins.SFProcessor {
	logger.Trace.Println("Calling NewSysFlowProc")
	p := new(SysFlowProcessor)
	return p
}

// GetName returns the plugin name.
func (s *SysFlowProcessor) GetName() string {
	return pluginName
}

// NewCtxSysFlowChan creates a new processor channel instance.
func NewCtxSysFlowChan(size int) interface{} {
	return &plugins.CtxSFChannel{In: make(chan *plugins.CtxSysFlow, size)}
}

// Register registers plugin to plugin cache.
func (s *SysFlowProcessor) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(pluginName, NewSysFlowProcessor)
	pc.AddChannel(channelName, NewCtxSysFlowChan)
	sPCOnce.Do(func() {
		sPluginCache = pc
	})
}

// Init initializes the processor with a configuration map.
func (s *SysFlowProcessor) Init(conf map[string]interface{}) (err error) {
	hdlCache := GetHandlerCacheInstance(sPluginCache)
	s.hdl, err = hdlCache.GetHandler(conf)
	if err != nil {
		return errors.Wrap(err, "couldn't obtain the processor handler from cache")
	}
	if err = s.hdl.Init(conf); err != nil {
		return errors.Wrap(err, "couldn't initialize processor handler")
	}
	return nil
}

// SetOutChan sets the output channel of the plugin.
func (s *SysFlowProcessor) SetOutChan(ch []interface{}) {
	s.hdl.SetOutChan(ch)
}

// Process implements the main processor method of the plugin.
func (s *SysFlowProcessor) Process(ch interface{}, wg *sync.WaitGroup) {
	entEnabled := s.hdl.IsEntityEnabled()
	cha := ch.(*plugins.CtxSFChannel)
	record := cha.In
	defer wg.Done()
	logger.Trace.Println("Starting SysFlow Processor...")
	for {
		sf, ok := <-record
		if !ok {
			logger.Trace.Println("SysFlow Processor channel closed. Shutting down.")
			break
		}
		switch sf.Rec.UnionType {
		case sfgo.SF_HEADER:
			if entEnabled {
				s.hdl.HandleHeader(sf, sf.Header)
			}
		case sfgo.SF_CONT:
			if entEnabled {
				s.hdl.HandleContainer(sf, sf.Container)
			}
		case sfgo.SF_PROCESS:
			if entEnabled {
				s.hdl.HandleProcess(sf, sf.Process)
			}
		case sfgo.SF_FILE:
			if entEnabled {
				s.hdl.HandleFile(sf, sf.File)
			}
		case sfgo.SF_PROC_EVT:
			pe := sf.Rec.ProcessEvent
			s.hdl.HandleProcEvt(sf, pe)
		case sfgo.SF_NET_FLOW:
			nf := sf.Rec.NetworkFlow
			s.hdl.HandleNetFlow(sf, nf)
		case sfgo.SF_FILE_FLOW:
			ff := sf.Rec.FileFlow
			s.hdl.HandleFileFlow(sf, ff)
		case sfgo.SF_FILE_EVT:
			fe := sf.Rec.FileEvent
			s.hdl.HandleFileEvt(sf, fe)
		case sfgo.SF_PROC_FLOW:
		case sfgo.SF_NET_EVT:
		default:
			logger.Warn.Println("Error unsupported SysFlow Type: ", sf.Rec.UnionType)
		}
	}
}

// Cleanup tears down the plugin resources.
func (s *SysFlowProcessor) Cleanup() {
	logger.Trace.Println("Exiting ", pluginName)
	s.hdl.Cleanup()
}
