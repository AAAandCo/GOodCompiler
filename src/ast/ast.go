package ast

import (
	"token"
)

type NodeAst interface {
	astNode()
//	Accept(v Visitor)
}

type Declaration interface {
	NodeAst
	declNode()
}

type Statement interface {
	NodeAst
	stmtNode()
}

type Expression interface {
	NodeAst
	exprNode()
}

// Expressions
type (
	BadExpr struct {
		Msg string
	}

	BasicLit struct {
		T token.Token
	}

	Ident struct {
		Name string
		T token.Token
		Obj *Object
	}

	UnaryExpr struct {
		X Expression
		Op token.TokenType
		OpT token.Token
	}

	BinaryExpr struct {
		X Expression
		Op token.TokenType
		OpT token.Token
		Y Expression
	}

	ArrayType struct {
		Index int
		At Expression  // base_type
	}
)

func (p *BadExpr) astNode() {}
func (p *BasicLit) astNode() {}
func (p *Ident) astNode() {}
func (p *UnaryExpr) astNode() {}
func (p *BinaryExpr) astNode() {}
func (p *ArrayType) astNode() {}

func (p *BadExpr) exprNode() {}
func (p *BasicLit) exprNode() {}
func (p *Ident) exprNode() {}
func (p *UnaryExpr) exprNode() {}
func (p *BinaryExpr) exprNode() {}
func (p *ArrayType) exprNode() {}

// Declarations
type (
	VarDecl struct {
		Name *Ident
		Type Expression
	}

	FuncDecl struct {
		Name *Ident
		Params []Field
		RetType Expression
		Body *BlockStmt
	}
)

type Field struct {
	Name *Ident
	Type Expression
}

func (p *VarDecl) astNode() {}
func (p *FuncDecl) astNode() {}

func (p *VarDecl) declNode() {}
func (p *FuncDecl) declNode() {}


// Statements
type (
	DeclStmt struct {
		Decl Declaration
	}

	EmtpyStmt struct {}

	ExprStmt struct {
		X Expression
	}

	AssignStmt struct {
		LLst []Expression
		Op token.TokenType
		RList []Expression
	}

	ReturnStmt struct {
		X Expression
	}

	BlockStmt struct {
		List []Expression
	}

	IfStmt struct {
		Cond Expression
		Body *BlockStmt
		Else *BlockStmt
	}

	ForStmt struct {
		X Expression
		Body *BlockStmt
	}
)

func (p *DeclStmt) astNode() {}
func (p *EmtpyStmt) astNode() {}
func (p *ExprStmt) astNode() {}
func (p *AssignStmt) astNode() {}
func (p *ReturnStmt) astNode() {}
func (p *BlockStmt) astNode() {}
func (p *IfStmt) astNode() {}
func (p *ForStmt) astNode() {}

func (p *DeclStmt) stmtNode() {}
func (p *EmtpyStmt) stmtNode() {}
func (p *ExprStmt) stmtNode() {}
func (p *AssignStmt) stmtNode() {}
func (p *ReturnStmt) stmtNode() {}
func (p *BlockStmt) stmtNode() {}
func (p *IfStmt) stmtNode() {}
func (p *ForStmt) stmtNode() {}

// File
type FileAst struct {
	Decls []Declaration
}

func (p *FileAst) astNode() {}