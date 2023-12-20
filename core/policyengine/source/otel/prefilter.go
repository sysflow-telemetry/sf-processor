package otel

import (
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

type Prefilter struct{}

func NewPrefilter() source.Prefilter[*ResourceLogs] {
	return &Prefilter{}
}

func (p *Prefilter) IsApplicable(r *ResourceLogs, rule policy.Rule[*ResourceLogs]) bool {
	return true
}
