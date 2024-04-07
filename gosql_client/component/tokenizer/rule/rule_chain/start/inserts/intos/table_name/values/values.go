package values

import (
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type ValueChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *ValueChain) Exec(toks []string) bool {

	c.curTok = nil

	var isSuccess bool = false
	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		c.curTok = &firstTok
		c.remainToks = toks[1:]
	}

	return isSuccess
}

func (c *ValueChain) Validate(toks []string) bool {
	var tok string = toks[0]
	isValueRule := c.pool.Get(constants.RULE_IS_VALUES)
	isValue := isValueRule.Validate(rule_input.SingleTok{Tok: tok})

	if !isValue {
		var msg string = isValueRule.ErrorMsg()
		c.errMsg = &msg
	}

	return isValue
}

func (c *ValueChain) EmitTok() *string {
	return c.curTok
}

func (c *ValueChain) TokType() alias.TokType {
	return constants.TOKTYPE_KEYWORD
}

func (c *ValueChain) RemainToks() []string {
	return c.remainToks
}

func (c *ValueChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *ValueChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *ValueChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &ValueChain{
		pool: pool,
	}
}
