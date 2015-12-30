package main

import (
	"fmt"
	"parser"
	"os"
)

func main() {
	if (len(os.Args) <= 1) {
		fmt.Println("Enter Filename with source code as first command line argument")
	}

	filename := os.Args[1]
	ast, err := parser.ParseFile(filename)

	fmt.Println(ast)
	fmt.Println(err)
}

