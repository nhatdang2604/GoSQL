package inserts

import (
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type InsertChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *InsertChain) Exec(toks []string) bool {

	c.curTok = nil

	var isSuccess bool = false
	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		c.curTok = &firstTok
		c.remainToks = toks[1:]
	}

	return isSuccess
}

func (c *InsertChain) Validate(toks []string) bool {
	var tok string = toks[0]
	isInsertRule := c.pool.Get(constants.RULE_IS_INSERT)
	isInsert := isInsertRule.Validate(rule_input.SingleTok{Tok: tok})

	if !isInsert {
		var msg string = isInsertRule.ErrorMsg()
		c.errMsg = &msg
	}

	return isInsert
}

func (c *InsertChain) EmitTok() *string {
	return c.curTok
}

func (c *InsertChain) RemainToks() []string {
	return c.remainToks
}

func (c *InsertChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *InsertChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *InsertChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &InsertChain{
		pool: pool,
	}
}
