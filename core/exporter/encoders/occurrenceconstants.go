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

// Severity type for enumeration.
type Severity int

// Severity enumeration.
const (
	SeverityLow Severity = iota
	SMedium
	SHigh
)

// String returns the string representation of a severity instance.
func (s Severity) String() string {
	return [...]string{"LOW", "MEDIUM", "HIGH"}[s]
}

// Certainty type for enumeration.
type Certainty int

// Certainty enumeration.
const (
	CertaintyLow Certainty = iota
	CertaintyMedium
	CertaintyHigh
)

// String returns the string representation of a severity instance.
func (s Certainty) String() string {
	return [...]string{"LOW", "MEDIUM", "HIGH"}[s]
}
