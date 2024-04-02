package constants

import "gosql_client/component/tokenizer/alias"

const (
	KEYWORD_START  alias.Keyword = "start_lexical"
	KEYWORD_SELECT alias.Keyword = "select"
	KEYWORD_FROM   alias.Keyword = "from"
	KEYWORD_AS     alias.Keyword = "as"
	KEYWORD_INSERT alias.Keyword = "insert"
	KEYWORD_INTO   alias.Keyword = "into"
	KEYWORD_VALUES alias.Keyword = "values"
)

var KEYWORDS []alias.Keyword = []alias.Keyword{
	KEYWORD_START,
	KEYWORD_SELECT,
	KEYWORD_FROM,
	KEYWORD_AS,
	KEYWORD_INSERT,
	KEYWORD_INTO,
	KEYWORD_VALUES,
}
