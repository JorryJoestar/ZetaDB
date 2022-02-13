package parser

//---------------------------------------- AST ----------------------------------------
type ASTEnum uint8

const (
	AST_DDL ASTEnum = 1
	AST_DML ASTEnum = 2
	AST_DCL ASTEnum = 3
	AST_DQL ASTEnum = 4
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
	DDL_TABLE_CREATE     DDLEnum = 1
	DDL_TABLE_DROP       DDLEnum = 2
	DDL_TABLE_ALTER_ADD  DDLEnum = 3
	DDL_TABLE_ALTER_DROP DDLEnum = 4
	DDL_ASSERT_CREATE    DDLEnum = 5
	DDL_ASSERT_DROP      DDLEnum = 6
	DDL_VIEW_CREATE      DDLEnum = 7
	DDL_VIEW_DROP        DDLEnum = 8
	DDL_INDEX_CREATE     DDLEnum = 9
	DDL_INDEX_DROP       DDLEnum = 10
	DDL_TRIGGER_CREATE   DDLEnum = 11
	DDL_TRIGGER_DROP     DDLEnum = 12
	DDL_PSM_CREATE       DDLEnum = 13
	DDL_PSM_DROP         DDLEnum = 14
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

type SelectListEntryEnum uint8

const (
	SELECT_LIST_ENTRY_ATTRIBUTE_NAME SelectListEntryEnum = 1 //AttriNameWithTableName
	SELECT_LIST_ENTRY_AGGREGATION    SelectListEntryEnum = 2 //Aggregation
	SELECT_LIST_ENTRY_EXPRESSION     SelectListEntryEnum = 3 //Expression
)

type SelectListEntry struct {
	Type                   SelectListEntryEnum
	AliasValid             bool
	Alias                  string
	AttriNameWithTableName *AttriNameWithTableNameNode
	Aggregation            *AggregationNode
	Expression             *ExpressionNode
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
	DCL_TRANSACTION_BEGIN    DCLEnum = 1
	DCL_TRANSACTION_COMMIT   DCLEnum = 2
	DCL_TRANSACTION_ROLLBACK DCLEnum = 3
	DCL_CONNECTION           DCLEnum = 4
)

//data control language
//transaction, connection
type DCLNode struct {
	Type       DCLEnum
	ServerName string
	UserName   string
	Password   string
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
	Type   DMLEnum
	Update *UpdateNode
	Insert *InsertNode
	Delete *DeleteNode
}

//update
type UpdateNode struct {
	TableName  string
	Condition  *ConditionNode
	UpdateList []*UpdateListEntry
}

type UpdateListEntryEnum uint8

const (
	UPDATE_LIST_ENTRY_EXPRESSION       UpdateListEntryEnum = 1
	UPDATE_LIST_ENTRY_ELEMENTARY_VALUE UpdateListEntryEnum = 2
)

type UpdateListEntry struct {
	AttributeName   string
	ElementaryValue *ElementaryValueNode
	Expression      *ExpressionNode
}

//insert
type InsertEnum uint8

const (
	INSERT_FROM_SUBQUERY  InsertEnum = 1
	INSERT_FROM_VALUELIST InsertEnum = 2
)

type InsertNode struct {
	Type                InsertEnum
	TableName           string
	Subquery            *QueryNode
	AttriNameListValid  bool
	AttriNameList       []string
	ElementaryValueList []*ElementaryValueNode
}

//delete
type DeleteNode struct {
	TableName string
	Condition *ConditionNode
}

//---------------------------------------- public ----------------------------------------

//domain
type DomainEnum uint8

const (
	DOMAIN_CHAR            DomainEnum = 1
	DOMAIN_VARCHAR         DomainEnum = 2 //VARCHAR(n)
	DOMAIN_BIT             DomainEnum = 3 //BIT(n)
	DOMAIN_BITVARYING      DomainEnum = 4 //BITVARYING(n)
	DOMAIN_BOOLEAN         DomainEnum = 5
	DOMAIN_INT             DomainEnum = 6
	DOMAIN_INTEGER         DomainEnum = 7
	DOMAIN_SHORTINT        DomainEnum = 8
	DOMAIN_FLOAT           DomainEnum = 9
	DOMAIN_REAL            DomainEnum = 10
	DOMAIN_DOUBLEPRECISION DomainEnum = 11
	DOMAIN_DECIMAL         DomainEnum = 12 //DECIMAL(n,d)
	DOMAIN_NUMERIC         DomainEnum = 13 //NUMERIC(n,d)
	DOMAIN_DATE            DomainEnum = 14
	DOMAIN_TIME            DomainEnum = 15
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
type ConstraintEnum uint8

const (
	CONSTRAINT_UNIQUE      ConstraintEnum = 1
	CONSTRAINT_PRIMARY_KEY ConstraintEnum = 2
	CONSTRAINT_FOREIGN_KEY ConstraintEnum = 3
	CONSTRAINT_NOT_NULL    ConstraintEnum = 4
	CONSTRAINT_DEFAULT     ConstraintEnum = 5
	CONSTRAINT_CHECK       ConstraintEnum = 6
)

type ConstraintDeferrableEnum uint8

const (
	CONSTRAINT_NOT_DEFERRABLE      ConstraintDeferrableEnum = 1
	CONSTRAINT_INITIALLY_DEFERRED  ConstraintDeferrableEnum = 2
	CONSTRAINT_INITIALLY_IMMEDIATE ConstraintDeferrableEnum = 3
)

type ConstraintUpdateSetEnum uint8

const (
	CONSTRAINT_UPDATE_SET_NULL    ConstraintUpdateSetEnum = 1
	CONSTRAINT_UPDATE_SET_CASCADE ConstraintUpdateSetEnum = 2
)

type ConstraintDeleteSetEnum uint8

const (
	CONSTRAINT_DELETE_SET_NULL    ConstraintDeleteSetEnum = 1
	CONSTRAINT_DELETE_SET_CASCADE ConstraintDeleteSetEnum = 2
)

type ConstraintNode struct {
	Type                 ConstraintEnum
	ConstraintNameValid  bool
	ConstraintName       string
	AttriNameList        []string                 //CONSTRAINT_UNIQUE,CONSTRAINT_PRIMARY_KEY
	Condition            *ConditionNode           //CONSTRAINT_CHECK
	AttributeNameLocal   string                   //CONSTRAINT_FOREIGN_KEY
	AttributeNameForeign string                   //CONSTRAINT_FOREIGN_KEY
	ForeignTableName     string                   //CONSTRAINT_FOREIGN_KEY
	Deferrable           ConstraintDeferrableEnum //CONSTRAINT_FOREIGN_KEY
	UpdateSet            ConstraintUpdateSetEnum  //CONSTRAINT_FOREIGN_KEY
	DeleteSet            ConstraintDeleteSetEnum  //CONSTRAINT_FOREIGN_KEY
}

//elementary value
type ElementaryValueEnum uint8

const (
	ELEMENTARY_VALUE_INT     ElementaryValueEnum = 1
	ELEMENTARY_VALUE_FLOAT   ElementaryValueEnum = 2
	ELEMENTARY_VALUE_STRING  ElementaryValueEnum = 3
	ELEMENTARY_VALUE_BOOLEAN ElementaryValueEnum = 4
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
type CompareMarkEnum uint8

const (
	COMPAREMARK_EQUAL        CompareMarkEnum = 1 // =
	COMPAREMARK_NOTEQUAL     CompareMarkEnum = 2 // <>
	COMPAREMARK_LESS         CompareMarkEnum = 3 // <
	COMPAREMARK_GREATER      CompareMarkEnum = 4 // >
	COMPAREMARK_LESSEQUAL    CompareMarkEnum = 5 // <=
	COMPAREMARK_GREATEREQUAL CompareMarkEnum = 6 // >=
)

type PredicateEnum uint8

const (
	PREDICATE_COMPARE_ELEMENTARY_VALUE PredicateEnum = 1  //AttriNameWithTableNameL, ElementaryValue, CompareMark
	PREDICATE_COMPARE_ATTRIBUTE_NAME   PredicateEnum = 2  //AttriNameWithTableNameL, AttriNameWithTableNameR, CompareMark
	PREDICATE_LIKE_STRING_VALUE        PredicateEnum = 3  //AttriNameWithTableNameL, ElementaryValue(string)
	PREDICATE_IN_SUBQUERY              PredicateEnum = 4  //AttriNameWithTableNameL, Subquery
	PREDICATE_NOT_IN_SUBQUERY          PredicateEnum = 5  //AttriNameWithTableNameL, Subquery
	PREDICATE_IN_TABLE                 PredicateEnum = 6  //AttriNameWithTableNameL, TableName
	PREDICATE_NOT_IN_TABLE             PredicateEnum = 7  //AttriNameWithTableNameL, TableName
	PREDICATE_COMPARE_ALL_SUBQUERY     PredicateEnum = 8  //AttriNameWithTableNameL, Subquery, CompareMark
	PREDICATE_COMPARE_NOT_ALL_SUBQUERY PredicateEnum = 9  //AttriNameWithTableNameL, Subquery, CompareMark
	PREDICATE_COMPARE_ANY_SUBQUERY     PredicateEnum = 10 //AttriNameWithTableNameL, Subquery, CompareMark
	PREDICATE_COMPARE_NOT_ANY_SUBQUERY PredicateEnum = 11 //AttriNameWithTableNameL, Subquery, CompareMark
	PREDICATE_COMPARE_ALL_TABLE        PredicateEnum = 12 //AttriNameWithTableNameL, TableName, CompareMark
	PREDICATE_COMPARE_NOT_ALL_TABLE    PredicateEnum = 13 //AttriNameWithTableNameL, TableName, CompareMark
	PREDICATE_COMPARE_ANY_TABLE        PredicateEnum = 14 //AttriNameWithTableNameL, TableName, CompareMark
	PREDICATE_COMPARE_NOT_ANY_TABLE    PredicateEnum = 15 //AttriNameWithTableNameL, TableName, CompareMark
	PREDICATE_IS_NULL                  PredicateEnum = 16 //AttriNameWithTableNameL
	PREDICATE_IS_NOT_NULL              PredicateEnum = 17 //AttriNameWithTableNameL
	PREDICATE_TUPLE_IN_SUBQUERY        PredicateEnum = 18 //AttributeTuple, Subquery
	PREDICATE_TUPLE_NOT_IN_SUBQUERY    PredicateEnum = 19 //AttributeTuple, Subquery
	PREDICATE_TUPLE_IN_TABLE           PredicateEnum = 20 //AttributeTuple, TableName
	PREDICATE_TUPLE_NOT_IN_TABLE       PredicateEnum = 21 //AttributeTuple, TableName
)

type PredicateNode struct {
	Type                    PredicateEnum
	CompareMark             CompareMarkEnum
	ElementaryValue         *ElementaryValueNode
	AttriNameWithTableNameL *AttriNameWithTableNameNode
	AttriNameWithTableNameR *AttriNameWithTableNameNode
	AttributeTuple          []*AttriNameWithTableNameNode
	Subquery                *QueryNode
	TableName               string
}

//expression
type ExpressionOperatorEnum uint8

const (
	EXPRESSION_OPERATOR_PLUS          ExpressionOperatorEnum = 1
	EXPRESSION_OPERATOR_MINUS         ExpressionOperatorEnum = 2
	EXPRESSION_OPERATOR_DIVISION      ExpressionOperatorEnum = 3
	EXPRESSION_OPERATOR_MULTIPLY      ExpressionOperatorEnum = 4
	EXPRESSION_OPERATOR_CONCATENATION ExpressionOperatorEnum = 5
)

type ExpressionNode struct {
	Type             ExpressionOperatorEnum
	ExpressionEntryL *ExpressionEntryNode
	ExpressionEntryR *ExpressionEntryNode
}

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
	DistinctValid          bool                        //invalid when AGGREGATION_COUNT_ALL
	AttriNameWithTableName *AttriNameWithTableNameNode //invalid when AGGREGATION_COUNT_ALL
}
