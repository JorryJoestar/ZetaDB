package execution

import (
	"ZetaDB/container"
	"ZetaDB/parser"
	"ZetaDB/storage"
	"sync"
)

type ExecutionEngine struct {
	se     *storage.StorageEngine
	parser *parser.Parser
}

//use GetExecutionEngine to get the unique ExecutionEngine
var eeInstance *ExecutionEngine
var eeOnce sync.Once

func GetExecutionEngine(se *storage.StorageEngine, parser *parser.Parser) *ExecutionEngine {

	eeOnce.Do(func() {
		eeInstance = &ExecutionEngine{
			se:     se,
			parser: parser}
	})
	return eeInstance
}

func (ee *ExecutionEngine) ExecutePhysicalPlan(pp *container.PhysicalPlan) error {

	return nil
}

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
