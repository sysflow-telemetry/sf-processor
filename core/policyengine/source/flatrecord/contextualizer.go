package flatrecord

import "github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"

type Contextualizer struct{}

func (s *Contextualizer) AddRules(r *Record, rules ...policy.Rule[*Record]) {
	//TBD
}

func (s *Contextualizer) GetRules(r *Record) []policy.Rule[*Record] {
	//TBD
	return nil
}

func (s *Contextualizer) AddTags(r *Record, tags ...string) {
	//TBD
}

func (s *Contextualizer) GetTags(r *Record) []string {
	//TBD
	return nil
}
