package main

import (
	"fmt"
	"parser"
	"os"
	"ast"
)

func main() {
	if (len(os.Args) <= 1) {
		fmt.Println("Enter Filename with source code as first command line argument")
		return
	}

	filename := os.Args[1]
	fileAst, _ := parser.ParseFile(filename)

	v := &ast.PrintVisitor{}

	v.VisitVarDecl(&ast.VarDecl{
		Name: &ast.Ident{
			Name: "Test Identifier",
		},
	})

	fileAst.Accept(v)
	//fmt.Println("String represetation of the AST:")
	fmt.Println(fileAst)
}