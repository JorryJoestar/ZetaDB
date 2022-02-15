package main

import (
	"ZetaDB/parser"
	"fmt"
)

type Tester struct{}

func (tester *Tester) PrintAST(ast *parser.ASTNode) {
	if ast == nil {
		fmt.Println("ast nil")
	} else {
		fmt.Printf("AST TYPE: %v", ast.Type)
	}

}
