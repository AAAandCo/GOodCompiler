package main

import (
	"fmt"
	"lexer"
)

func main() {
	lexerObj := new(lexer.Lexer)
    tokens, errors := lexerObj.ParseTokens(" var = !ert /*df*/ {} [] ()  1 + 3 4534 78 sin(3)")

	fmt.Println(tokens)
	fmt.Println(errors)
}

