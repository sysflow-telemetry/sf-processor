package engine_test

import (
	"os"
	"testing"

	"github.ibm.com/sysflow/goutils/logger"
)

func TestMain(m *testing.M) {
	logger.InitLoggers(logger.TRACE)
	os.Exit(m.Run())
}
