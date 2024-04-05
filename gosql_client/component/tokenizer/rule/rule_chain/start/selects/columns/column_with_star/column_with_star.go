package column_with_star

import (
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type ColumnWithStarChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *ColumnWithStarChain) Exec(toks []string) bool {
	var isSuccess bool = false

	c.remainToks = toks
	c.curTok = nil

	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		c.curTok = &firstTok
		c.remainToks = toks[1:]
	}

	return isSuccess
}

func (c *ColumnWithStarChain) Validate(toks []string) bool {
	var tok string = toks[0]
	isStarRule := c.pool.Get(constants.RULE_IS_STAR)
	var isStar bool = isStarRule.Validate(rule_input.SingleTok{Tok: tok})

	if !isStar {
		var msg string = isStarRule.ErrorMsg()
		c.errMsg = &msg
	}

	return isStar
}

func (c *ColumnWithStarChain) ToRuleChain() rule_chain.RuleChain {
	return c
}

func (c *ColumnWithStarChain) EmitTok() *string {
	return c.curTok
}

func (c *ColumnWithStarChain) RemainToks() []string {
	return c.remainToks
}

func (c *ColumnWithStarChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *ColumnWithStarChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *ColumnWithStarChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) *ColumnWithStarChain {
	return &ColumnWithStarChain{
		pool: pool,
	}
}
