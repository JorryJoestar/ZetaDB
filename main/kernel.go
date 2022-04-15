package main

import (
	"ZetaDB/container"
	"ZetaDB/execution"
	"ZetaDB/parser"
	its "ZetaDB/physicalPlan"
	"ZetaDB/storage"
	"ZetaDB/utility"
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
	//ee := execution.GetExecutionEngine()
	//tm := execution.GetTableManipulator()
	//Parse := parser.GetParser()
	//rewriter := optimizer.GetRewriter()

	predicate := &container.Predicate{
		PredicateType:      1,
		CompareMark:        6,
		CompareValueType:   1,
		CompareIntValue:    2,
		LeftAttributeIndex: 0,
	}
	condition := &container.Condition{
		Predicate:     predicate,
		ConditionType: container.CONDITION_PREDICATE,
	}

	tableId, schema, _ := ktm.Query_k_tableId_schema_FromTableName("student")
	headPageId, _, _, _, _ := ktm.Query_k_table(tableId)
	sfr := its.NewSequentialFileReaderIterator(headPageId, schema)
	sfr.Open(nil, nil)

	si := its.NewSelectionIterator(condition)
	si.Open(sfr, nil)
	for si.HasNext() {
		tuple, _ := si.GetNext()
		fmt.Println("tupleId: ", tuple.TupleGetTupleId())
	}

	transaction := storage.GetTransaction()
	//ktm.InitializeSystem()

	//sql := "create table student(id int, name varchar(20));"
	//sql := "drop table student;"
	//sql := "insert into student values (976, 'Alex');"
	//astNode, _ := Parse.ParseSql(sql)
	//pp := rewriter.ASTNodeToPhysicalPlan(1, astNode, sql)
	//result := ee.ExecutePhysicalPlan(pp)
	//fmt.Println(result)

	transaction.PushTransactionIntoDisk()

	//PrintTable(2)
	//PrintTable(8)
	//PrintTable(9)
	//PrintTable(15)

	PrintTableByName("student")
}

func getNewTuple(id int32, name string) *container.Tuple {
	ktm := execution.GetKeytableManager()

	schema0 := ktm.GetKeyTableSchema(0)

	field0, _ := container.NewFieldFromBytes(utility.INTToBytes(id))
	fieldsByte, _ := utility.VARCHARToBytes(name)
	field1, _ := container.NewFieldFromBytes(fieldsByte)
	var fields []*container.Field
	fields = append(fields, field0)
	fields = append(fields, field1)
	tuple, _ := container.NewTuple(0, 0, schema0, fields)
	return tuple
}

//socket := network.GetSocket()
//socket.Listen()
//ee.CreateTableOperator(10, "create table x(id int primary key, longString varchar(100000));")

/* 	ee.CreateTableOperator(10, "create table m(id int primary key, longString varchar(100000));") */

/* 	var longS string
   	for i := 1; i <= 5000; i++ {
   		longS += "b"
   	}

   	newTuple := getNewTuple(2, "ClaireMao")
   	tm.InsertTupleIntoTable(17, newTuple) */

//tm.DeleteTupleFromTable(17, 13)
