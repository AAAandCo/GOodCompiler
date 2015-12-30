package ast

import "token"

type NodeAst interface {
	astNode()
}

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


