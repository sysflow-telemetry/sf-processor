package otel

import (
	v1 "go.opentelemetry.io/proto/otlp/common/v1"
)

/* for now just wrapping around the OTEL types in
case something more elaborate is needed */

type ResourceLogs = v1.ResourceLogs

type KeyValue = v1.KeyValue

type StringValue = v1.AnyValue_StringValue

type ArrayValue = v1.AnyValue_ArrayValue

type BoolValue = v1.AnyValue_BoolValue

type BytesValue = v1.AnyValue_BytesValue

type DoubleValue = v1.AnyValue_DoubleValue

type IntValue = v1.AnyValue_IntValue

type KvListValue = v1.AnyValue_KvlistValue
