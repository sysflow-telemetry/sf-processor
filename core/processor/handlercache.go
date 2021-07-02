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
	"errors"
	"fmt"
	"os"
	"plugin"
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-processor/core/flattener"
)

var sHCInstance *HandlerCache
var sHCOnce sync.Once

const (
	cHandlerSym     string = "Handler" // variable required to load handler dynamically from shared object.
	cHandlerName    string = "handler"
	cHandlerLibPath string = "handlerlibpath"
)

// HandlerCache defines a data strucure for managing handlers.
type HandlerCache struct {
	hdlFuncMap  map[string]interface{}
	pluginCache plugins.SFPluginCache
}

// newHandlerCache creates a new HandlerCache instance.
func newHandlerCache(pc plugins.SFPluginCache) *HandlerCache {
	hdl := &HandlerCache{
		hdlFuncMap:  make(map[string]interface{}),
		pluginCache: pc}
	hdl.init()
	return hdl
}

// GetHandlerCacheInstance returns HandlerCache singleton instance
func GetHandlerCacheInstance(pc plugins.SFPluginCache) *HandlerCache {
	sHCOnce.Do(func() {
		sHCInstance = newHandlerCache(pc)
	})
	return sHCInstance
}

// initializes plugin cache.
func (p *HandlerCache) init() {
	flat := &flattener.Flattener{}
	flat.RegisterHandler(p)
	flat.RegisterChannel(p.pluginCache)
}

// LoadHandler loads dynamic handlers to handler cache from dir path.
func (p *HandlerCache) loadHandler(name string, path string) (plugins.SFHandler, error) {
	var plug *plugin.Plugin
	dynPlugin := path + "/" + name + ".so"
	if _, err := os.Stat(dynPlugin); err == nil {
		if plug, err = plugin.Open(dynPlugin); err != nil {
			return nil, err
		}
		sym, err := plug.Lookup(cHandlerSym)
		if err != nil {
			return nil, err
		}
		if hdlr, ok := sym.(plugins.SFHandler); ok {
			hdlr.RegisterHandler(p)
			hdlr.RegisterChannel(p.pluginCache)
			return hdlr, nil
		}
	} else {
		return nil, errors.New("error trying load plugin at: " + dynPlugin)
	}
	return nil, fmt.Errorf("unable to dynamicly load Handler '%s' from library %s", name, dynPlugin)
}

// AddHandler adds a handler method to the handler cache
func (p *HandlerCache) AddHandler(name string, factory interface{}) {
	p.hdlFuncMap[name] = factory
}

// GetHandler retrieves a cached plugin handler by name.
func (p *HandlerCache) GetHandler(conf map[string]interface{}) (plugins.SFHandler, error) {
	if name, ok := conf[cHandlerName].(string); ok {
		if val, ok := p.hdlFuncMap[name]; ok {
			funct := val.(func() plugins.SFHandler)
			return funct(), nil
		}
		if path, o := conf[cHandlerLibPath].(string); o {
			return p.loadHandler(name, path)
		}
		return nil, fmt.Errorf("handler '%s' not found in built-in handlers, and no attribute 'handlerlib' for dynamic library defined", name)
	}
	return nil, fmt.Errorf("attribute 'handler' missing from sysflow processor's configuration")
}
