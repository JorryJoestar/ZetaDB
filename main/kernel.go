package main

import (
	//. "ZetaDB/execution/querySubOperator"
	"ZetaDB/container"
	execution "ZetaDB/execution/querySubOperator"
	"ZetaDB/parser"
	. "ZetaDB/storage"
	. "ZetaDB/utility"
	"fmt"
	"sync"
)

type Kernel struct {
	parser *parser.Parser
}

//for singleton pattern
var instance *Kernel
var once sync.Once

//to get kernel, call this function
func GetInstance() *Kernel {
	once.Do(func() {
		instance = &Kernel{
			parser: parser.GetInstance()}
	})
	return instance
}

func main() {
	kernel := GetInstance()

	s := "select a,b,c from b;"

	ast := kernel.parser.ParseSql(s)
	fmt.Println(ASTToString(ast))

	se := GetStorageEngine(DEFAULT_DATAFILE_LOCATION, DEFAULT_INDEXFILE_LOCATION, DEFAULT_LOGFILE_LOCATION)

	domain, _ := container.NewDomain("name", container.VARCHAR, 5000, 0)
	var domainList []*container.Domain
	domainList = append(domainList, domain)
	schema, _ := container.NewSchema("tab", domainList, nil)

	p72 := NewDataPageMode0(72, 72, 72, 72)
	p66 := NewDataPageMode0(66, 66, 66, 66)

	longString := ""
	for i := 0; i < 100; i++ {
		longString += "A"
	}
	p72.InsertTuple(CreateNewTuple("longString", 72, 0, schema))
	p72.InsertTuple(CreateNewTuple("Claire", 72, 1, schema))
	p72.InsertTuple(CreateNewTuple("simeon", 72, 2, schema))
	p72.InsertTuple(CreateNewTuple("simeon", 72, 3, schema))
	p72.InsertTuple(CreateNewTuple("simeon", 72, 4, schema))
	p72.InsertTuple(CreateNewTuple("Claire", 72, 5, schema))
	p72.InsertTuple(CreateNewTuple("simeon", 72, 6, schema))
	p72.InsertTuple(CreateNewTuple("alse", 72, 7, schema))

	p66.InsertTuple(CreateNewTuple("Claire", 66, 0, schema))
	p66.InsertTuple(CreateNewTuple("simeon", 66, 1, schema))
	p66.InsertTuple(CreateNewTuple("alse", 66, 2, schema))
	p66.InsertTuple(CreateNewTuple("alex", 66, 3, schema))
	p66.InsertTuple(CreateNewTuple(longString, 66, 4, schema))

	se.InsertDataPage(p72)
	se.InsertDataPage(p66)

	se.SwapDataPage(72)
	se.SwapDataPage(66)

	seqIt72 := execution.NewSequentialFileReaderIterator(se, 72, schema)
	seqIt66 := execution.NewSequentialFileReaderIterator(se, 66, schema)

	seqIt72.Open(nil, nil)
	seqIt66.Open(nil, nil)

	setIntersectionIt := execution.NewSetIntersectionIterator()
	setIntersectionIt.Open(seqIt72, seqIt66)

	for setIntersectionIt.HasNext() {
		tup, _ := setIntersectionIt.GetNext()
		tupBytes, _ := tup.TupleGetFieldValue(0)
		fmt.Println(BytesToVARCHAR(tupBytes))
	}

}

func CreateNewTuple(name string, tableId uint32, tupleId uint32, schema *container.Schema) *container.Tuple {
	nameBytes, _ := VARCHARToBytes(name)
	field, _ := container.NewFieldFromBytes(nameBytes)
	var fields []*container.Field
	fields = append(fields, field)
	newTuple, _ := container.NewTuple(tableId, tupleId, schema, fields)
	return newTuple
}
