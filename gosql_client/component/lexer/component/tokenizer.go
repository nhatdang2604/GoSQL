package component

import "strings"

type Tokenizable interface {
	Tokenize(input string) []string
}

type Tokenizer struct{}

func (Tokenizer) Tokenize(input string) []string {
	var tokens []string = strings.Split(input, " ")
	return tokens
}
