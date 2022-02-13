%{
package parser

import (
    "fmt"
)

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

%}

%union {
    NodePt *Node
}

%type <NodePt> ast

%type <NodePt> ddl

%type <NodePt> createTableStmt

%token LPAREN

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
    
    -------------------------------- createTableStmt ------------------------------

    createTableStmt
        CREATE TABLE ID LPAREN attributeDeclarationList RPAREN SEMICOLON
		CREATE TABLE ID LPAREN attributeDeclarationList COMMA constraintPhraseList LPAREN SEMICOLON

    attributeDeclarationList
        attributeDeclaration
        attributeDeclarationList

    attributeDeclaration
        ID domain
        ID domain constraintList

    constraintList
        constraint
        constraintList constraint

    -------------------------------------------------------------------------------------




   ------------------------------------------------------------------------------------- */
ddl
    :createTableStmt {
        
    }

createTableStmt
    :CREATE TABLE ID LPAREN attributeDeclarationList RPAREN SEMICOLON {

    }
    |CREATE TABLE ID LPAREN attributeDeclarationList COMMA constraintList LPAREN SEMICOLON {

    }
    ;

attributeDeclarationList
    :



/*  --------------------------------------------------------------------------------
    |                                      public                                  |
    --------------------------------------------------------------------------------

    ---------------------------------- constraint ----------------------------------
    
    constraint
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

    	CONSTRAINT ID DEFAULT elementaryValue

		CONSTRAINT ID UNIQUE
		CONSTRAINT ID PRIMARYKEY
		CONSTRAINT ID NOT NULLMARK

		CONSTRAINT ID CHECK LPAREN condition RPAREN

		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN

		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN setDeferrable
		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN onUpdateSet
		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN onDeleteSet

		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN setDeferrable onUpdateSet
		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN setDeferrable onDeleteSet
		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN onUpdateSet setDeferrable
		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN onUpdateSet onDeleteSet
		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN onDeleteSet setDeferrable
		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN onDeleteSet onUpdateSet

		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN setDeferrable onUpdateSet onDeleteSet
		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN setDeferrable onDeleteSet onUpdateSet
		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN onUpdateSet setDeferrable onDeleteSet
		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN onUpdateSet onDeleteSet setDeferrable
		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN onDeleteSet setDeferrable onUpdateSet
		CONSTRAINT ID REFERENCES ID LPAREN ID RPAREN onDeleteSet onUpdateSet setDeferrable

    --------------------------------- elementaryValue -----------------------------
    
    elementaryValue
        INTVALUE
        FLOATVALUE
        STRINGVALUE
        BOOLEANVALUE

   -------------------------------------------------------------------------------- */



%%