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

// Package pipeline implements a pluggable data processing pipeline infrastructure.
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
	running     bool
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
func (pl *Pipeline) Load(driverName string) error {
	conf, err := pl.pluginCache.GetConfig()
	if err != nil {
		logger.Error.Println("Unable to load pipeline config: ", err)
		return err
	}
	setManifestInfo(conf)
	if err := pl.pluginCache.LoadDrivers(pl.driverDir); err != nil {
		logger.Error.Println("Unable to load dynamic driver: ", err)
		return err
	}
	if pl.driver, err = pl.pluginCache.GetDriver(driverName); err != nil {
		logger.Error.Println("Unable to load driver: ", err)
		return err
	}
	var in interface{}
	var out interface{}
	for _, p := range conf.Pipeline {
		var prc plugins.SFProcessor
		if val, ok := p[ProcConfig].(string); ok {
			prc, err = pl.pluginCache.GetProcessor(pl.pluginDir, val)
			if err != nil {
				logger.Error.Println(err)
				return err
			}
			tp := fmt.Sprintf("%T", prc)
			logger.Trace.Println(tp)
			err = prc.Init(p)
			if err != nil {
				logger.Error.Println(err)
				return err
			}
		} else {
			logger.Error.Println("processor tag must exist in plugin config")
			return err
		}
		if v, o := p[InChanConfig].(string); o {
			in, err = pl.pluginCache.GetChan(v, ChanSize)
			if err != nil {
				logger.Error.Println(err)
				return err
			}
			pl.channels = append(pl.channels, in)
			chp := fmt.Sprintf("%T", in)
			logger.Trace.Println(chp)
		} else {
			logger.Error.Println("in tag must exist in plugin config")
			return errors.New("in tag must exist in plugin config")
		}
		if v, o := p[OutChanConfig]; o {
			var channels []interface{}
			switch t := v.(type) {
			case []interface{}:
				for _, channel := range t {
					out, err = pl.pluginCache.GetChan(channel.(string), ChanSize)
					if err != nil {
						logger.Error.Println(err)
						return err
					}
					channels = append(channels, out)
					chp := fmt.Sprintf("%T", out)
					pl.channels = append(pl.channels, out)
					logger.Trace.Println(chp)
				}
			case string:
				out, err = pl.pluginCache.GetChan(t, ChanSize)
				if err != nil {
					logger.Error.Println(err)
					return err
				}
				channels = append(channels, out)
				chp := fmt.Sprintf("%T", out)
				pl.channels = append(pl.channels, out)
				logger.Trace.Println(chp)
			}
			prc.SetOutChan(channels)
		}
		pl.processors = append(pl.processors, prc)
		pl.wg.Add(1)
		go pl.process(prc, in)
	}
	return nil
}

// Init initializes the pipeline
func (pl *Pipeline) Init(path string) error {
	logger.Info.Println("Starting the processing pipeline")
	// initialize driver
	err := pl.driver.Init(pl)
	if err != nil {
		logger.Error.Println("Driver initialization error: " + err.Error())
		return err
	}
	// start processing
	pl.running = true
	err = pl.driver.Run(path, &(pl.running))
	if err != nil {
		pl.running = false
		logger.Error.Println("Cannot start the driver: " + err.Error())
		return err
	}
	return nil
}

// Shutdown stops the pipeline
func (pl *Pipeline) Shutdown() error {
	logger.Info.Println("Stopping the processing pipeline")
	pl.running = false
	pl.driver.Cleanup()
	return nil
}

// GetRootChannel returns the first channel in the pipeline
func (pl *Pipeline) GetRootChannel() interface{} {
	if len(pl.channels) > 0 {
		return pl.channels[0]
	}
	return nil
}

// Print outputs summary information about the loaded pipeline
func (pl *Pipeline) Print() {
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
