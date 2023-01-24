package engine

import "github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"

// Prefilter interface
type Prefilter[R any] interface {
	IsApplicable(r R, rule policy.Rule[R]) bool
}
