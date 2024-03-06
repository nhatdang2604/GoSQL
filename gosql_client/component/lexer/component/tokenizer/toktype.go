package tokenizer

import "gosql_client/component/lexer/component/tokenizer/alias"

const (
	TOKTYPE_KEYWORD    alias.TokType = "KEYWORD"
	TOKTYPE_SYMBOL     alias.TokType = "SYMBOL"
	TOKTYPE_IDENTIFIER alias.TokType = "IDENTIFIER"
	TOKTYPE_INT_CONST  alias.TokType = "INT_CONST"
	TOKTYPE_STR_CONST  alias.TokType = "STR_CONST"
	TOKTYPE_BOO_CONST  alias.TokType = "BOO_CONST"
)
