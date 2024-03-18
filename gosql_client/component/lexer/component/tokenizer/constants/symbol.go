package constants

import "gosql_client/component/lexer/component/tokenizer/alias"

const (
	SYMBOL_STAR  alias.Symbol = '*'
	SYMBOL_DOT   alias.Symbol = '.'
	SYMBOL_COMMA alias.Symbol = ','
)

var SYMBOLS []alias.Symbol = []alias.Symbol{
	SYMBOL_STAR,
	SYMBOL_DOT,
	SYMBOL_STAR,
}
