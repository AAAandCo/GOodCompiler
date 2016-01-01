package ast

import "token"

type Node interface {
	astNode()
}

type Declaration interface {
	Node
	declNode()
}

type Statement interface {
	Node
	stmtNode()
}

type Expression interface {
	Node
	exprNode()
}

type (
	VarDecl struct {
		Name string

	}

	FuncDecl struct {
		Name string
		// todo: add parameter list
		// todo: add function body
	}
)



type (

	FileAst struct {
		decl []DeclarationAst
	}

	DeclarationAst struct {}

	NumberExprAst struct {
		Val float32
	}

	VariableExprAst struct {
		Name string
	}

	BinaryExprAst struct {
		Op token.Token
		Left NodeAst
		Right NodeAst
	}

	CallExprAst struct {
		Callee string
		Args []NodeAst
	}
)

func (p *DeclarationAst) astNode() {}
func (p *FileAst) astNode() {}
func (p *NumberExprAst) astNode() {}
func (p *VariableExprAst) astNode() {}
func (p *BinaryExprAst) astNode() {}
func (p *CallExprAst) astNode() {}


