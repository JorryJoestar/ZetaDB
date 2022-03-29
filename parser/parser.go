package parser

import (
	"sync"
)

type Parser struct {
	AST *ASTNode
}

var instance *Parser
var once sync.Once

//to get parser, call this function
func GetParser() *Parser {
	once.Do(func() {
		instance = &Parser{}
	})
	return instance
}

func (parser *Parser) ParseSql(sqlString string) *ASTNode {
	calcParse(newCalcLexer([]byte(sqlString)))

	return parser.AST
}
