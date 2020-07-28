//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
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
