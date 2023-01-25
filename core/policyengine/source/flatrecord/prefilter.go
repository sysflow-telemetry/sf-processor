package flatrecord

import (
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

// Prefilter defines a prefilter object
type Prefilter struct{}

func NewPrefilter() source.Prefilter[*Record] {
	return &Prefilter{}
}

func (s *Prefilter) IsApplicable(r *Record, rule policy.Rule[*Record]) bool {
	if len(rule.Prefilter) == 0 {
		return true
	}
	rtype := Mapper.MapStr(SF_TYPE)(r)
	for _, pf := range rule.Prefilter {
		if rtype == pf {
			return true
		}
	}
	return false
}
