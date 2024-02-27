package main

import (
	"fmt"

	"gosql/src/client/component"
)

func main() {
	var stringToTest string = "test 0 12 23321321"
	var tokenizer component.Tokenizer = &component.Tokenizer{}

	fmt.Printf("Test: v", tokenizer.Tokenize(stringToTest))
}
