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
package pipeline

import (
	"errors"
	"fmt"
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
)

// Pipeline represents a loaded plugin pipeline
type Pipeline struct {
	wg          *sync.WaitGroup
	driver      plugins.SFDriver
	processors  []plugins.SFProcessor
	channels    []interface{}
	handlers    []plugins.SFHandler
	pluginCache *PluginCache
	config      string
	pluginDir   string
	driverDir   string
}

// New creates a new pipeline object
func New(driverDir string, pluginDir string, config string) *Pipeline {
	return &Pipeline{config: config,
		driverDir:   driverDir,
		pluginDir:   pluginDir,
		wg:          new(sync.WaitGroup),
		pluginCache: NewPluginCache(config),
	}

}

// GetNumChannels returns the number of channels in the pipeline
func (pl *Pipeline) GetNumChannels() int {
	return len(pl.channels)
}

// GetNumProcessors returns the number of processors in the pipeline
func (pl *Pipeline) GetNumProcessors() int {
	return len(pl.processors)
}

// GetNumHandlers returns the number of handlers in the pipeline
func (pl *Pipeline) GetNumHandlers() int {
	return len(pl.handlers)
}

// GetPluginCache returns the plugin cache for the pipeline
func (pl *Pipeline) GetPluginCache() plugins.SFPluginCache {
	return pl.pluginCache
}

// AddChannel adds a channel to the plugin cache
func (pl *Pipeline) AddChannel(channelName string, channel interface{}) {
	pl.pluginCache.AddChannel(channelName, channel)
}

// Load loads and enables the pipeline
func (pl *Pipeline) Load(driverName string) (plugins.SFDriver, error) {
	if err := pl.pluginCache.LoadDrivers(pl.driverDir); err != nil {
		logger.Error.Println("Unable to load dynamic driver: ", err)
		return nil, err
	}
	if err := pl.pluginCache.LoadPlugins(pl.pluginDir); err != nil {
		logger.Error.Println("Unable to load dynamic plugins: ", err)
		return nil, err
	}
	conf, err := pl.pluginCache.GetConfig()
	if err != nil {
		logger.Error.Println("Unable to load pipeline config: ", err)
		return nil, err
	}
	setManifestInfo(conf)
	var driver plugins.SFDriver
	if driver, err = pl.pluginCache.GetDriver(driverName); err != nil {
		logger.Error.Println("Unable to load driver: ", err)
		return nil, err
	}
	var in interface{}
	var out interface{}
	for _, p := range conf.Pipeline {
		hdler := false
		var hdl plugins.SFHandler
		if val, ok := p[HdlConfig]; ok {
			hdl, err = pl.pluginCache.GetHandler(val)
			if err != nil {
				logger.Error.Println(err)
				return nil, err
			}
			pl.handlers = append(pl.handlers, hdl)
			xType := fmt.Sprintf("%T", hdl)
			logger.Trace.Println(xType)
			hdler = true
		}
		var prc plugins.SFProcessor
		if val, ok := p[ProcConfig]; ok {
			prc, err = pl.pluginCache.GetProcessor(val, hdl, hdler)
			if err != nil {
				logger.Error.Println(err)
				return nil, err
			}
			tp := fmt.Sprintf("%T", prc)
			logger.Trace.Println(tp)
			err = prc.Init(p)
			if err != nil {
				logger.Error.Println(err)
				return nil, err
			}
		} else {
			logger.Error.Println("processor or handler tag must exist in plugin config")
			return nil, err
		}
		if v, o := p[InChanConfig]; o {
			in, err = pl.pluginCache.GetChan(v, ChanSize)
			pl.channels = append(pl.channels, in)
			chp := fmt.Sprintf("%T", in)
			logger.Trace.Println(chp)
		} else {
			logger.Error.Println("in tag must exist in plugin config")
			return nil, errors.New("in tag must exist in plugin config")
		}
		if v, o := p[OutChanConfig]; o {
			out, err = pl.pluginCache.GetChan(v, ChanSize)
			chp := fmt.Sprintf("%T", out)
			pl.channels = append(pl.channels, out)
			logger.Trace.Println(chp)
			prc.SetOutChan(out)
		}
		pl.processors = append(pl.processors, prc)
		pl.wg.Add(1)
		go pl.process(prc, in)
	}
	return driver, nil
}

// GetRootChannel returns the first channel in the pipeline
func (pl *Pipeline) GetRootChannel() interface{} {
	if len(pl.channels) > 0 {
		return pl.channels[0]
	}
	return nil
}

// PrintPipeline outputs summary information about the loaded pipeline
func (pl *Pipeline) PrintPipeline() {
	logger.Trace.Printf("Loaded %d stages\n", len(pl.processors))
	logger.Trace.Printf("Loaded %d channels\n", len(pl.channels))
	logger.Trace.Printf("Loaded %d handlers\n", len(pl.handlers))
}

// Wait calls on pipeline's waitgroup
func (pl *Pipeline) Wait() {
	pl.wg.Wait()
}

// Proxy function for handling transparent cleanup of resources
func (pl *Pipeline) process(prc plugins.SFProcessor, in interface{}) {
	prc.Process(in, pl.wg)
	prc.Cleanup()
}
