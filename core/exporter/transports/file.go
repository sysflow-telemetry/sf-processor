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
	"os"

	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
)

// TextFileProto implements the TransportProtocol interface for a text file.
type TextFileProto struct {
	config  commons.Config
	fhandle *os.File
}

//  NewTextFileProto creates a new text file protcol object.
func NewTextFileProto(conf commons.Config) TransportProtocol {
	return &TextFileProto{config: conf}
}

// Init initializes the text file.
func (s *TextFileProto) Init() error {
	os.Remove(s.config.Path)
	f, err := os.OpenFile(s.config.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	s.fhandle = f
	return err
}

// Export writes the buffer to the open file.
func (s *TextFileProto) Export(data []commons.EncodedData) (err error) {
	for _, d := range data {
		if buf, ok := d.([]byte); ok {
			if _, err = s.fhandle.Write(buf); err != nil {
				return err
			}
			s.fhandle.WriteString("\n")
		} else if buf, err := json.Marshal(d); err == nil {
			if _, err = s.fhandle.Write(buf); err != nil {
				return err
			}
			s.fhandle.WriteString("\n")
		} else {
			if _, err = s.fhandle.WriteString(fmt.Sprintf("%v\n", d)); err != nil {
				return err
			}
		}
	}
	return
}

// Register registers the text file proto object with the exporter.
func (s *TextFileProto) Register(eps map[commons.Transport]TransportProtocolFactory) {
	eps[commons.FileTransport] = NewTextFileProto
}

// Cleanup closes the text file.
func (s *TextFileProto) Cleanup() {
	if s.fhandle != nil {
		s.fhandle.Close()
	}
}
