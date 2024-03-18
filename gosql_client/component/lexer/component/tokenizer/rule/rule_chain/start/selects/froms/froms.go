package froms

import (
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_pool"
)

type FromChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *FromChain) Exec(toks []string) bool {

	var isSuccess bool = false

	//TODO

	return isSuccess
}

func (c *FromChain) EmitTok() string {
	return c.curTok
}

func (c *FromChain) RemainToks() []string {
	return c.remainToks
}

func (c *FromChain) ErrorMsg() string {
	return c.errMsg
}

func (c *FromChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &FromChain{
		pool: pool,
	}
}
