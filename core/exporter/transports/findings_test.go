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
			FindingsAccountID:  "<accountid>",
			FindingsAPIKey:     "<apikey>",
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
