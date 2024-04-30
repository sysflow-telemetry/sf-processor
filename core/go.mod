//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
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
module github.com/sysflow-telemetry/sf-processor/core

go 1.19

require (
	github.com/RackSec/srslog v0.0.0-20180709174129-a4725f04ec91
	github.com/antlr/antlr4 v0.0.0-20200417160354-8c50731894e0
	github.com/bradleyjkemp/sigma-go v0.5.1
	github.com/cespare/xxhash/v2 v2.1.2
	github.com/confluentinc/confluent-kafka-go/v2 v2.3.0
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210427093042-01613f93a7ae
	github.com/fsnotify/fsnotify v1.5.1
	github.com/mailru/easyjson v0.7.6
	github.com/paulbellamy/ratecounter v0.2.0
	github.com/pkg/errors v0.9.1
	github.com/satta/gommunityid v0.0.0-20210315182841-1cdcb73ce408
	github.com/stretchr/testify v1.9.0
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20240224034728-382270fa878c
	github.com/tidwall/gjson v1.14.1
	go.opentelemetry.io/proto/otlp v0.19.0
	golang.org/x/exp v0.0.0-20230206171751-46f607a40771
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/actgardner/gogen-avro/v7 v7.3.1 // indirect
	github.com/alecthomas/participle v0.7.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
