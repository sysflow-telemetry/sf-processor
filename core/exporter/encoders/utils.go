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
package encoders

import (
	"time"
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

// Max returns the larger of two integers, x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// converts a unix time value in ns to UTC time and returns an RFC3399 string
func toIsoTimeStr(ts int64) string {
	ts_sec := int64(ts / 1E+9)
	ts_ns := int64(ts % 1E+9)
	t := time.Unix(ts_sec, ts_ns).In(time.UTC)
	return t.Format(time.RFC3339Nano)
}
