package dot_at_first

import (
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
	"gosql_client/component/tokenizer/rule/rule_unit"
)

type DotAtFirstChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *DotAtFirstChain) Exec(toks []string) bool {
	var isSuccess bool = false

	c.remainToks = toks
	c.curTok = nil

	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		var dot string = string(constants.SYMBOL_DOT)

		c.curTok = &dot
		toks[0] = c.rmDotFromTok(firstTok) //replace the ".b" with "b"
		c.remainToks = toks
	}
	return isSuccess
}

func (c *DotAtFirstChain) Validate(toks []string) bool {
	var tok string = toks[0]
	var hasDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_DOT)
	var hasOnlyOneDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_ONLY_ONE_DOT)
	var hasDotAtFirstRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_DOT_AT_FIRST)

	var hasDot = hasDotRule.Validate(rule_input.SingleTok{Tok: tok})
	var hasOnlyOneDot = hasOnlyOneDotRule.Validate(rule_input.SingleTok{Tok: tok})
	var hasDotAtFirst = hasDotAtFirstRule.Validate(rule_input.SingleTok{Tok: tok})

	if !hasDot {
		var msg string = hasDotRule.ErrorMsg()
		c.errMsg = &msg
	} else if !hasOnlyOneDot {
		var msg string = hasOnlyOneDotRule.ErrorMsg()
		c.errMsg = &msg
	} else if !hasDotAtFirst {
		var msg string = hasDotAtFirstRule.ErrorMsg()
		c.errMsg = &msg
	}

	return hasDot && hasOnlyOneDot && hasDotAtFirst
}

func (c *DotAtFirstChain) rmDotFromTok(tok string) string {
	return tok[1:]
}

func (c *DotAtFirstChain) EmitTok() *string {
	return c.curTok
}

func (c *DotAtFirstChain) TokType() alias.TokType {
	return constants.TOKTYPE_SYMBOL
}

func (c *DotAtFirstChain) RemainToks() []string {
	return c.remainToks
}

func (c *DotAtFirstChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *DotAtFirstChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *DotAtFirstChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &DotAtFirstChain{
		pool: pool,
	}
}
