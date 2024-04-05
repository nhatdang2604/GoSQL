package start

import (
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type StartChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *StartChain) Exec(toks []string) bool {

	c.curTok = nil
	c.remainToks = toks
	var isSuccess bool = c.Validate(toks)
	return isSuccess
}

func (c *StartChain) Validate(toks []string) bool {
	return true
}

func (c *StartChain) EmitTok() *string {
	return c.curTok
}

func (c *StartChain) RemainToks() []string {
	return c.remainToks
}

func (c *StartChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *StartChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *StartChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &StartChain{
		pool: pool,
	}
}
