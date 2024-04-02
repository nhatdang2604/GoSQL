package interfaces

import (
	"gosql_client/component/tokenizer/alias"
)

type Tokenizable interface {
	HasMoreTokens() bool
	TokenType() alias.TokType
	Advance()
	CurrentError() error

	KeywordVal() alias.Keyword
	SymbolVal() (byte, error)
	IdentifierVal() (string, error)
	IntVal() (int, error)
	StrVal() (string, error)
	BooVal() (bool, error)
}
