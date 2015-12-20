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
	"unicode"
	"math"
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
%token <val> DIGIT LETTER FUNC_EXPR

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */
%left FUNC_EXPR


%%

list	: /* empty */
	| statement
	;

statement	:    expr
		{
			fmt.Printf( "result: %d\n", $1 );
		}
	|    LETTER '=' expr
		{
			regs[$1]  =  $3
		}
	;


expr	:    '(' expr ')'
		{ $$  =  $2 }
	|    expr '+' expr
		{ $$  =  $1 + $3 }
	|    expr '-' expr
		{ $$  =  $1 - $3 }
	|    expr '*' expr
		{ $$  =  $1 * $3 }
	|    expr '/' expr
		{ $$  =  $1 / $3 }
	|    expr '%' expr
		{ $$  =  $1 % $3 }
	|    expr '&' expr
		{ $$  =  $1 & $3 }
	|    FUNC_EXPR number
	     {
	     	fmt.Printf("func := %d;\n", $2)
	     	$$ = int(math.Sin(float64($2)))
	     	fmt.Printf("float64 := %f;\n", float64($2))
	     	fmt.Printf("math sin := %f;\n", math.Sin(float64($2)))
	     	fmt.Printf("resINt := %d;\n", $$)
	     }
	|    expr '|' expr
		{ $$  =  $1 | $3 }
	|    '-'  expr        %prec  UMINUS
		{ $$  = -$2  }
	|    LETTER
		{ $$  = regs[$1] }
	|    number
		{
		    fmt.Printf("expr <- number\n")
			$$ = $1
		}
	;

number	:    DIGIT
		{
		    fmt.Printf("val := %d;\n", $1)
			$$ = $1;
			if $1==0 {
				base = 8
			} else {
				base = 10
			}
		}
	|    number DIGIT
		{
			$$ = base * $1 + $2
		    fmt.Printf("val := %d;\n", $$)
		}
	;

%%      /*  start  of  programs  */

type CalcLex struct {
	s string
	pos int
}


func (l *CalcLex) Lex(lval *CalcSymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.s) {
			return 0
		}
		c = rune(l.s[l.pos])
		l.pos += 1
	}
	if l.pos == len(l.s) {
		return 0
	}

	if (c == 's') {
		fmt.Printf("s consumed")
		return FUNC_EXPR
	}

	if unicode.IsDigit(c) {
		lval.val = int(c) - '0'
		return DIGIT
	} else if unicode.IsLower(c) {
		lval.val = int(c) - 'a'
		return LETTER
	}
	return int(c)
}

func (l *CalcLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func main() {
	fi := bufio.NewReader(os.Stdin)

	for {
		var eqn string
		var ok bool

		fmt.Printf("equation: ")
		if eqn, ok = readline(fi); ok {
			CalcParse(&CalcLex{s: eqn})
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