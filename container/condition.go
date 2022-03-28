package container

type Condition struct {

	/*
		CONDITION_PREDICATE ConditionType = 0
		CONDITION_AND       ConditionType = 1
		CONDITION_OR        ConditionType = 2
	*/
	ConditionType uint8
	Predicate     Predicate
	ConditionL    *Condition
	ConditionR    *Condition
}
