package main

import (
	//. "ZetaDB/execution/querySubOperator"

	"ZetaDB/container"
	"ZetaDB/execution"
	"ZetaDB/parser"
	"ZetaDB/storage"
	"sync"
)

type Kernel struct {
	parser          *parser.Parser
	storageEngine   *storage.StorageEngine
	executionEngine *execution.ExecutionEngine
	transaction     *container.Transaction
}

//for singleton pattern
var instance *Kernel
var once sync.Once

//to get kernel, call this function
func GetInstance() *Kernel {
	once.Do(func() {
		instance = &Kernel{
			parser:        parser.GetParser(),
			storageEngine: storage.GetStorageEngine(),
			transaction:   container.GetTransaction()}
	})
	instance.executionEngine = execution.GetExecutionEngine()
	return instance
}

func main() {

}
