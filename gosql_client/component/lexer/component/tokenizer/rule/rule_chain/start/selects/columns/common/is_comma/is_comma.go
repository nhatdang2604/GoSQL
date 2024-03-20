package is_comma

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_pool"
)

type IsCommaChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *IsCommaChain) Exec(toks []string) bool {
	c.remainToks = toks
	var isSuccess bool = false
	var firstTok string = toks[0]

	if c.isComma(firstTok) {
		c.curTok = firstTok
		c.remainToks = toks[1:]
		isSuccess = c.setNextRule(c.remainToks)
	}
	return isSuccess
}

func (c *IsCommaChain) isComma(tok string) bool {
	var isCommaRule rule_pool.Rule = c.pool.Get(constants.RULE_IS_COMMA)
	var isComma bool = isCommaRule.Validate(tok)

	return isComma
}

func (c *IsCommaChain) setNextRule(toks []string) bool {
	var isSuccess bool = false
	var nextTok string = toks[0]

	//TODO:

	if !isSuccess {
		c.errMsg = fmt.Sprintf("Unexpected keyword '%s' after `%b`", nextTok, constants.SYMBOL_COMMA)
	}

	return isSuccess
}

func (c *IsCommaChain) EmitTok() string {
	return c.curTok
}

func (c *IsCommaChain) RemainToks() []string {
	return c.remainToks
}

func (c *IsCommaChain) ErrorMsg() string {
	return c.errMsg
}

func (c *IsCommaChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &IsCommaChain{
		pool: pool,
	}
}
