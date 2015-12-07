package lexer

import (
	"regexp"
	"token"
)

type Token token.Token

type Error struct {
	RowIndex int
	ColumtIndex int
	Value string
}

type Lexer struct {
	commentReg *regexp.Regexp
	intReg *regexp.Regexp
	floatReg *regexp.Regexp
	stringReg *regexp.Regexp
	
    addOpReg *regexp.Regexp
	subOpReg *regexp.Regexp
	mulOpReg *regexp.Regexp
	quoOpReg *regexp.Regexp
	remOpReg *regexp.Regexp
	
	lparenReg *regexp.Regexp
	lbrackReg *regexp.Regexp
	lbraceReg *regexp.Regexp
	
	rparenReg *regexp.Regexp
	rbrackReg *regexp.Regexp
	rbraceReg *regexp.Regexp
	semicolonReg *regexp.Regexp
	colonReg *regexp.Regexp
	
	breakReg *regexp.Regexp
	continueReg *regexp.Regexp
	
	ifReg *regexp.Regexp
	elseReg *regexp.Regexp
	forReg *regexp.Regexp
	
	funcReg *regexp.Regexp
	returnReg *regexp.Regexp
	
	typeReg *regexp.Regexp
	varReg *regexp.Regexp
	
	literalReg *regexp.Regexp
    numberReg *regexp.Regexp
	
	assignedReg *regexp.Regexp
	
	whiteSpacesReg *regexp.Regexp
	absorbErrorReg *regexp.Regexp
	
	rowIndex int
	columtIndex int
	tokens []Token
	errors []Error
}

func (self *Lexer) trim(expression string) string {	
	slice := self.whiteSpacesReg.FindStringIndex(expression)
	if slice != nil {
		self.columtIndex += slice[1]
		return expression[slice[1]:]
	}
	return expression
}

func (self *Lexer) absorbError(expression string) (string, string) {	
	slice := self.absorbErrorReg.FindStringIndex(expression)
	if slice != nil {
		self.columtIndex += slice[1]
		return expression[slice[1]:], expression[slice[0]:slice[1]]
	}
	return expression, ""
}

func (self *Lexer) appendError(expression string) string {	
	var error Error
	error.ColumtIndex = self.columtIndex
	error.RowIndex = self.rowIndex
	expression, error.Value = self.absorbError(expression)
	self.errors = append(self.errors, error)

	return expression
}

func (self *Lexer) parseTokenByExp(expressionSrc string, re regexp.Regexp, tokenType token.TokenType) (expression string, token Token) {	
	slice := re.FindStringIndex(expressionSrc)
	expression = expressionSrc
	if slice != nil {
		self.columtIndex += slice[1]
		token.ColumtIndex = self.columtIndex
		token.RowIndex = self.rowIndex
		token.TokenType = tokenType
		token.Value = expression[slice[0]:slice[1]]
		expression = expression[slice[1]:]
		return expression, token
	}
	return expression, token
}

func (self *Lexer) parseToken(expressionSrc string) (expression string, tokenElem Token) {	
    expression = self.trim(expressionSrc)
	
	expression, tokenElem = self.parseTokenByExp(expression, *self.numberReg, token.NUMBER)
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.commentReg, token.COMMENT)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.intReg, token.INT)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.floatReg, token.FLOAT)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.stringReg, token.STRING)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.addOpReg, token.ADD)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.subOpReg, token.SUB)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.mulOpReg, token.MUL)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.quoOpReg, token.QUO)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.remOpReg, token.REM)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.lparenReg, token.LPAREN)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.lbrackReg, token.LBRACK)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.lbraceReg, token.LBRACE)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.rparenReg, token.RPAREN)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.rbrackReg, token.RBRACK)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.rbraceReg, token.RBRACE)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.semicolonReg, token.SEMICOLON)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.colonReg, token.COLON)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.breakReg, token.BREAK)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.continueReg, token.CONTINUE)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.ifReg, token.IF)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.elseReg, token.ELSE)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.forReg, token.FOR)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.funcReg, token.FUNC)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.returnReg, token.RETURN)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.typeReg, token.TYPE)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.varReg, token.VAR)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.assignedReg, token.ASSIGNED)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression, tokenElem = self.parseTokenByExp(expression, *self.literalReg, token.LITERAL)
	}
	if tokenElem.TokenType == token.ILLEGAL {
		expression = self.appendError(expression)		
	}

	return expression, tokenElem
}

func (self *Lexer) ParseTokens(expression string) ([]Token, []Error) {	
	self.rowIndex = 0
	self.columtIndex = 0
	
	self.commentReg = regexp.MustCompile(`^((\/\*).*(\*\/))`)
	self.intReg = regexp.MustCompile(`^(int)`)
	self.floatReg = regexp.MustCompile(`^(float)`)
	self.stringReg = regexp.MustCompile(`^(string)`)
	
    self.addOpReg = regexp.MustCompile(`^(\+)`)
	self.subOpReg = regexp.MustCompile(`^(\-)`)
	self.mulOpReg = regexp.MustCompile(`^(\*)`)
	self.quoOpReg = regexp.MustCompile(`^(\/)`)
	self.remOpReg = regexp.MustCompile(`^(\%)`)
	
	self.lparenReg = regexp.MustCompile(`^(\()`)
	self.lbrackReg = regexp.MustCompile(`^(\[)`)
	self.lbraceReg = regexp.MustCompile(`^(\{)`)
	
	self.rparenReg = regexp.MustCompile(`^(\))`)
	self.rbrackReg = regexp.MustCompile(`^(\])`)
	self.rbraceReg = regexp.MustCompile(`^(\})`)
	self.semicolonReg = regexp.MustCompile(`^(;)`)
	self.colonReg = regexp.MustCompile(`^(:)`)
	
	self.breakReg = regexp.MustCompile(`^(break)`)
	self.continueReg = regexp.MustCompile(`^(continue)`)
	
	self.ifReg = regexp.MustCompile(`^(if)`)
	self.elseReg = regexp.MustCompile(`^(else)`)
	self.forReg = regexp.MustCompile(`^(for)`)
	
	self.funcReg = regexp.MustCompile(`^(func)`)
	self.returnReg = regexp.MustCompile(`^(return)`)
	
	self.typeReg = regexp.MustCompile(`^(type)`)
	self.varReg = regexp.MustCompile(`^(var)`)
	
	self.assignedReg = regexp.MustCompile(`^(\=)`)
	
	self.numberReg = regexp.MustCompile(`^(([0-9]*\.[0-9]+)|([0-9]+\.[0-9]*)|([0-9]+))`)
	self.literalReg = regexp.MustCompile(`^([a-zA-Z]+[a-zA-Z0-9]*)`)
	
	self.whiteSpacesReg = regexp.MustCompile(`^(\s+)`)
	self.absorbErrorReg = regexp.MustCompile(`^(.[^\s]+)`)

	var token Token
	
	for expression != "" {
		expression, token = self.parseToken(expression)
		self.tokens = append(self.tokens, token)
	}
	
	return self.tokens, self.errors
}