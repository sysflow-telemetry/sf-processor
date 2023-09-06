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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64(t *testing.T) {
	transformer := NewTransformer()
	msg := "http://"
	v, err := transformer.Transform(msg, Base64Flag)
	assert.NoError(t, err)
	assert.Equal(t, []string{"aHR0cDovLw=="}, v)
}

func TestWinDash(t *testing.T) {
	transformer := NewTransformer()
	msg := "my-windows-variant"
	v, err := transformer.Transform(msg, WinDashFlag)
	assert.NoError(t, err)
	assert.Equal(t, []string{"my/windows/variant"}, v)
}

func TestWinDashBase64(t *testing.T) {
	transformer := NewTransformer()
	msg := "my-windows-variant"
	v, err := transformer.Transform(msg, WinDashFlag.Set(Base64Flag))
	assert.NoError(t, err)
	assert.Equal(t, []string{"bXkvd2luZG93cy92YXJpYW50"}, v)
}
