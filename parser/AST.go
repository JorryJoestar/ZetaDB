package parser

//---------------------------------------- AST ----------------------------------------
type ASTEnum uint8

const (
	AST_DDL ASTEnum = 1
	AST_DML ASTEnum = 2
	AST_DCL ASTEnum = 3
	AST_DQL ASTEnum = 4
)

type ASTNode struct {
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
//table, view, Psm, trigger, constraint, assert, index
type DDLNode struct {
	Type    DDLEnum
	Table   *TableNode
	Assert  *AssertNode
	View    *ViewNode
	Index   *IndexNode
	Trigger *TriggerNode
	Psm     *PsmNode
}

//table
type TableNode struct {
	TableName           string
	ConstraintName      string //alter drop constraint
	AttributeNameList   []string
	DomainList          []*DomainNode
	ConstraintList      []*ConstraintNode
	ConstraintListValid bool
}

//assert
type AssertNode struct {
	AssertName string
	Condition  *ConditionNode
}

//view
type ViewNode struct {
	ViewName               string
	Query                  *QueryNode
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
type TriggerBeforeAfterEnum uint8

const (
	BEFORE_UPDATE_OF  TriggerBeforeAfterEnum = 1
	BEFORE_UPDATE     TriggerBeforeAfterEnum = 2
	AFTER_UPDATE_OF   TriggerBeforeAfterEnum = 3
	AFTER_UPDATE      TriggerBeforeAfterEnum = 4
	INSTEAD_UPDATE_OF TriggerBeforeAfterEnum = 5
	INSTEAD_UPDATE    TriggerBeforeAfterEnum = 6
	BEFORE_INSERT     TriggerBeforeAfterEnum = 7
	AFTER_INSERT      TriggerBeforeAfterEnum = 8
	INSTEAD_INSERT    TriggerBeforeAfterEnum = 9
	BEFORE_DELETE     TriggerBeforeAfterEnum = 10
	AFTER_DELETE      TriggerBeforeAfterEnum = 11
	INSTEAD_DELETE    TriggerBeforeAfterEnum = 12
)

type TriggerForEachEnum uint8

const (
	FOR_EACH_ROW       TriggerForEachEnum = 1
	FOR_EACH_STATEMENT TriggerForEachEnum = 2
)

type TriggerNode struct {
	TriggerName string

	BeforeAfterType      TriggerBeforeAfterEnum
	BeforeAfterAttriName string
	BeforeAfterTableName string

	ReferencingValid bool
	OldNewList       []*TriggerOldNewEntryNode

	ForEachType TriggerForEachEnum

	WhenValid bool
	Condition *ConditionNode

	DmlList []*DMLNode
}

type TriggerOldNewEnum uint8

const (
	OLD_ROW_AS   TriggerOldNewEnum = 1
	NEW_ROW_AS   TriggerOldNewEnum = 2
	OLD_TABLE_AS TriggerOldNewEnum = 3
	NEW_TABLE_AS TriggerOldNewEnum = 4
)

type TriggerOldNewEntryNode struct {
	Type TriggerOldNewEnum
	Name string
}

//psm (function, procedure)
type PsmEnum uint8

const (
	PSM_PROCEDURE PsmEnum = 1
	PSM_FUNCTION  PsmEnum = 2
)

type PsmNode struct {
	Type              PsmEnum
	PsmName           string
	PsmValueListValid bool            //call
	PsmValueList      []*PsmValueNode //call

	PsmParameterListValid bool
	PsmParameterList      []*PsmParameterEntryNode

	PsmLocalDeclarationListValid bool
	PsmLocalDeclarationList      []*PsmLocalDeclarationEntryNode

	PsmBody []*PsmExecEntryNode
}

type PsmParameterEntryEnum uint8

const (
	PSM_PARAMETER_IN    PsmParameterEntryEnum = 1
	PSM_PARAMETER_OUT   PsmParameterEntryEnum = 2
	PSM_PARAMETER_INOUT PsmParameterEntryEnum = 3
)

type PsmParameterEntryNode struct {
	Type   PsmParameterEntryEnum
	Name   string
	Domain *DomainNode
}

type PsmLocalDeclarationEntryNode struct {
	Name   string
	Domain *DomainNode
}

type PsmExecEntryEnum uint8

const (
	PSM_EXEC_RETURN   PsmExecEntryEnum = 1
	PSM_EXEC_SET      PsmExecEntryEnum = 2
	PSM_EXEC_FOR_LOOP PsmExecEntryEnum = 3
	PSM_EXEC_BRANCH   PsmExecEntryEnum = 4
	PSM_EXEC_DML      PsmExecEntryEnum = 5
)

type PsmExecEntryNode struct {
	Type         PsmExecEntryEnum
	PsmValue     *PsmValueNode
	VariableName string //PSM_EXEC_SET
	PsmForLoop   *PsmForLoopNode
	PsmBranch    *PsmBranchNode
	Dml          *DMLNode
}

type PsmForLoopNode struct {
	LoopName    string
	CursorName  string
	Query       *QueryNode
	PsmExecList []*PsmExecEntryNode
}

type PsmBranchNode struct {
	Condition     *ConditionNode
	IfPsmExecList []*PsmExecEntryNode

	PsmElseifListValid bool
	PsmElseifList      []*PsmElseifEntryNode

	ElseValid       bool
	ElsePsmExecList []*PsmExecEntryNode
}

type PsmElseifEntryNode struct {
	Condition   *ConditionNode
	PsmExecList []*PsmExecEntryNode
}

type PsmValueEnum uint8

const (
	PSMVALUE_ELEMENTARY_VALUE PsmValueEnum = 1
	PSMVALUE_CALL             PsmValueEnum = 2
	PSMVALUE_EXPRESSION       PsmValueEnum = 3
	PSMVALUE_ID               PsmValueEnum = 4
)

type PsmValueNode struct {
	Type PsmValueEnum

	ElementaryValue *ElementaryValueNode
	PsmCall         *PsmNode
	Expression      *ExpressionNode
	Id              string
}

//---------------------------------------- DQL ----------------------------------------

//data query language
//SELECT

type DQLEnum uint8

const (
	DQL_SINGLE_QUERY DQLEnum = 1
	DQL_UNION        DQLEnum = 2
	DQL_DIFFERENCE   DQLEnum = 3
	DQL_INTERSECTION DQLEnum = 4
)

type DQLNode struct {
	Type  DQLEnum
	DqlL  *DQLNode
	DqlR  *DQLNode
	Query *QueryNode
}

//query
type QueryNode struct {
	//select
	StarValid     bool
	DistinctValid bool
	SelectList    []*SelectListEntryNode

	//from

	FromListValid bool //true then FromList valid, false then Join valid
	FromList      []*FromListEntryNode
	Join          *JoinNode

	//where
	WhereValid     bool
	WhereCondition *ConditionNode

	//group by
	GroupByValid bool
	GroupByList  []*AttriNameOptionTableNameNode

	//having
	HavingValid     bool
	HavingCondition *ConditionNode

	//order by
	OrderByValid bool
	OrderByList  []*OrderByListEntryNode

	//limit
	LimitValid bool
	InitialPos int
	OffsetPos  int
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

type JoinNode struct {
	Type           JoinEnum
	JoinTableNameL string
	JoinTableNameR string
	OnList         []*OnListEntryNode
}

type SelectListEntryEnum uint8

const (
	SELECT_LIST_ENTRY_ATTRIBUTE_NAME SelectListEntryEnum = 1 //AttriNameOptionTableName
	SELECT_LIST_ENTRY_AGGREGATION    SelectListEntryEnum = 2 //Aggregation
	SELECT_LIST_ENTRY_EXPRESSION     SelectListEntryEnum = 3 //Expression
)

type SelectListEntryNode struct {
	Type                     SelectListEntryEnum
	AliasValid               bool
	Alias                    string
	AttriNameOptionTableName *AttriNameOptionTableNameNode
	Aggregation              *AggregationNode
	Expression               *ExpressionNode
}

type FromListEntryEnum uint8

const (
	FROM_LIST_ENTRY_SUBQUERY FromListEntryEnum = 1
	FROM_LIST_ENTRY_TABLE    FromListEntryEnum = 2
)

type FromListEntryNode struct {
	Type       FromListEntryEnum
	TableName  string
	Query      *QueryNode
	AliasValid bool
	Alias      string
}

type OnListEntryNode struct {
	AttriNameWithTableNameL *AttriNameOptionTableNameNode
	AttriNameWithTableNameR *AttriNameOptionTableNameNode
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

type OrderByListEntryNode struct {
	Type                     OrderByListEntryEnum
	Trend                    OrderTrendEnum
	Expression               *ExpressionNode
	AttriNameOptionTableName *AttriNameOptionTableNameNode
}

//---------------------------------------- DCL ----------------------------------------
type DCLEnum uint8

const (
	DCL_TRANSACTION_BEGIN    DCLEnum = 1
	DCL_TRANSACTION_COMMIT   DCLEnum = 2
	DCL_TRANSACTION_ROLLBACK DCLEnum = 3
	DCL_SHOW_TABLES          DCLEnum = 4
	DCL_SHOW_ASSERTIONS      DCLEnum = 5
	DCL_SHOW_VIEWS           DCLEnum = 6
	DCL_SHOW_INDEXS          DCLEnum = 7
	DCL_SHOW_TRIGGERS        DCLEnum = 8
	DCL_SHOW_FUNCTIONS       DCLEnum = 9
	DCL_SHOW_PROCEDURES      DCLEnum = 10
	DCL_CREATE_USER          DCLEnum = 11
	DCL_LOG_USER             DCLEnum = 12
	DCL_PSMCALL              DCLEnum = 13
	DCL_INIT                 DCLEnum = 14
	DCL_DROP_USER            DCLEnum = 15
	DCL_HALT                 DCLEnum = 16
)

//data control language
//transaction, connection
type DCLNode struct {
	Type     DCLEnum
	UserName string
	Password string
	PsmCall  *PsmNode
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
	UpdateList []*UpdateListEntryNode
}

type UpdateListEntryEnum uint8

const (
	UPDATE_LIST_ENTRY_EXPRESSION       UpdateListEntryEnum = 1
	UPDATE_LIST_ENTRY_ELEMENTARY_VALUE UpdateListEntryEnum = 2
)

type UpdateListEntryNode struct {
	Type            UpdateListEntryEnum
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
	Query               *QueryNode
	AttriNameListValid  bool
	AttriNameList       []string
	ElementaryValueList []*ElementaryValueNode
}

//delete
type DeleteNode struct {
	TableName string
	Condition *ConditionNode
}

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

//attriNameOptionTableName
type AttriNameOptionTableNameNode struct {
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
	ElementaryValue      *ElementaryValueNode     //CONSTRAINT_DEFAULT
	Condition            *ConditionNode           //CONSTRAINT_CHECK
	AttributeNameLocal   string                   //CONSTRAINT_FOREIGN_KEY, CONSTRAINT_NOT_NULL,CONSTRAINT_DEFAULT
	AttributeNameForeign string                   //CONSTRAINT_FOREIGN_KEY
	ForeignTableName     string                   //CONSTRAINT_FOREIGN_KEY
	DeferrableValid      bool                     //CONSTRAINT_FOREIGN_KEY
	Deferrable           ConstraintDeferrableEnum //CONSTRAINT_FOREIGN_KEY
	UpdateSetValid       bool                     //CONSTRAINT_FOREIGN_KEY
	UpdateSet            ConstraintUpdateSetEnum  //CONSTRAINT_FOREIGN_KEY
	DeleteSetValid       bool                     //CONSTRAINT_FOREIGN_KEY
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
	PREDICATE_LIKE_STRING_VALUE        PredicateEnum = 2  //AttriNameWithTableNameL, ElementaryValue(string)
	PREDICATE_IN_SUBQUERY              PredicateEnum = 3  //AttriNameWithTableNameL, Query
	PREDICATE_NOT_IN_SUBQUERY          PredicateEnum = 4  //AttriNameWithTableNameL, Query
	PREDICATE_IN_TABLE                 PredicateEnum = 5  //AttriNameWithTableNameL, TableName
	PREDICATE_NOT_IN_TABLE             PredicateEnum = 6  //AttriNameWithTableNameL, TableName
	PREDICATE_COMPARE_ALL_SUBQUERY     PredicateEnum = 7  //AttriNameWithTableNameL, Query, CompareMark
	PREDICATE_COMPARE_NOT_ALL_SUBQUERY PredicateEnum = 8  //AttriNameWithTableNameL, Query, CompareMark
	PREDICATE_COMPARE_ANY_SUBQUERY     PredicateEnum = 9  //AttriNameWithTableNameL, Query, CompareMark
	PREDICATE_COMPARE_NOT_ANY_SUBQUERY PredicateEnum = 10 //AttriNameWithTableNameL, Query, CompareMark
	PREDICATE_COMPARE_ALL_TABLE        PredicateEnum = 11 //AttriNameWithTableNameL, TableName, CompareMark
	PREDICATE_COMPARE_NOT_ALL_TABLE    PredicateEnum = 12 //AttriNameWithTableNameL, TableName, CompareMark
	PREDICATE_COMPARE_ANY_TABLE        PredicateEnum = 13 //AttriNameWithTableNameL, TableName, CompareMark
	PREDICATE_COMPARE_NOT_ANY_TABLE    PredicateEnum = 14 //AttriNameWithTableNameL, TableName, CompareMark
	PREDICATE_IS_NULL                  PredicateEnum = 15 //AttriNameWithTableNameL
	PREDICATE_IS_NOT_NULL              PredicateEnum = 16 //AttriNameWithTableNameL
	PREDICATE_TUPLE_IN_SUBQUERY        PredicateEnum = 17 //AttriNameOptionTableNameList, Query
	PREDICATE_TUPLE_NOT_IN_SUBQUERY    PredicateEnum = 18 //AttriNameOptionTableNameList, Query
	PREDICATE_TUPLE_IN_TABLE           PredicateEnum = 19 //AttriNameOptionTableNameList, TableName
	PREDICATE_TUPLE_NOT_IN_TABLE       PredicateEnum = 20 //AttriNameOptionTableNameList, TableName
	PREDICATE_SUBQUERY_EXISTS          PredicateEnum = 21 //Query
	PREDICATE_SUBQUERY_NOT_EXISTS      PredicateEnum = 22 //Query
)

type PredicateNode struct {
	Type                         PredicateEnum
	CompareMark                  CompareMarkEnum
	ElementaryValue              *ElementaryValueNode
	AttriNameWithTableNameL      *AttriNameOptionTableNameNode
	AttriNameWithTableNameR      *AttriNameOptionTableNameNode
	AttriNameOptionTableNameList []*AttriNameOptionTableNameNode
	Query                        *QueryNode
	TableName                    string
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
	Type                     ExpressionEntryEnum
	ElementaryValue          *ElementaryValueNode          //EXPRESSION_ENTRY_ELEMENTARY_VALUE
	AttriNameOptionTableName *AttriNameOptionTableNameNode //EXPRESSION_ENTRY_ATTRIBUTE_NAME
	Aggregation              *AggregationNode              //EXPRESSION_ENTRY_AGGREGATION
	Expression               *ExpressionNode               //EXPRESSION_ENTRY_EXPRESSION
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
	Type                     AggregationEnum
	DistinctValid            bool                          //invalid when AGGREGATION_COUNT_ALL
	AttriNameOptionTableName *AttriNameOptionTableNameNode //invalid when AGGREGATION_COUNT_ALL
}
