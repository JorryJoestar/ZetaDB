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

	s := DEFAULT_KEY_TABLE_0_SCHEMA

	ast := kernel.parser.ParseSql(s)
	fmt.Println(ASTToString(ast))

	se := GetStorageEngine(DEFAULT_DATAFILE_LOCATION, DEFAULT_INDEXFILE_LOCATION, DEFAULT_LOGFILE_LOCATION)

	domain0, _ := container.NewDomain("id", container.INT, 0, 0)
	domain1, _ := container.NewDomain("name", container.VARCHAR, 20, 0)
	domain2, _ := container.NewDomain("age", container.INT, 0, 0)
	domain3, _ := container.NewDomain("height", container.FLOAT, 0, 0)
	domain4, _ := container.NewDomain("birthday", container.DATE, 0, 0)

	var domainList []*container.Domain
	domainList = append(domainList, domain0)
	domainList = append(domainList, domain1)
	domainList = append(domainList, domain2)
	domainList = append(domainList, domain3)
	domainList = append(domainList, domain4)
	schema, _ := container.NewSchema("testTable", domainList, nil)

	p31 := NewDataPageMode0(31, 31, 31, 31)
	tuple0 := CreateNewTuple(3, "simeon", 24, 167.1, "1997-11-12", 31, 0, schema)
	tuple1 := CreateNewTuple(7, "alex", 26, 199.2, "1998-03-02", 31, 0, schema)
	tuple2 := CreateNewTuple(12, "claire", 71, 160.3, "1997-10-20", 31, 0, schema)
	tuple3 := CreateNewTuple(6, "jojo", 31, 76.2, "1997-01-03", 31, 0, schema)
	tuple4 := CreateNewTuple(5, "woozie", 55, 155.2, "1997-02-08", 31, 0, schema)
	tuple5 := CreateNewTuple(9, "ruby", 17, 198.1, "1997-09-12", 31, 0, schema)
	p31.InsertTuple(tuple0)
	p31.InsertTuple(tuple1)
	p31.InsertTuple(tuple2)
	p31.InsertTuple(tuple3)
	p31.InsertTuple(tuple4)
	p31.InsertTuple(tuple5)
	se.InsertDataPage(p31)
	se.SwapDataPage(31)

	sfrit := execution.NewSequentialFileReaderIterator(se, 31, schema)

	var proIndexs []int
	proIndexs = append(proIndexs, 1)
	proIndexs = append(proIndexs, 3)
	proIt := execution.NewProjectionIterator(proIndexs)

	sfrit.Open(nil, nil)
	proIt.Open(sfrit, nil)
	for proIt.HasNext() {
		tuple, _ := proIt.GetNext()
		nameBytes, _ := tuple.TupleGetFieldValue(0)
		name, _ := BytesToVARCHAR(nameBytes)
		heightBytes, _ := tuple.TupleGetFieldValue(1)
		height, _ := BytesToFLOAT(heightBytes)
		fmt.Print(name)
		fmt.Print(" ")
		fmt.Println(height)

	}

}

func CreateNewTuple(id int32, name string, age int32, height float32, birthday string, tableId uint32, tupleId uint32, schema *container.Schema) *container.Tuple {
	//id
	idBytes := INTToBytes(id)
	idField, _ := container.NewFieldFromBytes(idBytes)

	//name
	nameBytes, _ := VARCHARToBytes(name)
	nameField, _ := container.NewFieldFromBytes(nameBytes)

	//age
	ageBytes := INTToBytes(age)
	ageField, _ := container.NewFieldFromBytes(ageBytes)

	//height
	heightBytes := FLOATToBytes(height)
	heightField, _ := container.NewFieldFromBytes(heightBytes)

	//birthday
	birthdayBytes, _ := DATEToBytes(birthday)
	birthdayField, _ := container.NewFieldFromBytes(birthdayBytes)

	var fields []*container.Field
	fields = append(fields, idField)
	fields = append(fields, nameField)
	fields = append(fields, ageField)
	fields = append(fields, heightField)
	fields = append(fields, birthdayField)
	newTuple, _ := container.NewTuple(tableId, tupleId, schema, fields)
	return newTuple
}
