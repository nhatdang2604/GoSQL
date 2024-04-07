package table_name

import (
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type TableNameChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *TableNameChain) Exec(toks []string) bool {

	c.remainToks = toks
	c.curTok = nil

	var isSuccess bool = false
	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		c.curTok = &firstTok
		c.remainToks = toks[1:]
	}

	return isSuccess
}

// Always return true in this case
func (c *TableNameChain) Validate(toks []string) bool {
	return true
}

func (c *TableNameChain) EmitTok() *string {
	return c.curTok
}

func (c *TableNameChain) TokType() alias.TokType {
	return constants.TOKTYPE_IDENTIFIER
}

func (c *TableNameChain) RemainToks() []string {
	return c.remainToks
}

func (c *TableNameChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *TableNameChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *TableNameChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func (c *TableNameChain) SetRulePool(pool rule_pool.RulePool) {
	c.pool = pool
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &TableNameChain{
		pool: pool,
	}
}
