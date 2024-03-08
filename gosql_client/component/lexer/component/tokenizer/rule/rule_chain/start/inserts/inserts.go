package inserts

import (
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_pool"
)

type InsertChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *InsertChain) Exec(toks []string) bool {

	var isSuccess bool = false

	//TODO:

	return isSuccess
}

func (c *InsertChain) EmitTok() string {
	return c.curTok
}

func (c *InsertChain) RemainToks() []string {
	return c.remainToks
}

func (c *InsertChain) ErrorMsg() string {
	return c.errMsg
}

func (c *InsertChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &InsertChain{
		pool: pool,
	}
}
