// +build !with_tests

package main

import (
	"fmt"
	"lexer"
)

func main1() {
	lexerObj := new(lexer.Lexer)
    tokens, errors := lexerObj.ParseTokens(" \"string value return\" var = !ert df {} [] ()  1 + 3 4534 78")

	fmt.Println(tokens)
	fmt.Println(errors)
}

