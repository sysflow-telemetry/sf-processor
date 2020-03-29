package engine

// Predicate defines an interface for a logical predicate.
type Predicate interface {
	Eval(...interface{}) bool
}

// Foo tests predicates.
func Foo(p Predicate, args ...interface{}) bool {
	return p.Eval(args)
}

func Boo(p func(...interface{}) bool, args ...interface{}) bool {
	return p(args)
}

// type And struct {
// 	predicates []interface{}
// }

// func NewAnd(predicates []interface{}) *And {
// 	return &And{newPredicate(), predicates}
// }

// // And returns an AndPredicate with the given predicate.
// func And(predicates ...interface{}) interface{} {
// 	return predicate.NewAnd(predicates)
// }

// // Not returns a NotPredicate with the given predicate.
// func Not(predicates interface{}) interface{} {
// 	return predicate.NewNot(predicates)
// }

// // Or returns an OrPredicate with the given predicate.
// func Or(predicates ...interface{}) interface{} {
// 	return predicate.NewOr(predicates)
// }
