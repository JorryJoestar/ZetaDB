package container

type Predicate struct {
	/*
		CompareValue           PredicateType = 1
		LikeString             PredicateType = 2
		InRelation             PredicateType = 3
		NotInRelation          PredicateType = 4
		CompareAllRelation     PredicateType = 5
		CompareNotAllRelation  PredicateType = 6
		CompareAnyRelation     PredicateType = 7
		CompareNotAnyRelation  PredicateType = 8
		IsNull                 PredicateType = 9
		IsNotNull              PredicateType = 10
		AttriListInRelation    PredicateType = 11
		AttriListNotInRelation PredicateType = 12
		Exists                 PredicateType = 13
		NotExists              PredicateType = 14
	*/
	PredicateType uint8

	/*
		=  CompareMark = 1
		<> CompareMark = 2
		<  CompareMark = 3
		>  CompareMark = 4
		<= CompareMark = 5
		>= CompareMark = 6
	*/
	CompareMark uint8

	/*
		CompareIntValue     CompareValueType = 1
		CompareFloatValue   CompareValueType = 2
		CompareStringValue  CompareValueType = 3
		CompareBooleanValue CompareValueType = 4
	*/
	CompareValueType    uint8
	CompareIntValue     int
	CompareFloatValue   float64
	CompareStringValue  string
	CompareBooleanValue bool

	LeftAttributeIndex  int
	RightAttributeIndex int
	AttributeIndexList  []int
	Relation            LogicalPlanNode
}
