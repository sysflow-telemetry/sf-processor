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

package source

// TransformerFlags defines a bitmap for transformer options.
type TransformerFlags uint16

// NoFlags represents the zero value for transformer flags.
const NoFlags TransformerFlags = 0

// Transformer flags.
const (
	Base64Flag        TransformerFlags = 1 << iota // Base64 flag
	Base64Offset1Flag                              // Base64 1-byte offset flag
	Base64Offset2Flag                              // Base64 2-byte offset flag
	UTF16LEFlag                                    // UTF16-LE flag
	UTF16BEFlag                                    // UTF16-LE flag
	UTF16BOMFlag                                   // Byte order mask flag
	WinDashFlag                                    // WinDash flag
	CIDRFlag                                       // CIDR flag
)

// Set sets the bitmap flag.
func (s TransformerFlags) Set(flag TransformerFlags) TransformerFlags { return s | flag }

// Has checks if flag is set in the bitmap.
func (s TransformerFlags) Has(flag TransformerFlags) bool { return s&flag != NoFlags }

// Transformer defines a set of functions that transform source attribute values.
type Transformer struct{}

func NewTransformer() *Transformer {
	return &Transformer{}
}

func (s *Transformer) TransformToString(src []byte, flags TransformerFlags) string {
	dst := make([]byte, len(src))
	s.Transform(dst, src, flags)
	return string(dst)
}

func (s *Transformer) Transform(dst, src []byte, flags TransformerFlags) {
	if flags.Has(Base64Flag) {
		s.base64(dst, src, 0)
	} else if flags.Has(Base64Offset1Flag) {
		s.base64(dst, src, 1)
	} else if flags.Has(Base64Offset2Flag) {
		s.base64(dst, src, 2)
	}
	if flags.Has(UTF16BEFlag) {
		s.utf16be(dst, src, flags.Has(UTF16BOMFlag))
	} else if flags.Has(UTF16LEFlag) {
		s.utf16le(dst, src)
	}
}

func (s *Transformer) base64(dst, src []byte, offset int) {
	panic("Missing implementation")
}

func (s *Transformer) utf16le(dst, src []byte) {
	panic("Missing implementation")
}

func (s *Transformer) utf16be(dst, src []byte, useBOM bool) {
	panic("Missing implementation")
}
