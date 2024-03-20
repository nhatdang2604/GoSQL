package column_rule_chain

import "gosql_client/component/lexer/component/tokenizer/rule/rule_chain"

type ColumnRuleChain interface {
	IsValid(tok string) bool
	ToRuleChain() rule_chain.RuleChain
}
