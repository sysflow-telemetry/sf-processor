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
package engine

import (
	"bytes"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

func trimBoundingQuotes(s string) string {
	if len(s) > 0 && (s[0] == '"' || s[0] == '\'') {
		s = s[1:]
	}
	if len(s) > 0 && (s[len(s)-1] == '"' || s[len(s)-1] == '\'') {
		s = s[:len(s)-1]
	}
	return s
}

//  Adds quotes to a string if necessary
func TryAddQuotes(v string, buf *bytes.Buffer) {
	l := len(v)
	if l > 0 && (v[0] == '"' || v[0] == '\'') {
		buf.WriteString(v)
	} else {
		buf.WriteByte('"')
		buf.WriteString(v)
		buf.WriteByte('"')
	}
}

// GetRecType returns the record type of the record
func GetRecType(r *Record, src sfgo.Source) string {
	switch r.GetInt(sfgo.SF_REC_TYPE, src) {
	case sfgo.PROC:
		return TyP
	case sfgo.FILE:
		return TyF
	case sfgo.CONT:
		return TyC
	case sfgo.PROC_EVT:
		return TyPE
	case sfgo.FILE_EVT:
		return TyFE
	case sfgo.FILE_FLOW:
		return TyFF
	case sfgo.NET_FLOW:
		return TyNF
	case sfgo.HEADER:
		return TyH
	default:
		return TyUnknow
	}

}
