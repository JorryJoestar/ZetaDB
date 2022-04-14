package container

//elementary value
type ConditionEnum uint8

const (
	CONDITION_PREDICATE ConditionEnum = 1
	CONDITION_AND       ConditionEnum = 2
	CONDITION_OR        ConditionEnum = 3
)

type Condition struct {
	ConditionType ConditionEnum
	Predicate     *Predicate
	ConditionL    *Condition
	ConditionR    *Condition
}
