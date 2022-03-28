package container

type Predicate interface {
	ConductPredicate() bool
}

func PredicateFactory() {

}

//
type CompareValuePredicate struct {
}
