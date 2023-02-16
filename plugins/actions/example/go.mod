//
// Copyright (C) 2021 IBM Corporation.
//
// Authors:
// Andreas Schade <san@zurich.ibm.com>
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
module github.com/sysflow-telemetry/sf-processor/plugins/actions/example

go 1.19

require github.com/sysflow-telemetry/sf-processor/core v0.0.0-20220221021811-25c7181c2904

require (
	github.com/actgardner/gogen-avro/v7 v7.3.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20230213025119-faaf3336095c // indirect
	github.com/tidwall/gjson v1.14.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	golang.org/x/exp v0.0.0-20230206171751-46f607a40771 // indirect
)

replace github.com/sysflow-telemetry/sf-processor/core => ../../../core
