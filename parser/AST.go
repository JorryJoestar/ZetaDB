package parser

//---------------------------------------- AST ----------------------------------------
type ASTEnum uint8

const (
	AST_DDL ASTEnum = 0
	AST_DML ASTEnum = 1
	AST_DCL ASTEnum = 2
	AST_DQL ASTEnum = 3
)

type AST struct {
	Type ASTEnum
	Ddl  *DDLNode
	Dml  *DMLNode
	Dcl  *DCLNode
	Dql  *DQLNode
}

//---------------------------------------- DDL ----------------------------------------
type DDLEnum uint8

const (
	DDL_TABLE_CREATE     DDLEnum = 0
	DDL_TABLE_DROP       DDLEnum = 1
	DDL_TABLE_ALTER_ADD  DDLEnum = 2
	DDL_TABLE_ALTER_DROP DDLEnum = 3
	DDL_ASSERT_CREATE    DDLEnum = 4
	DDL_ASSERT_DROP      DDLEnum = 5
	DDL_VIEW_CREATE      DDLEnum = 6
	DDL_VIEW_DROP        DDLEnum = 7
	DDL_INDEX_CREATE     DDLEnum = 8
	DDL_INDEX_DROP       DDLEnum = 9
	DDL_TRIGGER_CREATE   DDLEnum = 10
	DDL_TRIGGER_DROP     DDLEnum = 11
	DDL_PSM_CREATE       DDLEnum = 12
	DDL_PSM_DROP         DDLEnum = 13
)

//data definition language
//CREATE, ALTER, DROP
//table, view, PSM, trigger, constraint, assert, index
type DDLNode struct {
	DdlType DDLEnum
	Table   *TableNode
	Assert  *AssertNode
	View    *ViewNode
	Index   *IndexNode
	Trigger *TableNode
	PSM     *PSMNode
}

//table
type TableNode struct {
	TableName     string //used by create, drop, alter
	AttributeName string // used by alter add, alter drop

	//create table
	AttributeNameList   []string
	DomainList          []*DomainNode
	ConstraintList      []*ConstraintNode
	ConstraintListValid bool

	//alter table add
	Domain     *DomainNode
	Constraint *ConstraintNode
}

//assert
type AssertNode struct {
	AssertName string
	Condition  *ConditionNode
}

//view
type ViewNode struct {
	ViewName               string
	Query                  QueryNode
	AttributeNameList      []string
	AttributeNameListValid bool
}

//index
type IndexNode struct {
	IndexName         string
	TableName         string
	AttributeNameList []string
}

//trigger
type TriggerNode struct {
}

//PSM (function, procedure)
type PSMNode struct {
}

//---------------------------------------- DQL ----------------------------------------

//data query language
//SELECT

type DQLEnum uint8

const (
	DQL_SINGLE_QUERY DQLEnum = 0 //use only QueryL
	DQL_UNION        DQLEnum = 1
	DQL_DIFFERENCE   DQLEnum = 2
	DQL_INTERSECTION DQLEnum = 3
)

type DQLNode struct {
	Type   DQLEnum
	QueryL *QueryNode
	QueryR *QueryNode
}

type JoinEnum uint8

const (
	CROSS_JOIN               JoinEnum = 1
	JOIN_ON                  JoinEnum = 2
	NATURAL_JOIN             JoinEnum = 3
	NATURAL_FULL_OUTER_JOIN  JoinEnum = 4
	NATURAL_LEFT_OUTER_JOIN  JoinEnum = 5
	NATURAL_RIGHT_OUTER_JOIN JoinEnum = 6
	FULL_OUTER_JOIN_ON       JoinEnum = 7
	LEFT_OUTER_JOIN_ON       JoinEnum = 8
	RIGHT_OUTER_JOIN_ON      JoinEnum = 9
)

//query
type QueryNode struct {
	//select
	StarValid     bool
	DistinctValid bool
	SelectList    []*SelectListEntry

	//from
	FromListValid  bool //true then FromList valid, false then Join valid
	FromList       []*FromListEntry
	JoinType       JoinEnum
	JoinTableNameL string
	JoinTableNameR string
	OnList         []*OnListEntry

	//where
	WhereValid     bool
	WhereCondition *ConditionNode

	//group by
	GroupByValid bool
	GroupByList  []*AttriNameWithTableNameNode

	//having
	HavingValid     bool
	HavingCondition *ConditionNode

	//order by
	OrderByValid bool
	OrderByList  []*OrderByListEntry

	//limit
	LimitValid bool
	InitialPos uint32
	OffsetPos  uint32
}

type SelectListEntry struct {
}

type FromListEntryEnum uint8

const (
	FROM_LIST_ENTRY_SUBQUERY FromListEntryEnum = 1
	FROM_LIST_ENTRY_TABLE    FromListEntryEnum = 2
)

type FromListEntry struct {
	Type       FromListEntryEnum
	TableName  string
	Subquery   *QueryNode
	AliasValid bool
	Alias      string
}

type OnListEntry struct {
	AttriNameWithTableNameL *AttriNameWithTableNameNode
	AttriNameWithTableNameR *AttriNameWithTableNameNode
}

type OrderByListEntryEnum uint8

const (
	ORDER_BY_LIST_ENTRY_EXPRESSION OrderByListEntryEnum = 1
	ORDER_BY_LIST_ENTRY_ATTRIBUTE  OrderByListEntryEnum = 2
)

type OrderTrendEnum uint8

const (
	ORDER_BY_LIST_ENTRY_ASC  OrderTrendEnum = 1
	ORDER_BY_LIST_ENTRY_DESC OrderTrendEnum = 2
)

type OrderByListEntry struct {
	Type                   OrderByListEntryEnum
	Trend                  OrderTrendEnum
	Expression             *ExpressionNode
	AttriNameWithTableName *AttriNameWithTableNameNode
}

//---------------------------------------- DCL ----------------------------------------
type DCLEnum uint8

const (
	DCL_TRANSACTION_BEGIN    DCLEnum = 0
	DCL_TRANSACTION_COMMIT   DCLEnum = 1
	DCL_TRANSACTION_ROLLBACK DCLEnum = 2
)

//data control language
//transaction, connection
type DCLNode struct {
	Type DCLEnum
}

//---------------------------------------- DML ----------------------------------------
type DMLEnum uint8

const (
	DML_INSERT DMLEnum = 0
	DML_UPDATE DMLEnum = 1
	DML_DELETE DMLEnum = 2
)

//data manipulation language
//UPDATE, INSERT, DELETE
type DMLNode struct {
	Type DMLEnum
}

//---------------------------------------- public ----------------------------------------

//domain
type DomainEnum uint8

const (
	DOMAIN_CHAR            DomainEnum = 0
	DOMAIN_VARCHAR         DomainEnum = 1 //VARCHAR(n)
	DOMAIN_BIT             DomainEnum = 2 //BIT(n)
	DOMAIN_BITVARYING      DomainEnum = 3 //BITVARYING(n)
	DOMAIN_BOOLEAN         DomainEnum = 4
	DOMAIN_INT             DomainEnum = 5
	DOMAIN_INTEGER         DomainEnum = 6
	DOMAIN_SHORTINT        DomainEnum = 7
	DOMAIN_FLOAT           DomainEnum = 8
	DOMAIN_REAL            DomainEnum = 9
	DOMAIN_DOUBLEPRECISION DomainEnum = 10
	DOMAIN_DECIMAL         DomainEnum = 11 //DECIMAL(n,d)
	DOMAIN_NUMERIC         DomainEnum = 12 //NUMERIC(n,d)
	DOMAIN_DATE            DomainEnum = 13
	DOMAIN_TIME            DomainEnum = 14
)

type DomainNode struct {
	Type DomainEnum
	N    int
	D    int
}

//(TableName.)AttributeName
type AttriNameWithTableNameNode struct {
	TableNameValid bool
	AttributeName  string
	TableName      string
}

//constraint
type ConstraintNode struct {
}

//elementary value
type ElementaryValueEnum uint8

const (
	ELEMENTARY_VALUE_INT     ElementaryValueEnum = 0
	ELEMENTARY_VALUE_FLOAT   ElementaryValueEnum = 1
	ELEMENTARY_VALUE_STRING  ElementaryValueEnum = 2
	ELEMENTARY_VALUE_BOOLEAN ElementaryValueEnum = 3
)

type ElementaryValueNode struct {
	Type         ElementaryValueEnum
	IntValue     int     //ELEMENTARY_VALUE_INT
	FloatValue   float64 //ELEMENTARY_VALUE_FLOAT
	StringValue  string  //ELEMENTARY_VALUE_STRING
	BooleanValue bool    //ELEMENTARY_VALUE_BOOLEAN
}

//condition
type ConditionEnum uint8

const (
	CONDITION_PREDICATE ConditionEnum = 0
	CONDITION_AND       ConditionEnum = 1
	CONDITION_OR        ConditionEnum = 2
)

type ConditionNode struct {
	Type       ConditionEnum
	Predicate  *PredicateNode //CONDITION_PREDICATE
	ConditionL *ConditionNode
	ConditionR *ConditionNode
}

//predicate
type PredicateEnum uint8

const ()

type PredicateNode struct {
}

//expression
type ExpressionNode struct{}

type ExpressionEntryEnum uint8

const (
	EXPRESSION_ENTRY_ELEMENTARY_VALUE ExpressionEntryEnum = 1
	EXPRESSION_ENTRY_ATTRIBUTE_NAME   ExpressionEntryEnum = 2
	EXPRESSION_ENTRY_AGGREGATION      ExpressionEntryEnum = 3
	EXPRESSION_ENTRY_EXPRESSION       ExpressionEntryEnum = 4
)

type ExpressionEntryNode struct {
	Type                   ExpressionEntryEnum
	ElementaryValue        *ElementaryValueNode        //EXPRESSION_ENTRY_ELEMENTARY_VALUE
	AttriNameWithTableName *AttriNameWithTableNameNode //EXPRESSION_ENTRY_ATTRIBUTE_NAME
	Aggregation            *AggregationNode            //EXPRESSION_ENTRY_AGGREGATION
	Expression             *ExpressionNode             //EXPRESSION_ENTRY_EXPRESSION
}

//aggregation
type AggregationEnum uint8

const (
	AGGREGATION_SUM       AggregationEnum = 1
	AGGREGATION_AVG       AggregationEnum = 2
	AGGREGATION_MIN       AggregationEnum = 3
	AGGREGATION_MAX       AggregationEnum = 4
	AGGREGATION_COUNT     AggregationEnum = 5
	AGGREGATION_COUNT_ALL AggregationEnum = 6
)

type AggregationNode struct {
	Type                   AggregationEnum
	DistinctValid          bool
	AttriNameWithTableName *AttriNameWithTableNameNode
}
