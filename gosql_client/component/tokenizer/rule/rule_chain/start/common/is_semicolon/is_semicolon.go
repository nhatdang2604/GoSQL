package is_semicolon

import (
	"fmt"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/columns/interfaces/column_rule_chain"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type IsSemicolonChain struct {
	nextRuleChain       rule_chain.RuleChain
	pool                rule_pool.RulePool
	curTok              string
	errMsg              string
	remainToks          []string
	columnRuleChainPool []column_rule_chain.ColumnRuleChain
}

func (c *IsSemicolonChain) Exec(toks []string) bool {
	c.remainToks = toks
	var isSuccess bool = false
	var firstTok string = toks[0]

	if c.isSemicolon(firstTok) {
		c.curTok = firstTok
		c.remainToks = toks[1:]
		isSuccess = c.setNextRule(c.remainToks)
	}
	return isSuccess
}

func (c *IsSemicolonChain) isSemicolon(tok string) bool {
	var isSemicolonRule rule_pool.Rule = c.pool.Get(constants.RULE_IS_COMMA)
	var isSemicolon bool = isSemicolonRule.Validate(tok)

	return isSemicolon
}

func (c *IsSemicolonChain) setNextRule(toks []string) bool {
	var isSuccess bool = false
	var nextTok string = ""

	if len(toks) == 0 {
		isSuccess = true
		c.nextRuleChain = nil
	} else {
		nextTok = toks[0]
	}

	if !isSuccess {
		c.errMsg = fmt.Sprintf("Unexpected keyword '%s' after `%b`", nextTok, constants.SYMBOL_SEMICOLON)
	}

	return isSuccess
}

func (c *IsSemicolonChain) EmitTok() string {
	return c.curTok
}

func (c *IsSemicolonChain) RemainToks() []string {
	return c.remainToks
}

func (c *IsSemicolonChain) ErrorMsg() string {
	return c.errMsg
}

func (c *IsSemicolonChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) *IsSemicolonChain {
	return &IsSemicolonChain{
		columnRuleChainPool: []column_rule_chain.ColumnRuleChain{},
		pool:                pool,
	}
}
