//
// Copyright (C) 2021 IBM Corporation.
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

// Package utils implements common helpers for the exporter.
package utils

var exists = struct{}{}

// Set defines a set data structure.
type Set struct {
	m map[string]struct{}
}

// NewSet creates a new set.
func NewSet() *Set {
	s := &Set{}
	s.m = make(map[string]struct{})
	return s
}

// Add adds an element to the set.
func (s *Set) Add(value string) {
	s.m[value] = exists
}

// Remove remoces an element from the set.
func (s *Set) Remove(value string) {
	delete(s.m, value)
}

// Contains checks if value is in the set.
func (s *Set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

// Len returns the number of elements in the set.
func (s *Set) Len() int {
	return len(s.m)
}

// IsSubset checks if set s is a subset of l.
func (s *Set) IsSubset(l *Set) bool {
	for k := range s.m {
		if !l.Contains(k) {
			return false
		}
	}
	return true
}
