package main

import (
	"expression_parser"
	"fmt"
	"lexer"
)

func main() {
	source := "1+(2*3)-4;(1+2)*(6/2);1+2*3-4/2;"
	my_lexer := lexer.NewLexer(source)
	parser := expression_parser.NewExpressionParser(my_lexer)
	parser.Parse()
	fmt.Println("\nparsing end here")
}
