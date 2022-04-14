package container

type PhysicalPlanType uint8

const (
	INITIALIZE_SYSTEM PhysicalPlanType = 1
	INSERT            PhysicalPlanType = 2
	DELETE            PhysicalPlanType = 3
	UPDATE            PhysicalPlanType = 4
	QUERY             PhysicalPlanType = 5
	CREATE_TABLE      PhysicalPlanType = 6
	DROP_TABLE        PhysicalPlanType = 7
	ALTER_TABLE_ADD   PhysicalPlanType = 8
	ALTER_TABLE_DROP  PhysicalPlanType = 9
	CREATE_ASSERT     PhysicalPlanType = 10
	DROP_ASSERT       PhysicalPlanType = 11
	CREATE_VIEW       PhysicalPlanType = 12
	DROP_VIEW         PhysicalPlanType = 13
	CREATE_INDEX      PhysicalPlanType = 14
	DROP_INDEX        PhysicalPlanType = 15
	CREATE_TRIGGER    PhysicalPlanType = 16
	DROP_TRIGGER      PhysicalPlanType = 17
	CREATE_PSM        PhysicalPlanType = 18
	DROP_PSM          PhysicalPlanType = 19
	SHOW_TABLES       PhysicalPlanType = 20
	SHOW_ASSERTIONS   PhysicalPlanType = 21
	SHOW_VIEWS        PhysicalPlanType = 22
	SHOW_INDEXS       PhysicalPlanType = 23
	SHOW_TRIGGERS     PhysicalPlanType = 24
	SHOW_FUNCTIONS    PhysicalPlanType = 25
	SHOW_PROCEDURES   PhysicalPlanType = 26
	CREATE_USER       PhysicalPlanType = 27
	LOG_USER          PhysicalPlanType = 28
	PSM_CALL          PhysicalPlanType = 29
)

type PhysicalPlan struct {
	PlanType      PhysicalPlanType
	Parameter     []string       //used by all operators except query
	QueryTreeRoot *QueryTreeNode //used by query operator
}

//PhysicalPlan generator
func NewPhysicalPlan(planType PhysicalPlanType, parameter []string, queryTreeRoot *QueryTreeNode) *PhysicalPlan {
	return &PhysicalPlan{
		PlanType:      planType,
		Parameter:     parameter,
		QueryTreeRoot: queryTreeRoot,
	}
}

type QueryTreeNode struct {
}
