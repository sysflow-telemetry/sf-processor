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
module github.com/sysflow-telemetry/sf-processor/driver

go 1.14

require (
	github.com/actgardner/gogen-avro/v7 v7.3.1
	github.com/kr/pretty v0.2.0 // indirect
	github.com/linkedin/goavro v2.1.0+incompatible
	github.com/spf13/viper v1.6.3
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20210929184451-1b092dc0cfc7
	github.com/sysflow-telemetry/sf-processor/core v0.0.0-20210709164444-484da05c2a85
)

replace github.com/sysflow-telemetry/sf-processor/core => ../core

// replace github.com/sysflow-telemetry/sf-apis/go => ../../sf-apis/go

replace (
	github.com/Shopify/sarama => github.com/elastic/sarama v1.19.1-0.20200629123429-0e7b69039eec
	github.com/dop251/goja => github.com/andrewkroh/goja v0.0.0-20190128172624-dd2ac4456e20
)
