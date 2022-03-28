package execution

import (
	"ZetaDB/container"
	"ZetaDB/storage"
	"sync"
)

type ExecutionEngine struct {
	se *storage.StorageEngine
}

//use GetExecutionEngine to get the unique ExecutionEngine
var eeInstance *ExecutionEngine
var eeOnce sync.Once

func GetExecutionEngine() *ExecutionEngine {

	eeOnce.Do(func() {
		eeInstance = &ExecutionEngine{}
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
