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
	ast := kernel.parser.ParseSql("alter table x add m int references t(x) on delete set cascade;")
	fmt.Println(ASTToString(ast))
}
