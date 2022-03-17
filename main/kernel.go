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
	lPage, _ := se.FetchLogPage(1)
	fmt.Println(lPage)
	fmt.Println(lPage.LogPageGetLogNum())
	for _, l := range lPage.LogPageGetLogs() {
		fmt.Println(l)
	}

}
