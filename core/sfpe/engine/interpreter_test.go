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
	pi.Compile("../tests/policies/macro_test.yaml")
}

func TestVisitor(t *testing.T) {
	var m = make(map[string]interface{})
	m["sf.proc.pid"] = 1234
	m["sf.proc.name"] = "/bin/bash"
	logger.Trace.Printf("pid: %d name: %s\n", m["sf.proc.pid"], m["sf.proc.name"])
}
