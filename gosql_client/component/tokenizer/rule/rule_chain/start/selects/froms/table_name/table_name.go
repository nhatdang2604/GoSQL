package table_name

import (
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
	"strings"
)

type TableNameChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *TableNameChain) Exec(toks []string) bool {
	c.curTok = nil
	var isSuccess bool = false
	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		if c.hasSemicolon(firstTok) {
			c.setAsNextRuleChainIsSemicolon(toks)
		} else {
			c.setAsNextRuleChainIsAs(toks)
		}
	}

	return isSuccess
}

func (c *TableNameChain) Validate(toks []string) bool {
	return true
}

func (c *TableNameChain) hasSemicolon(tok string) bool {
	hasSemicolonRule := c.pool.Get(constants.RULE_HAS_SEMICOLON)
	var hasSemicolon bool = hasSemicolonRule.Validate(rule_input.SingleTok{Tok: tok})
	return hasSemicolon
}

func (c *TableNameChain) setAsNextRuleChainIsSemicolon(toks []string) bool {
	var tok string = toks[0]
	var semicolon string = string(constants.SYMBOL_SEMICOLON)
	var splits []string = strings.Split(tok, semicolon)
	var tableName string = splits[0]

	toks[0] = semicolon // transform token 'b;' into ';'
	c.curTok = &tableName
	c.remainToks = toks
	return true
}

func (c *TableNameChain) setAsNextRuleChainIsAs(toks []string) bool {
	var tableName = toks[0]
	c.curTok = &tableName
	c.remainToks = toks[1:]

	return true
}

func (c *TableNameChain) EmitTok() *string {
	return c.curTok
}

func (c *TableNameChain) RemainToks() []string {
	return c.remainToks
}

func (c *TableNameChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *TableNameChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *TableNameChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &TableNameChain{
		pool: pool,
	}
}
