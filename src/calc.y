// Copyright 2011 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// based off of Appendix A from http://dinosaur.compilertools.net/yacc/

%{

package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"strconv"
	"lexer"
	"token"
)

var regs = make([]int, 26)
var base int

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	val int
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <val> expr number

// same for terminals
%token <val> INT_NUMBER ADD SUB MUL QUO REM SIN LPAREN RPAREN ASSIGNED IDENTIFIER

%left ADD  SUB
%left MUL  QUO  REM
%left UMINUS      /*  supplies  precedence  for  unary  minus  */
%left SIN

%%

list	: /* empty */
	| statement
	;

statement	:    expr
		{
			fmt.Printf( "result: %d\n", $1 );
		}
	|    IDENTIFIER ASSIGNED expr
		{
			regs[$1]  =  $3
		}
	;


expr	:    LPAREN expr RPAREN
		{ $$  =  $2 }
	|    expr ADD expr
		{ $$  =  $1 + $3 }
	|    expr SUB expr
		{ $$  =  $1 - $3 }
	|    expr MUL expr
		{ $$  =  $1 * $3 }
	|    expr QUO expr
		{ $$  =  $1 / $3 }
	|    expr REM expr
		{ $$  =  $1 % $3 }
	|    SIN LPAREN expr RPAREN
	     {
	     	fmt.Printf("func := %d;\n", $3)
	     	$$ = int(math.Sin(float64($3)))
	     	fmt.Printf("float64 := %f;\n", float64($3))
	     	fmt.Printf("math sin := %f;\n", math.Sin(float64($3)))
	     	fmt.Printf("resINt := %d;\n", $$)
	     }
	|    SUB  expr        %prec  UMINUS
		{ $$  = -$2  }
	|    IDENTIFIER
		{ $$  = regs[$1] }
	|    number
		{
		    fmt.Printf("expr <- number\n")
			$$ = $1
		}
	;

number	:   INT_NUMBER;

%%      /*  start  of  programs  */

type CalcLex struct {
	tokens []token.Token
	tokenMap map[int]int
	pos int	
}


func (l *CalcLex) Lex(lval *CalcSymType) int {	
	if l.pos == len(l.tokens) {
		return 0
	}
	var c token.Token = l.tokens[l.pos]
	l.pos += 1

	if (c.TokenType == token.INT_NUMBER) {
		numb, _ := strconv.ParseInt(c.Value, 10, 0)
		lval.val = int(numb)
		return INT_NUMBER
	}

	return int(l.tokenMap[int(c.TokenType)])
}

func (l *CalcLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func prepareTokenMap() map[int]int {
	var tokenMap map[int]int
	tokenMap = make(map[int]int)

	tokenMap[int(token.ADD)] = ADD
	tokenMap[int(token.SUB)] = SUB
	tokenMap[int(token.MUL)] = MUL
	tokenMap[int(token.QUO)] = QUO
	tokenMap[int(token.REM)] = REM
	tokenMap[int(token.SIN)] = SIN
	tokenMap[int(token.LPAREN)] = LPAREN
	tokenMap[int(token.RPAREN)] = RPAREN
	tokenMap[int(token.ASSIGNED)] = ASSIGNED
	tokenMap[int(token.INT_NUMBER)] = INT_NUMBER
	tokenMap[int(token.IDENTIFIER)] = IDENTIFIER
	
	return tokenMap
}

func main() {
	fi := bufio.NewReader(os.Stdin)

	for {
		var eqn string
		var ok bool

		fmt.Printf("equation: ")
		if eqn, ok = readline(fi); ok {
			lexerObj := new(lexer.Lexer)
			tokensParsed, errors := lexerObj.ParseTokens(eqn)
			if len(errors) > 0 {
				panic("Unexpected tokens")
			}
			tokenMapPrepared := prepareTokenMap()
			CalcParse(&CalcLex{tokens: tokensParsed, tokenMap: tokenMapPrepared})
		} else {
			break
		}
	}
}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}