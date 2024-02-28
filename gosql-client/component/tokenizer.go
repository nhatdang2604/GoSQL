package component

import (
	"strings"
)

type Tokenizer interface {
	Tokenize(input string) []string
}

type Lexer struct{}

func (Lexer) Tokenize(input string) []string {
	var tokens []string = strings.Split(input, " ")
	return tokens
}
