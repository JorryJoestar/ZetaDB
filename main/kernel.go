package main

import (
	"ZetaDB/container"
	"ZetaDB/execution"
	"ZetaDB/parser"
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
	//ktm := execution.GetKeytableManager()
	ee := execution.GetExecutionEngine()
	//tm := execution.GetTableManipulator()
	Parse := parser.GetParser()
	rewriter := execution.GetRewriter()

	transaction := storage.GetTransaction()
	//ktm.InitializeSystem()

	//sql := "create table student(id int, name varchar(20));"
	//sql := "drop table student;"
	//sql := "insert into student values (976, 'Alex');"
	//sql := "delete from student where id = 976;"
	sql := "update student set name = 'jack' where id = 976;"
	astNode, _ := Parse.ParseSql(sql)
	pp := rewriter.ASTNodeToExecutionPlan(1, astNode, sql)
	result := ee.Execute(pp)
	fmt.Println(result)

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
