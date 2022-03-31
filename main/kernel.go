package main

import (
	"ZetaDB/container"
	its "ZetaDB/execution/querySubOperator"

	"ZetaDB/execution"
	"ZetaDB/parser"
	"ZetaDB/storage"
	"fmt"
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

	/* 	ee := execution.GetExecutionEngine()
	   	ee.InitializeSystem() */

	ktm := execution.GetKeytableManager()
	schema, _ := ktm.GetKeyTableSchema(15)

	seqIt := its.NewSequentialFileReaderIterator(15, schema)
	seqIt.Open(nil, nil)

	var tuples []*container.Tuple

	for seqIt.HasNext() {
		tuple, _ := seqIt.GetNext()
		tuples = append(tuples, tuple)
	}

	result := TableToString(schema, tuples)
	fmt.Println(result)

	transaction.PushTransactionIntoDisk()
}
