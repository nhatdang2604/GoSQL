package start

import (
	"fmt"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_chain/start/inserts"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type StartChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *StartChain) Exec(toks []string) bool {

	c.curTok = constants.KEYWORD_START
	c.remainToks = toks
	var isSuccess bool = false

	if c.isStart(c.curTok) {

		if len(toks) == 0 {
			var firstTok = ""
			return c.setAsInvalid(firstTok)
		}

		var firstTok = toks[0]
		if c.isNextRuleSelect(firstTok) {
			c.nextRuleChain = selects.New(c.pool)
			isSuccess = true
		} else if c.isNextRuleInsert(firstTok) {
			c.nextRuleChain = inserts.New(c.pool)
			isSuccess = true
		} else {
			isSuccess = c.setAsInvalid(firstTok)
		}
	}

	return isSuccess
}

func (c *StartChain) isStart(tok string) bool {
	startRule := c.pool.Get(constants.RULE_IS_START)
	var isSuccess bool = startRule.Validate(rule_input.SingleTok{Tok: tok})
	return isSuccess
}

func (c *StartChain) isNextRuleSelect(tok string) bool {
	selectRule := c.pool.Get(constants.RULE_IS_SELECT)
	var isSelect bool = selectRule.Validate(rule_input.SingleTok{Tok: tok})
	return isSelect
}

func (c *StartChain) isNextRuleInsert(tok string) bool {
	insertRule := c.pool.Get(constants.RULE_IS_INSERT)
	var isInsert bool = insertRule.Validate(rule_input.SingleTok{Tok: tok})
	return isInsert
}

func (c *StartChain) setAsInvalid(tok string) bool {
	c.errMsg = fmt.Sprintf("Expected `%s` or `%s` but found `%s`", constants.KEYWORD_SELECT, constants.KEYWORD_INSERT, tok)
	c.nextRuleChain = nil
	return false
}

func (c *StartChain) EmitTok() string {
	return c.curTok
}

func (c *StartChain) RemainToks() []string {
	return c.remainToks
}

func (c *StartChain) ErrorMsg() string {
	return c.errMsg
}

func (c *StartChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &StartChain{
		pool: pool,
	}
}
