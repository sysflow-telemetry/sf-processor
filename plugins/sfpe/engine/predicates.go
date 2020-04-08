package engine

import (
	"github.ibm.com/sysflow/sf-processor/plugins/flattener/types"
)

// Predicate defines the type of a functional predicate.
type Predicate func(types.FlatRecord) bool

// True defines a functional predicate that always returns true.
var True = Criterion{func(r types.FlatRecord) bool { return true }}

// False defines a functional predicate that always returns false.
var False = Criterion{func(r types.FlatRecord) bool { return false }}

// Criterion defines an interface for functional predicate operations.
type Criterion struct {
	Pred Predicate
}

// Eval evaluates a functional predicate.
func (c Criterion) Eval(r types.FlatRecord) bool {
	return c.Pred(r)
}

// And computes the conjunction of two functional predicates.
func (c Criterion) And(cr Criterion) Criterion {
	var p Predicate = func(r types.FlatRecord) bool { return c.Eval(r) && cr.Eval(r) }
	return Criterion{p}
}

// Or computes the conjunction of two functional predicates.
func (c Criterion) Or(cr Criterion) Criterion {
	var p Predicate = func(r types.FlatRecord) bool { return c.Eval(r) || cr.Eval(r) }
	return Criterion{p}
}

// Not computes the negation of the function predicate.
func (c Criterion) Not() Criterion {
	var p Predicate = func(r types.FlatRecord) bool { return !c.Eval(r) }
	return Criterion{p}
}

// All derives the conjuctive clause of all predicates in a slice of predicates.
func All(criteria []Criterion) Criterion {
	all := True
	for _, c := range criteria {
		all = all.And(c)
	}
	return all
}

// Any derives the disjuntive clause of all predicates in a slice of predicates.
func Any(criteria []Criterion) Criterion {
	any := False
	for _, c := range criteria {
		any = any.Or(c)
	}
	return any
}

// Exists creates a criterion for an existential predicate.
func Exists(attr string) Criterion {
	//var p Predicate = func(r types.FlatRecord) bool { return r }
	//return Criterion{p}
	return False
}

// Eq creates a criterion for an equality predicate.
func Eq(lattr string, rattr string) Criterion {
	return False
}

// NEq creates a criterion for an inequality predicate.
func NEq(lattr string, rattr string) Criterion {
	return NEq(lattr, rattr).Not()
}

// Ge creates a criterion for a greater-or-equal predicate.
func Ge(lattr string, rattr string) Criterion {
	return False
}

// Gt creates a criterion for a greater-than predicate.
func Gt(lattr string, rattr string) Criterion {
	return False
}

// Le creates a criterion for a lower-or-equal predicate.
func Le(lattr string, rattr string) Criterion {
	return Gt(lattr, rattr).Not()
}

// Lt creates a criterion for a lower-than predicate.
func Lt(lattr string, rattr string) Criterion {
	return Ge(lattr, rattr).Not()
}

// StartsWith creates a criterion for a starts-with predicate.
func StartsWith(lattr string, rattr string) Criterion {
	return False
}

// Contains creates a criterion for a contains predicate.
func Contains(lattr string, rattr string) Criterion {
	return False
}

// IContains creates a criterion for a case-insensitive contains predicate.
func IContains(lattr string, rattr string) Criterion {
	return False
}

// In creates a criterion for a list-inclusion predicate.
func In(attr string, list []string) Criterion {
	return False
}

// PMatch creates a criterion for a list-pattern-matching predicate.
func PMatch(attr string, list []string) Criterion {
	return False
}
