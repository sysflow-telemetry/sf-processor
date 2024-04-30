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
	"fmt"

	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
	otelcommon "go.opentelemetry.io/proto/otlp/common/v1"
	otellogs "go.opentelemetry.io/proto/otlp/logs/v1"
)

type Contextualizer struct{}

func NewContextualizer() source.Contextualizer[*ResourceLogs] {
	return &Contextualizer{}
}

func (c *Contextualizer) AddRules(logs *ResourceLogs, rules ...policy.Rule[*ResourceLogs]) {

	// get or create scope logs for sf-processor enrichment
	attrs := c.getOrCreateProcessorScopeAttributes(logs)

	// get or create kv attributes for rule names and tags
	kvRule := c.getOrCreateAttribute(attrs, SF_PROCESSOR_RULES)
	kvTag := c.getOrCreateAttribute(attrs, SF_PROCESSOR_TAGS)

	// set rules and retrieve tags
	tags := make(map[string]int)
	for _, rule := range rules {
		v := kvRule.Value.Value
		arrayValue := v.(*ArrayValue)
		array := arrayValue.ArrayValue
		ruleValue := &otelcommon.AnyValue_StringValue{StringValue: rule.Name}
		array.Values = append(array.Values, &otelcommon.AnyValue{Value: ruleValue})
		for _, tag := range rule.Tags {
			switch v := tag.(type) {
			case []string:
				for _, s := range v {
					tags[s] = 1
				}
			default:
				s := string(fmt.Sprintf("%v", v))
				tags[s] = 1
			}
		}
	}

	// set tags
	for tag := range tags {
		v := kvTag.Value.Value
		arrayValue := v.(*ArrayValue)
		array := arrayValue.ArrayValue
		tagValue := &otelcommon.AnyValue_StringValue{StringValue: tag}
		array.Values = append(array.Values, &otelcommon.AnyValue{Value: tagValue})
	}
}

func (c *Contextualizer) GetRules(r *ResourceLogs) []policy.Rule[*ResourceLogs] {
	return nil
}

func (c *Contextualizer) AddTags(r *ResourceLogs, tags ...string) {}

func (c *Contextualizer) GetTags(r *ResourceLogs) []string {
	return nil
}

func (c *Contextualizer) getOrCreateProcessorScopeAttributes(logs *ResourceLogs) *[]*otelcommon.KeyValue {
	for _, scopeLog := range logs.ScopeLogs {
		if scopeLog.Scope != nil && scopeLog.Scope.Name == SF_PROCESSOR_SCOPE_NAME {
			return &scopeLog.Scope.Attributes
		}
	}
	var sl *otellogs.ScopeLogs = &otellogs.ScopeLogs{}
	sl.Scope = &otelcommon.InstrumentationScope{Name: SF_PROCESSOR_SCOPE_NAME}
	sl.Scope.Attributes = make([]*otelcommon.KeyValue, 0)
	logs.ScopeLogs = append(logs.ScopeLogs, sl)
	return &sl.Scope.Attributes
}

func (c *Contextualizer) getOrCreateAttribute(attrs *[]*otelcommon.KeyValue, key string) *otelcommon.KeyValue {
	if len(*attrs) > 1 {
		lastAttr := (*attrs)[len(*attrs)-2]
		if lastAttr.Key == key {
			return lastAttr
		}
	}
	if len(*attrs) > 0 {
		lastAttr := (*attrs)[len(*attrs)-1]
		if lastAttr.Key == key {
			return lastAttr
		}
	}
	arrayValue := &otelcommon.ArrayValue{Values: make([]*otelcommon.AnyValue, 0)}
	anyArrayValue := &otelcommon.AnyValue_ArrayValue{ArrayValue: arrayValue}
	kvAttr := &otelcommon.KeyValue{Key: key, Value: &otelcommon.AnyValue{Value: anyArrayValue}}
	*attrs = append(*attrs, kvAttr)
	return kvAttr
}
