package execution

import (
	"ZetaDB/container"
	"ZetaDB/optimizer"
	"ZetaDB/parser"
	"ZetaDB/storage"
	"ZetaDB/utility"
	"sync"
)

type ExecutionEngine struct {
	se       *storage.StorageEngine
	parser   *parser.Parser
	rewriter *optimizer.Rewriter
	ktm      *KeytableManager
	tm       *TableManipulator
}

//use GetExecutionEngine to get the unique ExecutionEngine
var eeInstance *ExecutionEngine
var eeOnce sync.Once

func GetExecutionEngine() *ExecutionEngine {

	eeOnce.Do(func() {
		eeInstance = &ExecutionEngine{
			se:       storage.GetStorageEngine(),
			parser:   parser.GetParser(),
			rewriter: &optimizer.Rewriter{}}
	})
	eeInstance.tm = GetTableManipulator()
	eeInstance.ktm = GetKeytableManager()

	return eeInstance
}

//initialze the whole system, create key tables and insert necessary tuples into these tables
func (ee *ExecutionEngine) InitializeSystem() {
	ee.ktm.InitializeSystem()
}

//create a new table
//insert a tuple into key table 9: k_table
//insert a tuple into key table 2: k_tableId_userId
//insert a tuple into key table 8: k_tableId_schema
//assign an empty headPage for this table
func (ee *ExecutionEngine) CreateTableOperator(userId int32, schemaString string) {
	transaction := storage.GetTransaction()

	//insert a tuple into key table 9: k_table
	//new tableId is lastTupleId+1 in k_table
	_, _, lastTupleId9, _, _ := ee.ktm.Query_k_table(9)
	newTableId := lastTupleId9 + 1
	newHeadPageId := ee.ktm.GetVacantDataPageId()
	ee.ktm.Insert_k_table(newTableId, newHeadPageId, newHeadPageId, 0, 0)

	//update newHeadPage
	newHeadPage := storage.NewDataPageMode0(newHeadPageId, newTableId, newHeadPageId, newHeadPageId)
	transaction.InsertDataPage(newHeadPage)
	ee.se.InsertDataPage(newHeadPage)

	//insert a tuple into key table 2: k_tableId_userId
	_, _, lastTupleId2, _, _ := ee.ktm.Query_k_table(2)
	schema2 := ee.ktm.GetKeyTableSchema(2)
	fields20, _ := container.NewFieldFromBytes(utility.INTToBytes(int32(newTableId)))
	fields21, _ := container.NewFieldFromBytes(utility.INTToBytes(userId))
	var fields2 []*container.Field
	fields2 = append(fields2, fields20)
	fields2 = append(fields2, fields21)
	tuple2, _ := container.NewTuple(2, lastTupleId2+1, schema2, fields2)
	ee.tm.InsertTupleIntoTable(2, tuple2)

	//insert a tuple into key table 8: k_tableId_schema
	_, _, lastTupleId8, _, _ := ee.ktm.Query_k_table(8)
	schema8 := ee.ktm.GetKeyTableSchema(8)
	fields80, _ := container.NewFieldFromBytes(utility.INTToBytes(int32(newTableId)))
	fields81Bytes, _ := utility.VARCHARToBytes(schemaString)
	fields81, _ := container.NewFieldFromBytes(fields81Bytes)
	var fields8 []*container.Field
	fields8 = append(fields8, fields80)
	fields8 = append(fields8, fields81)
	tuple8, _ := container.NewTuple(8, lastTupleId8+1, schema8, fields8)
	ee.tm.InsertTupleIntoTable(8, tuple8)
}

func (ee *ExecutionEngine) DeleteOperator() {}
func (ee *ExecutionEngine) InsertOperator() {}
func (ee *ExecutionEngine) UpdateOperator() {}

//TODO
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
