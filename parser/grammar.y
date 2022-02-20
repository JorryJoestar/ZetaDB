%{
package parser

import (
    "fmt"
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
    EXPRESSION_ENTRY                NodeEnum = 28

/* predicate */
    COMPAREMARK_ENUM                NodeEnum = 34

/* subquery */
    SUBQUERY_NODE                   NodeEnum = 35

/* createTable */
    ATTRIBUTE_DECLARATION_NODE      NodeEnum = 36

/* createTrigger */
    TRIGGER_FOR_EACH_ENUM           NodeEnum = 38
    TRIGGER_OLDNEW_ENTRY            NodeEnum = 39
    TRIGGER_BEFOREAFTER_NODE        NodeEnum = 40

/* expression */
    EXPRESSION_NODE                 NodeEnum = 41
    EXPRESSION_ENTRY_NODE           NodeEnum = 42

/* aggregation */
    AGGREGATION_NODE                NodeEnum = 43

/* psm */
    PSM_NODE                        NodeEnum = 8
    PSM_VALUE_NODE                  NodeEnum = 44
    PSM_EXEC_ENTRY_NODE             NodeEnum = 45
    PSM_FOR_LOOP_NODE               NodeEnum = 46
    PSM_BRANCH_NODE                 NodeEnum = 47
    PSM_ELSEIF_ENTRY_NODE           NodeEnum = 48
    PSM_PARAMETER_ENTRY_NODE        NodeEnum = 49
    PSM_LOCAL_DECLARATION_ENTRY_NODE NodeEnum = 50

/* delete */
    DELETE_NODE                     NodeEnum = 51

/* insert */
    INSERT_NODE                     NodeEnum = 52

/* update */
    UPDATE_LIST_ENTRY               NodeEnum = 53
    UPDATE_NODE                     NodeEnum = 54
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

/* dql */
    Query            *QueryNode
    SelectListEntry  *SelectListEntryNode
    FromListEntry    *FromListEntryNode
    OnListEntry      *OnListEntryNode
    OrderByListEntry *OrderByListEntryNode

/* createTable */
    AttributeDeclaration *AttributeDeclarationNode

/* constraint */
    ConstraintDeferrable ConstraintDeferrableEnum
    ConstraintUpdateSet  ConstraintUpdateSetEnum
    ConstraintDeleteSet  ConstraintDeleteSetEnum
    Constraint           *ConstraintNode
    ForeignKeyParameter  *ForeignKeyParameterNode

/* aggregation */
    Aggregation             *AggregationNode

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

/* trigger */
    TriggerForEach           TriggerForEachEnum
    TriggerOldNewEntry       *TriggerOldNewEntryNode
    TriggerBeforeAfterStmt   *TriggerBeforeAfterStmtNode

/* expression */
    Expression               *ExpressionNode
    ExpressionEntry          *ExpressionEntryNode

/* psm */
    Psm                      *PsmNode
    PsmValue                 *PsmValueNode
    PsmExecEntry             *PsmExecEntryNode
    PsmForLoop               *PsmForLoopNode
    PsmBranch                *PsmBranchNode
    PsmElseifEntry           *PsmElseifEntryNode
    PsmParameterEntry        *PsmParameterEntryNode
    PsmLocalDeclarationEntry *PsmLocalDeclarationEntryNode

/* delete */
    Delete                   *DeleteNode

/* insert */
    Insert                   *InsertNode

/* update */
    Update                  *UpdateNode
    UpdateListEntry         *UpdateListEntryNode

/* public */
    Subquery                 *QueryNode
}

// -------------------- List --------------------
type ListEnum uint8

const (
    CONSTRAINT_AFTER_ATTRIBUTE_LIST ListEnum = 1
    CONSTRAINT_LIST                 ListEnum = 2
    ATTRINAME_OPTION_TABLENAME_LIST ListEnum = 3
    ATTRIBUTE_DECLARATION_LIST      ListEnum = 4
    TRIGGER_OLDNEW_LIST             ListEnum = 5
    DML_LIST                        ListEnum = 6
    PSM_VALUE_LIST                  ListEnum = 7
    PSM_EXEC_LIST                   ListEnum = 8
    PSM_ELSEIF_LIST                 ListEnum = 9
    PSM_PARAMETER_LIST              ListEnum = 10
    PSM_LOCAL_DECLARATION_LIST      ListEnum = 11
    ELEMENTARY_VALUE_LIST           ListEnum = 12
    UPDATE_LIST                     ListEnum = 13

)

type List struct {
    Type                         ListEnum
    ConstraintAfterAttributeList []*ConstraintNode
    ConstraintList               []*ConstraintNode
    AttriNameOptionTableNameList []*AttriNameOptionTableNameNode
    AttributeDeclarationList     []*AttributeDeclarationNode
    TriggerOldNewList            []*TriggerOldNewEntryNode
    DmlList                      []*DMLNode
    PsmValueList                 []*PsmValueNode
    PsmExecList                  []*PsmExecEntryNode
    PsmElseifList                []*PsmElseifEntryNode
    PsmParameterList             []*PsmParameterEntryNode
    PsmLocalDeclarationList      []*PsmLocalDeclarationEntryNode
    ElementaryValueList          []*ElementaryValueNode
    UpdateList                   []*UpdateListEntryNode
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

// triggerBeforeAfterStmt
type TriggerBeforeAfterStmtNode struct {
    BeforeAfterType      TriggerBeforeAfterEnum
	BeforeAfterAttriName string
	BeforeAfterTableName string
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

// dml
%type <NodePt> dml

// delete
%type <NodePt> deleteStmt
%token FROM WHERE

// insert
%type <NodePt> insertStmt
%type <List> elementaryValueList
%token INSERTINTO VALUES

// update
%type <NodePt> updateStmt
%type <List> updateList
%type <NodePt> updateListEntry

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

// createTrigger
%type <NodePt> createTriggerStmt
%type <NodePt> triggerBeforeAfterStmt
%type <List>   triggerOldNewList
%type <NodePt> oldNewEntry
%type <NodePt> triggerForEachEnum
%type <NodePt> triggerWhenCondition
%type <List>   triggerExecStmt
%type <List>   dmlList
%token TRIGGER REFERENCING BEFORE UPDATE OF AFTER INSTEAD INSERT DELETE
%token OLD ROW NEW FOR EACH STATEMENT WHEN BEGINTOKEN END

// dropTrigger
%type <NodePt> dropTriggerStmt

// aggregation
%type <NodePt> aggregation
%token STAR SUM AVG MIN MAX COUNT DISTINCT

// expression
%type <NodePt> expression
%type <NodePt> expressionEntry
%token PLUS SUBTRACT DIVISION CONCATENATION
%left PLUS SUBTRACT DIVISION CONCATENATION STAR

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

// psm
%type <NodePt> createPsmStmt
%type <NodePt> dropPsmStmt
%type <List> psmParameterList
%type <NodePt> psmParameterEntry
%type <List> psmLocalDeclarationList
%type <NodePt> psmLocalDeclarationEntry
%type <List> psmBody
%type <List> psmExecList
%type <NodePt> psmExecEntry
%type <NodePt> psmForLoop
%type <NodePt> psmBranch
%type <List> psmElseifList
%type <NodePt> psmElseifEntry
%type <NodePt> psmValue
%token ELSEIF THEN IF ELSE CURSOR DO RETURN SET OUT INOUT DECLARE PROCEDURE FUNCTION RETURNS

// psmCallStmt
%type <NodePt> psmCallStmt
%type <NodePt> psmCall
%type <List> psmValueList
%token CALL

// public
%type <List> attriNameOptionTableNameList
%type <StringList> attriNameList
%token LPAREN RPAREN NOT NULLMARK COMMA
%token <String> ID
%%

/*  --------------------------------------------------------------------------------
    |                                      ast                                     |
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
    |dml {
        $$ = &Node{}
        $$.Type = AST_NODE

        $$.Ast = &ASTNode{}
        $$.Ast.Type = AST_DML
        $$.Ast.Dml = $1.Dml

        GetInstance().AST = $$.Ast
    }
    |psmCallStmt { // TODO
        $$ = &Node{}
        fmt.Println("psmCallStmt")
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                     ddl                                      |
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
    |createTriggerStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_TRIGGER_CREATE
        $$.Ddl.Trigger = $1.Trigger
    }
    |dropTriggerStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_TRIGGER_DROP
        $$.Ddl.Trigger = $1.Trigger
    }
    |createPsmStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_PSM_CREATE
        $$.Ddl.Psm = $1.Psm
    }
    |dropPsmStmt {
        $$ = &Node{}
        $$.Type = DDL_NODE

        $$.Ddl = &DDLNode{}
        $$.Ddl.Type = DDL_PSM_DROP
        $$.Ddl.Psm = $1.Psm
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                     dml                                      |
    --------------------------------------------------------------------------------

        dml
            deleteStmt
            insertStmt
            updateStmt

    -------------------------------------------------------------------------------- */
dml
    :deleteStmt {
        $$ = &Node{}
        $$.Type = DML_NODE

        $$.Dml = &DMLNode{}
        $$.Dml.Type = DML_DELETE
        $$.Dml.Delete = $1.Delete
    }
    |insertStmt {
        $$ = &Node{}
        $$.Type = DML_NODE

        $$.Dml = &DMLNode{}
        $$.Dml.Type = DML_INSERT
        $$.Dml.Insert = $1.Insert
    }
    |updateStmt {
        $$ = &Node{}
        $$.Type = DML_NODE

        $$.Dml = &DMLNode{}
        $$.Dml.Type = DML_UPDATE
        $$.Dml.Update = $1.Update
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                   deleteStmt                                 |
    --------------------------------------------------------------------------------

        deleteStmt
            DELETE FROM ID WHERE condition SEMICOLON

    -------------------------------------------------------------------------------- */
deleteStmt
    :DELETE FROM ID WHERE condition SEMICOLON {
        $$ = &Node{}
        $$.Type = DELETE_NODE

        $$.Delete = &DeleteNode{}
        $$.Delete.TableName = $3
        $$.Delete.Condition = $5.Condition
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                   insertStmt                                 |
    --------------------------------------------------------------------------------

        insertStmt
			INSERTINTO ID VALUES subQuery SEMICOLON
			INSERTINTO ID VALUES LPAREN elementaryValueList RPAREN SEMICOLON
			INSERTINTO ID LPAREN attriNameList RPAREN VALUES subQuery SEMICOLON
			INSERTINTO ID LPAREN attriNameList RPAREN VALUES LPAREN elementaryValueList RPAREN SEMICOLON

        elementaryValueList
            elementaryValueList COMMA elementaryValue
            elementaryValue

    -------------------------------------------------------------------------------- */

/*  ---------------------------------- insertStmt ---------------------------------- */
insertStmt
    :INSERTINTO ID VALUES subQuery SEMICOLON {
        $$ = &Node{}
        $$.Type = INSERT_NODE

        $$.Insert = &InsertNode{}

        $$.Insert.Type = INSERT_FROM_SUBQUERY
        $$.Insert.TableName = $2
        $$.Insert.Subquery = $4.Subquery
        $$.Insert.AttriNameListValid = false
    }
    |INSERTINTO ID VALUES LPAREN elementaryValueList RPAREN SEMICOLON {
        $$ = &Node{}
        $$.Type = INSERT_NODE

        $$.Insert = &InsertNode{}

        $$.Insert.Type = INSERT_FROM_VALUELIST
        $$.Insert.TableName = $2
        $$.Insert.AttriNameListValid = false
        $$.Insert.ElementaryValueList = $5.ElementaryValueList
    }
    |INSERTINTO ID LPAREN attriNameList RPAREN VALUES subQuery SEMICOLON {
        $$ = &Node{}
        $$.Type = INSERT_NODE

        $$.Insert = &InsertNode{}

        $$.Insert.Type = INSERT_FROM_SUBQUERY
        $$.Insert.TableName = $2
        $$.Insert.Subquery = $7.Subquery
        $$.Insert.AttriNameListValid = true
        $$.Insert.AttriNameList = $4
    }
    |INSERTINTO ID LPAREN attriNameList RPAREN VALUES LPAREN elementaryValueList RPAREN SEMICOLON {
        $$ = &Node{}
        $$.Type = INSERT_NODE

        $$.Insert = &InsertNode{}

        $$.Insert.Type = INSERT_FROM_VALUELIST
        $$.Insert.TableName = $2
        $$.Insert.ElementaryValueList = $8.ElementaryValueList
        $$.Insert.AttriNameListValid = true
        $$.Insert.AttriNameList = $4
    }
    ;

/*  ----------------------------- elementaryValueList ------------------------------ */
elementaryValueList
    :elementaryValueList COMMA elementaryValue {
        $$ = $1
        $$.ElementaryValueList = append($$.ElementaryValueList,$3.ElementaryValue)
    }
    |elementaryValue {
        $$ = List{}
        $$.Type = ELEMENTARY_VALUE_LIST

        $$.ElementaryValueList = append($$.ElementaryValueList,$1.ElementaryValue)
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                   updateStmt                                 |
    --------------------------------------------------------------------------------
        
        updateStmt
			UPDATE ID SET updateList WHERE condition SEMICOLON

        updateList
			updateListEntry
			updateList COMMA updateListEntry

        updateListEntry
        	ID EQUAL expression
			ID EQUAL elementaryValue

    -------------------------------------------------------------------------------- */

/*  ---------------------------------- updateStmt ---------------------------------- */
updateStmt
	:UPDATE ID SET updateList WHERE condition SEMICOLON {
        $$ = &Node{}
        $$.Type = UPDATE_NODE

        $$.Update = &UpdateNode{}
        $$.Update.TableName = $2
        $$.Update.Condition = $6.Condition
        $$.Update.UpdateList = $4.UpdateList
    }
    ;

/*  ---------------------------------- updateList ---------------------------------- */
updateList
	:updateListEntry {
        $$ = List{}
        $$.Type = UPDATE_LIST
        
        $$.UpdateList = append($$.UpdateList,$1.UpdateListEntry)
    }
	|updateList COMMA updateListEntry {
        $$ = $1
        $$.UpdateList = append($$.UpdateList,$3.UpdateListEntry)
    }
    ;

/*  ------------------------------- updateListEntry -------------------------------- */
updateListEntry
    :ID EQUAL expression {
        $$ = &Node{}
        $$.Type = UPDATE_LIST_ENTRY

        $$.UpdateListEntry = &UpdateListEntryNode{}
        $$.UpdateListEntry.Type = UPDATE_LIST_ENTRY_EXPRESSION
        $$.UpdateListEntry.AttributeName = $1
        $$.UpdateListEntry.Expression = $3.Expression
    }
    |ID EQUAL elementaryValue {
        $$ = &Node{}
        $$.Type = UPDATE_LIST_ENTRY

        $$.UpdateListEntry = &UpdateListEntryNode{}
        $$.UpdateListEntry.Type = UPDATE_LIST_ENTRY_ELEMENTARY_VALUE
        $$.UpdateListEntry.AttributeName = $1
        $$.UpdateListEntry.ElementaryValue = $3.ElementaryValue
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
    |                               createTriggerStmt                              |
    --------------------------------------------------------------------------------

        createTriggerStmt
            CREATE TRIGGER ID triggerBeforeAfterStmt REFERENCING triggerOldNewList triggerForEachEnum triggerWhenCondition triggerExecStmt
            CREATE TRIGGER ID triggerBeforeAfterStmt triggerForEachEnum triggerWhenCondition triggerExecStmt
            CREATE TRIGGER ID triggerBeforeAfterStmt REFERENCING triggerOldNewList triggerForEachEnum triggerExecStmt
            CREATE TRIGGER ID triggerBeforeAfterStmt triggerForEachEnum triggerExecStmt

        triggerBeforeAfterStmt
            BEFORE UPDATE ON ID
			BEFORE UPDATE OF ID ON ID
			AFTER UPDATE ON ID
			AFTER UPDATE OF ID ON ID
			INSTEAD OF UPDATE ON ID
			INSTEAD OF UPDATE OF ID ON ID
			BEFORE INSERT ON ID
			AFTER INSERT ON ID
			INSTEAD OF INSERT ON ID
			BEFORE DELETE ON ID
			AFTER DELETE ON ID
			INSTEAD OF DELETE ON ID

        triggerOldNewList
            oldNewEntry
			triggerOldNewList COMMA oldNewEntry
        
        oldNewEntry
            OLD ROW AS ID
			NEW ROW AS ID
			OLD TABLE AS ID
			NEW TABLE AS ID
        
        triggerForEachEnum
            FOR EACH ROW
            FOR EACH STATEMENT

        triggerWhenCondition
            WHEN condition
        
        triggerExecStmt
            BEGINTOKEN dmlList END SEMICOLON

        dmlList
            dml
            dmlList dml

    -------------------------------------------------------------------------------- */

/*  ------------------------------ createTriggerStmt ------------------------------- */
createTriggerStmt
    :CREATE TRIGGER ID triggerBeforeAfterStmt REFERENCING triggerOldNewList triggerForEachEnum triggerWhenCondition triggerExecStmt {
        $$ = &Node{}
        $$.Type = TRIGGER_NODE

        $$.Trigger = &TriggerNode{}
        $$.Trigger.TriggerName = $3

        $$.Trigger.BeforeAfterType = $4.TriggerBeforeAfterStmt.BeforeAfterType
        $$.Trigger.BeforeAfterAttriName = $4.TriggerBeforeAfterStmt.BeforeAfterAttriName
        $$.Trigger.BeforeAfterTableName = $4.TriggerBeforeAfterStmt.BeforeAfterTableName

        $$.Trigger.ReferencingValid = true
        $$.Trigger.OldNewList = $6.TriggerOldNewList

        $$.Trigger.ForEachType = $7.TriggerForEach

        $$.Trigger.WhenValid = true
        $$.Trigger.Condition = $8.Condition

        $$.Trigger.DmlList = $9.DmlList
    }
    |CREATE TRIGGER ID triggerBeforeAfterStmt triggerForEachEnum triggerWhenCondition triggerExecStmt {
        $$ = &Node{}
        $$.Type = TRIGGER_NODE

        $$.Trigger = &TriggerNode{}
        $$.Trigger.TriggerName = $3

        $$.Trigger.BeforeAfterType = $4.TriggerBeforeAfterStmt.BeforeAfterType
        $$.Trigger.BeforeAfterAttriName = $4.TriggerBeforeAfterStmt.BeforeAfterAttriName
        $$.Trigger.BeforeAfterTableName = $4.TriggerBeforeAfterStmt.BeforeAfterTableName

        $$.Trigger.ReferencingValid = false

        $$.Trigger.ForEachType = $5.TriggerForEach

        $$.Trigger.WhenValid = true
        $$.Trigger.Condition = $6.Condition

        $$.Trigger.DmlList = $7.DmlList
    }
    |CREATE TRIGGER ID triggerBeforeAfterStmt REFERENCING triggerOldNewList triggerForEachEnum triggerExecStmt {
        $$ = &Node{}
        $$.Type = TRIGGER_NODE

        $$.Trigger = &TriggerNode{}
        $$.Trigger.TriggerName = $3

        $$.Trigger.BeforeAfterType = $4.TriggerBeforeAfterStmt.BeforeAfterType
        $$.Trigger.BeforeAfterAttriName = $4.TriggerBeforeAfterStmt.BeforeAfterAttriName
        $$.Trigger.BeforeAfterTableName = $4.TriggerBeforeAfterStmt.BeforeAfterTableName

        $$.Trigger.ReferencingValid = true
        $$.Trigger.OldNewList = $6.TriggerOldNewList

        $$.Trigger.ForEachType = $7.TriggerForEach

        $$.Trigger.WhenValid = false

        $$.Trigger.DmlList = $8.DmlList
    }
    |CREATE TRIGGER ID triggerBeforeAfterStmt triggerForEachEnum triggerExecStmt {
        $$ = &Node{}
        $$.Type = TRIGGER_NODE

        $$.Trigger = &TriggerNode{}
        $$.Trigger.TriggerName = $3

        $$.Trigger.BeforeAfterType = $4.TriggerBeforeAfterStmt.BeforeAfterType
        $$.Trigger.BeforeAfterAttriName = $4.TriggerBeforeAfterStmt.BeforeAfterAttriName
        $$.Trigger.BeforeAfterTableName = $4.TriggerBeforeAfterStmt.BeforeAfterTableName

        $$.Trigger.ReferencingValid = false

        $$.Trigger.ForEachType = $5.TriggerForEach

        $$.Trigger.WhenValid = false

        $$.Trigger.DmlList = $6.DmlList
    }
    ;

/*  --------------------------- triggerBeforeAfterStmt ----------------------------- */
triggerBeforeAfterStmt
    :BEFORE UPDATE ON ID {
        $$ = &Node{}
        $$.Type = TRIGGER_BEFOREAFTER_NODE

        $$.TriggerBeforeAfterStmt = &TriggerBeforeAfterStmtNode{}
        $$.TriggerBeforeAfterStmt.BeforeAfterType = BEFORE_UPDATE
        $$.TriggerBeforeAfterStmt.BeforeAfterTableName = $4
    }
	|BEFORE UPDATE OF ID ON ID {
        $$ = &Node{}
        $$.Type = TRIGGER_BEFOREAFTER_NODE

        $$.TriggerBeforeAfterStmt = &TriggerBeforeAfterStmtNode{}
        $$.TriggerBeforeAfterStmt.BeforeAfterType = BEFORE_UPDATE_OF
        $$.TriggerBeforeAfterStmt.BeforeAfterAttriName = $4
        $$.TriggerBeforeAfterStmt.BeforeAfterTableName = $6
    }
	|AFTER UPDATE ON ID {
        $$ = &Node{}
        $$.Type = TRIGGER_BEFOREAFTER_NODE

        $$.TriggerBeforeAfterStmt = &TriggerBeforeAfterStmtNode{}
        $$.TriggerBeforeAfterStmt.BeforeAfterType = AFTER_UPDATE
        $$.TriggerBeforeAfterStmt.BeforeAfterTableName = $4
    }
	|AFTER UPDATE OF ID ON ID {
        $$ = &Node{}
        $$.Type = TRIGGER_BEFOREAFTER_NODE

        $$.TriggerBeforeAfterStmt = &TriggerBeforeAfterStmtNode{}
        $$.TriggerBeforeAfterStmt.BeforeAfterType = AFTER_UPDATE_OF
        $$.TriggerBeforeAfterStmt.BeforeAfterAttriName = $4
        $$.TriggerBeforeAfterStmt.BeforeAfterTableName = $6
    }
	|INSTEAD OF UPDATE ON ID {
        $$ = &Node{}
        $$.Type = TRIGGER_BEFOREAFTER_NODE

        $$.TriggerBeforeAfterStmt = &TriggerBeforeAfterStmtNode{}
        $$.TriggerBeforeAfterStmt.BeforeAfterType = INSTEAD_UPDATE
        $$.TriggerBeforeAfterStmt.BeforeAfterTableName = $5
    }
	|INSTEAD OF UPDATE OF ID ON ID {
        $$ = &Node{}
        $$.Type = TRIGGER_BEFOREAFTER_NODE

        $$.TriggerBeforeAfterStmt = &TriggerBeforeAfterStmtNode{}
        $$.TriggerBeforeAfterStmt.BeforeAfterType = INSTEAD_UPDATE_OF
        $$.TriggerBeforeAfterStmt.BeforeAfterAttriName = $5
        $$.TriggerBeforeAfterStmt.BeforeAfterTableName = $7
    }
	|BEFORE INSERT ON ID {
        $$ = &Node{}
        $$.Type = TRIGGER_BEFOREAFTER_NODE

        $$.TriggerBeforeAfterStmt = &TriggerBeforeAfterStmtNode{}
        $$.TriggerBeforeAfterStmt.BeforeAfterType = BEFORE_INSERT
        $$.TriggerBeforeAfterStmt.BeforeAfterTableName = $4
    }
	|AFTER INSERT ON ID {
        $$ = &Node{}
        $$.Type = TRIGGER_BEFOREAFTER_NODE

        $$.TriggerBeforeAfterStmt = &TriggerBeforeAfterStmtNode{}
        $$.TriggerBeforeAfterStmt.BeforeAfterType = AFTER_INSERT
        $$.TriggerBeforeAfterStmt.BeforeAfterTableName = $4
    }
	|INSTEAD OF INSERT ON ID {
        $$ = &Node{}
        $$.Type = TRIGGER_BEFOREAFTER_NODE

        $$.TriggerBeforeAfterStmt = &TriggerBeforeAfterStmtNode{}
        $$.TriggerBeforeAfterStmt.BeforeAfterType = INSTEAD_INSERT
        $$.TriggerBeforeAfterStmt.BeforeAfterTableName = $5
    }
	|BEFORE DELETE ON ID {
        $$ = &Node{}
        $$.Type = TRIGGER_BEFOREAFTER_NODE

        $$.TriggerBeforeAfterStmt = &TriggerBeforeAfterStmtNode{}
        $$.TriggerBeforeAfterStmt.BeforeAfterType = BEFORE_DELETE
        $$.TriggerBeforeAfterStmt.BeforeAfterTableName = $4
    }
	|AFTER DELETE ON ID {
        $$ = &Node{}
        $$.Type = TRIGGER_BEFOREAFTER_NODE

        $$.TriggerBeforeAfterStmt = &TriggerBeforeAfterStmtNode{}
        $$.TriggerBeforeAfterStmt.BeforeAfterType = AFTER_DELETE
        $$.TriggerBeforeAfterStmt.BeforeAfterTableName = $4
    }
	|INSTEAD OF DELETE ON ID {
        $$ = &Node{}
        $$.Type = TRIGGER_BEFOREAFTER_NODE

        $$.TriggerBeforeAfterStmt = &TriggerBeforeAfterStmtNode{}
        $$.TriggerBeforeAfterStmt.BeforeAfterType = INSTEAD_DELETE
        $$.TriggerBeforeAfterStmt.BeforeAfterTableName = $5
    }
    ;

/*  ----------------------------- triggerOldNewList -------------------------------- */
triggerOldNewList
    :oldNewEntry {
        $$ = List{}
        $$.Type = TRIGGER_OLDNEW_LIST

        $$.TriggerOldNewList = append($$.TriggerOldNewList,$1.TriggerOldNewEntry)
    }
	|triggerOldNewList COMMA oldNewEntry {
        $$ = $1

        $$.TriggerOldNewList = append($$.TriggerOldNewList,$3.TriggerOldNewEntry)
    }
    ;

/*  --------------------------------- oldNewEntry ---------------------------------- */
oldNewEntry
    :OLD ROW AS ID {
        $$ = &Node{}
        $$.Type = TRIGGER_OLDNEW_ENTRY

        $$.TriggerOldNewEntry = &TriggerOldNewEntryNode{}
        $$.TriggerOldNewEntry.Type = OLD_ROW_AS
        $$.TriggerOldNewEntry.Name = $4
    }
	|NEW ROW AS ID {
        $$ = &Node{}
        $$.Type = TRIGGER_OLDNEW_ENTRY

        $$.TriggerOldNewEntry = &TriggerOldNewEntryNode{}
        $$.TriggerOldNewEntry.Type = NEW_ROW_AS
        $$.TriggerOldNewEntry.Name = $4
    }
	|OLD TABLE AS ID {
        $$ = &Node{}
        $$.Type = TRIGGER_OLDNEW_ENTRY

        $$.TriggerOldNewEntry = &TriggerOldNewEntryNode{}
        $$.TriggerOldNewEntry.Type = OLD_TABLE_AS
        $$.TriggerOldNewEntry.Name = $4
    }
	|NEW TABLE AS ID {
        $$ = &Node{}
        $$.Type = TRIGGER_OLDNEW_ENTRY

        $$.TriggerOldNewEntry = &TriggerOldNewEntryNode{}
        $$.TriggerOldNewEntry.Type = NEW_TABLE_AS
        $$.TriggerOldNewEntry.Name = $4
    }
    ;

/*  ------------------------------ triggerForEachEnum ------------------------------ */
triggerForEachEnum
    :FOR EACH ROW {
        $$ = &Node{}
        $$.Type = TRIGGER_FOR_EACH_ENUM
        $$.TriggerForEach = FOR_EACH_ROW
    }
    |FOR EACH STATEMENT {
        $$ = &Node{}
        $$.Type = TRIGGER_FOR_EACH_ENUM
        $$.TriggerForEach = FOR_EACH_STATEMENT
    }
    ;

/*  ---------------------------- triggerWhenCondition ------------------------------ */
triggerWhenCondition
    :WHEN condition {
        $$ = $2
    }
    ;

/*  -------------------------------- triggerExecStmt ------------------------------- */
triggerExecStmt
    :BEGINTOKEN dmlList END SEMICOLON {
        $$ = $2
    }
    ;

/*  ---------------------------------- dmlList ------------------------------------- */
dmlList
    :dml {
        $$ = List{}
        $$.Type = DML_LIST
        
        $$.DmlList = append($$.DmlList,$1.Dml)
    }
    |dmlList dml {
        $$ = $1
        $$.DmlList = append($$.DmlList,$2.Dml)
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                dropTriggerStmt                               |
    --------------------------------------------------------------------------------
        
        dropTriggerStmt
            DROP TRIGGER ID SEMICOLON

    -------------------------------------------------------------------------------- */
dropTriggerStmt
    :DROP TRIGGER ID SEMICOLON {
        $$ = &Node{}
        $$.Type = TRIGGER_NODE

        $$.Trigger = &TriggerNode{}
        $$.Trigger.TriggerName = $3
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                 createPsmStmt                                |
    --------------------------------------------------------------------------------

        createPsmStmt
            CREATE PROCEDURE ID LPAREN psmParameterList RPAREN psmLocalDeclarationList psmBody
            CREATE PROCEDURE ID LPAREN psmParameterList RPAREN psmBody
            CREATE PROCEDURE ID psmLocalDeclarationList psmBody
            CREATE PROCEDURE ID psmBody
            CREATE FUNCTION ID LPAREN psmParameterList RPAREN RETURNS domain psmLocalDeclarationList psmBody
            CREATE FUNCTION ID LPAREN psmParameterList RPAREN RETURNS domain psmBody
            CREATE FUNCTION ID RETURNS domain psmLocalDeclarationList psmBody
            CREATE FUNCTION ID RETURNS domain psmBody

        psmParameterList
            psmParameterList COMMA psmParameterEntry
            psmParameterEntry

        psmParameterEntry
            IN ID domain
			OUT ID domain
			INOUT ID domain

        psmLocalDeclarationList
            psmLocalDeclarationList psmLocalDeclarationEntry
            psmLocalDeclarationEntry

        psmLocalDeclarationEntry
            DECLARE ID domain SEMICOLON
        
        psmBody
            BEGINTOKEN psmExecList END SEMICOLON

        psmExecList
            psmExecList psmExecEntry
            psmExecEntry

        psmExecEntry
            RETURN psmValue SEMICOLON
			SET ID EQUAL psmValue SEMICOLON
			psmForLoop
			psmBranch
            dml

        psmForLoop
            FOR ID AS ID CURSOR FOR subQuery DO psmExecList END FOR SEMICOLON

        psmBranch
            IF condition THEN psmExecList psmElseifList ELSE psmExecList END IF SEMICOLON
            IF condition THEN psmExecList psmElseifList END IF SEMICOLON
            IF condition THEN psmExecList ELSE psmExecList END IF SEMICOLON
            IF condition THEN psmExecList END IF SEMICOLON

        psmElseifList
            psmElseifList psmElseifEntry
            psmElseifEntry

        psmElseifEntry
            ELSEIF condition THEN psmExecList

        psmValue
            elementaryValue
			psmCall
			expression
			ID

    -------------------------------------------------------------------------------- */

/*  ------------------------------- createPsmStmt ---------------------------------- */
createPsmStmt
    :CREATE PROCEDURE ID LPAREN psmParameterList RPAREN psmLocalDeclarationList psmBody {
        $$ = &Node{}
        $$.Type = PSM_NODE

        $$.Psm = &PsmNode{}
        $$.Psm.Type = PSM_PROCEDURE
        $$.Psm.PsmName = $3
        $$.Psm.PsmParameterListValid = true
        $$.Psm.PsmParameterList = $5.PsmParameterList
        $$.Psm.PsmLocalDeclarationListValid = true
        $$.Psm.PsmLocalDeclarationList = $7.PsmLocalDeclarationList
        $$.Psm.PsmBody = $8.PsmExecList

    }
    |CREATE PROCEDURE ID LPAREN psmParameterList RPAREN psmBody {
        $$ = &Node{}
        $$.Type = PSM_NODE

        $$.Psm = &PsmNode{}
        $$.Psm.Type = PSM_PROCEDURE
        $$.Psm.PsmName = $3
        $$.Psm.PsmParameterListValid = true
        $$.Psm.PsmParameterList = $5.PsmParameterList
        $$.Psm.PsmLocalDeclarationListValid = false
        $$.Psm.PsmBody = $7.PsmExecList
    }
    |CREATE PROCEDURE ID psmLocalDeclarationList psmBody {
        $$ = &Node{}
        $$.Type = PSM_NODE

        $$.Psm = &PsmNode{}
        $$.Psm.Type = PSM_PROCEDURE
        $$.Psm.PsmName = $3
        $$.Psm.PsmParameterListValid = false
        $$.Psm.PsmLocalDeclarationListValid = true
        $$.Psm.PsmLocalDeclarationList = $4.PsmLocalDeclarationList
        $$.Psm.PsmBody = $5.PsmExecList
    }
    |CREATE PROCEDURE ID psmBody {
        $$ = &Node{}
        $$.Type = PSM_NODE

        $$.Psm = &PsmNode{}
        $$.Psm.Type = PSM_PROCEDURE
        $$.Psm.PsmName = $3
        $$.Psm.PsmParameterListValid = false
        $$.Psm.PsmLocalDeclarationListValid = false
        $$.Psm.PsmBody = $4.PsmExecList
    }
    |CREATE FUNCTION ID LPAREN psmParameterList RPAREN RETURNS domain psmLocalDeclarationList psmBody {
        $$ = &Node{}
        $$.Type = PSM_NODE

        $$.Psm = &PsmNode{}
        $$.Psm.Type = PSM_FUNCTION
        $$.Psm.PsmName = $3
        $$.Psm.PsmParameterListValid = true
        $$.Psm.PsmParameterList = $5.PsmParameterList
        $$.Psm.PsmLocalDeclarationListValid = true
        $$.Psm.PsmLocalDeclarationList = $9.PsmLocalDeclarationList
        $$.Psm.PsmBody = $10.PsmExecList
    }
    |CREATE FUNCTION ID LPAREN psmParameterList RPAREN RETURNS domain psmBody {
        $$ = &Node{}
        $$.Type = PSM_NODE

        $$.Psm = &PsmNode{}
        $$.Psm.Type = PSM_FUNCTION
        $$.Psm.PsmName = $3
        $$.Psm.PsmParameterListValid = true
        $$.Psm.PsmParameterList = $5.PsmParameterList
        $$.Psm.PsmLocalDeclarationListValid = false
        $$.Psm.PsmBody = $9.PsmExecList
    }
    |CREATE FUNCTION ID RETURNS domain psmLocalDeclarationList psmBody {
        $$ = &Node{}
        $$.Type = PSM_NODE

        $$.Psm = &PsmNode{}
        $$.Psm.Type = PSM_FUNCTION
        $$.Psm.PsmName = $3
        $$.Psm.PsmParameterListValid = false
        $$.Psm.PsmLocalDeclarationListValid = true
        $$.Psm.PsmLocalDeclarationList = $6.PsmLocalDeclarationList
        $$.Psm.PsmBody = $7.PsmExecList
    }
    |CREATE FUNCTION ID RETURNS domain psmBody {
        $$ = &Node{}
        $$.Type = PSM_NODE

        $$.Psm = &PsmNode{}
        $$.Psm.Type = PSM_FUNCTION
        $$.Psm.PsmName = $3
        $$.Psm.PsmParameterListValid = false
        $$.Psm.PsmLocalDeclarationListValid = false
        $$.Psm.PsmBody = $6.PsmExecList
    }
    ;

/*  ------------------------------- psmParameterList ------------------------------- */
psmParameterList
    :psmParameterList COMMA psmParameterEntry {
        $$ = $1
        $$.PsmParameterList = append($$.PsmParameterList,$3.PsmParameterEntry)
    }
    |psmParameterEntry {
        $$ = List{}
        $$.Type = PSM_PARAMETER_LIST

        $$.PsmParameterList = append($$.PsmParameterList,$1.PsmParameterEntry)
    }
    ;

/*  ------------------------------ psmParameterEntry-------------------------------- */
psmParameterEntry
    :IN ID domain {
        $$ = &Node{}
        $$.Type = PSM_PARAMETER_ENTRY_NODE

        $$.PsmParameterEntry = &PsmParameterEntryNode{}
        $$.PsmParameterEntry.Type = PSM_PARAMETER_IN
        $$.PsmParameterEntry.Name = $2
        $$.PsmParameterEntry.Domain = $3.Domain


    }
	|OUT ID domain {
        $$ = &Node{}
        $$.Type = PSM_PARAMETER_ENTRY_NODE

        $$.PsmParameterEntry = &PsmParameterEntryNode{}
        $$.PsmParameterEntry.Type = PSM_PARAMETER_OUT
        $$.PsmParameterEntry.Name = $2
        $$.PsmParameterEntry.Domain = $3.Domain
    }
	|INOUT ID domain {
        $$ = &Node{}
        $$.Type = PSM_PARAMETER_ENTRY_NODE

        $$.PsmParameterEntry = &PsmParameterEntryNode{}
        $$.PsmParameterEntry.Type = PSM_PARAMETER_INOUT
        $$.PsmParameterEntry.Name = $2
        $$.PsmParameterEntry.Domain = $3.Domain
    }
    ;

/*  -------------------------- psmLocalDeclarationList ----------------------------- */
psmLocalDeclarationList
    :psmLocalDeclarationList psmLocalDeclarationEntry {
        $$ = $1
        $$.PsmLocalDeclarationList = append($$.PsmLocalDeclarationList,$2.PsmLocalDeclarationEntry)
    }
    |psmLocalDeclarationEntry {
        $$ = List{}
        $$.Type = PSM_LOCAL_DECLARATION_LIST

        $$.PsmLocalDeclarationList = append($$.PsmLocalDeclarationList,$1.PsmLocalDeclarationEntry)
    }
    ;

/*  -------------------------- psmLocalDeclarationEntry ---------------------------- */
psmLocalDeclarationEntry
    :DECLARE ID domain SEMICOLON {
        $$ = &Node{}
        $$.Type = PSM_LOCAL_DECLARATION_ENTRY_NODE

        $$.PsmLocalDeclarationEntry = &PsmLocalDeclarationEntryNode{}
        $$.PsmLocalDeclarationEntry.Name = $2
        $$.PsmLocalDeclarationEntry.Domain = $3.Domain
    }
    ;

/*  ----------------------------------- psmBody ------------------------------------ */
psmBody
    :BEGINTOKEN psmExecList END SEMICOLON {
        $$ = $2
    }
    ;

/*  --------------------------------- psmExecList ---------------------------------- */
psmExecList
    :psmExecList psmExecEntry {
        $$ = $1
        $$.PsmExecList = append($$.PsmExecList,$2.PsmExecEntry)

    }
    |psmExecEntry {
        $$ = List{}
        $$.Type = PSM_EXEC_LIST

        $$.PsmExecList = append($$.PsmExecList,$1.PsmExecEntry)
    }
    ;

/*  -------------------------------- psmExecEntry ---------------------------------- */
psmExecEntry
    :RETURN psmValue SEMICOLON {
        $$ = &Node{}
        $$.Type = PSM_EXEC_ENTRY_NODE

        $$.PsmExecEntry = &PsmExecEntryNode{}
        $$.PsmExecEntry.Type = PSM_EXEC_RETURN
        $$.PsmExecEntry.PsmValue = $2.PsmValue
    }
    |SET ID EQUAL psmValue SEMICOLON {
        $$ = &Node{}
        $$.Type = PSM_EXEC_ENTRY_NODE

        $$.PsmExecEntry = &PsmExecEntryNode{}
        $$.PsmExecEntry.Type = PSM_EXEC_SET
        $$.PsmExecEntry.VariableName = $2
        $$.PsmExecEntry.PsmValue = $4.PsmValue
    }
	|psmForLoop {
        $$ = &Node{}
        $$.Type = PSM_EXEC_ENTRY_NODE

        $$.PsmExecEntry = &PsmExecEntryNode{}
        $$.PsmExecEntry.Type = PSM_EXEC_FOR_LOOP
        $$.PsmExecEntry.PsmForLoop = $1.PsmForLoop
    }
	|psmBranch {
        $$ = &Node{}
        $$.Type = PSM_EXEC_ENTRY_NODE

        $$.PsmExecEntry = &PsmExecEntryNode{}
        $$.PsmExecEntry.Type = PSM_EXEC_BRANCH
        $$.PsmExecEntry.PsmBranch = $1.PsmBranch
    }
    |dml {
        $$ = &Node{}
        $$.Type = PSM_EXEC_ENTRY_NODE

        $$.PsmExecEntry = &PsmExecEntryNode{}
        $$.PsmExecEntry.Type = PSM_EXEC_DML
        $$.PsmExecEntry.Dml = $1.Dml
    }
    ;

/*  ---------------------------------- psmForLoop ---------------------------------- */
psmForLoop
    :FOR ID AS ID CURSOR FOR subQuery DO psmExecList END FOR SEMICOLON {
        $$ = &Node{}
        $$.Type = PSM_FOR_LOOP_NODE

        $$.PsmForLoop = &PsmForLoopNode{}
        $$.PsmForLoop.LoopName = $2
        $$.PsmForLoop.CursorName = $4
        $$.PsmForLoop.Subquery = $7.Query
        $$.PsmForLoop.PsmExecList = $9.PsmExecList
    }
    ;

/*  ---------------------------------- psmBranch ----------------------------------- */
psmBranch
    :IF condition THEN psmExecList psmElseifList ELSE psmExecList END IF SEMICOLON {
        $$ = &Node{}
        $$.Type = PSM_BRANCH_NODE

        $$.PsmBranch = &PsmBranchNode{}

        $$.PsmBranch.Condition = $2.Condition
        $$.PsmBranch.IfPsmExecList = $4.PsmExecList
        $$.PsmBranch.PsmElseifListValid = true
        $$.PsmBranch.PsmElseifList = $5.PsmElseifList
        $$.PsmBranch.ElseValid = true
        $$.PsmBranch.ElsePsmExecList = $7.PsmExecList
    }
    |IF condition THEN psmExecList psmElseifList END IF SEMICOLON {
        $$ = &Node{}
        $$.Type = PSM_BRANCH_NODE

        $$.PsmBranch = &PsmBranchNode{}

        $$.PsmBranch.Condition = $2.Condition
        $$.PsmBranch.IfPsmExecList = $4.PsmExecList
        $$.PsmBranch.PsmElseifListValid = true
        $$.PsmBranch.PsmElseifList = $5.PsmElseifList
        $$.PsmBranch.ElseValid = false
    }
    |IF condition THEN psmExecList ELSE psmExecList END IF SEMICOLON {
        $$ = &Node{}
        $$.Type = PSM_BRANCH_NODE

        $$.PsmBranch = &PsmBranchNode{}

        $$.PsmBranch.Condition = $2.Condition
        $$.PsmBranch.IfPsmExecList = $4.PsmExecList
        $$.PsmBranch.PsmElseifListValid = false
        $$.PsmBranch.ElseValid = true
        $$.PsmBranch.ElsePsmExecList = $6.PsmExecList
    }
    |IF condition THEN psmExecList END IF SEMICOLON {
        $$ = &Node{}
        $$.Type = PSM_BRANCH_NODE

        $$.PsmBranch = &PsmBranchNode{}

        $$.PsmBranch.Condition = $2.Condition
        $$.PsmBranch.IfPsmExecList = $4.PsmExecList
        $$.PsmBranch.PsmElseifListValid = false
        $$.PsmBranch.ElseValid = false
    }
    ;

/*  -------------------------------- psmElseifList --------------------------------- */
psmElseifList
    :psmElseifList psmElseifEntry {
        $$ = $1
        $$.PsmElseifList = append($$.PsmElseifList,$2.PsmElseifEntry)
    }
    |psmElseifEntry {
        $$ = List{}
        $$.Type = PSM_ELSEIF_LIST

        $$.PsmElseifList = append($$.PsmElseifList,$1.PsmElseifEntry)
    }
    ;

/*  ------------------------------- psmElseifEntry --------------------------------- */
psmElseifEntry
    :ELSEIF condition THEN psmExecList {
        $$ = &Node{}
        $$.Type = PSM_ELSEIF_ENTRY_NODE

        $$.PsmElseifEntry = &PsmElseifEntryNode{}
        $$.PsmElseifEntry.Condition = $2.Condition
        $$.PsmElseifEntry.PsmExecList = $4.PsmExecList
    }
    ;

/*  ------------------------------------ psmValue ---------------------------------- */
psmValue
    :elementaryValue {
        $$ = &Node{}
        $$.Type = PSM_VALUE_NODE

        $$.PsmValue = &PsmValueNode{}
        $$.PsmValue.Type = PSMVALUE_ELEMENTARY_VALUE
        $$.PsmValue.ElementaryValue = $1.ElementaryValue
    }
	|psmCall {
        $$ = &Node{}
        $$.Type = PSM_VALUE_NODE

        $$.PsmValue = &PsmValueNode{}
        $$.PsmValue.Type = PSMVALUE_CALL
        $$.PsmValue.PsmCall = $1.Psm
    }
	|expression {
        $$ = &Node{}
        $$.Type = PSM_VALUE_NODE

        $$.PsmValue = &PsmValueNode{}
        $$.PsmValue.Type = PSMVALUE_EXPRESSION
        $$.PsmValue.Expression = $1.Expression
    }
	|ID {
        $$ = &Node{}
        $$.Type = PSM_VALUE_NODE

        $$.PsmValue = &PsmValueNode{}
        $$.PsmValue.Type = PSMVALUE_ID
        $$.PsmValue.Id = $1
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                  dropPsmStmt                                 |
    --------------------------------------------------------------------------------

        dropPsmStmt
            DROP FUNCTION ID SEMICOLON
            DROP PROCEDURE ID SEMICOLON
    
    -------------------------------------------------------------------------------- */
dropPsmStmt
    :DROP FUNCTION ID SEMICOLON {
        $$ = &Node{}
        $$.Type = PSM_NODE

        $$.Psm = &PsmNode{}
        $$.Psm.Type = PSM_FUNCTION
        $$.Psm.PsmName = $3
    }
    |DROP PROCEDURE ID SEMICOLON {
        $$ = &Node{}
        $$.Type = PSM_NODE

        $$.Psm = &PsmNode{}
        $$.Psm.Type = PSM_PROCEDURE
        $$.Psm.PsmName = $3
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                  psmCallStmt                                 |
    --------------------------------------------------------------------------------

        psmCallStmt
            psmCall SEMICOLON
        
        psmCall
            CALL ID LPAREN psmValueList RPAREN
            CALL ID LPAREN RPAREN

        psmValueList
            psmValueList COMMA psmValue
            psmValue

    -------------------------------------------------------------------------------- */

/*  --------------------------------- psmCallStmt ---------------------------------- */
psmCallStmt
    :psmCall SEMICOLON {
        $$ = $1
    }
    ;

/*  ------------------------------------ psmCall ----------------------------------- */
psmCall
    :CALL ID LPAREN psmValueList RPAREN {
        $$ = &Node{}
        $$.Type = PSM_NODE

        $$.Psm = &PsmNode{}
        $$.Psm.PsmName = $2
        $$.Psm.PsmValueListValid = true
        $$.Psm.PsmValueList = $4.PsmValueList
    }
    |CALL ID LPAREN RPAREN {
        $$ = &Node{}
        $$.Type = PSM_NODE

        $$.Psm = &PsmNode{}
        $$.Psm.PsmName = $2
        $$.Psm.PsmValueListValid = false
    }
    ;

/*  --------------------------------- psmValueList --------------------------------- */
psmValueList
    :psmValueList COMMA psmValue {
        $$ = $1
        $$.PsmValueList = append($$.PsmValueList,$3.PsmValue)
    }
    |psmValue {
        $$ = List{}
        $$.Type = PSM_VALUE_LIST
        $$.PsmValueList = append($$.PsmValueList,$1.PsmValue)
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                  expression                                  |
    --------------------------------------------------------------------------------

        expression
            expressionEntry PLUS expressionEntry
            expressionEntry SUBTRACT expressionEntry
            expressionEntry STAR expressionEntry
            expressionEntry DIVISION expressionEntry
            expressionEntry CONCATENATION expressionEntry

        expressionEntry
            elementaryValue
            attriNameOptionTableName
            aggregation
            expression
            LPAREN expression RPAREN

    -------------------------------------------------------------------------------- */
        
/*  ---------------------------------- expression ---------------------------------- */
expression
    :expressionEntry PLUS expressionEntry {
        $$ = &Node{}
        $$.Type = EXPRESSION_NODE

        $$.Expression = &ExpressionNode{}
        $$.Expression.Type = EXPRESSION_OPERATOR_PLUS
        $$.Expression.ExpressionEntryL = $1.ExpressionEntry
        $$.Expression.ExpressionEntryR = $3.ExpressionEntry
    }
    |expressionEntry SUBTRACT expressionEntry {
        $$ = &Node{}
        $$.Type = EXPRESSION_NODE

        $$.Expression = &ExpressionNode{}
        $$.Expression.Type = EXPRESSION_OPERATOR_MINUS
        $$.Expression.ExpressionEntryL = $1.ExpressionEntry
        $$.Expression.ExpressionEntryR = $3.ExpressionEntry
    }
    |expressionEntry STAR expressionEntry {
        $$ = &Node{}
        $$.Type = EXPRESSION_NODE

        $$.Expression = &ExpressionNode{}
        $$.Expression.Type = EXPRESSION_OPERATOR_MULTIPLY
        $$.Expression.ExpressionEntryL = $1.ExpressionEntry
        $$.Expression.ExpressionEntryR = $3.ExpressionEntry
    }
    |expressionEntry DIVISION expressionEntry {
        $$ = &Node{}
        $$.Type = EXPRESSION_NODE

        $$.Expression = &ExpressionNode{}
        $$.Expression.Type = EXPRESSION_OPERATOR_DIVISION
        $$.Expression.ExpressionEntryL = $1.ExpressionEntry
        $$.Expression.ExpressionEntryR = $3.ExpressionEntry
    }
    |expressionEntry CONCATENATION expressionEntry {
        $$ = &Node{}
        $$.Type = EXPRESSION_NODE

        $$.Expression = &ExpressionNode{}
        $$.Expression.Type = EXPRESSION_OPERATOR_CONCATENATION
        $$.Expression.ExpressionEntryL = $1.ExpressionEntry
        $$.Expression.ExpressionEntryR = $3.ExpressionEntry
    }
    ;

/*  ------------------------------- expressionEntry -------------------------------- */
expressionEntry
    :elementaryValue {
        $$ = &Node{}
        $$.Type = EXPRESSION_ENTRY_NODE

        $$.ExpressionEntry = &ExpressionEntryNode{}
        $$.ExpressionEntry.Type = EXPRESSION_ENTRY_ELEMENTARY_VALUE

        $$.ExpressionEntry.ElementaryValue = $1.ElementaryValue
    }
    |attriNameOptionTableName {
        $$ = &Node{}
        $$.Type = EXPRESSION_ENTRY_NODE

        $$.ExpressionEntry = &ExpressionEntryNode{}
        $$.ExpressionEntry.Type = EXPRESSION_ENTRY_ATTRIBUTE_NAME

        $$.ExpressionEntry.AttriNameOptionTableName = $1.AttriNameOptionTableName
    }
    |aggregation {
        $$ = &Node{}
        $$.Type = EXPRESSION_ENTRY_NODE

        $$.ExpressionEntry = &ExpressionEntryNode{}
        $$.ExpressionEntry.Type = EXPRESSION_ENTRY_AGGREGATION

        $$.ExpressionEntry.Aggregation = $1.Aggregation
    }
    |expression {
        $$ = &Node{}
        $$.Type = EXPRESSION_ENTRY_NODE

        $$.ExpressionEntry = &ExpressionEntryNode{}
        $$.ExpressionEntry.Type = EXPRESSION_ENTRY_EXPRESSION

        $$.ExpressionEntry.Expression = $1.Expression
    }
    |LPAREN expression RPAREN {
        $$ = &Node{}
        $$.Type = EXPRESSION_ENTRY_NODE

        $$.ExpressionEntry = &ExpressionEntryNode{}
        $$.ExpressionEntry.Type = EXPRESSION_ENTRY_EXPRESSION

        $$.ExpressionEntry.Expression = $2.Expression
    }
    ;

/*  --------------------------------------------------------------------------------
    |                                 aggregation                                  |
    --------------------------------------------------------------------------------

        aggregation
            SUM LPAREN DISTINCT attriNameOptionTableName RPAREN
            SUM LPAREN attriNameOptionTableName RPAREN
            AVG LPAREN DISTINCT attriNameOptionTableName RPAREN
            AVG LPAREN attriNameOptionTableName RPAREN
            MIN LPAREN DISTINCT attriNameOptionTableName RPAREN
            MIN LPAREN attriNameOptionTableName RPAREN
            MAX LPAREN DISTINCT attriNameOptionTableName RPAREN
            MAX LPAREN attriNameOptionTableName RPAREN
            COUNT LPAREN DISTINCT attriNameOptionTableName RPAREN
            COUNT LPAREN attriNameOptionTableName RPAREN
            COUNT LPAREN STAR RPAREN

    -------------------------------------------------------------------------------- */
aggregation
    :SUM LPAREN DISTINCT attriNameOptionTableName RPAREN {
        $$ = &Node{}
        $$.Type = AGGREGATION_NODE

        $$.Aggregation = &AggregationNode{}
        $$.Aggregation.Type = AGGREGATION_SUM
        $$.Aggregation.DistinctValid = true
        $$.Aggregation.AttriNameOptionTableName = $4.AttriNameOptionTableName
    }
    |SUM LPAREN attriNameOptionTableName RPAREN {
        $$ = &Node{}
        $$.Type = AGGREGATION_NODE

        $$.Aggregation = &AggregationNode{}
        $$.Aggregation.Type = AGGREGATION_SUM
        $$.Aggregation.DistinctValid = false
        $$.Aggregation.AttriNameOptionTableName = $3.AttriNameOptionTableName
    }
    |AVG LPAREN DISTINCT attriNameOptionTableName RPAREN {
        $$ = &Node{}
        $$.Type = AGGREGATION_NODE

        $$.Aggregation = &AggregationNode{}
        $$.Aggregation.Type = AGGREGATION_AVG
        $$.Aggregation.DistinctValid = true
        $$.Aggregation.AttriNameOptionTableName = $4.AttriNameOptionTableName
    }
    |AVG LPAREN attriNameOptionTableName RPAREN {
        $$ = &Node{}
        $$.Type = AGGREGATION_NODE

        $$.Aggregation = &AggregationNode{}
        $$.Aggregation.Type = AGGREGATION_AVG
        $$.Aggregation.DistinctValid = false
        $$.Aggregation.AttriNameOptionTableName = $3.AttriNameOptionTableName
    }
    |MIN LPAREN DISTINCT attriNameOptionTableName RPAREN {
        $$ = &Node{}
        $$.Type = AGGREGATION_NODE

        $$.Aggregation = &AggregationNode{}
        $$.Aggregation.Type = AGGREGATION_MIN
        $$.Aggregation.DistinctValid = true
        $$.Aggregation.AttriNameOptionTableName = $4.AttriNameOptionTableName
    }
    |MIN LPAREN attriNameOptionTableName RPAREN {
        $$ = &Node{}
        $$.Type = AGGREGATION_NODE

        $$.Aggregation = &AggregationNode{}
        $$.Aggregation.Type = AGGREGATION_MIN
        $$.Aggregation.DistinctValid = false
        $$.Aggregation.AttriNameOptionTableName = $3.AttriNameOptionTableName
    }
    |MAX LPAREN DISTINCT attriNameOptionTableName RPAREN {
        $$ = &Node{}
        $$.Type = AGGREGATION_NODE

        $$.Aggregation = &AggregationNode{}
        $$.Aggregation.Type = AGGREGATION_MAX
        $$.Aggregation.DistinctValid = true
        $$.Aggregation.AttriNameOptionTableName = $4.AttriNameOptionTableName
    }
    |MAX LPAREN attriNameOptionTableName RPAREN {
        $$ = &Node{}
        $$.Type = AGGREGATION_NODE

        $$.Aggregation = &AggregationNode{}
        $$.Aggregation.Type = AGGREGATION_MAX
        $$.Aggregation.DistinctValid = false
        $$.Aggregation.AttriNameOptionTableName = $3.AttriNameOptionTableName
    }
    |COUNT LPAREN DISTINCT attriNameOptionTableName RPAREN {
        $$ = &Node{}
        $$.Type = AGGREGATION_NODE

        $$.Aggregation = &AggregationNode{}
        $$.Aggregation.Type = AGGREGATION_COUNT
        $$.Aggregation.DistinctValid = true
        $$.Aggregation.AttriNameOptionTableName = $4.AttriNameOptionTableName
    }
    |COUNT LPAREN attriNameOptionTableName RPAREN {
        $$ = &Node{}
        $$.Type = AGGREGATION_NODE

        $$.Aggregation = &AggregationNode{}
        $$.Aggregation.Type = AGGREGATION_COUNT
        $$.Aggregation.DistinctValid = false
        $$.Aggregation.AttriNameOptionTableName = $3.AttriNameOptionTableName
    }
    |COUNT LPAREN STAR RPAREN {
        $$ = &Node{}
        $$.Type = AGGREGATION_NODE

        $$.Aggregation = &AggregationNode{}
        $$.Aggregation.Type = AGGREGATION_COUNT_ALL
        $$.Aggregation.DistinctValid = false
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