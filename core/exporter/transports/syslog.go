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
	"crypto/tls"
	"errors"
	"fmt"

	syslog "github.com/RackSec/srslog"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
)

// SyslogProto implements the TransportProtocol interface for syslog.
type SyslogProto struct {
	sysl   *syslog.Writer
	config commons.Config
}

//  NewSyslogProto creates a new syslog protocol object.
func NewSyslogProto(conf commons.Config) TransportProtocol {
	return &SyslogProto{config: conf}
}

// Init initializes the syslog daemon connection.
func (s *SyslogProto) Init() error {
	var err error
	raddr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	if s.config.Proto == commons.TCPTLSProto {
		// TODO: verify connection with given trust certifications
		nopTLSConfig := &tls.Config{InsecureSkipVerify: true}
		s.sysl, err = syslog.DialWithTLSConfig("tcp+tls", raddr, syslog.LOG_ALERT|syslog.LOG_DAEMON, s.config.Tag, nopTLSConfig)
	} else {
		s.sysl, err = syslog.Dial(s.config.Proto.String(), raddr, syslog.LOG_ALERT|syslog.LOG_DAEMON, s.config.Tag)
	}
	if err == nil {
		s.sysl.SetFormatter(syslog.RFC5424Formatter)
		if s.config.LogSource != sfgo.Zeros.String {
			s.sysl.SetHostname(s.config.LogSource)
		}
	}
	return err
}

// Export sends buffer to syslog daemon as an alert.
func (s *SyslogProto) Export(data commons.EncodedData) error {
	if buf, ok := data.([]byte); ok {
		return s.sysl.Alert(unsafeBytesToString(buf))
	}
	return errors.New("Expected byte array as export data")
}

// Register registers the syslog proto object with the exporter.
func (s *SyslogProto) Register(eps map[commons.Transport]TransportProtocolFactory) {
	eps[commons.SyslogTransport] = NewSyslogProto
}

// Cleanup closes the syslog connection.
func (s *SyslogProto) Cleanup() {
	if s.sysl != nil {
		s.sysl.Close()
	}
}
