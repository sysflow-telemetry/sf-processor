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
	"os"
	"time"
)

// Configuration keys
const (
	ConfigKeyHostname = "hostname"
	ConfigKeyPeriod   = "period"
)

// Config defines a configuration object for the exporter.
type Config struct {
	Hostname        string
	Period          time.Duration
	OutputInterface string
}

// CreateConfig creates a new config object from config dictionary.
func CreateConfig(raw map[string]string) (conf *Config, err error) {
	conf = &Config{
		Hostname: "",
		Period:   time.Minute,
		OutputInterface: "log",
	} // default values

	if v, ok := raw[ConfigKeyHostname]; ok {
		conf.Hostname = v
	}
	if conf.Hostname == "" {
		if conf.Hostname, err = os.Hostname(); nil != err {
			return nil, err
		}
	}

	if v, ok := raw[ConfigKeyPeriod]; ok {
		if conf.Period, err = time.ParseDuration(v); nil != err {
			return nil, err
		}
	}
	return conf, nil
}
