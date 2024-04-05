package selects

import (
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type SelectChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *SelectChain) Exec(toks []string) bool {
	var isSuccess bool = false
	c.curTok = nil

	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		c.curTok = &firstTok
		c.remainToks = toks[1:]
	}

	return isSuccess
}

func (c *SelectChain) Validate(toks []string) bool {
	var tok string = toks[0]
	selectRule := c.pool.Get(constants.RULE_IS_SELECT)
	var isSelect bool = selectRule.Validate(rule_input.SingleTok{Tok: tok})

	if !isSelect {
		var msg string = selectRule.ErrorMsg()
		c.errMsg = &msg
	}

	return isSelect
}

func (c *SelectChain) EmitTok() *string {
	return c.curTok
}

func (c *SelectChain) TokType() alias.TokType {
	return constants.TOKTYPE_KEYWORD
}

func (c *SelectChain) RemainToks() []string {
	return c.remainToks
}

func (c *SelectChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *SelectChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *SelectChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &SelectChain{
		pool:   pool,
		errMsg: nil,
	}
}
