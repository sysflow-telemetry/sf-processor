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
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package engine

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy/falco"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy/sigma"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/flatrecord"
)

var pi *PolicyInterpreter[*flatrecord.Record]

func SetupInterpreter(m *testing.M) {
	pc := falco.NewPolicyCompiler(flatrecord.NewOperations())
	pi = NewPolicyInterpreter(Config{}, pc, nil, nil, nil)
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

func TestCompileSigma(t *testing.T) {
	logger.Trace.Println("Running test compile")
	pc := sigma.NewPolicyCompiler(flatrecord.NewOperations(), "../../../resources/policies/sigma/config/sysflow.yml")
	pi = NewPolicyInterpreter(Config{}, pc, nil, nil, nil)
	paths, err := ioutils.ListFilePaths("../../../resources/policies/sigma/rules/linux/process_creation/proc_creation_lnx_webshell_detection.yml", ".yml")
	assert.NoError(t, err)
	assert.NoError(t, pi.Compile(paths...))
	t.Logf("Rules: %d\n", len(pi.rules))
}
