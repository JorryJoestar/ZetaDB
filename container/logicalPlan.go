package container

type LogicalPlanNodeEnum uint8

const (
	BAG_DIFFERENCE         LogicalPlanNodeEnum = 1
	BAG_INTERSECTION       LogicalPlanNodeEnum = 2
	BAG_UNION              LogicalPlanNodeEnum = 3
	DUPLICATE_ELIMINATION  LogicalPlanNodeEnum = 4
	GROUPING               LogicalPlanNodeEnum = 5
	INDEX_FILEREADER       LogicalPlanNodeEnum = 6
	NATURAL_JOIN           LogicalPlanNodeEnum = 7
	PRODUCT                LogicalPlanNodeEnum = 8
	PROJECTION             LogicalPlanNodeEnum = 9
	RENAME                 LogicalPlanNodeEnum = 10
	SELECTION              LogicalPlanNodeEnum = 11
	SEQUENTIAL_FILE_READER LogicalPlanNodeEnum = 12
	SET_DIFFERENCE         LogicalPlanNodeEnum = 13
	SET_INTERSECTION       LogicalPlanNodeEnum = 14
	SET_UNION              LogicalPlanNodeEnum = 15
	THETA_JOIN             LogicalPlanNodeEnum = 16
)

//TODO unfinished
type LogicalPlanNode struct {
	NodeType LogicalPlanNodeEnum

	LeftNode  *LogicalPlanNode
	RightNode *LogicalPlanNode

	//SELECTION
	Condition *Condition

	//SEQUENTIAL_FILE_READER
	TableHeadPageId uint32
	Schema          *Schema
}
