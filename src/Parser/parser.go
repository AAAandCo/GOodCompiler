package parser

import (
	"token"
	"lexer"
	"ast"
)

type parser struct {
	currentToken token.Token
	tokens token.Tokens
}

func (p *parser) init(source string) {
	lexerObj := new(lexer.Lexer)
	tokens, _ := lexerObj.ParseTokens(source)
	p.tokens = tokens
}

func (p *parser) parseFile() *ast.FileAst {

	return &ast.FileAst{}
}