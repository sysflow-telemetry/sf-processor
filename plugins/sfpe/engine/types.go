package engine

// Action type for enumeration.
type Action string

// Action enumeration.
const (
	Alert Action = "alert"
	Tag   Action = "tag"
)

// EnrichmentTag denotes a type for enrichment tags.
type EnrichmentTag interface{}

// Rule type
type Rule struct {
	name      string
	desc      string
	condition Criterion
	actions   []Action
	tags      []EnrichmentTag
}

// Filter type
type Filter struct {
	name      string
	condition Criterion
}
