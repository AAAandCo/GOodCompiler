package parser

import (
	"ast"
	"io/ioutil"
	"fmt"
)

func readSource(filename string) (string, error) {
	fmt.Println("interface.readSource() exec")
	src, err := ioutil.ReadFile(filename)
	return string(src), err
}

func ParseFile(filename string) (f *ast.FileAst, err error) {
	fmt.Println("interface.ParseFile() exec")
	text, err := readSource(filename)
	var p parser
	p.init(text)

	resultAst := p.parseFile();

	return resultAst, err
}
