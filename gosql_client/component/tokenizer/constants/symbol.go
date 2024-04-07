package constants

import "gosql_client/component/tokenizer/alias"

const (
	SYMBOL_STAR         alias.Symbol = '*'
	SYMBOL_DOT          alias.Symbol = '.'
	SYMBOL_COMMA        alias.Symbol = ','
	SYMBOL_SEMICOLON    alias.Symbol = ';'
	SYMBOL_OPEN_BRACKET alias.Symbol = '('
)

var SYMBOLS []alias.Symbol = []alias.Symbol{
	SYMBOL_STAR,
	SYMBOL_DOT,
	SYMBOL_STAR,
	SYMBOL_SEMICOLON,
	SYMBOL_OPEN_BRACKET,
}
