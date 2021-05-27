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
package transports

import (
	"encoding/json"
	"fmt"

	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/utils"
)

// TerminalProto implements the TransportProtocol interface of a terminal output.
type TerminalProto struct{}

//  NewTerminalProto creates a new terminal protcol object.
func NewTerminalProto(conf commons.Config) TransportProtocol {
	return &TerminalProto{}
}

//Init initializes the terminal output object.
func (s *TerminalProto) Init() error {
	return nil
}

// Export exports the contents of buffer for the terminal.
func (s *TerminalProto) Export(data []commons.EncodedData) error {
	for _, d := range data {
		if buf, ok := d.([]byte); ok {
			fmt.Println(utils.UnsafeBytesToString(buf))
		} else if buf, err := json.Marshal(d); err == nil {
			fmt.Println(utils.UnsafeBytesToString(buf))
		} else {
			fmt.Printf("%v\n", data)
		}
	}
	return nil
}

// Register registers the terminal proto object with the exporter.
func (s *TerminalProto) Register(eps map[commons.Transport]TransportProtocolFactory) {
	eps[commons.StdOutTransport] = NewTerminalProto
}

// Cleanup cleans up the terminal output object.
func (s *TerminalProto) Cleanup() {}
