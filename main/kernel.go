package main

import (
	"ZetaDB/parser"
	"ZetaDB/storage"
	"fmt"
	"os"
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
	bf.GetPageSize()

	///
	var bytes []byte
	s1 := "hello simeon kkkkkkkk abc ddd"
	b := []byte(s1)
	bytes = append(bytes, b...)
	bytes[0] = 11
	bytes[1] = 27
	bytes[2] = 19
	bytes[3] = 64
	bytes[4] = 8
	bytes[5] = 4

	f, err := os.Create(fileLocation)
	if err != nil {
		fmt.Println(fileLocation, err)
		return
	}

	f.Write(bytes)

	buffer := make([]byte, 2)
	n, err := f.ReadAt(buffer, 1)
	fmt.Println(n)
	fmt.Println(buffer)

}
