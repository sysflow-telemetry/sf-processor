package engine_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.ibm.com/sysflow/sf-processor/core/flattener/types"
	. "github.ibm.com/sysflow/sf-processor/core/sfpe/engine"
)

func TestNot(t *testing.T) {
	c := False
	var r types.FlatRecord
	assert.Equal(t, true, c.Not().Eval(r))
}

func TestAnd(t *testing.T) {
	c := False
	var r types.FlatRecord
	assert.Equal(t, false, c.And(c).Eval(r))
	assert.Equal(t, false, c.And(c.Not()).Eval(r))
	assert.Equal(t, false, c.Not().And(c).Eval(r))
	assert.Equal(t, true, c.Not().And(c.Not()).Eval(r))
}

func TestOr(t *testing.T) {
	c := False
	var r types.FlatRecord
	assert.Equal(t, false, c.Or(c).Eval(r))
	assert.Equal(t, true, c.Or(c.Not()).Eval(r))
	assert.Equal(t, true, c.Not().Or(c).Eval(r))
	assert.Equal(t, true, c.Not().Or(c.Not()).Eval(r))
}
