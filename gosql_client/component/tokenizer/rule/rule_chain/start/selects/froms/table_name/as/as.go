package as

import (
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type AsChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *AsChain) Exec(toks []string) bool {

	var isSuccess bool = false
	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		c.curTok = &firstTok
		c.remainToks = toks[1:]
	}

	return isSuccess
}

func (c *AsChain) Validate(toks []string) bool {
	var tok string = toks[0]
	asRule := c.pool.Get(constants.RULE_IS_AS)
	var isAs bool = asRule.Validate(rule_input.SingleTok{Tok: tok})

	if !isAs {
		var msg string = asRule.ErrorMsg()
		c.errMsg = &msg
	}

	return isAs
}

func (c *AsChain) EmitTok() *string {
	return c.curTok
}

func (c *AsChain) TokType() alias.TokType {
	return constants.TOKTYPE_KEYWORD
}

func (c *AsChain) RemainToks() []string {
	return c.remainToks
}

func (c *AsChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *AsChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *AsChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &AsChain{
		pool: pool,
	}
}
