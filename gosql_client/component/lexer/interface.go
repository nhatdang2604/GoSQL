package lexer

type Tokenizer interface {
	Tokenize(input string) []string
}
