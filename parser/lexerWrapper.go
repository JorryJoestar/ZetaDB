package parser

//#include "token.h"
//#include "calc.lexer.h"
import "C"

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

type calcLex struct {
	yylineno int
	yytext   string
	lastErr  error
}

var _ calcLexer = (*calcLex)(nil)

func newCalcLexer(data []byte) *calcLex {
	p := new(calcLex)

	C.yy_scan_bytes(
		(*C.char)(C.CBytes(data)),
		C.yy_size_t(len(data)),
	)

	return p
}

// The parser calls this method to get each new token. This
// implementation returns operators and NUM.
func (p *calcLex) Lex(yylval *calcSymType) int {
	p.lastErr = nil

	var tok = C.yylex()

	p.yylineno = int(C.yylineno)
	p.yytext = C.GoString(C.yytext)
	switch tok {

	case C.INTVALUE:
		yylval.Int, _ = strconv.Atoi(p.yytext)
		return INTVALUE
	case C.FLOATVALUE:
		yylval.Float, _ = strconv.ParseFloat(p.yytext, 64)
		return FLOATVALUE
	case C.STRINGVALUE:
		yylval.String = strings.Trim(p.yytext, "\"'")
		return STRINGVALUE
	case C.BOOLVALUE:
		if p.yytext[0] == byte('t') {
			yylval.Boolean = true
		} else {
			yylval.Boolean = false
		}
		return BOOLVALUE
	}
	return 0 //end of statement
}

// The parser calls this method on a parse error.
func (p *calcLex) Error(s string) {
	p.lastErr = errors.New("yacc: " + s)
	if err := p.lastErr; err != nil {
		log.Println(err)
	}
}
