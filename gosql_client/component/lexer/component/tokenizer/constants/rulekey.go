package constants

import "gosql_client/component/lexer/component/tokenizer/alias"

const (
	RULE_START_TO_TOKENIZE alias.RuleKey = "start_to_tokenize"
	RULE_IS_SELECT         alias.RuleKey = "is_select"
	RULE_IS_INSERT         alias.RuleKey = "is_insert"
)
