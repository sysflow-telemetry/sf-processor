package otel

import (
	common "go.opentelemetry.io/proto/otlp/common/v1"
	logs "go.opentelemetry.io/proto/otlp/logs/v1"
)

/* for now just wrapping around the OTEL types in
case something more elaborate is needed */

type ResourceLogs = logs.ResourceLogs

type KeyValue = common.KeyValue

type StringValue = common.AnyValue_StringValue

type ArrayValue = common.AnyValue_ArrayValue

type BoolValue = common.AnyValue_BoolValue

type BytesValue = common.AnyValue_BytesValue

type DoubleValue = common.AnyValue_DoubleValue

type IntValue = common.AnyValue_IntValue

type KvListValue = common.AnyValue_KvlistValue
