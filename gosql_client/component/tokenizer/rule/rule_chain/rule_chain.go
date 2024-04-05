package rule_chain

import "gosql_client/component/tokenizer/alias"

type RuleChain interface {
	Exec(toks []string) bool
	Validate(toks []string) bool
	EmitTok() *string
	TokType() alias.TokType
	RemainToks() []string
	ErrorMsg() *string

	SetNextRuleChain(nextRuleChain RuleChain)
	NextRuleChain() RuleChain
}
