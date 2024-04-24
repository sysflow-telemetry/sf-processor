package otel

import (
	"fmt"

	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
	v1 "go.opentelemetry.io/proto/otlp/common/v1"
)

type Contextualizer struct{}

func NewContextualizer() source.Contextualizer[*ResourceLogs] {
	return &Contextualizer{}
}

func (c *Contextualizer) AddRules(logs *ResourceLogs, rules ...policy.Rule[*ResourceLogs]) {
	fmt.Println("Calling Add Rules...")
	attrs := logs.Resource.Attributes
	var kvRule *v1.KeyValue = nil

	fmt.Printf("The attributes len is %d\n", len(attrs))

	if len(attrs) > 1 {
		lastRecord := attrs[len(attrs)-2]
		if lastRecord.Key == "sf.processor.rules" {
			kvRule = lastRecord
		}
	}

	if kvRule == nil && len(attrs) > 0 {
		lastRecord := attrs[len(attrs)-1]
		if lastRecord.Key == "sf.processor.rules" {
			kvRule = lastRecord
		}
	}

	if kvRule == nil {
		arrayValue := &v1.ArrayValue{Values: make([]*v1.AnyValue, 0)}
		anyArrayValue := &v1.AnyValue_ArrayValue{ArrayValue: arrayValue}
		kvRule = &v1.KeyValue{Key: "sf.processor.rules", Value: &v1.AnyValue{Value: anyArrayValue}}
		attrs = append(attrs, kvRule)
	}

	for _, rule := range rules {
		//KeyValue->Value->Value(AnyValue)->Values(ArrayValue)
		v := kvRule.Value.Value
		arrayValue := v.(*ArrayValue)
		array := arrayValue.ArrayValue
		ruleValue := &v1.AnyValue_StringValue{StringValue: rule.Name}
		array.Values = append(array.Values, &v1.AnyValue{Value: ruleValue})
	}

	var kvTag *v1.KeyValue = nil

	if len(attrs) > 1 {
		lastRecord := attrs[len(attrs)-2]
		if lastRecord.Key == "sf.processor.tags" {
			kvTag = lastRecord
		}
	}

	if kvTag == nil && len(attrs) > 0 {
		lastRecord := attrs[len(attrs)-1]
		if lastRecord.Key == "sf.processor.tags" {
			kvRule = lastRecord
		}
	}

	if kvTag == nil {
		arrayValue := &v1.ArrayValue{Values: make([]*v1.AnyValue, 0)}
		anyArrayValue := &v1.AnyValue_ArrayValue{ArrayValue: arrayValue}
		kvTag = &v1.KeyValue{Key: "sf.processor.tags", Value: &v1.AnyValue{Value: anyArrayValue}}
		attrs = append(attrs, kvTag)
	}

	tags := make(map[string]int)

	for _, rule := range rules {
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

	for tag := range tags {
		//KeyValue->Value->Value(AnyValue)->Values(ArrayValue)
		v := kvTag.Value.Value
		arrayValue := v.(*ArrayValue)
		array := arrayValue.ArrayValue
		tagValue := &v1.AnyValue_StringValue{StringValue: tag}
		array.Values = append(array.Values, &v1.AnyValue{Value: tagValue})
	}
	logs.Resource.Attributes = attrs
	fmt.Printf("The attributes len after is %d\n", len(attrs))
	fmt.Printf("logs.Resource.Attributes: %v\n", logs.Resource.Attributes)

}

func (c *Contextualizer) GetRules(r *ResourceLogs) []policy.Rule[*ResourceLogs] {
	return nil
}

func (c *Contextualizer) AddTags(r *ResourceLogs, tags ...string) {
}

func (c *Contextualizer) GetTags(r *ResourceLogs) []string {
	return nil
}
