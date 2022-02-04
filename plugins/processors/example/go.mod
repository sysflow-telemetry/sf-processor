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
module github.com/sysflow-telemetry/sf-processor/plugins/processors/example

go 1.14

require (
	github.com/actgardner/gogen-avro v6.5.0+incompatible // indirect
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20220204025248-25d44ab7fe1f
	github.com/sysflow-telemetry/sf-processor/core v0.0.0-20201209134442-13e2a6e66430
)

replace github.com/sysflow-telemetry/sf-processor/core => ../../../core
