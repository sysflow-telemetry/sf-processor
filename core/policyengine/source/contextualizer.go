package source

import (
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
)

type Contextualizer[R any] interface {
	// AddRule adds one or more rules matching a record.
	AddRules(r R, rules ...policy.Rule[R])
	// GetRules retrieves the list of stored rules associated with a record.
	GetRules(r R) []policy.Rule[R]
	// Adds one or more tags to a record.
	AddTags(r R, tags ...string)
	// GetTags retrieves the list of tags associated with a record.
	GetTags(r R) []string
}

// // A prototype for an engine that uses the contextualizer
// type Engine[R any] struct {
// 	ctx Contextualizer[R]
// }

// func NewEngine[R any](c Contextualizer[R]) *Engine[R] {
// 	return &Engine[R]{c}
// }

// func (s *Engine[R]) Test() {}

// // A consumer that uses the engine (our Processor plugin, for example)
// func consumer() {
// 	e := NewEngine[*flatrecord.Record](&flatrecord.Contextualizer{})
// 	e.Test()
// }
