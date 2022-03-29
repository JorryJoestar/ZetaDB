package execution

import (
	"ZetaDB/container"
	"ZetaDB/optimizer"
	"ZetaDB/parser"
	"ZetaDB/storage"
	"sync"
)

type ExecutionEngine struct {
	se       *storage.StorageEngine
	parser   *parser.Parser
	rewriter *optimizer.Rewriter
	ktm      *KeytableManager
	dpm      *DataPageManipulator
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
	eeInstance.dpm = NewDataPageManipulator()
	eeInstance.ktm = GetKeytableManager()

	return eeInstance
}

//initialze the whole system, create key tables and insert necessary tuples into these tables
func (ee *ExecutionEngine) InitializeSystem() {
	ee.ktm.InitializeSystem()
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
