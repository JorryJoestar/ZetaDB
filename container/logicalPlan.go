package container

//default as bag operation
type LogicalPlan struct {
}

type LogicalPlanNode interface {
}

type UnionNode struct {
	NodeL LogicalPlanNode
	NodeR LogicalPlanNode
}

type IntersectionNode struct {
	NodeL LogicalPlanNode
	NodeR LogicalPlanNode
}

type DifferenceNode struct {
	NodeL LogicalPlanNode
	NodeR LogicalPlanNode
}

type SelectionNode struct {
	Node            LogicalPlanNode
	SelectCondition Condition
}

type ProjectionNode struct {
	Node                LogicalPlanNode
	ProjectionIndexList []int
}

type ProductNode struct {
	NodeL LogicalPlanNode
	NodeR LogicalPlanNode
}

type NaturalNode struct {
	NodeL LogicalPlanNode
	NodeR LogicalPlanNode
}

type ThetaNode struct {
	NodeL          LogicalPlanNode
	NodeR          LogicalPlanNode
	ThetaCondition Condition
}

type RenameNode struct {
	NewTableName         string
	NewAttributeNameList []string
}

type DuplicateNode struct {
	Node LogicalPlanNode
}

type RelationNode struct {
	TableId uint32
}
