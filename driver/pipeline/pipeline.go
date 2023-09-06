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
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
)

// Pipeline represents a loaded plugin pipeline
type Pipeline struct {
	wg          *sync.WaitGroup
	drivers     []plugins.SFDriver
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
	if len(driverName) > 0 {
		var driver plugins.SFDriver
		if driver, err = pl.pluginCache.GetDriver(driverName); err != nil {
			logger.Error.Println("Unable to load driver: ", err)
			return err
		}
		pl.drivers = append(pl.drivers, driver)
	} else {
		for _, d := range conf.Drivers {
			if val, ok := d[DrivConfig].(string); ok {
				var driver plugins.SFDriver
				if driver, err = pl.pluginCache.GetDriver(val); err != nil {
					logger.Error.Println("Unable to load driver: ", val, err)
					return err
				}
				logger.Trace.Println("Loading driver: " + driver.GetName())
				pl.drivers = append(pl.drivers, driver)
			}
		}
	}
	if len(pl.drivers) == 0 {
		return errors.New("No drivers configured on command line or in pipeline config")
	}
	for _, p := range conf.Pipeline {
		var out interface{}
		var inChannels []interface{}
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
		if v, o := p[InChanConfig]; o {
			switch t := v.(type) {
			case []interface{}:
				for _, channel := range t {
					in, err := pl.pluginCache.GetChan(channel.(string), ChanSize)
					if err != nil {
						logger.Error.Println(err)
						return err
					}
					inChannels = append(inChannels, in)
					chp := fmt.Sprintf("%T", in)
					pl.channels = append(pl.channels, in)
					logger.Trace.Println(chp)

				}
			case string:
				in, err := pl.pluginCache.GetChan(t, ChanSize)
				if err != nil {
					logger.Error.Println(err)
					return err
				}
				pl.channels = append(pl.channels, in)
				inChannels = append(inChannels, in)
				chp := fmt.Sprintf("%T", in)
				logger.Trace.Println(chp)
			}
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
		go pl.process(prc, inChannels)
	}
	pl.test()
	return nil
}

// Init initializes the pipeline
func (pl *Pipeline) Init(path string) error {
	logger.Info.Println("Starting the processing pipeline")
	numDrivers := len(pl.drivers)
	// initialize driver
	for i, d := range pl.drivers {
		conf := pl.GetDriverConfig(d.GetName())
		logger.Trace.Println("Initializing Driver with config", d.GetName())
		err := d.Init(pl, conf)
		if err != nil {
			logger.Error.Println("Driver initialization error: " + err.Error())
			return err
		}
		// start processing
		pl.running = true
		logger.Trace.Println("Calling Run on driver", d.GetName())
		if i == (numDrivers - 1) {
			err = d.Run(path, &(pl.running))
			if err != nil {
				pl.running = false
				logger.Error.Println("Cannot start the driver: " + err.Error())
				return err
			}
		} else {
			go d.Run(path, &(pl.running))
		}
	}
	return nil
}

// Shutdown stops the pipeline
func (pl *Pipeline) Shutdown() error {
	logger.Info.Println("Stopping the processing pipeline")
	pl.running = false
	for _, d := range pl.drivers {
		d.Cleanup()
	}
	return nil
}

// GetRootChannel returns the first channel in the pipeline
func (pl *Pipeline) GetRootChannel() interface{} {
	if len(pl.channels) > 0 {
		return pl.channels[0]
	}
	return nil
}

// GetChannel returns a channel given a "<id> <type>" string
func (pl *Pipeline) GetChannel(channel string) (interface{}, error) {
	return pl.pluginCache.GetChan(channel, ChanSize)
}

// GetDriverConfig returns a driver configuration from the pipeline.json
// given a driver name
func (pl *Pipeline) GetDriverConfig(driverName string) map[string]interface{} {
	conf, err := pl.pluginCache.GetConfig()
	if err != nil {
		logger.Error.Println("Unable to load pipeline config: ", err)
		return nil
	}
	for _, d := range conf.Drivers {
		if val, ok := d[DrivConfig].(string); ok {
			if val == driverName {
				return d
			}
		}
	}
	return nil
}

// Print outputs summary information about the loaded pipeline
func (pl *Pipeline) Print() {
	logger.Trace.Printf("Loaded %d stages", len(pl.processors))
	logger.Trace.Printf("Loaded %d handlers", len(pl.handlers))
	logger.Trace.Printf("Loaded %d channels", len(pl.channels))
}

// Wait calls on pipeline's waitgroup
func (pl *Pipeline) Wait() {
	pl.wg.Wait()
}

// Proxy function for handling transparent cleanup of resources
func (pl *Pipeline) process(prc plugins.SFProcessor, in []interface{}) {
	prc.Process(in, pl.wg)
	prc.Cleanup()
}

// Function for handling testable plugin checks.
func (pl *Pipeline) test() {
	ctx, cancel := context.WithTimeout(context.Background(), HealthChecksTimeout)
	defer cancel()

	c := make(chan error, 1)
	go func() {
		for _, prc := range pl.processors {
			if tprc, ok := prc.(plugins.SFTestableProcessor); ok {
				if _, err := tprc.Test(); err != nil {
					logger.Error.Printf("Health checks for plugin %s failed: %v", prc.GetName(), err)
					c <- err
					return
				}
			}
		}
		c <- nil
	}()

	select {
	case err := <-c:
		if err != nil {
			logger.Health.Println("Health checks: failed")
		} else {
			logger.Health.Println("Health checks: passed")
		}
		return
	case <-ctx.Done():
		logger.Error.Println("Health checks timed out: ", ctx.Err())
		logger.Health.Println("Health checks: failed")
		return
	}
}
