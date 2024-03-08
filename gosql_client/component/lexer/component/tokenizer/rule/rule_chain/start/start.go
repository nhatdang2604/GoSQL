package start

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/inserts"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_pool"
)

type StartChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *StartChain) Exec(toks []string) bool {

	c.remainToks = toks
	var firstTok = toks[0]
	var isSuccess bool = false

	if selectRule := c.pool.Get(constants.RULE_IS_SELECT); selectRule.Validate(rule_input.SingleTok{Tok: firstTok}) {
		isSuccess = c.setAsSelect(firstTok)
	} else if insertRule := c.pool.Get(constants.RULE_IS_INSERT); insertRule.Validate(rule_input.SingleTok{Tok: firstTok}) {
		isSuccess = c.setAsInsert(firstTok)
	} else {
		isSuccess = c.setAsInvalid(firstTok)
	}

	return isSuccess
}

func (c *StartChain) setAsSelect(tok string) bool {
	c.curTok = tok
	c.remainToks = c.remainToks[1:]
	c.nextRuleChain = selects.New(c.pool)
	return true
}

func (c *StartChain) setAsInsert(tok string) bool {
	c.curTok = tok
	c.remainToks = c.remainToks[1:]
	c.nextRuleChain = inserts.New(c.pool)
	return true
}

func (c *StartChain) setAsInvalid(tok string) bool {
	c.nextRuleChain = nil
	c.errMsg = fmt.Sprintf("Expected `%s` or `%s` but found `%s`", constants.KEYWORD_SELECT, constants.KEYWORD_INSERT, tok)
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
