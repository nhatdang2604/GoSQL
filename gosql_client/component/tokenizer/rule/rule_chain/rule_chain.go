package rule_chain

type RuleChain interface {
	Exec(toks []string) bool
	Validate(toks []string) bool
	EmitTok() *string
	RemainToks() []string
	ErrorMsg() *string

	SetNextRuleChain(nextRuleChain RuleChain)
	NextRuleChain() RuleChain
}
