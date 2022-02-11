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

//go:build linux

// Package sysflow implements pluggable drivers for SysFlow ingestion.
package sysflow

import (
	"fmt"
	"syscall"
)

// readMsgUnix reads a message from a unix socket (compatible with the Darwin architecture).
func (s *StreamingDriver) readMsgUnix(buf []byte, oobuf []byte) error {
	_, _, flags, _, err := s.conn.ReadMsgUnix(buf[:], oobuf[:])
	if flags != syscall.MSG_CMSG_CLOEXEC {
		return fmt.Errorf("ReadMsgUnix flags = %v, want %v (MSG_CMSG_CLOEXEC)", flags, syscall.MSG_CMSG_CLOEXEC)
	}
	return err
}
