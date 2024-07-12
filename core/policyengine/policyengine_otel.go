//go:build otel
// +build otel

//
// Copyright (C) 2024 IBM Corporation.
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

// Package policyengine implements a plugin for a rules engine for telemetry records.
package policyengine

import (
	logs "go.opentelemetry.io/proto/otlp/logs/v1"
)

// bypassPolicyEngine passes a record onto the exporter if there is no policy engine available.
// note any record transformations can be done here.
func (s *PolicyEngine) bypassPolicyEngine(rec *logs.ResourceLogs) {
	s.out(rec)
}

// processAsync processes a record in the policy engine.
// note any record transformations can be done here.
func (s *PolicyEngine) processAsync(rec *logs.ResourceLogs) {
	s.pi.ProcessAsync(rec)
}
