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

go 1.14

require (
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v20.10.11+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/sysflow-telemetry/sf-processor/core v0.0.0-20201209134442-13e2a6e66430
)

replace github.com/sysflow-telemetry/sf-processor/core => ../../../core
