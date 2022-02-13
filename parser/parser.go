package parser

type Parser struct {
}

func (parser *Parser) ParseSql(sqlString string) {
	calcParse(newCalcLexer([]byte(sqlString)))
}
