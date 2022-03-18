package main

import (
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
	/*	for i := 0; i < 5; i++ {
			np, _ := NewIndexPage(uint32(i), 1, 2)
			se.InsertIndexPage(np)
		}
		for i := 0; i < 5; i++ {
			se.SwapIndexPage(uint32(i))
		}
		p1, _ := NewIndexPage(5, 2, 4)
		p2, _ := NewIndexPage(6, 3, 7)
		se.InsertIndexPage(p1)
		se.InsertIndexPage(p2)
		se.SwapIndexPage(5)
		se.SwapIndexPage(6)
	*/

	/*	p, _ := se.GetIndexPage(1)
		fmt.Println(p.IndexPageSetElementAt(2, INTToBytes(19)))
		fmt.Println(p.IndexPageSetPointerPageIdAt(1, 99))
		fmt.Println(p.IndexPageToBytes())
		se.SwapIndexPage(1)
	*/

	p, _ := se.GetIndexPage(1)
	fmt.Println(p.IndexPageGetPointerPageIdAt(1))

}
