package engine_test

import (
	"os"
	"testing"

	"github.ibm.com/sysflow/sf-processor/common/logger"
	. "github.ibm.com/sysflow/sf-processor/core/sfpe/engine"
)

var pi PolicyInterpreter

func SetupInterpreter(m *testing.M) {
	pi = PolicyInterpreter{}
	os.Exit(m.Run())
}

func TestCompile(t *testing.T) {
	logger.Trace.Println("Running test compile")
	pi.Compile("../../../tests/policies/unit_test_macro.yaml")
}
