package rule_unit

import "gosql_client/component/tokenizer/alias"

type Rule interface {
	Key() alias.RuleKey
	Validate(i interface{}) bool
	ErrorMsg() string
}
