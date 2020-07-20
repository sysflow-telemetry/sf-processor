package engine_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.ibm.com/sysflow/sf-processor/core/policyengine/engine"
)

func TestNot(t *testing.T) {
	c := False
	var r *Record
	assert.Equal(t, true, c.Not().Eval(r))
}

func TestAnd(t *testing.T) {
	c := False
	var r *Record
	assert.Equal(t, false, c.And(c).Eval(r))
	assert.Equal(t, false, c.And(c.Not()).Eval(r))
	assert.Equal(t, false, c.Not().And(c).Eval(r))
	assert.Equal(t, true, c.Not().And(c.Not()).Eval(r))
}

func TestOr(t *testing.T) {
	c := False
	var r *Record
	assert.Equal(t, false, c.Or(c).Eval(r))
	assert.Equal(t, true, c.Or(c.Not()).Eval(r))
	assert.Equal(t, true, c.Not().Or(c).Eval(r))
	assert.Equal(t, true, c.Not().Or(c.Not()).Eval(r))
}

func TestAll(t *testing.T) {
	var r *Record
	assert.Equal(t, true, All([]Criterion{True, True}).Eval(r))
	assert.Equal(t, false, All([]Criterion{True, False}).Eval(r))
	assert.Equal(t, false, All([]Criterion{False, True}).Eval(r))
	assert.Equal(t, false, All([]Criterion{False, False}).Eval(r))
}

func TestAny(t *testing.T) {
	var r *Record
	assert.Equal(t, true, Any([]Criterion{True, True}).Eval(r))
	assert.Equal(t, true, Any([]Criterion{True, False}).Eval(r))
	assert.Equal(t, true, Any([]Criterion{False, True}).Eval(r))
	assert.Equal(t, false, Any([]Criterion{False, False}).Eval(r))
}
