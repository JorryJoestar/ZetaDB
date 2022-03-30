package main

import (
	//. "ZetaDB/execution/querySubOperator"

	"ZetaDB/execution"
	its "ZetaDB/execution/querySubOperator"
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
	ktm := execution.GetKeytableManager()
	schema, _ := ktm.Query_k_tableId_schema_FromTableId(9)
	seqIt := its.NewSequentialFileReaderIterator(9, schema)
	seqIt.Open(nil, nil)
	for seqIt.HasNext() {
		tup, _ := seqIt.GetNext()
		fmt.Println(tup.TupleGetMapKey())
	}
}
