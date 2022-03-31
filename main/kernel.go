package main

import (
	//. "ZetaDB/execution/querySubOperator"

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

	/* 	ee := execution.GetExecutionEngine()
	   	ee.InitializeSystem() */

	transaction.PushTransactionIntoDisk()
}
