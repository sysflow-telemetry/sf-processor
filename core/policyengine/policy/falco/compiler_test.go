package falco_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy/falco"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/flatrecord"
)

var rulesPath string = "../../../../resources/policies/runtimeintegrity"

func TestMain(m *testing.M) {
	logger.InitLoggers(logger.TRACE)
	os.Exit(m.Run())
}

func TestCompiler(t *testing.T) {
	pc := falco.NewPolicyCompiler(flatrecord.NewOperations())
	paths, err := ioutils.ListRecursiveFilePaths(rulesPath, ".yaml")
	assert.NoError(t, err)
	_, _, err = pc.Compile(paths...)
	assert.NoError(t, err)
}
