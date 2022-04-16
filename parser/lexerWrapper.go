package parser

//#include "token.h"
//#include "calc.lexer.h"
import "C"
import (
	"errors"
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
	//createTable
	case C.CREATE:
		return CREATE
	case C.TABLE:
		return TABLE

	//dropTable
	case C.DROP:
		return DROP

	//alterTableAdd
	case C.ALTER:
		return ALTER
	case C.ADD:
		return ADD

	//createAssert
	case C.ASSERTION:
		return ASSERTION

	//createView
	case C.VIEW:
		return VIEW
	case C.AS:
		return AS

	//createIndex
	case C.INDEX:
		return INDEX
	case C.ON:
		return ON

	//createTrigger
	case C.TRIGGER:
		return TRIGGER
	case C.REFERENCING:
		return REFERENCING
	case C.BEFORE:
		return BEFORE
	case C.UPDATE:
		return UPDATE
	case C.OF:
		return OF
	case C.AFTER:
		return AFTER
	case C.INSTEAD:
		return INSTEAD
	case C.INSERT:
		return INSERT
	case C.DELETE:
		return DELETE
	case C.OLD:
		return OLD
	case C.ROW:
		return ROW
	case C.NEW:
		return NEW
	case C.FOR:
		return FOR
	case C.EACH:
		return EACH
	case C.STATEMENT:
		return STATEMENT
	case C.WHEN:
		return WHEN
	case C.BEGINTOKEN:
		return BEGINTOKEN
	case C.END:
		return END

	//expression
	case C.PLUS:
		return PLUS
	case C.SUBTRACT:
		return SUBTRACT
	case C.DIVISION:
		return DIVISION
	case C.CONCATENATION:
		return CONCATENATION

	//aggregation
	case C.STAR:
		return STAR
	case C.SUM:
		return SUM
	case C.AVG:
		return AVG
	case C.MIN:
		return MIN
	case C.MAX:
		return MAX
	case C.COUNT:
		return COUNT
	case C.DISTINCT:
		return DISTINCT

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

	//condition
	case C.AND:
		return AND
	case C.OR:
		return OR

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

	//domain
	case C.CHAR:
		return CHAR
	case C.VARCHAR:
		return VARCHAR
	case C.BIT:
		return BIT
	case C.BITVARYING:
		return BITVARYING
	case C.BOOLEAN:
		return BOOLEAN
	case C.INT:
		return INT
	case C.INTEGER:
		return INTEGER
	case C.SHORTINT:
		return SHORTINT
	case C.FLOAT:
		return FLOAT
	case C.REAL:
		return REAL
	case C.DOUBLEPRECISION:
		return DOUBLEPRECISION
	case C.DECIMAL:
		return DECIMAL
	case C.NUMERIC:
		return NUMERIC
	case C.DATE:
		return DATE
	case C.TIME:
		return TIME

	//attriNameOptionTableName
	case C.DOT:
		return DOT

	//psm
	case C.CALL:
		return CALL
	case C.ELSEIF:
		return ELSEIF
	case C.THEN:
		return THEN
	case C.IF:
		return IF
	case C.ELSE:
		return ELSE
	case C.CURSOR:
		return CURSOR
	case C.DO:
		return DO
	case C.RETURN:
		return RETURN
	case C.SET:
		return SET
	case C.OUT:
		return OUT
	case C.INOUT:
		return INOUT
	case C.DECLARE:
		return DECLARE
	case C.PROCEDURE:
		return PROCEDURE
	case C.FUNCTION:
		return FUNCTION
	case C.RETURNS:
		return RETURNS

	//delete
	case C.FROM:
		return FROM
	case C.WHERE:
		return WHERE

	//insert
	case C.INSERTINTO:
		return INSERTINTO
	case C.VALUES:
		return VALUES

	//dcl
	case C.START:
		return START
	case C.TRANSACTION:
		return TRANSACTION
	case C.COMMIT:
		return COMMIT
	case C.ROLLBACK:
		return ROLLBACK
	case C.SHOW:
		return SHOW
	case C.TABLES:
		return TABLES
	case C.ASSERTIONS:
		return ASSERTIONS
	case C.VIEWS:
		return VIEWS
	case C.INDEXS:
		return INDEXS
	case C.TRIGGERS:
		return TRIGGERS
	case C.FUNCTIONS:
		return FUNCTIONS
	case C.PROCEDURES:
		return PROCEDURES
	case C.USER:
		return USER
	case C.PASSWORD:
		return PASSWORD
	case C.PASSWORDS:
		yylval.String = p.yytext
		return PASSWORDS
	case C.CONNECT:
		return CONNECT
	case C.INITIALIZE:
		return INITIALIZE
	case C.HALT:
		return HALT

	//dql
	case C.DESC:
		return DESC
	case C.ASC:
		return ASC
	case C.CROSS:
		return CROSS
	case C.JOIN:
		return JOIN
	case C.NATURAL:
		return NATURAL
	case C.FULL:
		return FULL
	case C.OUTER:
		return OUTER
	case C.LEFT:
		return LEFT
	case C.RIGHT:
		return RIGHT
	case C.SELECT:
		return SELECT
	case C.GROUPBY:
		return GROUPBY
	case C.HAVING:
		return HAVING
	case C.ORDERBY:
		return ORDERBY
	case C.LIMIT:
		return LIMIT
	case C.UNION:
		return UNION
	case C.DIFFERENCE:
		return DIFFERENCE
	case C.INTERSECTION:
		return INTERSECTION

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
	case C.SEMICOLON:
		return SEMICOLON
	}
	return 0 //end of statement
}

// The parser calls this method on a parse error.
func (p *calcLex) Error(s string) {
	p.lastErr = errors.New("yacc: " + s)
	if err := p.lastErr; err != nil {
		//log.Println(err)
		Parser := GetParser()
		Parser.err = err
	}
}
