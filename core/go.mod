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

go 1.17

require (
	github.com/IBM/go-sdk-core/v3 v3.3.1
	github.com/RackSec/srslog v0.0.0-20180709174129-a4725f04ec91
	github.com/actgardner/gogen-avro/v7 v7.3.1
	github.com/antlr/antlr4 v0.0.0-20200417160354-8c50731894e0
	github.com/cespare/xxhash v1.1.0
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210427093042-01613f93a7ae
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-openapi/strfmt v0.20.1
	github.com/ibm-cloud-security/security-advisor-sdk-go/v3 v3.0.0
	github.com/linkedin/goavro v2.1.0+incompatible
	github.com/mailru/easyjson v0.7.6
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6
	github.com/pkg/errors v0.9.1
	github.com/satta/gommunityid v0.0.0-20210315182841-1cdcb73ce408
	github.com/steakknife/bloomfilter v0.0.0-20180922174646-6819c0d2a570
	github.com/stretchr/testify v1.7.0
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20220204025248-25d44ab7fe1f
)

require (
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-openapi/errors v0.19.8 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
	go.mongodb.org/mongo-driver v1.5.1 // indirect
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/linkedin/goavro.v1 v1.0.5 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

//replace github.com/sysflow-telemetry/sf-apis/go => ../../sf-apis/go
