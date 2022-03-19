package main

import (
	"ZetaDB/container"
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

	/* 	for i := 20; i < 25; i++ {
	   		np := NewDataPageMode0(uint32(i), 12, 3, 6)
	   		fmt.Println(se.InsertDataPage(np))
	   	}
	   	for i := 20; i < 25; i++ {
	   		fmt.Println(se.SwapDataPage(uint32(i)))
	   	}
	   	p30M1 := NewDataPageMode1(30, 9, 12, 97, 23, INTToBytes(972))
	   	p90M2 := NewDataPageMode2(90, 78, 4, 30, 76, INTToBytes(42))
	   	fmt.Println(se.InsertDataPage(p30M1))
	   	fmt.Println(se.InsertDataPage(p90M2))
	   	fmt.Println(se.SwapDataPage(30))
	   	fmt.Println(se.SwapDataPage(90))

	   	pk2m1 := NewDataPageMode1(2, 2, 1, 3, 19, INTToBytes(999))
	   	pk6m2 := NewDataPageMode2(6, 6, 2, 2, 876, SHORTINTToBytes(87))
	   	fmt.Println(se.InsertDataPage(pk2m1))
	   	fmt.Println(se.InsertDataPage(pk6m2))
	   	fmt.Println(se.SwapDataPage(2))
	   	fmt.Println(se.SwapDataPage(6)) */

	var domainList []*container.Domain
	d1, _ := container.NewDomain("id", 6, 0, 0)
	domainList = append(domainList, d1)
	d2, _ := container.NewDomain("name", 2, 20, 0)
	domainList = append(domainList, d2)
	schema, _ := container.NewSchema("testTable", domainList, nil)

	pk2m1, err := se.GetDataPage(2, schema)
	fmt.Println(pk2m1, err)

}
