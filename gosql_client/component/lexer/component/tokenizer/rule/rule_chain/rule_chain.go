package rule_chain

type RuleChain interface {
	Exec(toks []string) bool
	EmitTok() string
	RemainToks() []string
	ErrorMsg() string

	NextRuleChain() RuleChain
}
