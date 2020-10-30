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
module github.ibm.com/sysflow/sf-processor/core

go 1.14

require (
	github.com/RackSec/srslog v0.0.0-20180709174129-a4725f04ec91
	github.com/antlr/antlr4 v0.0.0-20200417160354-8c50731894e0
	github.com/cespare/xxhash v1.1.0
	github.com/cespare/xxhash/v2 v2.1.1
	github.com/enriquebris/goconcurrentqueue v0.6.0
	github.com/mailru/easyjson v0.7.6
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6
	github.com/stretchr/testify v1.6.1
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20200618213240-a59f3a148871
)

replace github.com/sysflow-telemetry/sf-apis/go => ../modules/sf-apis/go
