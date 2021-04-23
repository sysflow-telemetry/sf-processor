package monitor

import (
	"errors"

	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
)

// PolicyMonitor is an interface representing policy monitor objects.
// Currently the interface supports a local directory policy monitor.
type PolicyMonitor interface {
	GetInterpreterChan() chan *engine.PolicyInterpreter
	StartMonitor() error
	StopMonitor() error
	CheckForPolicyUpdate() error
}

// NewPolicyMonitor creates a new policy monitor based on the engine configuration.
func NewPolicyMonitor(config engine.Config) (PolicyMonitor, error) {
	if config.Monitor == engine.LocalType {
		return NewLocalPolicyMonitor(config)
	}
	return nil, errors.New("Policy monitor of type: " + config.Monitor.String() + " is not supported.")
}
