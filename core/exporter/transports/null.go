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

import "github.com/sysflow-telemetry/sf-processor/core/exporter/commons"

// NullProto implements the TransportProtocol interface with not output
// for performance testing.
type NullProto struct {
}

// NewNullProto creates a new null protocol object.
func NewNullProto(conf commons.Config) TransportProtocol {
	return &NullProto{}
}

// Init intializes a new null protocol object.
func (s *NullProto) Init() error {
	return nil
}

// Export does nothing.
func (s *NullProto) Export(data commons.EncodedData) error {
	return nil
}

// Register registers the null protocol object with the exporter.
func (s *NullProto) Register(eps map[commons.Transport]TransportProtocolFactory) {
	eps[commons.NullTransport] = NewNullProto
}

//Cleanup cleans up the null protocol object.
func (s *NullProto) Cleanup() {}
