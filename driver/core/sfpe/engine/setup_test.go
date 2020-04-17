package engine_test

import (
	"os"
	"testing"

	"github.ibm.com/sysflow/sf-processor/common/logger"
)

func TestMain(m *testing.M) {
	logger.InitLoggers(logger.TRACE)
	os.Exit(m.Run())
}
