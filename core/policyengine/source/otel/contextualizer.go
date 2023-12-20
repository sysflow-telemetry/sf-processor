package otel

import (
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

type Contextualizer struct{}

func NewContextualizer() source.Contextualizer[*ResourceLogs] {
	return &Contextualizer{}
}

func (c *Contextualizer) AddRules(r *ResourceLogs, rules ...policy.Rule[*ResourceLogs]) {
}

func (c *Contextualizer) GetRules(r *ResourceLogs) []policy.Rule[*ResourceLogs] {
	return nil
}

func (c *Contextualizer) AddTags(r *ResourceLogs, tags ...string) {}

func (c *Contextualizer) GetTags(r *ResourceLogs) []string {
	return nil
}
