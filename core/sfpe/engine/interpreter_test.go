package engine_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.ibm.com/sysflow/goutils/ioutils"
	"github.ibm.com/sysflow/goutils/logger"
	. "github.ibm.com/sysflow/sf-processor/core/sfpe/engine"
)

var pi PolicyInterpreter

func SetupInterpreter(m *testing.M) {
	pi = PolicyInterpreter{}
	os.Exit(m.Run())
}

func TestCompile(t *testing.T) {
	logger.Trace.Println("Running test compile")
	paths, err := ioutils.ListFilePaths("../../../tests/policies", ".yaml")
	assert.NoError(t, err)
	assert.NoError(t, pi.Compile(paths...))
}
