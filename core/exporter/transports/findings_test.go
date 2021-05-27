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
package transports_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/encoders"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/transports"
)

func TestMain(m *testing.M) {
	logger.InitLoggers(logger.TRACE)
	os.Exit(m.Run())
}

func TestCreateOccurrence(t *testing.T) {
	config := commons.Config{
		FindingsConfig: commons.FindingsConfig{
			FindingsAccountID:   "79b06b84fc25fe1bd84a1e81d2e73cf0",
			FindingsApiKey:      "sBdD7pvDDIX7U9muci16gn81IVXa59wqLi3tloFZCcIX",
			FindingsProviderID:  "system-analytics-pipeline",
			FindingsSqlQueryCrn: "crn:v1:bluemix:public:sql-query:us-south:a/49f48a067ac4433a911740653049e83d:abdf1dc1-0232-4083-9f43-67eeaddd6d08::",
			FindingsRegion:      "us-south",
		},
	}
	occ := &encoders.Occurrence{ID: "notification-23243",
		LongDescr:  "This is a long description",
		ShortDescr: "Short description",
		Certainty:  encoders.CertaintyMedium,
		Severity:   encoders.SeverityLow,
		ResName:    "Container ABC",
		ResType:    "Docker",
		NoteID:     "notification",
	}
	proto := transports.NewFindingsApiProto(config)
	err := proto.Export([]commons.EncodedData{occ})
	assert.NoError(t, err)
}
