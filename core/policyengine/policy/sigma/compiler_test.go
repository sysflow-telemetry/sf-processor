package sigma_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy/sigma"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/flatrecord"
)

var configPath string = "../../../../resources/policies/sigma/config/sysflow.yml"
var rulesPath string = "../../../../resources/policies/sigma/rules/linux/process_creation"

func TestMain(m *testing.M) {
	logger.InitLoggers(logger.TRACE)
	os.Exit(m.Run())
}

func TestCompiler(t *testing.T) {
	pc := sigma.NewPolicyCompiler(flatrecord.NewOperations(), configPath)
	paths, err := ioutils.ListRecursiveFilePaths(rulesPath, ".yml")
	assert.NoError(t, err)
	_, _, err = pc.Compile(paths...)
	assert.NoError(t, err)
}
