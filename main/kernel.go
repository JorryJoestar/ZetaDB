package main

import (
	"ZetaDB/execution"
	"ZetaDB/parser"
	"ZetaDB/storage"
	"sync"
)

type Kernel struct {
	parser          *parser.Parser
	storageEngine   *storage.StorageEngine
	executionEngine *execution.ExecutionEngine
}

//for singleton pattern
var instance *Kernel
var once sync.Once

//to get kernel, call this function
func GetInstance() *Kernel {
	once.Do(func() {
		instance = &Kernel{
			parser:        parser.GetParser(),
			storageEngine: storage.GetStorageEngine()}
	})
	instance.executionEngine = execution.GetExecutionEngine()
	return instance
}

func main() {


	transaction := storage.GetTransaction()

	transaction.PushTransactionIntoDisk()

	
	PrintKeyTable(0)
}

/* func getNewTuple(tupleId uint32, userId int32, userName string) *container.Tuple {
	ktm := execution.GetKeytableManager()

	schema0 := ktm.GetKeyTableSchema(0)

	field0, _ := container.NewFieldFromBytes(utility.INTToBytes(userId))
	fieldsByte, _ := utility.VARCHARToBytes(userName)
	field1, _ := container.NewFieldFromBytes(fieldsByte)
	var fields []*container.Field
	fields = append(fields, field0)
	fields = append(fields, field1)
	tuple, _ := container.NewTuple(0, tupleId, schema0, fields)
	return tuple
} */
