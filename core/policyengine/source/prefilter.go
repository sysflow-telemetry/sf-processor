package source

import "github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"

// Prefilter interface
type Prefilter[R any] interface {
	IsApplicable(r R, rule policy.Rule[R]) bool
}

// DefaultPrefilter defines a prefilter object to be used as a default prefilter.
type DefaultPrefilter[R any] struct{}

func NewDefaultPrefilter[R any]() Prefilter[R] {
	return &DefaultPrefilter[R]{}
}

func (s *DefaultPrefilter[R]) IsApplicable(r R, rule policy.Rule[R]) bool {
	return true
}
