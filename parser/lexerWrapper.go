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

	//constraint
	case C.UNIQUE:
		return UNIQUE
	case C.PRIMARYKEY:
		return PRIMARYKEY
	case C.CHECK:
		return CHECK
	case C.FOREIGNKEY:
		return FOREIGNKEY
	case C.REFERENCES:
		return REFERENCES
	case C.NOT_DEFERRABLE:
		return NOT_DEFERRABLE
	case C.DEFERED_DEFERRABLE:
		return DEFERED_DEFERRABLE
	case C.IMMEDIATE_DEFERRABLE:
		return IMMEDIATE_DEFERRABLE
	case C.UPDATE_NULL:
		return UPDATE_NULL
	case C.UPDATE_CASCADE:
		return UPDATE_CASCADE
	case C.DELETE_NULL:
		return DELETE_NULL
	case C.DELETE_CASCADE:
		return DELETE_CASCADE
	case C.CONSTRAINT:
		return CONSTRAINT
	case C.DEFAULT:
		return DEFAULT

	//elemtary value
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

	//prdicate
	case C.LIKE:
		return LIKE
	case C.IN:
		return IN
	case C.ALL:
		return ALL
	case C.ANY:
		return ANY
	case C.IS:
		return IS
	case C.EXISTS:
		return EXISTS
	case C.NOTEQUAL:
		return NOTEQUAL
	case C.LESS:
		return LESS
	case C.GREATER:
		return GREATER
	case C.LESSEQUAL:
		return LESSEQUAL
	case C.GREATEREQUAL:
		return GREATEREQUAL
	case C.EQUAL:
		return EQUAL

	//attriNameOptionTableName
	case C.DOT:
		return DOT

	//public
	case C.LPAREN:
		return LPAREN
	case C.RPAREN:
		return RPAREN
	case C.COMMA:
		return COMMA
	case C.NOT:
		return NOT
	case C.NULLMARK:
		return NULLMARK
	case C.ID:
		yylval.String = p.yytext
		return ID
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
