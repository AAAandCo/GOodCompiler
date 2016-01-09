package ast

import (
	"fmt"
)

type Visitor interface {
	VisitVar(node *VarDecl) (w Visitor)
	VisitFunc(node *FuncDecl) (w Visitor)
}

type PrintVisitor struct {
	msg string
}

func (self *PrintVisitor) VisitVar(node *VarDecl) {
	fmt.Println("Variable Node. Name : " + node.Name.Name)
}

func (self *PrintVisitor) VisitFunc(node *FuncDecl) {
	fmt.Println("Some AST node.")
}