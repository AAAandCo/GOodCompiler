package main

import (
	"fmt"
	"math"
	"token"
	"lexer"
)

func main() {

	fmt.Printf(token.EOF.String());
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}

