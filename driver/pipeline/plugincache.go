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
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"github.ibm.com/sysflow/sf-processor/driver/sysflow"

	"github.com/spf13/viper"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.ibm.com/sysflow/goutils/ioutils"
	"github.ibm.com/sysflow/goutils/logger"
	"github.ibm.com/sysflow/sf-processor/core/exporter"
	"github.ibm.com/sysflow/sf-processor/core/policyengine"
	"github.ibm.com/sysflow/sf-processor/core/processor"
)

// PluginCache defines a data strucure for managing plugins.
type PluginCache struct {
	chanMap     map[string]interface{}
	driverMap   map[string]interface{}
	procFuncMap map[string]interface{}
	hdlFuncMap  map[string]interface{}
	chanFuncMap map[string]interface{}
	config      *viper.Viper
	configFile  string
}

// NewPluginCache creates a new PluginCache instance.
func NewPluginCache(conf string) *PluginCache {
	plug := &PluginCache{config: viper.New(),
		chanMap:     make(map[string]interface{}),
		driverMap:   make(map[string]interface{}),
		procFuncMap: make(map[string]interface{}),
		hdlFuncMap:  make(map[string]interface{}),
		chanFuncMap: make(map[string]interface{}),
		configFile:  conf}
	plug.init()
	return plug
}

// initializes plugin cache.
func (p *PluginCache) init() {
	(&processor.SysFlowProcessor{}).Register(p)
	(&policyengine.PolicyEngine{}).Register(p)
	(&exporter.Exporter{}).Register(p)
	//(&sysflow.FileDriver{}).Register(p)
	(&sysflow.StreamingDriver{}).Register(p)
}

// LoadPlugins loads dynamic plugins to plugin cache from dir path.
func (p *PluginCache) LoadPlugins(dir string) error {
	var plug *plugin.Plugin
	if paths, err := ioutils.ListFilePaths(dir, ".so"); err == nil {
		for _, path := range paths {
			if plug, err = plugin.Open(path); err != nil {
				return err
			}
			sym, err := plug.Lookup(plugins.PlugSym)
			if err != nil {
				return err
			}
			if proc, ok := sym.(plugins.SFProcessor); ok {
				// p.pluginMap[proc.GetName()] = plug
				proc.Register(p)
			}
		}
	}
	return nil
}

// LoadDrivers dynamic drivers to plugin cache from dir path.
func (p *PluginCache) LoadDrivers(dir string) error {
	var plug *plugin.Plugin
	if paths, err := ioutils.ListFilePaths(dir, ".so"); err == nil {
		for _, path := range paths {
			if plug, err = plugin.Open(path); err != nil {
				return err
			}
			sym, err := plug.Lookup(plugins.DriverSym)
			if err != nil {
				return err
			}
			if driver, ok := sym.(plugins.SFDriver); ok {
				driver.Register(p)
			}
		}
	}
	return nil
}

// AddDriver adds a driver factory method to the plugin cache.
func (p *PluginCache) AddDriver(name string, factory interface{}) {
	p.driverMap[name] = factory
}

// AddProcessor adds a processor factory method to the plugin cache.
func (p *PluginCache) AddProcessor(name string, factory interface{}) {
	p.procFuncMap[name] = factory
}

// AddHandler adds a handler factory method to the plugin cache.
func (p *PluginCache) AddHandler(name string, factory interface{}) {
	p.hdlFuncMap[name] = factory
}

// AddChannel adds a channel factory method to the plugin cache.
func (p *PluginCache) AddChannel(name string, factory interface{}) {
	p.chanFuncMap[name] = factory
}

// GetConfig reads the PluginCache configuration.
func (p *PluginCache) GetConfig() (*Config, error) {
	s, err := os.Stat(p.configFile)
	if os.IsNotExist(err) {
		return nil, err
	}
	if s.IsDir() {
		return nil, errors.New("pipeline config file is not a file")
	}
	dir := filepath.Dir(p.configFile)
	p.config.SetConfigName(strings.TrimSuffix(filepath.Base(p.configFile), filepath.Ext(p.configFile)))
	p.config.SetConfigType("json")
	p.config.AddConfigPath(dir)

	conf := new(Config)
	err = p.config.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = p.config.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	p.updateConfigFromEnv(conf)
	return conf, nil
}

// updateConfigFromEnv updates config object with environment variables if set.
// It assumes the following convention:
// - Environment variables follow the naming schema <PROCESSOR NAME>_<CONFIG ATTRIBUTE NAME>
// - Processor name in pipeline.json is all lower case
func (p *PluginCache) updateConfigFromEnv(config *Config) {
	for _, c := range config.Pipeline {
		if proc, ok := c[ProcConfig]; ok {
			for k, v := range p.getEnv(proc) {
				c[k] = v
			}
		}
	}
}

// getEnv returns the environemnt config settings for processor proc.
func (p *PluginCache) getEnv(proc string) map[string]string {
	var conf = make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		key := strings.SplitN(strings.ToLower(pair[0]), "_", 2)
		if len(key) == 2 && key[0] == proc {
			conf[key[1]] = pair[1]
		}
	}
	return conf
}

// GetHandler retrieves a cached plugin handler by name.
func (p *PluginCache) GetHandler(name string) (plugins.SFHandler, error) {
	if val, ok := p.hdlFuncMap[name]; ok {
		funct := val.(func() plugins.SFHandler)
		return funct(), nil
	}
	return nil, fmt.Errorf("Handler '%s' not found in plugin cache", name)
}

// GetChan retrieves a cached plugin channel by name.
func (p *PluginCache) GetChan(ch string, size int) (interface{}, error) {
	fields := strings.Fields(ch)
	if len(fields) != 2 {
		return nil, errors.New("Channel must be of the form <identifier> <type>")
	}
	if val, ok := p.chanMap[fields[0]]; ok {
		logger.Trace.Println("Found existing channel ", fields[0])
		return val, nil
	}
	if val, ok := p.chanFuncMap[fields[1]]; ok {
		funct := val.(func(int) interface{})
		c := funct(size)
		p.chanMap[fields[0]] = c
		return c, nil
	}
	return nil, fmt.Errorf("Channel '%s' not found in plugin cache", fields[0])
}

// GetProcessor retrieves a cached plugin processor by name.
func (p *PluginCache) GetProcessor(name string, hdl plugins.SFHandler, hdlr bool) (plugins.SFProcessor, error) {
	if val, ok := p.procFuncMap[name]; ok {
		logger.Trace.Println("Found processor in function map: ", name)
		var prc plugins.SFProcessor
		if hdlr {
			funct := val.(func(plugins.SFHandler) plugins.SFProcessor)
			prc = funct(hdl)
		} else {
			funct := val.(func() plugins.SFProcessor)
			prc = funct()
		}
		return prc, nil
	}
	return nil, fmt.Errorf("Plugin '%s' not found in plugin cache", name)
}

// GetDriver retrieves a cached plugin driver by name.
func (p *PluginCache) GetDriver(name string) (plugins.SFDriver, error) {
	if val, ok := p.driverMap[name]; ok {
		logger.Trace.Println("Found driver in function map: ", name)
		funct := val.(func() plugins.SFDriver)
		drv := funct()
		return drv, nil
	}
	return nil, fmt.Errorf("Driver '%s' not found in plugin cache", name)
}
