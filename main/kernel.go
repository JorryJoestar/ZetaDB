package main

import (
	"ZetaDB/container"
	"ZetaDB/execution"
	"ZetaDB/parser"
	"ZetaDB/storage"
	"ZetaDB/utility"
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
	//ktm := execution.GetKeytableManager()
	//ee := execution.GetExecutionEngine()
	tm := execution.GetTableManipulator()

	transaction := storage.GetTransaction()

	//ee.CreateTableOperator(0, "create table testTable(id int primary key, name varchar(20));")
	//ktm.InitializeSystem()
	//ee.DropTableOperator("testTable")

	/* 	for i := 1; i <= 1000; i++ {
		newTuple := getNewTuple(uint32(i), int32(i), "ClaireMao")
		tm.InsertTupleIntoTable(17, newTuple)
	} */

	for i := 11; i <= 1000; i++ {
		tm.DeleteTupleFromTable(17, uint32(i))
	}

	transaction.PushTransactionIntoDisk()

	for i := 0; i < 18; i++ {
		PrintTable(uint32(i))
	}
}

func getNewTuple(tupleId uint32, id int32, name string) *container.Tuple {
	ktm := execution.GetKeytableManager()

	schema0 := ktm.GetKeyTableSchema(0)

	field0, _ := container.NewFieldFromBytes(utility.INTToBytes(id))
	fieldsByte, _ := utility.VARCHARToBytes(name)
	field1, _ := container.NewFieldFromBytes(fieldsByte)
	var fields []*container.Field
	fields = append(fields, field0)
	fields = append(fields, field1)
	tuple, _ := container.NewTuple(0, tupleId, schema0, fields)
	return tuple
}
