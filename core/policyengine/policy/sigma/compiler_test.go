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
package sigma_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy/sigma"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/flatrecord"
)

var configPath string = "../../../../resources/policies/sigma/config/sysflow.yml"
var rulesPath string = "../../../../resources/policies/sigma/rules/linux/process_creation"

func TestMain(m *testing.M) {
	logger.InitLoggers(logger.TRACE)
	os.Exit(m.Run())
}

func TestCompiler(t *testing.T) {
	pc := sigma.NewPolicyCompiler(flatrecord.NewOperations(), configPath)
	paths, err := ioutils.ListRecursiveFilePaths(rulesPath, ".yml")
	assert.NoError(t, err)
	_, _, err = pc.Compile(paths...)
	assert.NoError(t, err)
}
