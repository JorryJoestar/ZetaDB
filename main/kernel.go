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
	s := ""
	s += "create trigger t\n"
	s += "instead of update on m\n"
	s += "referencing\n"
	s += "\told row as k,\n"
	s += "\told row as k\n"
	s += "for each row\n"
	s += "when (x < 12)\n"
	s += "begin\n"
	s += ".;.;"
	s += "end;"

	ast := kernel.parser.ParseSql(s)
	fmt.Println(ASTToString(ast))

	ast = kernel.parser.ParseSql("a+(b*c/2)+6")
}
