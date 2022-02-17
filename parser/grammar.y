%{
package parser

import (
)

// -------------------- Node --------------------
type NodeEnum uint8

const (
/* ast */
    AST_NODE                        NodeEnum = 1

/* ddl */
    DDL_NODE                        NodeEnum = 2
    TABLE_NODE                      NodeEnum = 3
    ASSERT_NODE                     NodeEnum = 4
    VIEW_NODE                       NodeEnum = 5
    INDEX_NODE                      NodeEnum = 6
    TRIGGER_NODE                    NodeEnum = 7
    PSM_NODE                        NodeEnum = 8

/* dql */
    DQL_NODE                        NodeEnum = 9
    QUERY_NODE                      NodeEnum = 10
    SELECT_LIST_ENTRY               NodeEnum = 11
    FROM_LIST_ENTRY                 NodeEnum = 12
    ON_LIST_ENTRY                   NodeEnum = 13
    ORDERBY_LIST_ENTRY              NodeEnum = 14

/* dcl */
    DCL_NODE                        NodeEnum = 15

/* dml */
    DML_NODE                        NodeEnum = 16
    UPDATE_NODE                     NodeEnum = 17
    UPDATE_LIST_ENTRY               NodeEnum = 18
    INSERT_NODE                     NodeEnum = 19
    DELETE_NODE                     NodeEnum = 20

/* constraint */
    CONSTRAINT_NODE                 NodeEnum = 23
    CONSTRAINT_DEFERRABLE_ENUM      NodeEnum = 30
    CONSTRAINT_UPDATE_SET_ENUM      NodeEnum = 31
    CONSTRAINT_DELETE_SET_ENUM      NodeEnum = 32
    FOREIGNKEY_PARAMETER_NODE       NodeEnum = 37

/* AttriNameOptionTableName */
    ATTRINAME_OPTION_TABLENAME_NODE NodeEnum = 33

/* elementaryValue */
    ELEMENTARY_VALUE_NODE           NodeEnum = 24

/* domain */
    DOMAIN_NODE                     NodeEnum = 21

    CONDITION_NODE                  NodeEnum = 25
    PREDICATE_NODE                  NodeEnum = 26
    EXPRESSION_NODE                 NodeEnum = 27
    EXPRESSION_ENTRY                NodeEnum = 28
    AGGREGATION_NODE                NodeEnum = 29

/* predicate */
    COMPAREMARK_ENUM                NodeEnum = 34

/* subquery */
    SUBQUERY_NODE                   NodeEnum = 35

/* createTable */
    ATTRIBUTE_DECLARATION_NODE      NodeEnum = 36

// NodeEnum = 38

)

type Node struct {
    Type NodeEnum

/* ast */
    Ast *ASTNode
    Ddl *DDLNode
    Dql *DQLNode
    Dcl *DCLNode
    Dml *DMLNode

/* ddl */
    Table   *TableNode
    Assert  *AssertNode
    View    *ViewNode
    Index   *IndexNode
    Trigger *TriggerNode
    Psm     *PsmNode

/* dql */
    Query            *QueryNode
    SelectListEntry  *SelectListEntryNode
    FromListEntry    *FromListEntryNode
    OnListEntry      *OnListEntryNode
    OrderByListEntry *OrderByListEntryNode

/* dml */
    Update          *UpdateNode
    UpdateListEntry *UpdateListEntryNode
    Insert          *InsertNode
    DeleteNode      *DeleteNode

/* createTable */
    AttributeDeclaration *AttributeDeclarationNode

/* constraint */
    ConstraintDeferrable ConstraintDeferrableEnum
    ConstraintUpdateSet  ConstraintUpdateSetEnum
    ConstraintDeleteSet  ConstraintDeleteSetEnum
    Constraint           *ConstraintNode
    ForeignKeyParameter  *ForeignKeyParameterNode

/* condition */
    Condition                *ConditionNode

/* predicate */
    Predicate                *PredicateNode
    CompareMark              CompareMarkEnum

/* attriNameOptionTableName */
    AttriNameOptionTableName *AttriNameOptionTableNameNode

/* elementaryValue */
    ElementaryValue          *ElementaryValueNode

/* domain */
    Domain                   *DomainNode

/* public */

    Expression               *ExpressionNode
    ExpressionEntry          *ExpressionEntryNode
    Aggregation              *AggregationNode
    Subquery                 *QueryNode
}

// -------------------- List --------------------
type ListEnum uint8

const (
    CONSTRAINT_AFTER_ATTRIBUTE_LIST ListEnum = 1
    CONSTRAINT_LIST                 ListEnum = 2
    ATTRINAME_OPTION_TABLENAME_LIST ListEnum = 3
    ATTRIBUTE_DECLARATION_LIST      ListEnum = 4
)

type List struct {
    Type                         ListEnum
    ConstraintAfterAttributeList []*ConstraintNode
    ConstraintList               []*ConstraintNode
    AttriNameOptionTableNameList []*AttriNameOptionTableNameNode
    AttributeDeclarationList     []*AttributeDeclarationNode
}

// -------------------- temporary struct --------------------
// temporary struct, not included in AST, assistant grammar.y to generate AST

// attributeDeclaration
type AttributeDeclarationNode struct {
    AttributeName                     string
    Domain                            *DomainNode
    ConstraintAfterAttributeListValid bool
    ConstraintAfterAttributeList      []*ConstraintNode
}

// foreignKeyParameter
type ForeignKeyParameterNode struct {
    DeferrableValid      bool
	Deferrable           ConstraintDeferrableEnum
	UpdateSetValid       bool
	UpdateSet            ConstraintUpdateSetEnum
	DeleteSetValid       bool
	DeleteSet            ConstraintDeleteSetEnum
}

%}

%union {
    NodePt     *Node
    List       List
    Int        int
    Float      float64
    String     string
    StringList []string
    Boolean    bool
}

// ast
%type <NodePt> ast

// ddl
%type <NodePt> ddl

// createTable
%type <NodePt> createTableStmt
%type <List> attributeDeclarationList
%type <NodePt> attributeDeclaration
%token CREATE TABLE SEMICOLON

// dropTable
%type <NodePt> dropTableStmt
%token DROP

// alterTable
%type <NodePt> alterTableAddStmt
%type <NodePt> alterTableDropStmt
%token ALTER ADD

// createAssert
%type <NodePt> createAssertStmt
%token ASSERTION

// dropAssert
%type <NodePt> dropAssertStmt

// createView
%type <NodePt> createViewStmt
%token VIEW AS

// dropView
%type <NodePt> dropViewStmt

// createIndex
%type <NodePt> createIndexStmt
%token INDEX ON

// dropIndex
%type <NodePt> dropIndexStmt

// constraint
%type <List> constraintAfterAttributeList
%type <NodePt> constraintAfterAttribute
%type <NodePt> constraintAfterAttributeWithName
%type <List> constraintList
%type <NodePt> constraint
%type <NodePt> constraintWithName
%type <NodePt> setDeferrable
%type <NodePt> onUpdateSet
%type <NodePt> onDeleteSet
%type <NodePt> foreignKeyParameter
%token DEFAULT UNIQUE PRIMARYKEY CHECK FOREIGNKEY REFERENCES
%token NOT_DEFERRABLE DEFERED_DEFERRABLE IMMEDIATE_DEFERRABLE
%token UPDATE_NULL UPDATE_CASCADE
%token DELETE_NULL DELETE_CASCADE
%token DEFERRED IMMEDIATE CONSTRAINT

// condition
%type <NodePt> condition
%token AND OR
%left AND OR

// predicate
%type <NodePt> predicate
%token LIKE IN ALL ANY IS EXISTS
%type <NodePt> compareMark
%token NOTEQUAL LESS GREATER LESSEQUAL GREATEREQUAL EQUAL

// domain
%type <NodePt> domain
%token CHAR VARCHAR BIT BITVARYING BOOLEAN INT INTEGER SHORTINT 
%token FLOAT REAL DOUBLEPRECISION DECIMAL NUMERIC DATE TIME

// attriNameOptionTableName
%type <NodePt> attriNameOptionTableName
%token DOT

// subQuery
%type <NodePt> subQuery

// elementaryValue
%type <NodePt> elementaryValue
%token <Int> INTVALUE 
%token <Float> FLOATVALUE 
%token <String> STRINGVALUE 
%token <Boolean> BOOLVALUE

// public
%type <List> attriNameOptionTableNameList
%type <StringList> attriNameList
%token LPAREN RPAREN NOT NULLMARK COMMA
%token <String> ID
%%

/*  --------------------------------------------------------------------------------
    |                                      AST                                     |
    --------------------------------------------------------------------------------
    
        ast
            ddl
            dml
            dcl
            dql

    -------------------------------------------------------------------------------- */

ast
    :ddl {
        $$ = &Node{}
        $$.Type = AST_NODE

        $$.Ast = &ASTNode{}
        $$.Ast.Type = AST_DDL
        $$.Ast.Ddl = $1.Ddl

        GetInstance().AST = $$.Ast
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                     DDL                                      |
    --------------------------------------------------------------------------------

        ddl
            createTableStmt
            dropTableStmt
            alterTableAddStmt
            alterTableDropStmt
            createAssertStmt
            dropAssertStmt
            createViewStmt
            dropViewStmt
            createIndexStmt
            dropIndexStmt
            createTriggerStmt
            dropTriggerStmt
            createPsmStmt
            dropPsmStmt
    
    -------------------------------------------------------------------------------- */
ddl
    :createTableStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE
        
        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_TABLE_CREATE
        $$.Ddl.Table = $1.Table
    }
    |dropTableStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_TABLE_DROP
        $$.Ddl.Table = $1.Table
    }
    |alterTableAddStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_TABLE_ALTER_ADD
        $$.Ddl.Table = $1.Table
    }
    |alterTableDropStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_TABLE_ALTER_DROP
        $$.Ddl.Table = $1.Table
    }
    |createAssertStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_ASSERT_CREATE
        $$.Ddl.Assert = $1.Assert
    }
    |dropAssertStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_ASSERT_DROP
        $$.Ddl.Assert = $1.Assert
    }
    |createViewStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_VIEW_CREATE
        $$.Ddl.View = $1.View        
    }
    |dropViewStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_VIEW_DROP
        $$.Ddl.View = $1.View
    }
    |createIndexStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_INDEX_CREATE
        $$.Ddl.Index = $1.Index
    }
    |dropIndexStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_INDEX_DROP
        $$.Ddl.Index = $1.Index
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                 createTableStmt                              |
    --------------------------------------------------------------------------------

    createTableStmt
        CREATE TABLE ID LPAREN attributeDeclarationList RPAREN SEMICOLON
		CREATE TABLE ID LPAREN attributeDeclarationList COMMA constraintList RPAREN SEMICOLON

    attributeDeclarationList
        attributeDeclaration
        attributeDeclarationList COMMA attributeDeclaration

    attributeDeclaration
        ID domain
        ID domain constraintAfterAttributeList

    -------------------------------------------------------------------------------- */

/*  -------------------------------- createTableStmt ------------------------------- */
createTableStmt
    :CREATE TABLE ID LPAREN attributeDeclarationList RPAREN SEMICOLON {
        $$ = &Node{}
        $$.Type = TABLE_NODE
        $$.Table = &TableNode{}
        
        $$.Table.TableName = $3
        for _,v := range $5.AttributeDeclarationList {
            $$.Table.AttributeNameList = append($$.Table.AttributeNameList,v.AttributeName)
            $$.Table.DomainList = append($$.Table.DomainList,v.Domain)
            if v.ConstraintAfterAttributeListValid {
                $$.Table.ConstraintListValid = true
                $$.Table.ConstraintList = append($$.Table.ConstraintList,v.ConstraintAfterAttributeList...)
            }
        }
    }
	|CREATE TABLE ID LPAREN attributeDeclarationList COMMA constraintList RPAREN SEMICOLON {
        $$ = &Node{}
        $$.Type = TABLE_NODE
        $$.Table = &TableNode{}
        
        $$.Table.TableName = $3
        for _,v := range $5.AttributeDeclarationList {
            $$.Table.AttributeNameList = append($$.Table.AttributeNameList,v.AttributeName)
            $$.Table.DomainList = append($$.Table.DomainList,v.Domain)
            if v.ConstraintAfterAttributeListValid {
                $$.Table.ConstraintListValid = true
                $$.Table.ConstraintList = append($$.Table.ConstraintList,v.ConstraintAfterAttributeList...)
            }
        }

        $$.Table.ConstraintListValid = true
        $$.Table.ConstraintList = append($$.Table.ConstraintList,$7.ConstraintList...)
    }
    ;

/*  ----------------------------- attributeDeclarationList ------------------------- */
attributeDeclarationList
    :attributeDeclaration {
        $$ = List{}
        $$.Type = ATTRIBUTE_DECLARATION_LIST
        $$.AttributeDeclarationList = append($$.AttributeDeclarationList,$1.AttributeDeclaration)
    }
    |attributeDeclarationList COMMA attributeDeclaration {
        $$ = $1
        $$.AttributeDeclarationList = append($$.AttributeDeclarationList,$3.AttributeDeclaration)
    }
    ;

/*  ------------------------------- attributeDeclaration --------------------------- */    
attributeDeclaration
    :ID domain {
        $$ = &Node{}
        $$.Type = ATTRIBUTE_DECLARATION_NODE

        $$.AttributeDeclaration = &AttributeDeclarationNode{}
        $$.AttributeDeclaration.AttributeName = $1
        $$.AttributeDeclaration.Domain = $2.Domain
        $$.AttributeDeclaration.ConstraintAfterAttributeListValid = false
    }
    |ID domain constraintAfterAttributeList {
        $$ = &Node{}
        $$.Type = ATTRIBUTE_DECLARATION_NODE

        $$.AttributeDeclaration = &AttributeDeclarationNode{}
        $$.AttributeDeclaration.AttributeName = $1
        $$.AttributeDeclaration.Domain = $2.Domain
        $$.AttributeDeclaration.ConstraintAfterAttributeListValid = true
        $$.AttributeDeclaration.ConstraintAfterAttributeList = $3.ConstraintAfterAttributeList

        for _,v := range $$.AttributeDeclaration.ConstraintAfterAttributeList {
            switch v.Type {
                case CONSTRAINT_UNIQUE:
                    v.AttriNameList = append(v.AttriNameList,$1)
                case CONSTRAINT_PRIMARY_KEY:
                    v.AttriNameList = append(v.AttriNameList,$1)
                case CONSTRAINT_FOREIGN_KEY:
                    v.AttributeNameLocal = $1
                case CONSTRAINT_NOT_NULL:
                    v.AttributeNameLocal = $1
                case CONSTRAINT_DEFAULT:
                    v.AttributeNameLocal = $1
                default:
            }
        }
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                 dropTableStmt                                |
    --------------------------------------------------------------------------------
        
        dropTableStmt
            DROP TABLE ID SEMICOLON

    -------------------------------------------------------------------------------- */
dropTableStmt
    :DROP TABLE ID SEMICOLON {
        $$ = &Node{}
        $$.Type = TABLE_NODE

        $$.Table = &TableNode{}
        $$.Table.TableName = $3
    }
    ;

/*  --------------------------------------------------------------------------------
    |                               alterTableAddStmt                              |
    --------------------------------------------------------------------------------
		
        alterTableAddStmt
			ALTER TABLE ID ADD attributeDeclaration SEMICOLON
			ALTER TABLE ID ADD constraintWithName SEMICOLON

    -------------------------------------------------------------------------------- */
alterTableAddStmt
	:ALTER TABLE ID ADD attributeDeclaration SEMICOLON {
        $$ = &Node{}
        $$.Type = TABLE_NODE

        $$.Table = &TableNode{}
        $$.Table.TableName = $3

        $$.Table.AttributeNameList = append($$.Table.AttributeNameList,$5.AttributeDeclaration.AttributeName)
        $$.Table.DomainList = append($$.Table.DomainList,$5.AttributeDeclaration.Domain)
        if $5.AttributeDeclaration.ConstraintAfterAttributeListValid {
            $$.Table.ConstraintListValid = true
            $$.Table.ConstraintList = append($$.Table.ConstraintList,$5.AttributeDeclaration.ConstraintAfterAttributeList...)
        }

    }
	|ALTER TABLE ID ADD constraintWithName SEMICOLON {
        $$ = &Node{}
        $$.Type = TABLE_NODE

        $$.Table = &TableNode{}
        $$.Table.TableName = $3

        $$.Table.ConstraintListValid = true
        $$.Table.ConstraintList = append($$.Table.ConstraintList,$5.Constraint)
    }
    ;

/*  --------------------------------------------------------------------------------
    |                               alterTableDropStmt                             |
    --------------------------------------------------------------------------------
		
        alterTableDropStmt
			ALTER TABLE ID DROP ID SEMICOLON
			ALTER TABLE ID DROP CONSTRAINT ID SEMICOLON

    -------------------------------------------------------------------------------- */
alterTableDropStmt
    :ALTER TABLE ID DROP ID SEMICOLON {
        $$ = &Node{}
        $$.Type = TABLE_NODE

        $$.Table = &TableNode{}
        $$.Table.TableName = $3
        $$.Table.AttributeNameList = append($$.Table.AttributeNameList,$5)
    }
	|ALTER TABLE ID DROP CONSTRAINT ID SEMICOLON {
        $$ = &Node{}
        $$.Type = TABLE_NODE

        $$.Table = &TableNode{}
        $$.Table.TableName = $3
        $$.Table.ConstraintName = $6
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                createAssertStmt                              |
    --------------------------------------------------------------------------------
	    
        createAssertStmt
			CREATE ASSERTION ID CHECK LPAREN condition RPAREN SEMICOLON

    -------------------------------------------------------------------------------- */
createAssertStmt
    :CREATE ASSERTION ID CHECK LPAREN condition RPAREN SEMICOLON {
        $$ = &Node{}
        $$.Type = ASSERT_NODE

        $$.Assert = &AssertNode {}
        $$.Assert.AssertName = $3
        $$.Assert.Condition = $6.Condition
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                 dropAssertStmt                               |
    --------------------------------------------------------------------------------
	    
        dropAssertStmt
			DROP ASSERTION ID SEMICOLON

    -------------------------------------------------------------------------------- */
dropAssertStmt
    :DROP ASSERTION ID SEMICOLON {
        $$ = &Node{}
        $$.Type = ASSERT_NODE
        
        $$.Assert = &AssertNode {}
        $$.Assert.AssertName = $3
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                  createViewStmt                              |
    --------------------------------------------------------------------------------

        createViewStmt
            CREATE VIEW ID AS subQuery SEMICOLON
			CREATE VIEW ID LPAREN attriNameList RPAREN AS subQuery SEMICOLON

    -------------------------------------------------------------------------------- */
createViewStmt
    :CREATE VIEW ID AS subQuery SEMICOLON {
        $$ = &Node{}
        $$.Type = VIEW_NODE

        $$.View = &ViewNode{}
        $$.View.ViewName = $3
        $$.View.Query = $5.Query
        $$.View.AttributeNameListValid = false
    }
	|CREATE VIEW ID LPAREN attriNameList RPAREN AS subQuery SEMICOLON {
        $$ = &Node{}
        $$.Type = VIEW_NODE

        $$.View = &ViewNode{}
        $$.View.ViewName = $3
        $$.View.Query = $8.Query
        $$.View.AttributeNameListValid = true
        $$.View.AttributeNameList = $5
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                   dropViewStmt                               |
    --------------------------------------------------------------------------------
        
        dropViewStmt
            DROP VIEW ID SEMICOLON

    -------------------------------------------------------------------------------- */
dropViewStmt
    :DROP VIEW ID SEMICOLON {
        $$ = &Node{}
        $$.Type = VIEW_NODE
        
        $$.View = &ViewNode{}
        $$.View.ViewName = $3
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                 createIndexStmt                              |
    --------------------------------------------------------------------------------

        createIndexStmt
            CREATE INDEX ID ON ID LPAREN attriNameList RPAREN SEMICOLON

    -------------------------------------------------------------------------------- */
createIndexStmt
    :CREATE INDEX ID ON ID LPAREN attriNameList RPAREN SEMICOLON {
        $$ = &Node{}
        $$.Type = INDEX_NODE
        
        $$.Index = &IndexNode{}
        $$.Index.IndexName = $3
        $$.Index.TableName = $5
        $$.Index.AttributeNameList = $7
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                  dropIndexStmt                               |
    --------------------------------------------------------------------------------

        dropIndexStmt
            DROP INDEX ID SEMICOLON

    -------------------------------------------------------------------------------- */
dropIndexStmt
    :DROP INDEX ID SEMICOLON {
        $$ = &Node{}
        $$.Type = INDEX_NODE

        $$.Index = &IndexNode{}
        $$.Index.IndexName = $3
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                   constraint                                 |
    --------------------------------------------------------------------------------

        constraintAfterAttributeList
            constraintAfterAttributeWithName
            constraintAfterAttribute
            constraintAfterAttributeList constraintAfterAttributeWithName
            constraintAfterAttributeList constraintAfterAttribute

        constraintList
            constraintWithName
            constraint
            constraintList COMMA constraintWithName
            constraintList COMMA constraint

        constraintWithName
            CONSTRAINT ID constraint
        
        constraint
            UNIQUE LPAREN attriNameList RPAREN
            PRIMARYKEY LPAREN attriNameList RPAREN
            CHECK LPAREN condition RPAREN
            FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN
            FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN foreignKeyParameter

        constraintAfterAttributeWithName
            CONSTRAINT ID constraintAfterAttribute

        constraintAfterAttribute
            DEFAULT elementaryValue
            UNIQUE
            PRIMARYKEY
            NOT NULLMARK
            REFERENCES ID LPAREN ID RPAREN
            REFERENCES ID LPAREN ID RPAREN foreignKeyParameter

        foreignKeyParameter
            setDeferrable
            onUpdateSet
            onDeleteSet
            setDeferrable onUpdateSet
            onUpdateSet setDeferrable
            setDeferrable onDeleteSet
            onDeleteSet setDeferrable
            onUpdateSet onDeleteSet
            onDeleteSet onUpdateSet
            setDeferrable onUpdateSet onDeleteSet
            setDeferrable onDeleteSet onUpdateSet
            onUpdateSet setDeferrable onDeleteSet
            onUpdateSet onDeleteSet setDeferrable
            onDeleteSet setDeferrable onUpdateSet
            onDeleteSet onUpdateSet setDeferrable

        setDeferrable
            NOT_DEFERRABLE
            DEFERED_DEFERRABLE
            IMMEDIATE_DEFERRABLE

        onUpdateSet
            UPDATE_NULL
            UPDATE_CASCADE

        onDeleteSet
            DELETE_NULL
            DELETE_CASCADE

    --------------------------------------------------------------------------------

/*  ----------------------- constraintAfterAttributeList --------------------------- */
constraintAfterAttributeList
    :constraintAfterAttributeWithName {
        $$ = List{}
        $$.Type = CONSTRAINT_AFTER_ATTRIBUTE_LIST
        $$.ConstraintAfterAttributeList = append($$.ConstraintAfterAttributeList,$1.Constraint)        
    }
    |constraintAfterAttribute {
        $$ = List{}
        $$.Type = CONSTRAINT_AFTER_ATTRIBUTE_LIST
        $$.ConstraintAfterAttributeList = append($$.ConstraintAfterAttributeList,$1.Constraint)
    }
    |constraintAfterAttributeList constraintAfterAttributeWithName {
        $$ = $1
        $$.ConstraintAfterAttributeList = append($$.ConstraintAfterAttributeList,$2.Constraint)
    }
    |constraintAfterAttributeList constraintAfterAttribute {
        $$ = $1
        $$.ConstraintAfterAttributeList = append($$.ConstraintAfterAttributeList,$2.Constraint)        
    }
    ;

/*  ------------------------------ constraintList ---------------------------------- */
constraintList
    :constraintWithName {
        $$ = List{}
        $$.Type = CONSTRAINT_LIST
        $$.ConstraintList = append($$.ConstraintList,$1.Constraint)
    }
    |constraint {
        $$ = List{}
        $$.Type = CONSTRAINT_LIST
        $$.ConstraintList = append($$.ConstraintList,$1.Constraint)
    }
    |constraintList COMMA constraintWithName {
        $$ = $1
        $$.ConstraintList = append($$.ConstraintList,$3.Constraint)
    }
    |constraintList COMMA constraint {
        $$ = $1
        $$.ConstraintList = append($$.ConstraintList,$3.Constraint)
    }
    ;

/*  ------------------------------- constraintWithName ----------------------------- */
constraintWithName
    :CONSTRAINT ID constraint {
        $$ = $3
        $$.Constraint.ConstraintNameValid = true
        $$.Constraint.ConstraintName = $2
    }
    ;

/*  ---------------------------------- constraint ---------------------------------- */
constraint
    :UNIQUE LPAREN attriNameList RPAREN {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_UNIQUE
        $$.Constraint.AttriNameList = $3
    }
    |PRIMARYKEY LPAREN attriNameList RPAREN {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_PRIMARY_KEY
        $$.Constraint.AttriNameList = $3
    }
	|CHECK LPAREN condition RPAREN {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CHECK
        $$.Constraint.Condition = $3.Condition
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN foreignKeyParameter {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8

        if $10.ForeignKeyParameter.DeferrableValid {
            $$.Constraint.Deferrable = $10.ForeignKeyParameter.Deferrable
            $$.Constraint.DeferrableValid = true
        }
        if $10.ForeignKeyParameter.UpdateSetValid {
            $$.Constraint.UpdateSet = $10.ForeignKeyParameter.UpdateSet
            $$.Constraint.UpdateSetValid = true
        }
        if $10.ForeignKeyParameter.DeleteSetValid {
            $$.Constraint.DeleteSet = $10.ForeignKeyParameter.DeleteSet
            $$.Constraint.DeleteSetValid = true
        }
    }
    ;

/*  ---------------------- constraintAfterAttributeWithName ------------------------ */
constraintAfterAttributeWithName
    :CONSTRAINT ID constraintAfterAttribute {
        $$ = $3
        $$.Constraint.ConstraintNameValid = true
        $$.Constraint.ConstraintName = $2
    }
    ;

/*  -------------------------- constraintAfterAttribute ---------------------------- */
constraintAfterAttribute
    :DEFAULT elementaryValue {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_DEFAULT
        $$.Constraint.ElementaryValue = $2.ElementaryValue
    }
    |UNIQUE {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_UNIQUE
    }
    |PRIMARYKEY {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_PRIMARY_KEY
    }
    |NOT NULLMARK {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_NOT_NULL
    }
    |REFERENCES ID LPAREN ID RPAREN {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
    }
    |REFERENCES ID LPAREN ID RPAREN foreignKeyParameter {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4

        if $6.ForeignKeyParameter.DeferrableValid {
            $$.Constraint.Deferrable = $6.ForeignKeyParameter.Deferrable
            $$.Constraint.DeferrableValid = true
        }
        if $6.ForeignKeyParameter.UpdateSetValid {
            $$.Constraint.UpdateSet = $6.ForeignKeyParameter.UpdateSet
            $$.Constraint.UpdateSetValid = true
        }
        if $6.ForeignKeyParameter.DeleteSetValid {
            $$.Constraint.DeleteSet = $6.ForeignKeyParameter.DeleteSet
            $$.Constraint.DeleteSetValid = true
        }
    }
    ;

/*  ------------------------------ foreignKeyParameter ----------------------------- */
foreignKeyParameter
    :setDeferrable {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = true
	    $$.ForeignKeyParameter.Deferrable = $1.ConstraintDeferrable
	    $$.ForeignKeyParameter.UpdateSetValid = false
	    $$.ForeignKeyParameter.DeleteSetValid = false
    }
    |onUpdateSet {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = false
	    $$.ForeignKeyParameter.UpdateSetValid = true
	    $$.ForeignKeyParameter.UpdateSet = $1.ConstraintUpdateSet
	    $$.ForeignKeyParameter.DeleteSetValid = false
    }
    |onDeleteSet {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = false
	    $$.ForeignKeyParameter.UpdateSetValid = false
	    $$.ForeignKeyParameter.DeleteSetValid = true
	    $$.ForeignKeyParameter.DeleteSet = $1.ConstraintDeleteSet
    }
    |setDeferrable onUpdateSet {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = true
	    $$.ForeignKeyParameter.Deferrable = $1.ConstraintDeferrable
	    $$.ForeignKeyParameter.UpdateSetValid = true
	    $$.ForeignKeyParameter.UpdateSet = $2.ConstraintUpdateSet
	    $$.ForeignKeyParameter.DeleteSetValid = false
    }
    |onUpdateSet setDeferrable {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid  = true
	    $$.ForeignKeyParameter.Deferrable = $2.ConstraintDeferrable
	    $$.ForeignKeyParameter.UpdateSetValid = true
	    $$.ForeignKeyParameter.UpdateSet = $1.ConstraintUpdateSet
	    $$.ForeignKeyParameter.DeleteSetValid = false
    }
    |setDeferrable onDeleteSet {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = true
	    $$.ForeignKeyParameter.Deferrable = $1.ConstraintDeferrable
	    $$.ForeignKeyParameter.UpdateSetValid = false
	    $$.ForeignKeyParameter.DeleteSetValid = true
	    $$.ForeignKeyParameter.DeleteSet = $2.ConstraintDeleteSet
    }
    |onDeleteSet setDeferrable {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = true
	    $$.ForeignKeyParameter.Deferrable = $2.ConstraintDeferrable
	    $$.ForeignKeyParameter.UpdateSetValid = false
	    $$.ForeignKeyParameter.DeleteSetValid = true
	    $$.ForeignKeyParameter.DeleteSet = $1.ConstraintDeleteSet
    }
    |onUpdateSet onDeleteSet {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = false
	    $$.ForeignKeyParameter.UpdateSetValid = true
	    $$.ForeignKeyParameter.UpdateSet = $1.ConstraintUpdateSet
	    $$.ForeignKeyParameter.DeleteSetValid = true
	    $$.ForeignKeyParameter.DeleteSet = $2.ConstraintDeleteSet
    }
    |onDeleteSet onUpdateSet {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = false
	    $$.ForeignKeyParameter.UpdateSetValid = true
	    $$.ForeignKeyParameter.UpdateSet = $2.ConstraintUpdateSet
	    $$.ForeignKeyParameter.DeleteSetValid = true
	    $$.ForeignKeyParameter.DeleteSet = $1.ConstraintDeleteSet
    }
    |setDeferrable onUpdateSet onDeleteSet {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = true
	    $$.ForeignKeyParameter.Deferrable = $1.ConstraintDeferrable
	    $$.ForeignKeyParameter.UpdateSetValid = true
	    $$.ForeignKeyParameter.UpdateSet = $2.ConstraintUpdateSet
	    $$.ForeignKeyParameter.DeleteSetValid = true
	    $$.ForeignKeyParameter.DeleteSet = $3.ConstraintDeleteSet
    }
    |setDeferrable onDeleteSet onUpdateSet {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = true
	    $$.ForeignKeyParameter.Deferrable = $1.ConstraintDeferrable
	    $$.ForeignKeyParameter.UpdateSetValid = true
	    $$.ForeignKeyParameter.UpdateSet = $3.ConstraintUpdateSet
	    $$.ForeignKeyParameter.DeleteSetValid = true
	    $$.ForeignKeyParameter.DeleteSet = $2.ConstraintDeleteSet
    }
    |onUpdateSet setDeferrable onDeleteSet {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = true
	    $$.ForeignKeyParameter.Deferrable = $2.ConstraintDeferrable
	    $$.ForeignKeyParameter.UpdateSetValid = true
	    $$.ForeignKeyParameter.UpdateSet = $1.ConstraintUpdateSet
	    $$.ForeignKeyParameter.DeleteSetValid = true
	    $$.ForeignKeyParameter.DeleteSet = $3.ConstraintDeleteSet
    }
    |onUpdateSet onDeleteSet setDeferrable {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = true
	    $$.ForeignKeyParameter.Deferrable = $3.ConstraintDeferrable
	    $$.ForeignKeyParameter.UpdateSetValid = true
	    $$.ForeignKeyParameter.UpdateSet = $1.ConstraintUpdateSet
	    $$.ForeignKeyParameter.DeleteSetValid = true
	    $$.ForeignKeyParameter.DeleteSet = $2.ConstraintDeleteSet
    }
    |onDeleteSet setDeferrable onUpdateSet {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = true
	    $$.ForeignKeyParameter.Deferrable = $2.ConstraintDeferrable
	    $$.ForeignKeyParameter.UpdateSetValid = true
	    $$.ForeignKeyParameter.UpdateSet = $3.ConstraintUpdateSet
	    $$.ForeignKeyParameter.DeleteSetValid = true
	    $$.ForeignKeyParameter.DeleteSet = $1.ConstraintDeleteSet
    }
    |onDeleteSet onUpdateSet setDeferrable {
        $$ = &Node{}
        $$.Type = FOREIGNKEY_PARAMETER_NODE
        $$.ForeignKeyParameter = &ForeignKeyParameterNode{}

        $$.ForeignKeyParameter.DeferrableValid = true
	    $$.ForeignKeyParameter.Deferrable = $3.ConstraintDeferrable
	    $$.ForeignKeyParameter.UpdateSetValid = true
	    $$.ForeignKeyParameter.UpdateSet = $2.ConstraintUpdateSet
	    $$.ForeignKeyParameter.DeleteSetValid = true
	    $$.ForeignKeyParameter.DeleteSet = $1.ConstraintDeleteSet
    }
    ;

/*  ------------------------------- setDeferrable ---------------------------------- */
    setDeferrable
	    :NOT_DEFERRABLE {
            $$ = &Node{}
            $$.Type = CONSTRAINT_DEFERRABLE_ENUM
            $$.ConstraintDeferrable = CONSTRAINT_NOT_DEFERRABLE
        }
	    |DEFERED_DEFERRABLE {
            $$ = &Node{}
            $$.Type = CONSTRAINT_DEFERRABLE_ENUM
            $$.ConstraintDeferrable = CONSTRAINT_INITIALLY_DEFERRED
        }
	    |IMMEDIATE_DEFERRABLE {
            $$ = &Node{}
            $$.Type = CONSTRAINT_DEFERRABLE_ENUM
            $$.ConstraintDeferrable = CONSTRAINT_INITIALLY_IMMEDIATE
        }
        ;

/*  ------------------------------- onUpdateSet ------------------------------------ */
    onUpdateSet
	    :UPDATE_NULL {
            $$ = &Node{}
            $$.Type = CONSTRAINT_UPDATE_SET_ENUM
            $$.ConstraintUpdateSet = CONSTRAINT_UPDATE_SET_NULL
        }
	    |UPDATE_CASCADE {
            $$ = &Node{}
            $$.Type = CONSTRAINT_UPDATE_SET_ENUM
            $$.ConstraintUpdateSet = CONSTRAINT_UPDATE_SET_CASCADE
        }
        ;

/*  ------------------------------- onDeleteSet ------------------------------------ */
    onDeleteSet
	    :DELETE_NULL {
            $$ = &Node{}
            $$.Type = CONSTRAINT_DELETE_SET_ENUM
            $$.ConstraintDeleteSet = CONSTRAINT_DELETE_SET_NULL
        }
	    |DELETE_CASCADE {
            $$ = &Node{}
            $$.Type = CONSTRAINT_DELETE_SET_ENUM
            $$.ConstraintDeleteSet = CONSTRAINT_DELETE_SET_CASCADE
        }
        ;

/*  --------------------------------------------------------------------------------
    |                                   condition                                  |
    --------------------------------------------------------------------------------

        condition
			predicate
			LPAREN condition RPAREN
			condition AND condition
			condition OR condition

    -------------------------------------------------------------------------------- */
condition
    :predicate {
        $$ = &Node{}
        $$.Type = CONDITION_NODE

        $$.Condition = &ConditionNode{}
        $$.Condition.Type = CONDITION_PREDICATE
        $$.Condition.Predicate = $1.Predicate
    }
	|LPAREN condition RPAREN {
        $$ = $2
    }
	|condition AND condition {
        $$ = &Node{}
        $$.Type = CONDITION_NODE

        $$.Condition = &ConditionNode{}
        $$.Condition.Type = CONDITION_AND
        $$.Condition.ConditionL = $1.Condition
        $$.Condition.ConditionR = $3.Condition
    }
	|condition OR condition {
        $$ = &Node{}
        $$.Type = CONDITION_NODE

        $$.Condition = &ConditionNode{}
        $$.Condition.Type = CONDITION_OR
        $$.Condition.ConditionL = $1.Condition
        $$.Condition.ConditionR = $3.Condition
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                   predicate                                  |
    --------------------------------------------------------------------------------

        predicate
			attriNameOptionTableName compareMark elementaryValue
			attriNameOptionTableName LIKE STRINGVALUE
			attriNameOptionTableName IN subQuery
			attriNameOptionTableName NOT IN subQuery
			attriNameOptionTableName IN ID
			attriNameOptionTableName NOT IN ID
			attriNameOptionTableName compareMark ALL subQuery
			NOT attriNameOptionTableName compareMark ALL subQuery
			attriNameOptionTableName compareMark ANY subQuery
			NOT attriNameOptionTableName compareMark ANY subQuery
			attriNameOptionTableName compareMark ALL ID
			NOT attriNameOptionTableName compareMark ALL ID
			attriNameOptionTableName compareMark ANY ID
			NOT attriNameOptionTableName compareMark ANY ID
			attriNameOptionTableName IS NULLMARK
			attriNameOptionTableName IS NOT NULLMARK
			LPAREN attriNameOptionTableNameList RPAREN IN subQuery
			LPAREN attriNameOptionTableNameList RPAREN NOT IN subQuery
			LPAREN attriNameOptionTableNameList RPAREN IN ID
			LPAREN attriNameOptionTableNameList RPAREN NOT IN ID
			EXISTS subQuery
			NOT EXISTS subQuery

        compareMark
			EQUAL
    		NOTEQUAL
    		LESS
    		GREATER
    		LESSEQUAL
    		GREATEREQUAL

    -------------------------------------------------------------------------------- */

/*  ----------------------------------- predicate ---------------------------------- */
predicate
	:attriNameOptionTableName compareMark elementaryValue {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_COMPARE_ELEMENTARY_VALUE

        $$.Predicate.AttriNameWithTableNameL = $1.AttriNameOptionTableName
        $$.Predicate.CompareMark = $2.CompareMark
        $$.Predicate.ElementaryValue = $3.ElementaryValue
    }
	|attriNameOptionTableName LIKE STRINGVALUE {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_LIKE_STRING_VALUE

        $$.Predicate.AttriNameWithTableNameL = $1.AttriNameOptionTableName
        $$.Predicate.ElementaryValue = &ElementaryValueNode{}
        $$.Predicate.ElementaryValue.Type = ELEMENTARY_VALUE_STRING
        $$.Predicate.ElementaryValue.StringValue = $3
    }
	|attriNameOptionTableName IN subQuery {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_IN_SUBQUERY

        $$.Predicate.AttriNameWithTableNameL = $1.AttriNameOptionTableName
        $$.Predicate.Subquery = $3.Subquery
    }
	|attriNameOptionTableName NOT IN subQuery {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_NOT_IN_SUBQUERY

        $$.Predicate.AttriNameWithTableNameL = $1.AttriNameOptionTableName
        $$.Predicate.Subquery = $4.Subquery
    }
	|attriNameOptionTableName IN ID {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_IN_TABLE

        $$.Predicate.AttriNameWithTableNameL = $1.AttriNameOptionTableName
        $$.Predicate.TableName = $3
    }
	|attriNameOptionTableName NOT IN ID {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_NOT_IN_TABLE

        $$.Predicate.AttriNameWithTableNameL = $1.AttriNameOptionTableName
        $$.Predicate.TableName = $4
    }
	|attriNameOptionTableName compareMark ALL subQuery {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_COMPARE_ALL_SUBQUERY

        $$.Predicate.AttriNameWithTableNameL = $1.AttriNameOptionTableName
        $$.Predicate.CompareMark = $2.CompareMark
        $$.Predicate.Subquery = $4.Subquery
    }
	|NOT attriNameOptionTableName compareMark ALL subQuery {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_COMPARE_NOT_ALL_SUBQUERY

        $$.Predicate.AttriNameWithTableNameL = $2.AttriNameOptionTableName
        $$.Predicate.CompareMark = $3.CompareMark
        $$.Predicate.Subquery = $5.Subquery
    }
	|attriNameOptionTableName compareMark ANY subQuery {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_COMPARE_ANY_SUBQUERY

        $$.Predicate.AttriNameWithTableNameL = $1.AttriNameOptionTableName
        $$.Predicate.CompareMark = $2.CompareMark
        $$.Predicate.Subquery = $4.Subquery
    }
	|NOT attriNameOptionTableName compareMark ANY subQuery {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_COMPARE_NOT_ANY_SUBQUERY

        $$.Predicate.AttriNameWithTableNameL = $2.AttriNameOptionTableName
        $$.Predicate.CompareMark = $3.CompareMark
        $$.Predicate.Subquery = $5.Subquery
    }
	|attriNameOptionTableName compareMark ALL ID {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_COMPARE_ALL_TABLE

        $$.Predicate.AttriNameWithTableNameL = $1.AttriNameOptionTableName
        $$.Predicate.CompareMark = $2.CompareMark
        $$.Predicate.TableName = $4
    }
	|NOT attriNameOptionTableName compareMark ALL ID {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_COMPARE_NOT_ALL_TABLE

        $$.Predicate.AttriNameWithTableNameL = $2.AttriNameOptionTableName
        $$.Predicate.CompareMark = $3.CompareMark
        $$.Predicate.TableName = $5
    }
	|attriNameOptionTableName compareMark ANY ID {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_COMPARE_ANY_TABLE

        $$.Predicate.AttriNameWithTableNameL = $1.AttriNameOptionTableName
        $$.Predicate.CompareMark = $2.CompareMark
        $$.Predicate.TableName = $4
    }
	|NOT attriNameOptionTableName compareMark ANY ID {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_COMPARE_NOT_ANY_TABLE

        $$.Predicate.AttriNameWithTableNameL = $2.AttriNameOptionTableName
        $$.Predicate.CompareMark = $3.CompareMark
        $$.Predicate.TableName = $5
    }
	|attriNameOptionTableName IS NULLMARK {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_IS_NULL

        $$.Predicate.AttriNameWithTableNameL = $1.AttriNameOptionTableName
    }
	|attriNameOptionTableName IS NOT NULLMARK {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_IS_NOT_NULL

        $$.Predicate.AttriNameWithTableNameL = $1.AttriNameOptionTableName
    }
	|LPAREN attriNameOptionTableNameList RPAREN IN subQuery {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_TUPLE_IN_SUBQUERY

        $$.Predicate.AttriNameOptionTableNameList = $2.AttriNameOptionTableNameList
        $$.Predicate.Subquery = $5.Subquery
    }
	|LPAREN attriNameOptionTableNameList RPAREN NOT IN subQuery {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_TUPLE_NOT_IN_SUBQUERY

        $$.Predicate.AttriNameOptionTableNameList = $2.AttriNameOptionTableNameList
        $$.Predicate.Subquery = $6.Subquery
    }
	|LPAREN attriNameOptionTableNameList RPAREN IN ID {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_TUPLE_IN_TABLE

        $$.Predicate.AttriNameOptionTableNameList = $2.AttriNameOptionTableNameList
        $$.Predicate.TableName = $5
    }
	|LPAREN attriNameOptionTableNameList RPAREN NOT IN ID {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_TUPLE_NOT_IN_TABLE

        $$.Predicate.AttriNameOptionTableNameList = $2.AttriNameOptionTableNameList
        $$.Predicate.TableName = $6
    }
	|EXISTS subQuery {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_SUBQUERY_EXISTS

        $$.Predicate.Subquery = $2.Subquery
    }
	|NOT EXISTS subQuery {
        $$ = &Node{}
        $$.Type = PREDICATE_NODE
        
        $$.Predicate = &PredicateNode{}
        $$.Predicate.Type = PREDICATE_SUBQUERY_NOT_EXISTS

        $$.Predicate.Subquery = $3.Subquery
    }
    ;

/*  --------------------------------- compareMark ---------------------------------- */
compareMark
    :EQUAL {
        $$ = &Node{}
        $$.Type = COMPAREMARK_ENUM
        $$.CompareMark = COMPAREMARK_EQUAL
    }
    |NOTEQUAL {
        $$ = &Node{}
        $$.Type = COMPAREMARK_ENUM
        $$.CompareMark = COMPAREMARK_NOTEQUAL
    }
    |LESS {
        $$ = &Node{}
        $$.Type = COMPAREMARK_ENUM
        $$.CompareMark = COMPAREMARK_LESS
    }
    |GREATER {
        $$ = &Node{}
        $$.Type = COMPAREMARK_ENUM
        $$.CompareMark = COMPAREMARK_GREATER
    }
    |LESSEQUAL {
        $$ = &Node{}
        $$.Type = COMPAREMARK_ENUM
        $$.CompareMark = COMPAREMARK_LESSEQUAL
    }
    |GREATEREQUAL {
        $$ = &Node{}
        $$.Type = COMPAREMARK_ENUM
        $$.CompareMark = COMPAREMARK_GREATEREQUAL
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                     domain                                   |
    --------------------------------------------------------------------------------

        domain
            CHAR
            VARCHAR LPAREN INTVALUE RPAREN
            BIT LPAREN INTVALUE RPAREN
            BITVARYING LPAREN INTVALUE RPAREN
            BOOLEAN
            INT
            INTEGER
            SHORTINT
            FLOAT
            REAL
            DOUBLEPRECISION
            DECIMAL LPAREN INTVALUE COMMA INTVALUE RPAREN
            NUMERIC LPAREN INTVALUE COMMA INTVALUE RPAREN
            DATE
            TIME

    -------------------------------------------------------------------------------- */
domain
    :CHAR {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_CHAR
    }
    |VARCHAR LPAREN INTVALUE RPAREN {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_VARCHAR
        $$.Domain.N = $3
    }
    |BIT LPAREN INTVALUE RPAREN {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_BIT
        $$.Domain.N = $3
    }
    |BITVARYING LPAREN INTVALUE RPAREN {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_BITVARYING
        $$.Domain.N = $3
    }
    |BOOLEAN {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_BOOLEAN
    }
    |INT {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_INT
    }
    |INTEGER {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_INTEGER
    }
    |SHORTINT {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_SHORTINT
    }
    |FLOAT {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_FLOAT
    }
    |REAL {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_REAL
    }
    |DOUBLEPRECISION {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_DOUBLEPRECISION
    }
    |DECIMAL LPAREN INTVALUE COMMA INTVALUE RPAREN {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_DECIMAL
        $$.Domain.N = $3
        $$.Domain.D = $5
    }
    |NUMERIC LPAREN INTVALUE COMMA INTVALUE RPAREN {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_NUMERIC
        $$.Domain.N = $3
        $$.Domain.D = $5
    }
    DATE {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_DATE
    }
    |TIME {
        $$ = &Node{}
        $$.Type = DOMAIN_NODE

        $$.Domain = &DomainNode{}
        $$.Domain.Type = DOMAIN_TIME
    }
    ;

/*  --------------------------------------------------------------------------------
    |                            attriNameOptionTableName                          |
    --------------------------------------------------------------------------------

        attriNameOptionTableName
            ID
            ID DOT ID

   -------------------------------------------------------------------------------- */
attriNameOptionTableName
    :ID {
        $$ = &Node{}
        $$.Type = ATTRINAME_OPTION_TABLENAME_NODE
        $$.AttriNameOptionTableName = &AttriNameOptionTableNameNode{}
        $$.AttriNameOptionTableName.TableNameValid = false
        $$.AttriNameOptionTableName.AttributeName = $1
    }
    |ID DOT ID {
        $$ = &Node{}
        $$.Type = ATTRINAME_OPTION_TABLENAME_NODE
        $$.AttriNameOptionTableName = &AttriNameOptionTableNameNode{}
        $$.AttriNameOptionTableName.TableNameValid = true
        $$.AttriNameOptionTableName.TableName = $1
        $$.AttriNameOptionTableName.AttributeName = $3
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                    subQuery                                  |
    --------------------------------------------------------------------------------

        subQuery

   -------------------------------------------------------------------------------- */
// TODO
subQuery
    :DOT {
        $$ = &Node{}
        $$.Type = SUBQUERY_NODE
        $$.Subquery = &QueryNode{}
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                elementaryValue                               |
    --------------------------------------------------------------------------------

    elementaryValue
        INTVALUE
        FLOATVALUE
        STRINGVALUE
        BOOLVALUE

   -------------------------------------------------------------------------------- */

elementaryValue
    :INTVALUE {
        $$ = &Node{}
        $$.Type = ELEMENTARY_VALUE_NODE
        $$.ElementaryValue = &ElementaryValueNode{}
        $$.ElementaryValue.Type = ELEMENTARY_VALUE_INT
        $$.ElementaryValue.IntValue = $1
    }
    |FLOATVALUE {
        $$ = &Node{}
        $$.Type = ELEMENTARY_VALUE_NODE
        $$.ElementaryValue = &ElementaryValueNode{}
        $$.ElementaryValue.Type = ELEMENTARY_VALUE_FLOAT
        $$.ElementaryValue.FloatValue = $1
    }
    |STRINGVALUE {
        $$ = &Node{}
        $$.Type = ELEMENTARY_VALUE_NODE
        $$.ElementaryValue = &ElementaryValueNode{}
        $$.ElementaryValue.Type = ELEMENTARY_VALUE_STRING
        $$.ElementaryValue.StringValue = $1
    }
    |BOOLVALUE {
        $$ = &Node{}
        $$.Type = ELEMENTARY_VALUE_NODE
        $$.ElementaryValue = &ElementaryValueNode{}
        $$.ElementaryValue.Type = ELEMENTARY_VALUE_BOOLEAN
        $$.ElementaryValue.BooleanValue = $1
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                     public                                   |
    --------------------------------------------------------------------------------
    
    attriNameList
        attriNameList COMMA ID
        ID
    
    attriNameOptionTableNameList
        attriNameOptionTableNameList COMMA attriNameOptionTableName
        attriNameOptionTableName
    
    -------------------------------------------------------------------------------- */

/*  -------------------------------- attriNameList --------------------------------- */
    attriNameList
        :attriNameList COMMA ID {
            $$ = append($1,$3)
        }
        |ID {
            $$ = append($$,$1)
        }
        ;

/*  ------------------------- attriNameOptionTableNameList ------------------------- */
    attriNameOptionTableNameList
        :attriNameOptionTableNameList COMMA attriNameOptionTableName {
            $$ = $1
            $$.AttriNameOptionTableNameList = append($$.AttriNameOptionTableNameList,$3.AttriNameOptionTableName)
        }
        |attriNameOptionTableName {
            $$ = List{}
            $$.Type = ATTRINAME_OPTION_TABLENAME_LIST
            $$.AttriNameOptionTableNameList = append($$.AttriNameOptionTableNameList,$1.AttriNameOptionTableName)
        }
        ;

%%