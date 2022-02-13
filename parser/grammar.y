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

}

%}

%union {
    NodePt *Node
}

%type <NodePt> ast

%token LPAREN

%%
ast
    :LPAREN {
        fmt.Println("success")
        GetInstance().AST = &ASTNode{
            Type: AST_DQL,
	        Ddl: nil,
	        Dml: nil,
	        Dcl: nil,
	        Dql: nil}
    }

%%