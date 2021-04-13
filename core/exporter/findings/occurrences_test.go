package findings_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	exporter "github.com/sysflow-telemetry/sf-processor/core/exporter"
)

func TestMain(m *testing.M) {
	logger.InitLoggers(logger.TRACE)
	os.Exit(m.Run())
}

func TestCreateOccurrence(t *testing.T) {
	occ := &exporter.Occurrence{ID: "notification-23243",
		LongDescr:  "This is a long description",
		ShortDescr: "Short description",
		Certainty:  exporter.CertaintyMedium,
		Severity:   exporter.SeverityLow,
		ResName:    "Container ABC",
		ResType:    "Docker",
		NoteID:     "notification",
	}
	config := exporter.Config{
		SAAccountID:   "79b06b84fc25fe1bd84a1e81d2e73cf0",
		SAApiKey:      "sBdD7pvDDIX7U9muci16gn81IVXa59wqLi3tloFZCcIX",
		SAProviderID:  "system-analytics-pipeline",
		SASqlQueryCrn: "crn:v1:bluemix:public:sql-query:us-south:a/49f48a067ac4433a911740653049e83d:abdf1dc1-0232-4083-9f43-67eeaddd6d08::",
		Region:        "us-south",
	}
	err := exporter.CreateOccurrence(occ, config)
	assert.NoError(t, err)
}
