package main

import (
	"ZetaDB/parser"
	"fmt"
)

type Tester struct{}

func (tester *Tester) PrintAST(ast *parser.ASTNode) {
	fmt.Printf("AST TYPE: %v", ast.Type)
}
