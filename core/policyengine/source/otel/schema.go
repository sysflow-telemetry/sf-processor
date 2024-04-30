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

// Package otel implements an open telemetry backend for the policy compilers.
package otel

import (
	common "go.opentelemetry.io/proto/otlp/common/v1"
	logs "go.opentelemetry.io/proto/otlp/logs/v1"
)

type ResourceLogs = logs.ResourceLogs

type KeyValue = common.KeyValue

type StringValue = common.AnyValue_StringValue

type ArrayValue = common.AnyValue_ArrayValue

type BoolValue = common.AnyValue_BoolValue

type BytesValue = common.AnyValue_BytesValue

type DoubleValue = common.AnyValue_DoubleValue

type IntValue = common.AnyValue_IntValue

type KvListValue = common.AnyValue_KvlistValue
