package main

import (
	"fmt"

	component "gosql-client/component"
)

func main() {
	var stringToTest string = "test 0 12 23321321"
	var tokenizer component.Tokenizer = component.Lexer{}

	var tokens []string = tokenizer.Tokenize(stringToTest)
	for index, element := range tokens {
		fmt.Printf("Test %v: %v\r\n", index, element)
	}
}
