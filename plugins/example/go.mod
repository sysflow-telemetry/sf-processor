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
module github.ibm.com/sysflow/sf-processor/plugins/example

go 1.14

require (
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20200618213240-a59f3a148871
	github.ibm.com/sysflow/goutils v0.0.0-20200528201643-85683bbabbe4
)

replace github.ibm.com/sysflow/goutils => ../../modules/goutils

replace github.com/sysflow-telemetry/sf-apis/go => ../../modules/sf-apis/go