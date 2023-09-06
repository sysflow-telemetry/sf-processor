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

// Package sigma implements a frontend for Sigma rules engine.
package sigma

import (
	"encoding/base64"
	"strings"
)

// TransformerFlags defines a bitmap for transformer options.
type TransformerFlags uint16

// NoFlags represents the zero value for transformer flags.
const NoFlags TransformerFlags = 0

// Transformer flags.
const (
	Base64Flag       TransformerFlags = 1 << iota // Base64 flag
	Base64OffsetFlag                              // Base64 offset flag
	WinDashFlag                                   // WinDash flag
	CIDRFlag                                      // CIDR flag
)

// Set sets the bitmap flag.
func (s TransformerFlags) Set(flag TransformerFlags) TransformerFlags { return s | flag }

// Clear unsets the bitmap flag.
func (s TransformerFlags) Clear(flag TransformerFlags) TransformerFlags { return s &^ flag }

// Has checks if flag is set in the bitmap.
func (s TransformerFlags) Has(flag TransformerFlags) bool { return s&flag != NoFlags }

// Transformer defines a set of functions that transform source attribute values.
type Transformer struct{}

func NewTransformer() *Transformer {
	return &Transformer{}
}

func (s *Transformer) Transform(src string, flags TransformerFlags) (dst []string, err error) {
	if flags == NoFlags {
		return []string{src}, nil
	}
	if flags.Has(WinDashFlag) {
		return s.Transform(s.windash(src), flags.Clear(WinDashFlag))
	}
	if flags.Has(CIDRFlag) {
		for _, cidr := range s.cidr(src) {
			var r []string
			r, err = s.Transform(cidr, flags.Clear(CIDRFlag))
			if err != nil {
				return
			}
			dst = append(dst, r...)
		}
		return
	}
	if flags.Has(Base64Flag) {
		dst = append(dst, s.base64(src, 0))
		return
	}
	if flags.Has(Base64OffsetFlag) {
		dst = append(dst, s.base64(src, 0))
		dst = append(dst, s.base64(src, 1))
		dst = append(dst, s.base64(src, 2))
		return
	}
	return []string{src}, nil
}

func (s *Transformer) base64(src string, offset int) string {
	if offset > 0 {
		panic("Missing implementation for base64 offsets")
	}
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func (s *Transformer) windash(src string) string {
	return strings.ReplaceAll(src, "-", "/")
}

func (s *Transformer) cidr(src string) []string {
	panic("Missing implementation for CIDR transformer")
}
