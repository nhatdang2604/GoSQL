package froms

import (
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type FromChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *FromChain) Exec(toks []string) bool {

	c.curTok = nil

	var isSuccess bool = false
	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		c.curTok = &firstTok
		c.remainToks = toks[1:]
	}

	return isSuccess
}

func (c *FromChain) Validate(toks []string) bool {
	var tok string = toks[0]
	fromRule := c.pool.Get(constants.RULE_IS_FROM)
	var isFrom bool = fromRule.Validate(rule_input.SingleTok{Tok: tok})

	if !isFrom {
		var msg string = fromRule.ErrorMsg()
		c.errMsg = &msg
	}

	return isFrom
}

func (c *FromChain) EmitTok() *string {
	return c.curTok
}

func (c *FromChain) TokType() alias.TokType {
	return constants.TOKTYPE_KEYWORD
}

func (c *FromChain) RemainToks() []string {
	return c.remainToks
}

func (c *FromChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *FromChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *FromChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &FromChain{
		pool: pool,
	}
}
