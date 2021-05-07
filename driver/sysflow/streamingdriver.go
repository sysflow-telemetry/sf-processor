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
package sysflow

import (
	"bytes"
	"net"
	"os"

	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

const (
	streamDriverName = "socket"
)

const (
	// BuffSize represents the buffer size of the stream
	BuffSize = 16384
	// OOBuffSize represents the OO buffer size of the stream
	OOBuffSize = 1024
)

// StreamingDriver represents a streaming sysflow datasource
type StreamingDriver struct {
	pipeline plugins.SFPipeline
	conn     *net.UnixConn
}

// NewStreamingDriver creates a new streaming driver object
func NewStreamingDriver() plugins.SFDriver {
	return &StreamingDriver{}
}

// GetName returns the driver name.
func (s *StreamingDriver) GetName() string {
	return streamDriverName
}

// Register registers driver to plugin cache
func (s *StreamingDriver) Register(pc plugins.SFPluginCache) {
	pc.AddDriver(streamDriverName, NewStreamingDriver)
}

// Init initializes the driver
func (s *StreamingDriver) Init(pipeline plugins.SFPipeline) error {
	s.pipeline = pipeline
	return nil
}

// Run runs the driver
func (s *StreamingDriver) Run(path string, running *bool) error {
	channel := s.pipeline.GetRootChannel()
	sfChannel := channel.(*plugins.SFChannel)

	records := sfChannel.In
	if err := os.RemoveAll(path); err != nil {
		logger.Error.Println("Remove error: ", err)
		return err
	}

	l, err := net.ListenUnix("unixpacket", &net.UnixAddr{Name: path, Net: "unixpacket"})
	if err != nil {
		logger.Error.Println("Listen error: ", err)
		return err
	}
	defer l.Close()

	sFlow := sfgo.NewSysFlow()
	deser, err := compiler.CompileSchemaBytes([]byte(sFlow.Schema()), []byte(sFlow.Schema()))
	if err != nil {
		logger.Error.Println("Compilation error: ", err)
		return err
	}

	for *running {
		buf := make([]byte, BuffSize)
		oobuf := make([]byte, OOBuffSize)
		reader := bytes.NewReader(buf)
		s.conn, err = l.AcceptUnix()
		if err != nil {
			logger.Error.Println("Accept error: ", err)
			break
		}
		for *running {
			sFlow = sfgo.NewSysFlow()
			_, _, flags, _, err := s.conn.ReadMsgUnix(buf[:], oobuf[:])
			if err != nil {
				logger.Error.Println("Read error: ", err)
				break
			}
			if flags == 0 {
				reader.Reset(buf)
				err = vm.Eval(reader, deser, sFlow)
				if err != nil {
					logger.Error.Println("Deserialization error: ", err)
				}
				records <- sFlow
			} else {
				logger.Error.Println("Flag error ReadMsgUnix: ", flags)
			}
		}
		s.conn.Close()
	}
	logger.Trace.Println("Closing main channel")
	close(records)
	s.pipeline.Wait()
	return nil
}

// Cleanup tears down the driver resources.
func (s *StreamingDriver) Cleanup() {
	logger.Trace.Println("Exiting ", streamDriverName)
	if s.conn != nil {
		s.conn.Close()
	}
}
