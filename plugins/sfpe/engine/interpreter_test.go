package engine_test

import (
	"os"
	"testing"

	"github.ibm.com/sysflow/sf-processor/common/logger"
	. "github.ibm.com/sysflow/sf-processor/plugins/sfpe/engine"
)

func TestMain(m *testing.M) {
	logger.InitLoggers(logger.TRACE)
	os.Exit(m.Run())
}

func TestCompile(t *testing.T) {
	logger.Trace.Println("Running test compile")
	Compile("../tests/policies/macro_test.yaml")
}

func TestPredicates(t *testing.T) {
	var res = Boo(func(i, j int) bool {
		return i < j
	}, 1, 2)
}

// func TestRegex(t *testing.T) {
// 	re2 := regexp.MustCompile(`(^\[)(.*)(\]$?)`)
// 	fmt.Println(re2.ReplaceAllString("[asdf,c,asdf,12,uioe]", "$2"))

// }
