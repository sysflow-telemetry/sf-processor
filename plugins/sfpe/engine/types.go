package engine

// Action type for enumeration.
type Action int

// Action enumeration.
const (
	Alert Action = iota
	Tag
)

// String returns the string representation of an action instance.
func (a Action) String() string {
	return [...]string{"alert", "tag"}[a]
}

// EnrichmentTag denotes the type for enrichment tags.
type EnrichmentTag interface{}

// Context denotes the type for contextual information obtained during rule processing.
type Context map[string]interface{}

// Priority denotes the type for rule priority.
type Priority int

// Priority enumeration.
const (
	Low Priority = iota
	Medium
	High
)

// String returns the string representation of a priority instance.
func (p Priority) String() string {
	return [...]string{"low", "medium", "high"}[p]
}

// Rule type
type Rule struct {
	name      string
	desc      string
	condition Criterion
	actions   []Action
	tags      []EnrichmentTag
	priority  Priority
	ctx       Context
}

// Filter type
type Filter struct {
	name      string
	condition Criterion
}
