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

// Package encoders implements codecs for exporting records and events in different data formats.
package encoders

// Exporter constants
const (
	BUFFER_SIZE = 10240
)

// JSON schema constants
const (
	VERSION_STR       = "{\"" + VERSION_ATTR + "\":"
	GROUP_ID          = "{\"" + GROUP_ID_ATTR + "\":\""
	COMMA             = ','
	DOUBLE_QUOTE      = '"'
	QUOTE_COLON       = "\":"
	QUOTE_COLON_CURLY = "\":{"
	END_CURLY_COMMA   = "},"
	END_CURLY         = '}'
	END_SQUARE        = ']'
	BEGIN_SQUARE      = '['
	SPACE             = ' '
	POLICIES          = ",\"" + POLICIES_ATTR + "\":["
	ID_TAG            = "{\"" + ID_TAG_ATTR + "\":"
	DESC              = ",\"" + DESC_ATTR + "\":"
	PRIORITY          = ",\"" + PRIORITY_ATTR + "\":"
	TAGS              = ",\"" + TAGS_ATTR + "\":["
	PERIOD            = '.'
	EMPTY_STRING      = "\"\""
)

const chars = "0123456789abcdef"
