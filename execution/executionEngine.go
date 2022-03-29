package execution

import (
	"ZetaDB/container"
	its "ZetaDB/execution/querySubOperator"
	"ZetaDB/optimizer"
	"ZetaDB/parser"
	"ZetaDB/storage"
	. "ZetaDB/utility"
	"errors"
	"sync"
)

type ExecutionEngine struct {
	se                    *storage.StorageEngine
	parser                *parser.Parser
	rewriter              *optimizer.Rewriter
	initializationManager *InitializationManager
}

//use GetExecutionEngine to get the unique ExecutionEngine
var eeInstance *ExecutionEngine
var eeOnce sync.Once

func GetExecutionEngine(se *storage.StorageEngine, parser *parser.Parser) *ExecutionEngine {

	eeOnce.Do(func() {
		eeInstance = &ExecutionEngine{
			se:       se,
			parser:   parser,
			rewriter: &optimizer.Rewriter{}}
	})
	eeInstance.initializationManager = NewInitializationManager(eeInstance.se, eeInstance.parser, eeInstance.rewriter)

	return eeInstance
}

//initialze the whole system, create key tables and insert necessary tuples into these tables
func (ee *ExecutionEngine) InitializeSystem() {
	ee.initializationManager.InitializeSystem()
}

//fetch schema of a table from dataFile according to tableName, k_tableId_schema table 8
//throw error if no such table
func (ee *ExecutionEngine) GetSchemaFromFileByTableName(tableName string) (*container.Schema, error) {
	astOfCreateTable8 := ee.parser.ParseSql(DEFAULT_KEYTABLES_SCHEMA[8])
	schemaOfCreateTable8, err := ee.rewriter.ASTNodeToSchema(astOfCreateTable8)
	if err != nil {
		return nil, err
	}

	seqIt := its.NewSequentialFileReaderIterator(ee.se, 8, schemaOfCreateTable8)
	seqIt.Open(nil, nil)
	for seqIt.HasNext() {
		table8Tuple, err := seqIt.GetNext()
		if err != nil {
			return nil, err
		}

		//TODO mantain: if schema of table 8 is changed, this index could be asked to change
		schemaStringBytes, err := table8Tuple.TupleGetFieldValue(1)
		if err != nil {
			return nil, err
		}
		schemaString, err := BytesToVARCHAR(schemaStringBytes)
		if err != nil {
			return nil, err
		}

		//parse this string to get schema
		ast := ee.parser.ParseSql(schemaString)
		currentSchema, err := ee.rewriter.ASTNodeToSchema(ast)
		if err != nil {
			return nil, err
		}

		//found correct table schema
		if currentSchema.GetSchemaTableName() == tableName {
			return currentSchema, nil
		}

	}
	return nil, errors.New("execution/executionEngine.go    GetSchemaFromFileByTableName() no such table")
}

//insert a tuple into a table, if no enough space, then create a new page
//throw error if no such table
func (ee *ExecutionEngine) InsertTupleIntoTable(tuple *container.Tuple, tableId uint8) error {
	return nil
}

//TODO
func (ee *ExecutionEngine) CreateTableOperator()    {}
func (ee *ExecutionEngine) DropTableOperator()      {}
func (ee *ExecutionEngine) AlterTableAddOperator()  {}
func (ee *ExecutionEngine) AlterTableDropOperator() {}
func (ee *ExecutionEngine) CreateAssertOperator()   {}
func (ee *ExecutionEngine) DropAssertOperator()     {}
func (ee *ExecutionEngine) CreateViewOperator()     {}
func (ee *ExecutionEngine) DropViewOperator()       {}
func (ee *ExecutionEngine) CreateIndexOperator()    {}
func (ee *ExecutionEngine) DropIndexOperator()      {}
func (ee *ExecutionEngine) CreateTriggerOperator()  {}
func (ee *ExecutionEngine) DropTriggerOperator()    {}
func (ee *ExecutionEngine) CreatePsmOperator()      {}
func (ee *ExecutionEngine) DropPsmOperator()        {}
func (ee *ExecutionEngine) DeleteOperator()         {}
func (ee *ExecutionEngine) InsertOperator()         {}
func (ee *ExecutionEngine) UpdateOperator()         {}
func (ee *ExecutionEngine) ShowTablesOperator()     {}
func (ee *ExecutionEngine) ShowAssertionsOperator() {}
func (ee *ExecutionEngine) ShowViewsOperator()      {}
func (ee *ExecutionEngine) ShowIndexsOperator()     {}
func (ee *ExecutionEngine) ShowTriggersOperator()   {}
func (ee *ExecutionEngine) ShowFunctionsOperator()  {}
func (ee *ExecutionEngine) ShowProceduresOperator() {}
func (ee *ExecutionEngine) CreateUserOperator()     {}
func (ee *ExecutionEngine) LogUserOperator()        {}
func (ee *ExecutionEngine) PsmCallOperator()        {}

func (ee *ExecutionEngine) QueryOperator(pp *container.PhysicalPlan) {

}

func (ee *ExecutionEngine) GetSchemaByTableId() {}

/* func (ee *ExecutionEngine) GetKeyTableSchema(tableId uint32) (*container.Schema, error) {
	if tableId < 0 || tableId > 16 {
		return nil, errors.New("execution/executionEngine.go    GetKeyTableSchema() tableId invalid")
	}
	var schemaAst *parser.ASTNode
	switch tableId {
	case 0:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_0_SCHEMA)
	case 1:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_1_SCHEMA)
	case 2:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_2_SCHEMA)
	case 3:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_3_SCHEMA)
	case 4:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_4_SCHEMA)
	case 5:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_5_SCHEMA)
	case 6:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_6_SCHEMA)
	case 7:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_7_SCHEMA)
	case 8:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_8_SCHEMA)
	case 9:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_9_SCHEMA)
	case 10:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_10_SCHEMA)
	case 11:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_11_SCHEMA)
	case 12:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_12_SCHEMA)
	case 13:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_13_SCHEMA)
	case 14:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_14_SCHEMA)
	case 15:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_15_SCHEMA)
	case 16:
		schemaAst = ee.parser.ParseSql(DEFAULT_KEY_TABLE_16_SCHEMA)
	}

	keyTableSchema, err := ee.ASTNodeToSchema(schemaAst)

	return keyTableSchema, err
} */
