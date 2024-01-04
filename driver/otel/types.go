package otel

import "github.com/sysflow-telemetry/sf-processor/core/policyengine/source/otel"

type OTELChannel struct {
	In chan *otel.ResourceLogs
}
