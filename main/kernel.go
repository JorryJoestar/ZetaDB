package main

import (
	"ZetaDB/parser"
	"ZetaDB/storage"
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

	bf := storage.GetBufferPool()
	bf.GetDataPageSize()

	birth := "1997-11-12"
	bytes, _ := DATEToBytes(birth)

	storage.BytesToFile(bytes, 12, DEFAULT_DATAFILE_LOCATION)

}
