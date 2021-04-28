package findings_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/findings"
)

func TestMain(m *testing.M) {
	logger.InitLoggers(logger.TRACE)
	os.Exit(m.Run())
}

func TestCreateOccurrence(t *testing.T) {
	config := commons.Config{
		SAAccountID:   "79b06b84fc25fe1bd84a1e81d2e73cf0",
		SAApiKey:      "sBdD7pvDDIX7U9muci16gn81IVXa59wqLi3tloFZCcIX",
		SAProviderID:  "system-analytics-pipeline",
		SASqlQueryCrn: "crn:v1:bluemix:public:sql-query:us-south:a/49f48a067ac4433a911740653049e83d:abdf1dc1-0232-4083-9f43-67eeaddd6d08::",
		Region:        "us-south",
	}
	occ := &findings.Occurrence{ID: "notification-23243",
		LongDescr:  "This is a long description",
		ShortDescr: "Short description",
		Certainty:  findings.CertaintyMedium,
		Severity:   findings.SeverityLow,
		ResName:    "Container ABC",
		ResType:    "Docker",
		NoteID:     "notification",
	}
	sa := findings.NewSAClient(config)
	err := sa.CreateOccurrence(occ)
	assert.NoError(t, err)
}
