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
	s += "delete from m where (x < 12);"
	s += "end;"

	ast := kernel.parser.ParseSql(s)
	fmt.Println(ASTToString(ast))

	k := "create procedure p\n"
	k += "(in f varchar(12),out x shortint)\n"
	k += "declare m int;\n"
	k += "declare h decimal(22,1);\n"
	k += "begin\n"
	k += "set n = 2;\n"
	k += "end;\n"

	ast = kernel.parser.ParseSql(k)
	fmt.Println(ASTToString(ast))

	l := "delete from m where x > 12;"

	ast = kernel.parser.ParseSql(l)
	fmt.Println(ASTToString(ast))
}
