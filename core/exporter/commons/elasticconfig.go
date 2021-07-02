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

// Package commons defines common facilities for exporters.
package commons

import (
	"strconv"
	"strings"
	"time"
)

// Configuration keys.
const (
	ESAddressesConfigKey string = "es.addresses"
	ESIndexConfigKey     string = "es.index"
	ESUsernameConfigKey  string = "es.username"
	ESPasswordConfigKey  string = "es.password"
	ESWorkersConfigKey   string = "es.bulk.numWorkers"
	ESFBufferConfigKey   string = "es.bulk.flushBuffer"
	ESFTimeoutConfigKey  string = "es.bulk.flushTimeout"
)

// ESConfig holds Elastic specific configuration.
type ESConfig struct {
	ESAddresses    []string
	ESIndex        string
	ESUsername     string
	ESPassword     string
	ESNumWorkers   int
	ESFlushBuffer  int
	ESFlushTimeout time.Duration
}

// CreateElasticConfig creates a new config object from config dictionary.
func CreateElasticConfig(bc Config, conf map[string]interface{}) (c ESConfig, err error) {
	// default values
	c = ESConfig{
		ESNumWorkers:   0,
		ESFlushBuffer:  5e+6,
		ESFlushTimeout: 30 * time.Second}

	// parse config map
	if v, ok := conf[ESAddressesConfigKey].(string); ok {
		c.ESAddresses = strings.Split(v, ",")
	}
	if v, ok := conf[ESIndexConfigKey].(string); ok {
		c.ESIndex = v
	}
	if v, ok := conf[ESUsernameConfigKey].(string); ok {
		c.ESUsername = v
	} else if bc.VaultEnabled && bc.Transport == ESTransport {
		s, err := bc.GetSecret(ESUsernameConfigKey)
		if err != nil {
			return c, err
		}
		c.ESUsername = string(s)
	}
	if v, ok := conf[ESPasswordConfigKey].(string); ok {
		c.ESPassword = v
	} else if bc.VaultEnabled && bc.Transport == ESTransport {
		s, err := bc.GetSecret(ESPasswordConfigKey)
		if err != nil {
			return c, err
		}
		c.ESPassword = string(s)
	}
	if v, ok := conf[ESWorkersConfigKey].(string); ok {
		c.ESNumWorkers, err = strconv.Atoi(v)
		if err != nil {
			return c, err
		}
	}
	if v, ok := conf[ESFBufferConfigKey].(string); ok {
		c.ESFlushBuffer, err = strconv.Atoi(v)
		if err != nil {
			return c, err
		}
	}
	if v, ok := conf[ESFTimeoutConfigKey].(string); ok {
		c.ESFlushTimeout, err = time.ParseDuration(v)
		if err != nil {
			return c, err
		}
	}
	return
}
