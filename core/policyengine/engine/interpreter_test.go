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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
)

var pi *PolicyInterpreter

func SetupInterpreter(m *testing.M) {
	pi = NewPolicyInterpreter(Config{})
	os.Exit(m.Run())
}

func TestCompile(t *testing.T) {
	logger.Trace.Println("Running test compile")
	paths, err := ioutils.ListFilePaths("../../../resources/policies/tests", ".yaml")
	assert.NoError(t, err)
	assert.NoError(t, pi.Compile(paths...))
}

func TestCompileDist(t *testing.T) {
	logger.Trace.Println("Running test compile")
	paths, err := ioutils.ListFilePaths("../../../resources/policies/distribution/filter.yaml", ".yaml")
	assert.NoError(t, err)
	assert.NoError(t, pi.Compile(paths...))
}
