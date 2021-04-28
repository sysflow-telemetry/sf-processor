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
module github.com/sysflow-telemetry/sf-processor/core

go 1.14

require (
	github.com/IBM/go-sdk-core/v3 v3.3.1
	github.com/RackSec/srslog v0.0.0-20180709174129-a4725f04ec91
	github.com/antlr/antlr4 v0.0.0-20200417160354-8c50731894e0
	github.com/cespare/xxhash v1.1.0
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210427093042-01613f93a7ae
	github.com/enriquebris/goconcurrentqueue v0.6.0
	github.com/go-openapi/strfmt v0.19.4
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/ibm-cloud-security/security-advisor-sdk-go v1.1.1
	github.com/mailru/easyjson v0.7.6
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6
	github.com/stretchr/testify v1.6.1
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20201207153955-828257760aa4
	golang.org/x/net v0.0.0-20200222125558-5a598a2470a0 // indirect
	golang.org/x/sys v0.0.0-20200223170610-d5e6a3e2c0ae // indirect
)

replace github.com/sysflow-telemetry/sf-apis/go => ../../sf-apis/go
