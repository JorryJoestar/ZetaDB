package parser

import (
	"sync"
)

type Parser struct {
	AST *ASTNode
	err error
}

var parserInstance *Parser
var parserOnce sync.Once

//to get parser, call this function
func GetParser() *Parser {
	parserOnce.Do(func() {
		parserInstance = &Parser{}
	})
	return parserInstance
}

func (parser *Parser) ParseSql(sqlString string) (*ASTNode, error) {
	parser.err = nil
	calcParse(newCalcLexer([]byte(sqlString)))

	return parser.AST, parser.err
}
