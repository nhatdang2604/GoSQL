package froms

import (
	"fmt"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/froms/table_name"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type FromChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *FromChain) Exec(toks []string) bool {

	var isSuccess bool = false
	var firstTok string = toks[0]
	if c.isFrom(firstTok) {
		c.curTok = firstTok
		c.remainToks = toks[1:]
		isSuccess = c.setNextRule(c.remainToks)
	}

	if !isSuccess && c.errMsg != "" {
		c.errMsg = fmt.Sprintf("Expected '%s' keyword, found '%s'", constants.KEYWORD_FROM, firstTok)
	}

	return isSuccess
}

func (c *FromChain) isFrom(tok string) bool {
	fromRule := c.pool.Get(constants.RULE_IS_FROM)
	var isFrom bool = fromRule.Validate(rule_input.SingleTok{Tok: tok})
	return isFrom
}

func (c *FromChain) setNextRule(toks []string) bool {
	c.nextRuleChain = table_name.New(c.pool)
	return true
}

func (c *FromChain) EmitTok() string {
	return c.curTok
}

func (c *FromChain) RemainToks() []string {
	return c.remainToks
}

func (c *FromChain) ErrorMsg() string {
	return c.errMsg
}

func (c *FromChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &FromChain{
		pool: pool,
	}
}
