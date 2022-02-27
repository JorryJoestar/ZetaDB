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

	iom, err := storage.GetIOManipulator(DEFAULT_DATAFILE_LOCATION, DEFAULT_INDEXFILE_LOCATION)
	if err != nil {
		fmt.Println(err)
	}

	ch := "k"
	bytes, err := CHARToBytes(ch)

	iom.BytesToIndexFile(bytes, 5)

	ss := "this is woozie speaking"
	bytes, err = VARCHARToBytes(ss)

	fmt.Println(bytes)

	iom.BytesToIndexFile(bytes, 10)

	bytes, err = iom.BytesFromIndexFile(33, 1)

	fmt.Println(bytes, err)

}
