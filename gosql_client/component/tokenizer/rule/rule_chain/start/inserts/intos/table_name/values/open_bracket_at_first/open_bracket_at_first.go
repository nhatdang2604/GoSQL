package open_bracket_at_first

import (
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
	"gosql_client/component/tokenizer/rule/rule_unit"
)

type OpenBracketAtFirstChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *OpenBracketAtFirstChain) Exec(toks []string) bool {
	var isSuccess bool = false

	c.remainToks = toks
	c.curTok = nil

	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		var dot string = string(constants.SYMBOL_OPEN_BRACKET)

		c.curTok = &dot
		toks[0] = c.rmOpenBracketFromTok(firstTok) //replace the "(b" with "b"
		c.remainToks = toks
	}
	return isSuccess
}

func (c *OpenBracketAtFirstChain) Validate(toks []string) bool {
	var tok string = toks[0]
	var hasOpenBracketAtFirstRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_DOT_AT_FIRST)
	var hasOpenBracketAtFirst = hasOpenBracketAtFirstRule.Validate(rule_input.SingleTok{Tok: tok})

	if !hasOpenBracketAtFirst {
		var msg string = hasOpenBracketAtFirstRule.ErrorMsg()
		c.errMsg = &msg
	}

	return hasOpenBracketAtFirst
}

func (c *OpenBracketAtFirstChain) rmOpenBracketFromTok(tok string) string {
	return tok[1:]
}

func (c *OpenBracketAtFirstChain) EmitTok() *string {
	return c.curTok
}

func (c *OpenBracketAtFirstChain) TokType() alias.TokType {
	return constants.TOKTYPE_SYMBOL
}

func (c *OpenBracketAtFirstChain) RemainToks() []string {
	return c.remainToks
}

func (c *OpenBracketAtFirstChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *OpenBracketAtFirstChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *OpenBracketAtFirstChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &OpenBracketAtFirstChain{
		pool: pool,
	}
}
