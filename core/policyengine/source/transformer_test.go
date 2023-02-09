//
// Copyright (C) 2023 IBM Corporation.
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

package source_test

import (
	"encoding/base64"
	"testing"
)

func TestBase64Offset(t *testing.T) {
	msg := "http://"
	offset0 := base64.StdEncoding.EncodeToString([]byte(msg))
	t.Logf("%s, %s", msg, offset0)

	// for i := 0; i < 3; i++ {
	// 	d := append([]byte{" "}, )
	// 	offset := base64.StdEncoding.EncodeToString()
	// }

	msg1 := []byte(msg)[1 : len(msg)-3]
	offset1 := base64.StdEncoding.EncodeToString(append([]byte(" "), msg1...))
	t.Logf("%s, %s", msg1, offset1)

	msg2 := []byte(msg)[2:]
	offset2 := base64.StdEncoding.EncodeToString(append([]byte("  "), msg2...))
	t.Logf("%s, %s", msg2, offset2)

	// msg = "ttp://"
	// offset := base64.StdEncoding.EncodeToString([]byte(msg))
	// t.Logf("%s, %s", msg, offset)

	// msg = "tp://"
	// offset = base64.StdEncoding.EncodeToString([]byte(msg))
	// t.Logf("%s, %s", msg, offset)

	// start_offsets := []int{0, 2, 3}
	// end_offsets := []int{0, -3, -2}

}
