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
	"fmt"

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

func parseSymPath(idx sfgo.Source, attr sfgo.Attribute, r *Record) (string, string) {
	orig := r.GetStr(attr, idx)
	var src, dst uint64
	var targetPath string
	// Possible format: aabbccddeeff0011->aabbccddeeff0011 /path/to/target.file
	if _, err := fmt.Sscanf(orig, "%x->%x %s", &src, &dst, &targetPath); nil == err {
		return targetPath, fmt.Sprintf("%x->%x", src, dst)
	}
	// Possible format: ffff9ce02054c800-\u003effff9ce02054c000 /sock/sysflow.sock
	if _, err := fmt.Sscanf(orig, "%x-\\u%x %s", &src, &dst, &targetPath); nil == err {
		return targetPath, fmt.Sprintf("%x-\\u%x", src, dst)
	}
	return orig, ""
}
