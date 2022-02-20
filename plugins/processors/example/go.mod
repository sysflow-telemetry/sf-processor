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

go 1.17

require (
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20220220141421-08882228d4d8
	github.com/sysflow-telemetry/sf-processor/core v0.0.0-20220218215337-cb82d403e9f6
)

require (
	github.com/actgardner/gogen-avro/v7 v7.3.1 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6 // indirect
)

replace github.com/sysflow-telemetry/sf-processor/core => ../../../core
