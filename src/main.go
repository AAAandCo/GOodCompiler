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

	v.VisitVar(&ast.VarDecl{
		Name: &ast.Ident{
			Name: "Test Identifier",
		},
	})

	//fmt.Println("String represetation of the AST:")
	fmt.Println(fileAst)
}