package is_semicolon

import (
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type IsSemicolonChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *IsSemicolonChain) Exec(toks []string) bool {
	c.curTok = nil

	var isSuccess bool = false

	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		c.curTok = &firstTok
		c.remainToks = toks[1:]
	}
	return isSuccess
}

func (c *IsSemicolonChain) Validate(toks []string) bool {
	var tok string = toks[0]
	var isSemicolonRule rule_pool.Rule = c.pool.Get(constants.RULE_IS_COMMA)
	var isSemicolon bool = isSemicolonRule.Validate(tok)

	if !isSemicolon {
		var msg string = isSemicolonRule.ErrorMsg()
		c.errMsg = &msg
	}

	return isSemicolon
}

func (c *IsSemicolonChain) EmitTok() *string {
	return c.curTok
}

func (c *IsSemicolonChain) TokType() alias.TokType {
	return constants.TOKTYPE_SYMBOL
}

func (c *IsSemicolonChain) RemainToks() []string {
	return c.remainToks
}

func (c *IsSemicolonChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *IsSemicolonChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *IsSemicolonChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &IsSemicolonChain{
		pool: pool,
	}
}
