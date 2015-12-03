package token

import "strconv"

type Token int

const (
	ILLEGAL Token = iota
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

	BREAK:    "break",
	CONTINUE: "continue",

	ELSE:        "else",
	FOR:         "for",

	FUNC:   "func",

	RETURN:    "return",

	TYPE:   "type",
	VAR:    "var",
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}