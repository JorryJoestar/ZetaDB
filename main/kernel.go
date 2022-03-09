package main

import (
	"ZetaDB/parser"

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

	b1, _ := DATEToBytes("1997-11-12")
	b2, _ := DATEToBytes("1997-10-20")
	b3, _ := DATEToBytes("1998-01-01")
	b4, _ := DATEToBytes("1999-02-16")
	b5, _ := DATEToBytes("2004-01-07")
	b6, _ := DATEToBytes("1911-03-08")

	var ss [][]byte
	InsertToOrderedSlice(8, &ss, b1)
	InsertToOrderedSlice(8, &ss, b2)
	InsertToOrderedSlice(8, &ss, b3)
	InsertToOrderedSlice(8, &ss, b4)
	InsertToOrderedSlice(8, &ss, b5)
	fmt.Println(ss)
	left, right, _ := InsertToOrderedSliceSplit(8, &ss, b6)
	fmt.Println(b6)
	fmt.Println(left)
	fmt.Println(right)

}
