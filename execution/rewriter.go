package execution

import (
	"ZetaDB/container"
	"ZetaDB/parser"
	"errors"
	"strconv"
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

//turn a conditionNode in AST to a Condition struct
func (rw *Rewriter) ConditionNodeToCondition(conditionNode *parser.ConditionNode, tableSchema *container.Schema) (*container.Condition, error) {
	newCondition := &container.Condition{}

	switch conditionNode.Type {
	case parser.CONDITION_PREDICATE:
		newCondition.ConditionType = container.CONDITION_PREDICATE
		newPredicate, err := rw.PredicateNodeToPredicate(conditionNode.Predicate, tableSchema)
		if err != nil {
			return nil, err
		} else {
			newCondition.Predicate = newPredicate
		}
	case parser.CONDITION_AND:
		newCondition.ConditionType = container.CONDITION_AND
		newConditionL, errL := rw.ConditionNodeToCondition(conditionNode.ConditionL, tableSchema)
		if errL != nil {
			return nil, errL
		}
		newConditionR, errR := rw.ConditionNodeToCondition(conditionNode.ConditionR, tableSchema)
		if errR != nil {
			return nil, errR
		}
		newCondition.ConditionL = newConditionL
		newCondition.ConditionR = newConditionR

	case parser.CONDITION_OR:
		newCondition.ConditionType = container.CONDITION_OR
		newConditionL, errL := rw.ConditionNodeToCondition(conditionNode.ConditionL, tableSchema)
		if errL != nil {
			return nil, errL
		}
		newConditionR, errR := rw.ConditionNodeToCondition(conditionNode.ConditionR, tableSchema)
		if errR != nil {
			return nil, errR
		}
		newCondition.ConditionL = newConditionL
		newCondition.ConditionR = newConditionR
	}

	return newCondition, nil
}

//turn a predicateNode in AST to a predicate struct
//TODO unfinished RightAttributeIndex, AttributeIndexList, Relation
func (rw *Rewriter) PredicateNodeToPredicate(predicateNode *parser.PredicateNode, tableSchema *container.Schema) (*container.Predicate, error) {
	newPredicate := &container.Predicate{}

	//set PredicateType
	switch predicateNode.Type {
	case parser.PREDICATE_COMPARE_ELEMENTARY_VALUE:
		newPredicate.PredicateType = 1
	case parser.PREDICATE_LIKE_STRING_VALUE:
		newPredicate.PredicateType = 2
	case parser.PREDICATE_IN_SUBQUERY:
		newPredicate.PredicateType = 3
	case parser.PREDICATE_NOT_IN_SUBQUERY:
		newPredicate.PredicateType = 4
	case parser.PREDICATE_IN_TABLE:
		newPredicate.PredicateType = 3
	case parser.PREDICATE_NOT_IN_TABLE:
		newPredicate.PredicateType = 4
	case parser.PREDICATE_COMPARE_ALL_SUBQUERY:
		newPredicate.PredicateType = 5
	case parser.PREDICATE_COMPARE_NOT_ALL_SUBQUERY:
		newPredicate.PredicateType = 6
	case parser.PREDICATE_COMPARE_ANY_SUBQUERY:
		newPredicate.PredicateType = 7
	case parser.PREDICATE_COMPARE_NOT_ANY_SUBQUERY:
		newPredicate.PredicateType = 8
	case parser.PREDICATE_COMPARE_ALL_TABLE:
		newPredicate.PredicateType = 5
	case parser.PREDICATE_COMPARE_NOT_ALL_TABLE:
		newPredicate.PredicateType = 6
	case parser.PREDICATE_COMPARE_ANY_TABLE:
		newPredicate.PredicateType = 7
	case parser.PREDICATE_COMPARE_NOT_ANY_TABLE:
		newPredicate.PredicateType = 8
	case parser.PREDICATE_IS_NULL:
		newPredicate.PredicateType = 9
	case parser.PREDICATE_IS_NOT_NULL:
		newPredicate.PredicateType = 10
	case parser.PREDICATE_TUPLE_IN_SUBQUERY:
		newPredicate.PredicateType = 11
	case parser.PREDICATE_TUPLE_NOT_IN_SUBQUERY:
		newPredicate.PredicateType = 12
	case parser.PREDICATE_TUPLE_IN_TABLE:
		newPredicate.PredicateType = 11
	case parser.PREDICATE_TUPLE_NOT_IN_TABLE:
		newPredicate.PredicateType = 12
	case parser.PREDICATE_SUBQUERY_EXISTS:
		newPredicate.PredicateType = 13
	case parser.PREDICATE_SUBQUERY_NOT_EXISTS:
		newPredicate.PredicateType = 14
	}

	//set CompareMark
	switch predicateNode.CompareMark {
	case parser.COMPAREMARK_EQUAL:
		newPredicate.CompareMark = 1
	case parser.COMPAREMARK_NOTEQUAL:
		newPredicate.CompareMark = 2
	case parser.COMPAREMARK_LESS:
		newPredicate.CompareMark = 3
	case parser.COMPAREMARK_GREATER:
		newPredicate.CompareMark = 4
	case parser.COMPAREMARK_LESSEQUAL:
		newPredicate.CompareMark = 5
	case parser.COMPAREMARK_GREATEREQUAL:
		newPredicate.CompareMark = 6
	}

	//set CompareValueType, CompareIntValue, CompareFloatValue, CompareStringValue, CompareBooleanValue
	switch predicateNode.ElementaryValue.Type {
	case parser.ELEMENTARY_VALUE_INT:
		newPredicate.CompareValueType = 1
		newPredicate.CompareIntValue = predicateNode.ElementaryValue.IntValue
	case parser.ELEMENTARY_VALUE_FLOAT:
		newPredicate.CompareValueType = 2
		newPredicate.CompareFloatValue = predicateNode.ElementaryValue.FloatValue
	case parser.ELEMENTARY_VALUE_STRING:
		newPredicate.CompareValueType = 3
		newPredicate.CompareStringValue = predicateNode.ElementaryValue.StringValue
	case parser.ELEMENTARY_VALUE_BOOLEAN:
		newPredicate.CompareValueType = 4
		newPredicate.CompareBooleanValue = predicateNode.ElementaryValue.BooleanValue
	}

	//set LeftAttributeIndex
	comparedDomainName := predicateNode.AttriNameWithTableNameL.AttributeName
	for i, domain := range tableSchema.GetSchemaDomains() {
		if domain.GetDomainName() == comparedDomainName {
			newPredicate.LeftAttributeIndex = i
			return newPredicate, nil
		}
	}

	return nil, errors.New("rewriter.go    PredicateNodeToPredicate() can not find LeftAttributeIndex")
}

func (rw *Rewriter) ASTNodeToPhysicalPlan(userId int32, astNode *parser.ASTNode, sqlString string) *container.ExecutionPlan {
	switch astNode.Type {
	case parser.AST_DDL: //DDL
		switch astNode.Ddl.Type {
		case parser.DDL_TABLE_CREATE:
			var parameter []string
			parameter = append(parameter, strconv.Itoa(int(userId)))
			parameter = append(parameter, sqlString)
			return container.NewExecutionPlan(container.EP_CREATE_TABLE, parameter, nil)
		case parser.DDL_TABLE_DROP:
			var parameter []string
			parameter = append(parameter, astNode.Ddl.Table.TableName)
			return container.NewExecutionPlan(container.EP_DROP_TABLE, parameter, nil)
		case parser.DDL_TABLE_ALTER_ADD:
		case parser.DDL_TABLE_ALTER_DROP:
		case parser.DDL_ASSERT_CREATE:
		case parser.DDL_ASSERT_DROP:
		case parser.DDL_VIEW_CREATE:
		case parser.DDL_VIEW_DROP:
		case parser.DDL_INDEX_CREATE:
		case parser.DDL_INDEX_DROP:
		case parser.DDL_TRIGGER_CREATE:
		case parser.DDL_TRIGGER_DROP:
		case parser.DDL_PSM_CREATE:
		case parser.DDL_PSM_DROP:
		}
	case parser.AST_DML: //DML
		switch astNode.Dml.Type {
		case parser.DML_INSERT: //TODO
			var parameter []string

			//insert tableName first
			parameter = append(parameter, astNode.Dml.Insert.TableName)

			for _, value := range astNode.Dml.Insert.ElementaryValueList {
				switch value.Type {
				case parser.ELEMENTARY_VALUE_INT:
					parameter = append(parameter, strconv.Itoa(value.IntValue))
				case parser.ELEMENTARY_VALUE_FLOAT:
					parameter = append(parameter, strconv.FormatFloat(value.FloatValue, 'f', -1, 64))
				case parser.ELEMENTARY_VALUE_STRING:
					parameter = append(parameter, value.StringValue)
				case parser.ELEMENTARY_VALUE_BOOLEAN:
					if value.BooleanValue {
						parameter = append(parameter, "TRUE")
					} else {
						parameter = append(parameter, "FALSE")
					}
				}
			}

			return container.NewExecutionPlan(container.EP_INSERT, parameter, nil)
		case parser.DML_UPDATE:
		case parser.DML_DELETE:
			//insert tableName into parameter
			var parameter []string
			parameter = append(parameter, astNode.Dml.Delete.TableName)

			//create logicalPlan
			ktm := GetKeytableManager()
			tableId, tableSchema, _ := ktm.Query_k_tableId_schema_FromTableName(astNode.Dml.Delete.TableName)
			headPageId, _, _, _, _ := ktm.Query_k_table(tableId)

			Condition, _ := rw.ConditionNodeToCondition(astNode.Dml.Delete.Condition, tableSchema)

			leftNodeOfRoot := &container.LogicalPlanNode{
				NodeType:        container.SEQUENTIAL_FILE_READER,
				TableHeadPageId: headPageId,
				Schema:          tableSchema,
			}
			logicalPlanRoot := &container.LogicalPlanNode{
				NodeType:  container.SELECTION,
				Condition: Condition,
				LeftNode:  leftNodeOfRoot,
			}

			return container.NewExecutionPlan(container.EP_DELETE, parameter, logicalPlanRoot)
		}
	case parser.AST_DCL: //DCL
	case parser.AST_DQL: //DQL

	}
	return nil
}
