package transports_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/transports"
)

func TestFindings(t *testing.T) {
	config := commons.Config{
		FindingsConfig: commons.FindingsConfig{
			FindingsAccountID:  "79b06b84fc25fe1bd84a1e81d2e73cf0",
			FindingsAPIKey:     "dJaIXtX-b1Kmrh0pw9Ryx_xIQ2hNcIS5KpViEpOEfIwN",
			FindingsProviderID: "system-analytics-pipeline",
			FindingsRegion:     "us-south",
		},
	}
	proto := transports.NewFindingsAPIProto(config)
	if p, ok := proto.(transports.TestableTransportProtocol); ok {
		_, err := p.Test()
		assert.NoError(t, err)
	}
}
