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

	s := "insert into k(p,m,q) values (1,2,3,4);"

	ast := kernel.parser.ParseSql(s)
	fmt.Println(ASTToString(ast))
}
