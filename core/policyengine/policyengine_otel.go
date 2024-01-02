//go:build otel
// +build otel

package policyengine

import (
	logs "go.opentelemetry.io/proto/otlp/logs/v1"
)

func (s *PolicyEngine) bypassPolicyEngine(rec *logs.ResourceLogs) {
	s.out(rec)
}

func (s *PolicyEngine) processAsync(rec *logs.ResourceLogs) {
	s.pi.ProcessAsync(rec)
}
