package token

import "strconv"

type Token struct {
	RowIndex int
	ColumtIndex int
	TokenType TokenType
	Value string
}

type Tokens []Token

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	COMMENT

	RETURN
	VAR
	TYPE
	FUNC

	IDENTIFIER
	INT
	FLOAT
	STRING

	ADD
	SUB
	MUL
	QUO
	REM

	IF
	FOR
	ELSE
	BREAK
	CONTINUE

	LPAREN
	LBRACK
	LBRACE

	RPAREN
	RBRACK
	RBRACE

	SEMICOLON
	COLON
	
	NUMBER
	LITERAL
	ASSIGNED
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:     "EOF",
	COMMENT: "COMMENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	STRING: "STRING",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",
	
	NUMBER:   "NUMBER",

	BREAK:    "BREAK",
	CONTINUE: "CONTINUE",

	IF:			 "IF",
	ELSE:        "ELSE",
	FOR:         "FOR",

	FUNC:   "FUNC",

	RETURN:    "RETURN",

	TYPE:   "TYPE",
	VAR:    "VAR",
	
	LITERAL: "LITERAL",
	ASSIGNED: "=",
}

func (tok TokenType) String() string {
	s := ""
	if 0 <= tok && tok < TokenType(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}