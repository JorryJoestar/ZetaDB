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
	se.GetIndexPage(5)
	p, _ := se.GetIndexPage(1)
	ir := container.NewIndexRecord(Uint32ToBytes(1020), 18, 1)
	ir2 := container.NewIndexRecord(Uint32ToBytes(111), 32, 3)

	fmt.Println(p.IndexPageInsertNewIndexRecord(ir))
	fmt.Println(p.IndexPageInsertNewIndexRecord(ir2))
	fmt.Println("---------------")
	fmt.Println(p.IndexPageGetIndexRecordAt(0))
	fmt.Println(p.IndexPageGetIndexRecordAt(1))
	se.SwapIndexPage(1)

}
