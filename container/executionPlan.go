package container

type ExecutionPlanType uint8

const (
	EP_INSERT           ExecutionPlanType = 2
	EP_DELETE           ExecutionPlanType = 3
	EP_UPDATE           ExecutionPlanType = 4
	EP_QUERY            ExecutionPlanType = 5
	EP_CREATE_TABLE     ExecutionPlanType = 6
	EP_DROP_TABLE       ExecutionPlanType = 7
	EP_ALTER_TABLE_ADD  ExecutionPlanType = 8
	EP_ALTER_TABLE_DROP ExecutionPlanType = 9
	EP_CREATE_ASSERT    ExecutionPlanType = 10
	EP_DROP_ASSERT      ExecutionPlanType = 11
	EP_CREATE_VIEW      ExecutionPlanType = 12
	EP_DROP_VIEW        ExecutionPlanType = 13
	EP_CREATE_INDEX     ExecutionPlanType = 14
	EP_DROP_INDEX       ExecutionPlanType = 15
	EP_CREATE_TRIGGER   ExecutionPlanType = 16
	EP_DROP_TRIGGER     ExecutionPlanType = 17
	EP_CREATE_PSM       ExecutionPlanType = 18
	EP_DROP_PSM         ExecutionPlanType = 19
	EP_SHOW_TABLES      ExecutionPlanType = 20
	EP_SHOW_ASSERTIONS  ExecutionPlanType = 21
	EP_SHOW_VIEWS       ExecutionPlanType = 22
	EP_SHOW_INDEXS      ExecutionPlanType = 23
	EP_SHOW_TRIGGERS    ExecutionPlanType = 24
	EP_SHOW_FUNCTIONS   ExecutionPlanType = 25
	EP_SHOW_PROCEDURES  ExecutionPlanType = 26
	EP_CREATE_USER      ExecutionPlanType = 27
	EP_LOG_USER         ExecutionPlanType = 28
	EP_PSM_CALL         ExecutionPlanType = 29
	EP_INIT             ExecutionPlanType = 30
	EP_DROP_USER        ExecutionPlanType = 31
	EP_HALT             ExecutionPlanType = 32
)

type ExecutionPlan struct {
	PlanType        ExecutionPlanType
	UserId          int32
	Parameter       []string         //used by all operators except query
	LogicalPlanRoot *LogicalPlanNode //used by query operator
}

//ExecutionPlan generator
func NewExecutionPlan(planType ExecutionPlanType, userId int32, parameter []string, logicalPlanRoot *LogicalPlanNode) *ExecutionPlan {
	return &ExecutionPlan{
		PlanType:        planType,
		UserId:          userId,
		Parameter:       parameter,
		LogicalPlanRoot: logicalPlanRoot,
	}
}
