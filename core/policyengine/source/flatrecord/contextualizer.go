package flatrecord

import (
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source"
)

type Contextualizer struct{}

func NewContextualizer() source.Contextualizer[*Record] {
	return &Contextualizer{}
}

func (s *Contextualizer) AddRules(r *Record, rules ...policy.Rule[*Record]) {
	r.Ctx.AddRules(rules...)
}

func (s *Contextualizer) GetRules(r *Record) []policy.Rule[*Record] {
	return r.Ctx.GetRules()
}

func (s *Contextualizer) AddTags(r *Record, tags ...string) {
	r.Ctx.AddTags(tags...)
}

func (s *Contextualizer) GetTags(r *Record) []string {
	return r.Ctx.GetTags()
}
