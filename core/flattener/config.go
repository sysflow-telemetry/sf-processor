//
// Copyright (C) 2022 IBM Corporation.
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

// Package flattener flattens input telemetry in a flattened representation.
package flattener

import (
	"strconv"
	"time"
)

// Configuration keys.
const (
	FilterOnOffKey  string = "filter.enabled"
	FilterMaxAgeKey string = "filter.maxage"
)

// Config defines a configuration object for the engine.
type Config struct {
	FilterOnOff  OnOff
	FilterMaxAge time.Duration
}

// CreateConfig creates a new config object from config dictionary.
func CreateConfig(conf map[string]interface{}) (Config, error) {
	var c Config = Config{FilterOnOff: Off, FilterMaxAge: 24 * time.Hour} // default values
	var err error
	if v, ok := conf[FilterOnOffKey].(string); ok {
		c.FilterOnOff = parseOnOffType(v)
	}
	if v, ok := conf[FilterMaxAgeKey].(string); ok {
		var duration int
		duration, err = strconv.Atoi(v)
		if err == nil {
			c.FilterMaxAge = time.Duration(duration) * time.Second
		}
	}
	return c, err
}

// OnOff defines an On-Off state type.
type OnOff int32

// OnOff types.
const (
	Off OnOff = iota
	On
)

func (s OnOff) String() string {
	return [...]string{"off", "on"}[s]
}

func (s OnOff) Enabled() bool {
	return s == On
}

func parseOnOffType(s string) OnOff {
	if Off.String() == s {
		return Off
	}
	if On.String() == s {
		return On
	}
	return Off
}
