//go:build otel
// +build otel

//
// Copyright (C) 2023 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
// Anthony Saieva <anthony.saieva@ibm.com>
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

package common

import (
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/otel"
	driver "github.com/sysflow-telemetry/sf-processor/driver/otel"
)

type Record = otel.ResourceLogs

type Channel = driver.OTELChannel

var NewOperations = otel.NewOperations

var NewPrefilter = otel.NewPrefilter

var NewContextualizer = otel.NewContextualizer
