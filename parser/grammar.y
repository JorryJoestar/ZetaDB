%{
package parser

import (
    "fmt"
)

// Node
type NodeEnum uint8

const (
    AST_NODE                        NodeEnum = 1

    DDL_NODE                        NodeEnum = 2
    TABLE_NODE                      NodeEnum = 3
    ASSERT_NODE                     NodeEnum = 4
    VIEW_NODE                       NodeEnum = 5
    INDEX_NODE                      NodeEnum = 6
    TRIGGER_NODE                    NodeEnum = 7
    PSM_NODE                        NodeEnum = 8

    DQL_NODE                        NodeEnum = 9
    QUERY_NODE                      NodeEnum = 10
    SELECT_LIST_ENTRY               NodeEnum = 11
    FROM_LIST_ENTRY                 NodeEnum = 12
    ON_LIST_ENTRY                   NodeEnum = 13
    ORDERBY_LIST_ENTRY              NodeEnum = 14

    DCL_NODE                        NodeEnum = 15

    DML_NODE                        NodeEnum = 16
    UPDATE_NODE                     NodeEnum = 17
    UPDATE_LIST_ENTRY               NodeEnum = 18
    INSERT_NODE                     NodeEnum = 19
    DELETE_NODE                     NodeEnum = 20

    DOMAIN_NODE                     NodeEnum = 21
    ATTRINAME_OPTION_TABLENAME_NODE NodeEnum = 22
    CONSTRAINT_NODE                 NodeEnum = 23
    ELEMENTARY_VALUE_NODE           NodeEnum = 24
    CONDITION_NODE                  NodeEnum = 25
    PREDICATE_NODE                  NodeEnum = 26
    EXPRESSION_NODE                 NodeEnum = 27
    EXPRESSION_ENTRY                NodeEnum = 28
    AGGREGATION_NODE                NodeEnum = 29
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

/* public */
    Domain                   *DomainNode
    AttriNameOptionTableName *AttriNameOptionTableNameNode
    Constraint               *ConstraintNode
    ElementaryValue          *ElementaryValueNode
    Condition                *ConditionNode
    Predicate                *PredicateNode
    Expression               *ExpressionNode
    ExpressionEntry          *ExpressionEntryNode
    Aggregation              *AggregationNode
}

// List
type ListEnum uint8

const (
    CONSTRAINT_AFTER_ATTRIBUTE_LIST ListEnum = 1
    CONSTRAINT_LIST                 ListEnum = 2
)

type List struct {
    Type                         ListEnum
    ConstraintAfterAttributeList []*ConstraintNode
    ConstraintList               []*ConstraintNode
}

%}

%union {
    NodePt  *Node
    Int     int
    Float   float64
    String  string
    Boolean bool
}

%type <NodePt> ast

// %type <NodePt> ddl
// %type <NodePt> createTableStmt

%type <NodePt> constraintAfterAttribute

%type <NodePt> elementaryValue

%token UNIQUE PRIMARYKEY CHECK FOREIGNKEY REFERENCES
%token NOT_DEFERRABLE DEFERED_DEFERRABLE IMMEDIATE_DEFERRABLE
%token UPDATE_NULL UPDATE_CASCADE
%token DELETE_NULL DELETE_CASCADE
%token DEFERRED IMMEDIATE CONSTRAINT

%token <Int> INTVALUE 
%token <Float> FLOATVALUE 
%token <String> STRINGVALUE 
%token <Boolean> BOOLVALUE

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
    :elementaryValue {
        fmt.Println("k")
        fmt.Println($1.ElementaryValue.IntValue)


        GetInstance().AST = &ASTNode{
            Type: AST_DQL,
	        Ddl: nil,
	        Dml: nil,
	        Dcl: nil,
	        Dql: nil}
    }

/*  --------------------------------------------------------------------------------
    |                                     DDL                                      |
    --------------------------------------------------------------------------------

    ------------------------------------- DDL --------------------------------------
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
    
    -------------------------------- createTableStmt -------------------------------

    createTableStmt
        CREATE TABLE ID LPAREN attributeDeclarationList RPAREN SEMICOLON
		CREATE TABLE ID LPAREN attributeDeclarationList COMMA constraintPhraseList LPAREN SEMICOLON

    attributeDeclarationList
        attributeDeclaration
        attributeDeclarationList

    attributeDeclaration
        ID domain
        ID domain constraintAfterAttributeList

    --------------------------------------------------------------------------------




    -------------------------------------------------------------------------------- */

/*  --------------------------------------------------------------------------------
    |                                   constraint                                 |
    --------------------------------------------------------------------------------

/*  ----------------------- constraintAfterAttributeList ---------------------------

    constraintAfterAttributeList
        constraintAfterAttribute
        constraintAfterAttributeList constraintAfterAttribute

    -------------------------------------------------------------------------------- */
constraintAfterAttributeList
    :constraintAfterAttribute {
        $$ = &List{}
        $$.Type = CONSTRAINT_AFTER_ATTRIBUTE_LIST
        $$.ConstraintAfterAttributeList = append($$.ConstraintAfterAttributeList,$1.Constraint)
    }
    |constraintAfterAttributeList constraintAfterAttribute {
        $$ = $1
        $$.ConstraintAfterAttributeList = append($$.ConstraintAfterAttributeList,$1.Constraint)        
    }
    ;

/*  ------------------------------ constraintList ----------------------------------

    constraintList
        constraint
        constraintList constraint

    -------------------------------------------------------------------------------- */
constraintList
    :constraint {
        $$ = List{}
        $$.Type = CONSTRAINT_LIST
        $$.ConstraintList = append($$.ConstraintList,$1.Constraint)
    }
    |constraintList constraint {
        $$ = $1
        $$.ConstraintList = append($$.ConstraintList,$1.Constraint)
    }
    ;

/*  ------------------------------- constraintWithName -----------------------------
    
    constraintWithName
        CONSTRAINT ID constraint

    -------------------------------------------------------------------------------- */
constraintWithName
    :CONSTRAINT ID constraint {
        $$ = $3
        $$.Constraint.ConstraintNameValid = true
        $$.Constraint.ConstraintName = $2
    }
    ;

/*  ---------------------------------- constraint ----------------------------------
    
    constraint
        UNIQUE LPAREN attriNameList RPAREN
		PRIMARYKEY LPAREN attriNameList RPAREN
		CHECK LPAREN condition RPAREN

		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN

		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN setDeferrable
		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onUpdateSet
		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onDeleteSet

		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN setDeferrable onUpdateSet
		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN setDeferrable onDeleteSet
		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onUpdateSet setDeferrable
		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onUpdateSet onDeleteSet
		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onDeleteSet setDeferrable
		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onDeleteSet onUpdateSet

		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN setDeferrable onUpdateSet onDeleteSet
		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN setDeferrable onDeleteSet onUpdateSet
		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onUpdateSet setDeferrable onDeleteSet
		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onUpdateSet onDeleteSet setDeferrable
		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onDeleteSet setDeferrable onUpdateSet
		FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onDeleteSet onUpdateSet setDeferrable

    -------------------------------------------------------------------------------- */
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
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN setDeferrable {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.Deferrable = $10
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onUpdateSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.UpdateSet = $10
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onDeleteSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.DeleteSet = $10
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN setDeferrable onUpdateSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.Deferrable = $10
        $$.Constraint.UpdateSet = $11
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN setDeferrable onDeleteSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.Deferrable = $10
        $$.Constraint.DeleteSet = $11
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onUpdateSet setDeferrable {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.Deferrable = $11
        $$.Constraint.UpdateSet = $10
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onUpdateSet onDeleteSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.UpdateSet = $10
        $$.Constraint.DeleteSet = $11
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onDeleteSet setDeferrable {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.Deferrable = $11
        $$.Constraint.DeleteSet = $10
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onDeleteSet onUpdateSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.UpdateSet = $11
        $$.Constraint.DeleteSet = $10
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN setDeferrable onUpdateSet onDeleteSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.Deferrable = $10
        $$.Constraint.UpdateSet = $11
        $$.Constraint.DeleteSet = $12
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN setDeferrable onDeleteSet onUpdateSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.Deferrable = $10
        $$.Constraint.UpdateSet = $12
        $$.Constraint.DeleteSet = $11
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onUpdateSet setDeferrable onDeleteSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.Deferrable = $11
        $$.Constraint.UpdateSet = $10
        $$.Constraint.DeleteSet = $12
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onUpdateSet onDeleteSet setDeferrable {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.Deferrable = $12
        $$.Constraint.UpdateSet = $10
        $$.Constraint.DeleteSet = $11
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onDeleteSet setDeferrable onUpdateSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.Deferrable = $11
        $$.Constraint.UpdateSet = $12
        $$.Constraint.DeleteSet = $10
    }
	|FOREIGNKEY LPAREN ID RPAREN REFERENCES ID LPAREN ID RPAREN onDeleteSet onUpdateSet setDeferrable {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_FOREIGN_KEY
        $$.Constraint.AttributeNameLocal = $3
        $$.Constraint.ForeignTableName = $6
        $$.Constraint.AttributeNameForeign = $8
        $$.Constraint.Deferrable = $12
        $$.Constraint.UpdateSet = $11
        $$.Constraint.DeleteSet = $10
    }
    ;

/*  ---------------------- constraintAfterAttributeWithName ------------------------
    
    constraintAfterAttributeWithName
        CONSTRAINT ID constraintAfterAttribute

    -------------------------------------------------------------------------------- */
constraintAfterAttributeWithName
    :CONSTRAINT ID constraintAfterAttribute {
        $$ = $3
        $$.Constraint.ConstraintNameValid = true
        $$.Constraint.ConstraintName = $2
    }
    ;

/*  -------------------------- constraintAfterAttribute ----------------------------
    
    constraintAfterAttribute
        DEFAULT elementaryValue

		UNIQUE
		PRIMARYKEY
		NOT NULLMARK

		CHECK LPAREN condition RPAREN

		REFERENCES ID LPAREN ID RPAREN

		REFERENCES ID LPAREN ID RPAREN setDeferrable
		REFERENCES ID LPAREN ID RPAREN onUpdateSet
		REFERENCES ID LPAREN ID RPAREN onDeleteSet

		REFERENCES ID LPAREN ID RPAREN setDeferrable onUpdateSet
		REFERENCES ID LPAREN ID RPAREN setDeferrable onDeleteSet
		REFERENCES ID LPAREN ID RPAREN onUpdateSet setDeferrable
		REFERENCES ID LPAREN ID RPAREN onUpdateSet onDeleteSet
		REFERENCES ID LPAREN ID RPAREN onDeleteSet setDeferrable
		REFERENCES ID LPAREN ID RPAREN onDeleteSet onUpdateSet

		REFERENCES ID LPAREN ID RPAREN setDeferrable onUpdateSet onDeleteSet
		REFERENCES ID LPAREN ID RPAREN setDeferrable onDeleteSet onUpdateSet
		REFERENCES ID LPAREN ID RPAREN onUpdateSet setDeferrable onDeleteSet
		REFERENCES ID LPAREN ID RPAREN onUpdateSet onDeleteSet setDeferrable
		REFERENCES ID LPAREN ID RPAREN onDeleteSet setDeferrable onUpdateSet
		REFERENCES ID LPAREN ID RPAREN onDeleteSet onUpdateSet setDeferrable

    -------------------------------------------------------------------------------- */
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
    |CHECK LPAREN condition RPAREN {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CHECK
        $$.Constraint.Condition = $3.Condition
    }
    |REFERENCES ID LPAREN ID RPAREN {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
    }
    |REFERENCES ID LPAREN ID RPAREN setDeferrable {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.Deferrable = $6
    }
    |REFERENCES ID LPAREN ID RPAREN onUpdateSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.UpdateSet = $6
    }
    |REFERENCES ID LPAREN ID RPAREN onDeleteSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.DeleteSet = $6
    }
    |REFERENCES ID LPAREN ID RPAREN setDeferrable onUpdateSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.Deferrable = $6
        $$.Constraint.UpdateSet = $7
    }
    |REFERENCES ID LPAREN ID RPAREN setDeferrable onDeleteSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.Deferrable = $6
        $$.Constraint.DeleteSet = $7
    }
	|REFERENCES ID LPAREN ID RPAREN onUpdateSet setDeferrable {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.UpdateSet = $6
        $$.Constraint.Deferrable = $7
    }
	|REFERENCES ID LPAREN ID RPAREN onUpdateSet onDeleteSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.UpdateSet = $6
        $$.Constraint.DeleteSet = $7
    }
	|REFERENCES ID LPAREN ID RPAREN onDeleteSet setDeferrable {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.DeleteSet = $6
        $$.Constraint.Deferrable = $7
    }
	|REFERENCES ID LPAREN ID RPAREN onDeleteSet onUpdateSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.DeleteSet = $6
        $$.Constraint.UpdateSet = $7
    }
	|REFERENCES ID LPAREN ID RPAREN setDeferrable onUpdateSet onDeleteSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.Deferrable = $6
        $$.Constraint.UpdateSet = $7
        $$.Constraint.DeleteSet = $8
    }
	|REFERENCES ID LPAREN ID RPAREN setDeferrable onDeleteSet onUpdateSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.Deferrable = $6
        $$.Constraint.DeleteSet = $7
        $$.Constraint.UpdateSet = $8
    }
	|REFERENCES ID LPAREN ID RPAREN onUpdateSet setDeferrable onDeleteSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.UpdateSet = $6
        $$.Constraint.Deferrable = $7
        $$.Constraint.DeleteSet = $8
    }
	|REFERENCES ID LPAREN ID RPAREN onUpdateSet onDeleteSet setDeferrable {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.UpdateSet = $6
        $$.Constraint.DeleteSet = $7
        $$.Constraint.Deferrable = $8
    }
	|REFERENCES ID LPAREN ID RPAREN onDeleteSet setDeferrable onUpdateSet {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.DeleteSet = $6
        $$.Constraint.Deferrable = $7
        $$.Constraint.UpdateSet = $8
    }
	|REFERENCES ID LPAREN ID RPAREN onDeleteSet onUpdateSet setDeferrable {
        $$ = &Node{}
        $$.Type = CONSTRAINT_NODE
        $$.Constraint = &ConstraintNode{}
        $$.Constraint.ConstraintNameValid = false

        $$.Constraint.Type = CONSTRAINT_CONSTRAINT_FOREIGN_KEY
        $$.Constraint.ForeignTableName = $2
        $$.Constraint.AttributeNameForeign = $4
        $$.Constraint.DeleteSet = $6
        $$.Constraint.UpdateSet = $7
        $$.Constraint.Deferrable = $8
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                    public                                    |
    --------------------------------------------------------------------------------

/* --------------------------------- elementaryValue -----------------------------
    
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

%%