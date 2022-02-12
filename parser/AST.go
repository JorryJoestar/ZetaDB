package parser

//---------------------------------------- AST ----------------------------------------
type ASTType uint8

const (
	AST_DDL ASTType = 0
	AST_DML ASTType = 1
	AST_DCL ASTType = 2
	AST_DQL ASTType = 3
)

type AST struct {
	AstType ASTType
	Ddl     *DDL
	Dml     *DML
	Dcl     *DCL
	Dql     *DQL
}

//---------------------------------------- DDL ----------------------------------------
type DDLType uint8

const (
	DDL_TABLE_CREATE     DDLType = 0
	DDL_TABLE_DROP       DDLType = 1
	DDL_TABLE_ALTER_ADD  DDLType = 2
	DDL_TABLE_ALTER_DROP DDLType = 3
	DDL_ASSERT_CREATE    DDLType = 4
	DDL_ASSERT_DROP      DDLType = 5
	DDL_VIEW_CREATE      DDLType = 6
	DDL_VIEW_DROP        DDLType = 7
	DDL_INDEX_CREATE     DDLType = 8
	DDL_INDEX_DROP       DDLType = 9
)

//data definition language
//CREATE, ALTER, DROP
//table, view, PSM, trigger, constraint, assert, index
type DDL struct {
	DdlType DDLType
	
}

//table
type Table struct {
	TableName           string
	AttributeNameList   []string
	DomainList          []*Domain
	ConstraintList      []*Constraint
	ConstraintListValid bool
}

//---------------------------------------- DQL ----------------------------------------

//data query language
//SELECT
type DQL struct {
}

//---------------------------------------- DCL ----------------------------------------
type DCLType uint8

const (
	DCL_TRANSACTION_BEGIN    DCLType = 0
	DCL_TRANSACTION_COMMIT   DCLType = 1
	DCL_TRANSACTION_ROLLBACK DCLType = 2
)

//data control language
//transaction, connection
type DCL struct {
	DclType DCLType
}

//---------------------------------------- DML ----------------------------------------
type DMLType uint8

const (
	DML_INSERT DMLType = 0
	DML_UPDATE DMLType = 1
	DML_DELETE DMLType = 2
)

//data manipulation language
//UPDATE, INSERT, DELETE
type DML struct {
	DmlType DMLType
}

//---------------------------------------- common ----------------------------------------

//domain
type DomainType uint8

const (
	DOMAIN_CHAR            DomainType = 0
	DOMAIN_VARCHAR         DomainType = 1 //VARCHAR(n)
	DOMAIN_BIT             DomainType = 2 //BIT(n)
	DOMAIN_BITVARYING      DomainType = 3 //BITVARYING(n)
	DOMAIN_BOOLEAN         DomainType = 4
	DOMAIN_INT             DomainType = 5
	DOMAIN_INTEGER         DomainType = 6
	DOMAIN_SHORTINT        DomainType = 7
	DOMAIN_FLOAT           DomainType = 8
	DOMAIN_REAL            DomainType = 9
	DOMAIN_DOUBLEPRECISION DomainType = 0
	DOMAIN_DECIMAL         DomainType = 1 //DECIMAL(n,d)
	DOMAIN_NUMERIC         DomainType = 2 //NUMERIC(n,d)
	DOMAIN_DATE            DomainType = 3
	DOMAIN_TIME            DomainType = 4
)

type Domain struct {
	Type DomainType
	N    int
	D    int
}

//(TableName.)AttributeName
type AttriNameOptTableName struct {
	TableNameValid bool
	AttributeName  string
	TableName      string
}

//constraint
type Constraint struct {
}
