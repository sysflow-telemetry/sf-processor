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

// Configuration keys.
const (
	PathConfigKey string = "file.path"
)

// FileConfig holds file output specific configuration.
type FileConfig struct {
	Path string
}

// CreateFileConfig creates a new config object from config dictionary.
func CreateFileConfig(bc Config, conf map[string]interface{}) (c FileConfig, err error) {
	// default values
	c = FileConfig{Path: "./export.out"}

	// parse config map
	if v, ok := conf[PathConfigKey].(string); ok {
		c.Path = v
	}
	return
}
