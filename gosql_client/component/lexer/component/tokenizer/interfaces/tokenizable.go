package interfaces

import (
	"gosql_client/component/lexer/component/tokenizer/alias"
)

type Tokenizable interface {
	HasMoreTokens() bool
	TokenType() alias.TokType
	Advance()

	keywordVal() alias.Keyword
	symbolVal() byte
	identifierVal() string
	intVal() int
	strVal() string
	booVal() bool
}
