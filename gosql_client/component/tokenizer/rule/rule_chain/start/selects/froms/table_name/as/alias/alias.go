package alias

import (
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
	"strings"
)

type AliasChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *AliasChain) Exec(toks []string) bool {

	c.curTok = nil

	var isSuccess bool = false
	if isSuccess = c.Validate(toks); isSuccess {
		var tok string = toks[0]
		var semicolon string = string(constants.SYMBOL_SEMICOLON)
		var splits []string = strings.Split(tok, semicolon)
		var alias string = splits[0]

		toks[0] = semicolon // transform token 'b;' into ';'
		c.curTok = &alias
		c.remainToks = toks
	}

	return isSuccess
}

func (c *AliasChain) Validate(toks []string) bool {
	var tok string = toks[0]
	hasSemicolonRule := c.pool.Get(constants.RULE_HAS_SEMICOLON)
	var hasSemicolon bool = hasSemicolonRule.Validate(rule_input.SingleTok{Tok: tok})

	if !hasSemicolon {
		var msg string = hasSemicolonRule.ErrorMsg()
		c.errMsg = &msg
	}

	return hasSemicolon
}

func (c *AliasChain) EmitTok() *string {
	return c.curTok
}

func (c *AliasChain) RemainToks() []string {
	return c.remainToks
}

func (c *AliasChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *AliasChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *AliasChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &AliasChain{
		pool: pool,
	}
}
