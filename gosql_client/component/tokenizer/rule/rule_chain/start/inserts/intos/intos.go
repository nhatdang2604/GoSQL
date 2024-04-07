package intos

import (
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type IntoChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *IntoChain) Exec(toks []string) bool {

	c.curTok = nil

	var isSuccess bool = false
	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		c.curTok = &firstTok
		c.remainToks = toks[1:]
	}

	return isSuccess
}

func (c *IntoChain) Validate(toks []string) bool {
	var tok string = toks[0]
	isIntoRule := c.pool.Get(constants.RULE_IS_INTO)
	isInto := isIntoRule.Validate(rule_input.SingleTok{Tok: tok})

	if !isInto {
		var msg string = isIntoRule.ErrorMsg()
		c.errMsg = &msg
	}

	return isInto
}

func (c *IntoChain) EmitTok() *string {
	return c.curTok
}

func (c *IntoChain) TokType() alias.TokType {
	return constants.TOKTYPE_KEYWORD
}

func (c *IntoChain) RemainToks() []string {
	return c.remainToks
}

func (c *IntoChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *IntoChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *IntoChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &IntoChain{
		pool: pool,
	}
}
