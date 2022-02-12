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
)

//data definition language
//CREATE, ALTER, DROP
//table, view, PSM, trigger, constraint, assert, index
type DDLNode struct {
	DdlType DDLEnum
	Table   *TableNode
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
	

}

//---------------------------------------- DQL ----------------------------------------

//data query language
//SELECT
type DQLNode struct {
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

//---------------------------------------- common ----------------------------------------

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
	DOMAIN_DOUBLEPRECISION DomainEnum = 0
	DOMAIN_DECIMAL         DomainEnum = 1 //DECIMAL(n,d)
	DOMAIN_NUMERIC         DomainEnum = 2 //NUMERIC(n,d)
	DOMAIN_DATE            DomainEnum = 3
	DOMAIN_TIME            DomainEnum = 4
)

type DomainNode struct {
	Type DomainEnum
	N    int
	D    int
}

//(TableName.)AttributeName
type AttriNameOptTableNameNode struct {
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
	IntValue     int
	FloatValue   float64
	StringValue  string
	BooleanValue bool
}
