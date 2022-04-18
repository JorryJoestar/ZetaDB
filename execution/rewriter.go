package execution

import (
	"ZetaDB/container"
	"ZetaDB/parser"
	pp "ZetaDB/physicalPlan"
	. "ZetaDB/utility"
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

func (rw *Rewriter) ASTNodeToExecutionPlan(userId int32, astNode *parser.ASTNode, sqlString string) (*container.ExecutionPlan, error) {
	switch astNode.Type {

	//DDL
	case parser.AST_DDL:
		switch astNode.Ddl.Type {

		//CREATE TABLE
		case parser.DDL_TABLE_CREATE:
			//error: log in first
			if userId == -1 {
				return nil, errors.New("error: log in first")
			}

			//error: table already exist
			//TODO

			var parameter []string
			parameter = append(parameter, strconv.Itoa(int(userId)))
			parameter = append(parameter, sqlString)
			return container.NewExecutionPlan(container.EP_CREATE_TABLE, userId, parameter, nil), nil

		//DROP TABLE
		case parser.DDL_TABLE_DROP:
			var parameter []string
			parameter = append(parameter, astNode.Ddl.Table.TableName)
			return container.NewExecutionPlan(container.EP_DROP_TABLE, userId, parameter, nil), nil

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

	//DML
	case parser.AST_DML:
		switch astNode.Dml.Type {

		//INSERT INTO TABLE
		case parser.DML_INSERT:
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

			return container.NewExecutionPlan(container.EP_INSERT, userId, parameter, nil), nil

		//UPDATE TABLE
		case parser.DML_UPDATE:
			ktm := GetKeytableManager()
			tableId, tableSchema, _ := ktm.Query_k_tableId_schema_FromTableName(astNode.Dml.Update.TableName)

			var parameter []string
			parameter = append(parameter, astNode.Dml.Update.TableName)

			//insert domain index and values to update
			for _, updataListEntryNode := range astNode.Dml.Update.UpdateList {
				var domainIndex int
				for i, domain := range tableSchema.GetSchemaDomains() {
					if domain.GetDomainName() == updataListEntryNode.AttributeName {
						domainIndex = i
						break
					}
				}
				parameter = append(parameter, strconv.Itoa(domainIndex))

				switch updataListEntryNode.ElementaryValue.Type {
				case parser.ELEMENTARY_VALUE_INT:
					parameter = append(parameter, strconv.Itoa(updataListEntryNode.ElementaryValue.IntValue))
				case parser.ELEMENTARY_VALUE_FLOAT:
					parameter = append(parameter, strconv.FormatFloat(updataListEntryNode.ElementaryValue.FloatValue, 'f', -1, 64))
				case parser.ELEMENTARY_VALUE_STRING:
					parameter = append(parameter, updataListEntryNode.ElementaryValue.StringValue)
				case parser.ELEMENTARY_VALUE_BOOLEAN:
					if updataListEntryNode.ElementaryValue.BooleanValue {
						parameter = append(parameter, "TRUE")
					} else {
						parameter = append(parameter, "FALSE")
					}
				}
			}

			//create logicalPlan
			headPageId, _, _, _, _ := ktm.Query_k_table(tableId)

			Condition, conditionErr := rw.ConditionNodeToCondition(astNode.Dml.Update.Condition, tableSchema)

			if conditionErr != nil {
				return nil, errors.New("error: condition invalid")
			}

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

			return container.NewExecutionPlan(container.EP_UPDATE, userId, parameter, logicalPlanRoot), nil

		//DELETE FROM TABLE
		case parser.DML_DELETE:
			//insert tableName into parameter
			var parameter []string
			parameter = append(parameter, astNode.Dml.Delete.TableName)

			//create logicalPlan
			ktm := GetKeytableManager()
			tableId, tableSchema, _ := ktm.Query_k_tableId_schema_FromTableName(astNode.Dml.Delete.TableName)
			headPageId, _, _, _, _ := ktm.Query_k_table(tableId)

			Condition, conditionErr := rw.ConditionNodeToCondition(astNode.Dml.Delete.Condition, tableSchema)

			if conditionErr != nil {
				return nil, errors.New("error: condition invalid")
			}

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

			return container.NewExecutionPlan(container.EP_DELETE, userId, parameter, logicalPlanRoot), nil
		}

	//DCL
	case parser.AST_DCL:
		switch astNode.Dcl.Type {
		case parser.DCL_TRANSACTION_BEGIN:
		case parser.DCL_TRANSACTION_COMMIT:
		case parser.DCL_TRANSACTION_ROLLBACK:

		//SHOW TABLES
		case parser.DCL_SHOW_TABLES:

		case parser.DCL_SHOW_ASSERTIONS:
		case parser.DCL_SHOW_VIEWS:
		case parser.DCL_SHOW_INDEXS:
		case parser.DCL_SHOW_TRIGGERS:
		case parser.DCL_SHOW_FUNCTIONS:
		case parser.DCL_SHOW_PROCEDURES:

		//CREATE USER
		case parser.DCL_CREATE_USER:
			//get userName and password
			userName := astNode.Dcl.UserName
			password := astNode.Dcl.Password

			//insert userName and password into parameter
			var parameter []string
			parameter = append(parameter, userName)
			parameter = append(parameter, password)

			return container.NewExecutionPlan(container.EP_CREATE_USER, userId, parameter, nil), nil

		//LOG IN USER
		case parser.DCL_LOG_USER:
			//get userName and password
			userName := astNode.Dcl.UserName
			password := astNode.Dcl.Password

			//insert userName and password into parameter
			var parameter []string
			parameter = append(parameter, userName)
			parameter = append(parameter, password)

			return container.NewExecutionPlan(container.EP_LOG_USER, userId, parameter, nil), nil
		case parser.DCL_PSMCALL:

		//INITIALIZE
		case parser.DCL_INIT:
			//check if current user is administor
			if userId != 0 {
				return nil, errors.New("error: current user is not administor")
			}

			return container.NewExecutionPlan(container.EP_INIT, userId, nil, nil), nil

		//DROP USER
		case parser.DCL_DROP_USER:
			//error: current user is not administor
			if userId != 0 {
				return nil, errors.New("error: current user is not administor")
			}

			//push userName to drop into parameter
			var parameter []string
			parameter = append(parameter, astNode.Dcl.UserName)

			return container.NewExecutionPlan(container.EP_DROP_USER, userId, parameter, nil), nil
		case parser.DCL_HALT:
			//check if current user is administor
			if userId != 0 {
				return nil, errors.New("error: current user is not administor")
			}

			return container.NewExecutionPlan(container.EP_HALT, userId, nil, nil), nil
		}

	//DQL
	case parser.AST_DQL:
		switch astNode.Dql.Type {

		//SINGLE QUERY
		case parser.DQL_SINGLE_QUERY:

			//error: log in first
			if userId == -1 {
				return nil, errors.New("error: log in first")
			}

			var logicalPlanRoot *container.LogicalPlanNode

			ktm := GetKeytableManager()
			//get tableId, schema, headPageId
			tableName := astNode.Dql.Query.FromList[0].TableName
			tableId, schema, tableExistErr := ktm.Query_k_tableId_schema_FromTableName(tableName)
			if tableExistErr != nil { //check if tableName exist
				return nil, errors.New("error: no such table")
			}
			headPageId, _, _, _, _ := ktm.Query_k_table(tableId)

			//check if this table belongs to the user
			//administrator has rights to query all tables
			if userId != 0 {
				schema2 := ktm.GetKeyTableSchema(2)
				seqItCheckOwner := pp.NewSequentialFileReaderIterator(2, schema2)
				seqItCheckOwner.Open(nil, nil)
				for seqItCheckOwner.HasNext() {
					tup, _ := seqItCheckOwner.GetNext()
					tupTableIdBytes, _ := tup.TupleGetFieldValue(0)
					tupUserIdBytes, _ := tup.TupleGetFieldValue(1)
					tupTableId, _ := BytesToINT(tupTableIdBytes)
					tupUserId, _ := BytesToINT(tupUserIdBytes)
					if uint32(tupTableId) == tableId && tupUserId != userId {
						return nil, errors.New("error: lack table ownership")
					}
				}
			}

			//build logicalPlan tree
			logicalPlanFromFile := &container.LogicalPlanNode{
				NodeType:        container.SEQUENTIAL_FILE_READER,
				TableHeadPageId: headPageId,
				Schema:          schema,
			}

			if astNode.Dql.Query.WhereValid { //where clause valid

				Condition, conditionErr := rw.ConditionNodeToCondition(astNode.Dql.Query.WhereCondition, schema)

				if conditionErr != nil {
					return nil, errors.New("error: condition invalid")
				}

				logicalPlanRoot = &container.LogicalPlanNode{
					NodeType:  container.SELECTION,
					LeftNode:  logicalPlanFromFile,
					Condition: Condition,
				}

			} else {
				logicalPlanRoot = logicalPlanFromFile
			}

			return container.NewExecutionPlan(container.EP_QUERY, userId, nil, logicalPlanRoot), nil
		case parser.DQL_UNION:
		case parser.DQL_DIFFERENCE:
		case parser.DQL_INTERSECTION:
		}

	}
	return nil, nil
}

//generate a physicalPlan from a logicalPlan
func (rw *Rewriter) LogicalPLanToPhysicalPlan(logicalPlanRootNode *container.LogicalPlanNode) *pp.PhysicalPlan {
	return pp.NewPhysicalPlan(rw.LogicalPlanNodeToIterator(logicalPlanRootNode))
}

//TODO unifinished
func (rw *Rewriter) LogicalPlanNodeToIterator(logicalPlanNode *container.LogicalPlanNode) pp.Iterator {
	//recursive ending
	if logicalPlanNode == nil {
		return nil
	}

	var returnIterator pp.Iterator

	iterator1 := rw.LogicalPlanNodeToIterator(logicalPlanNode.LeftNode)
	iterator2 := rw.LogicalPlanNodeToIterator(logicalPlanNode.RightNode)

	switch logicalPlanNode.NodeType {
	case container.BAG_DIFFERENCE:
	case container.BAG_INTERSECTION:
	case container.BAG_UNION:
	case container.DUPLICATE_ELIMINATION:
	case container.GROUPING:
	case container.INDEX_FILEREADER:
	case container.NATURAL_JOIN:
	case container.PRODUCT:
	case container.PROJECTION:
	case container.RENAME:
	case container.SELECTION:
		returnIterator = pp.NewSelectionIterator(logicalPlanNode.Condition)
	case container.SEQUENTIAL_FILE_READER:
		returnIterator = pp.NewSequentialFileReaderIterator(logicalPlanNode.TableHeadPageId, logicalPlanNode.Schema)
	case container.SET_DIFFERENCE:
	case container.SET_INTERSECTION:
	case container.SET_UNION:
	case container.THETA_JOIN:
	}

	//open this iterator, ready for useage
	returnIterator.Open(iterator1, iterator2)

	return returnIterator
}

func (rw *Rewriter) TupleFieldsToStrings(tuple *container.Tuple) []string {
	var fieldStrings []string

	tableId := tuple.TupleGetTableId()

	ktm := GetKeytableManager()
	schema, _ := ktm.Query_k_tableId_schema_FromTableId(tableId)
	domains := schema.GetSchemaDomains()

	for fieldIndex, field := range tuple.TupleGetFields() {
		fieldBytes, nullErr := field.FieldToBytes()

		var fieldString string

		switch domains[fieldIndex].GetDomainType() {
		case container.CHAR:

			if nullErr == nil { //not null
				fieldString, _ = BytesToCHAR(fieldBytes)
			} else { //null
				fieldString = "NULL"
			}

		case container.VARCHAR:

			if nullErr == nil { //not null
				fieldString, _ = BytesToVARCHAR(fieldBytes)
			} else { //null
				fieldString = "NULL"
			}

		case container.BIT:

			if nullErr == nil { //not null
				fieldString = BytesToHexString(fieldBytes)
			} else { //null
				fieldString = "NULL"
			}

		case container.BITVARYING:

			if nullErr == nil { //not null
				fieldString = BytesToHexString(fieldBytes)
			} else { //null
				fieldString = "NULL"
			}

		case container.BOOLEAN:

			if nullErr == nil { //not null
				boolValue := ByteToBOOLEAN(fieldBytes[0])
				if boolValue {
					fieldString = "TRUE"
				} else {
					fieldString = "FALSE"
				}
			} else { //null
				fieldString = "NULL"
			}

		case container.INT:

			if nullErr == nil { //not null
				intValue, _ := BytesToINT(fieldBytes)

				fieldString = strconv.Itoa(int(intValue))

			} else { //null
				fieldString = "NULL"
			}

		case container.INTEGER:

			if nullErr == nil { //not null
				intValue, _ := BytesToInteger(fieldBytes)
				fieldString = strconv.Itoa(int(intValue))
			} else { //null
				fieldString = "NULL"
			}

		case container.SHORTINT:

			if nullErr == nil { //not null
				int16Value, _ := BytesToSHORTINT(fieldBytes)
				fieldString = strconv.Itoa(int(int16Value))
			} else { //null
				fieldString = "NULL"
			}

		case container.FLOAT:

			if nullErr == nil { //not null
				floatValue, _ := BytesToFLOAT(fieldBytes)
				float64Value := float64(floatValue)
				fieldString = strconv.FormatFloat(float64Value, 'E', -1, 64)
			} else { //null
				fieldString = "NULL"
			}

		case container.REAL:

			if nullErr == nil { //not null
				floatValue, _ := BytesToREAL(fieldBytes)
				float64Value := float64(floatValue)
				fieldString = strconv.FormatFloat(float64Value, 'E', -1, 64)
			} else { //null
				fieldString = "NULL"
			}

		case container.DOUBLEPRECISION:

			if nullErr == nil { //not null
				floatValue, _ := BytesToDOUBLEPRECISION(fieldBytes)
				fieldString = strconv.FormatFloat(floatValue, 'E', -1, 64)
			} else { //null
				fieldString = "NULL"
			}

		case container.DECIMAL:

			if nullErr == nil { //not null
				d, _ := domains[fieldIndex].GetDomainD()
				n, _ := domains[fieldIndex].GetDomainN()
				floatValue, _ := BytesToDECIMAL(fieldBytes, int(n), int(d))
				fieldString = strconv.FormatFloat(floatValue, 'E', -1, 64)
			} else { //null
				fieldString = "NULL"
			}

		case container.NUMERIC:

			if nullErr == nil { //not null
				d, _ := domains[fieldIndex].GetDomainD()
				n, _ := domains[fieldIndex].GetDomainN()
				floatValue, _ := BytesToNUMERIC(fieldBytes, int(n), int(d))
				fieldString = strconv.FormatFloat(floatValue, 'E', -1, 64)
			} else { //null
				fieldString = "NULL"
			}

		case container.DATE:

			if nullErr == nil { //not null
				fieldString, _ = BytesToDATE(fieldBytes)
			} else { //null
				fieldString = "NULL"
			}

		case container.TIME:

			if nullErr == nil { //not null
				fieldString, _ = BytesToTIME(fieldBytes)
			} else { //null
				fieldString = "NULL"
			}

		}

		fieldStrings = append(fieldStrings, fieldString)
	}

	return fieldStrings
}
