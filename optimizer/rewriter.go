package optimizer

import (
	"ZetaDB/container"
	"ZetaDB/parser"
	"errors"
	"sync"
)

type Rewriter struct {
}

//for singleton pattern
var instance *Rewriter
var once sync.Once

//to get Rewriter, call this function
func GetRewriter() *Rewriter {
	once.Do(func() {
		instance = &Rewriter{}
	})
	return instance
}

//throw error if ASTNode type is not AST_DDL
//throw error if DDLNode type is not DDL_TABLE_CREATE
//throw error if len of AttributeNameList is not equal to len of DomainList
func (rw *Rewriter) ASTNodeToSchema(schemaAst *parser.ASTNode) (*container.Schema, error) {
	//throw error if ASTNode type is not AST_DDL
	if schemaAst.Type != parser.AST_DDL {
		return nil, errors.New("execution/executionEngine.go    ASTNodeToSchema() ASTNode type invalid")
	}

	//throw error if DDLNode type is not DDL_TABLE_CREATE
	if schemaAst.Ddl.Type != parser.DDL_TABLE_CREATE {
		return nil, errors.New("execution/executionEngine.go    ASTNodeToSchema() DDLNode type invalid")
	}

	//throw error if len of AttributeNameList is not equal to len of DomainList
	if len(schemaAst.Ddl.Table.DomainList) != len(schemaAst.Ddl.Table.AttributeNameList) {
		return nil, errors.New("execution/executionEngine.go    ASTNodeToSchema() list length dismatch")
	}

	//get tableName
	tableName := schemaAst.Ddl.Table.TableName

	//get attributeNameList
	var attributeNameList []string
	for _, v := range schemaAst.Ddl.Table.AttributeNameList {
		attributeNameList = append(attributeNameList, v)
	}

	//generate domainList
	var domainList []*container.Domain
	for i, v := range schemaAst.Ddl.Table.DomainList {
		domainName := attributeNameList[i]
		var domainType container.DomainType
		var intN int32
		var intD int32
		switch v.Type {
		case parser.DOMAIN_CHAR:
			domainType = container.CHAR
		case parser.DOMAIN_VARCHAR:
			domainType = container.VARCHAR
			intN = int32(v.N)
		case parser.DOMAIN_BIT:
			domainType = container.BIT
			intN = int32(v.N)
		case parser.DOMAIN_BITVARYING:
			domainType = container.BITVARYING
			intN = int32(v.N)
		case parser.DOMAIN_BOOLEAN:
			domainType = container.BOOLEAN
		case parser.DOMAIN_INT:
			domainType = container.INT
		case parser.DOMAIN_INTEGER:
			domainType = container.INTEGER
		case parser.DOMAIN_SHORTINT:
			domainType = container.SHORTINT
		case parser.DOMAIN_FLOAT:
			domainType = container.FLOAT
		case parser.DOMAIN_REAL:
			domainType = container.REAL
		case parser.DOMAIN_DOUBLEPRECISION:
			domainType = container.DOUBLEPRECISION
		case parser.DOMAIN_DECIMAL:
			domainType = container.DECIMAL
			intN = int32(v.N)
			intD = int32(v.D)
		case parser.DOMAIN_NUMERIC:
			domainType = container.NUMERIC
			intN = int32(v.N)
			intD = int32(v.D)
		case parser.DOMAIN_DATE:
			domainType = container.DATE
		case parser.DOMAIN_TIME:
			domainType = container.TIME
		}

		newDomain, err := container.NewDomain(domainName, domainType, intN, intD)
		if err != nil {
			return nil, err
		}

		domainList = append(domainList, newDomain)
	}

	//generate constraintNameList
	var constraintList []*container.Constraint
	/* 	for _, v := range schemaAst.Ddl.Table.ConstraintList {
		//TODO
	} */

	newSchema, err := container.NewSchema(tableName, domainList, constraintList)
	if err != nil {
		return nil, err
	}

	return newSchema, nil
}
