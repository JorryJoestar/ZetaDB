package main

import (
	"ZetaDB/parser"
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

	s := "select a from b UNION ( select x from m INTERSECTION select name from student);"

	ast := kernel.parser.ParseSql(s)
	fmt.Println(ASTToString(ast))
}
