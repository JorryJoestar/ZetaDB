package main

import (
	"ZetaDB/parser"
)

type Kernel struct {
	parser parser.Parser
}

func NewKernel() Kernel {
	return Kernel{
		parser: parser.Parser{}}
}

func main() {
	kernel := NewKernel()
	kernel.parser.ParseSql("(")
}
