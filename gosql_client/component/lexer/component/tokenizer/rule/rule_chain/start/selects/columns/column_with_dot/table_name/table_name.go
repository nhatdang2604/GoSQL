package table_name

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/froms"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_pool"
	"strings"
)

type TableNameRule struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *TableNameRule) Exec(toks []string) bool {

	var isSuccess bool = false
	var firstTok string = toks[0]
	if hasCommaRule := c.pool.Get(constants.RULE_HAS_COMMA); hasCommaRule.Validate(firstTok) {
		isSuccess = c.setAsComma(firstTok)
	} else {
		isSuccess = c.setAsFrom(firstTok)
	}

	return isSuccess
}

func (c *TableNameRule) setAsComma(tok string) bool {
	var comma string = string(constants.SYMBOL_COMMA)
	var splits []string = strings.Split(tok, comma)

	//More than 1 comma check
	if len(splits) > 1 {
		c.errMsg = fmt.Sprintf("More than 1 comma in '%s'", tok)
		return false
	}

	var tableName string = splits[0]
	c.curTok = tableName

	//Remove the tok in the current tokens
	var remainToksWithoutCurTok []string = c.remainToks[1:]
	var remainToks []string = append([]string{comma}, remainToksWithoutCurTok...)
	c.remainToks = remainToks

	c.nextRuleChain = froms.New(c.pool)

	return true
}

func (c *TableNameRule) setAsFrom(tok string) bool {
	c.curTok = tok
	c.remainToks = c.remainToks[1:]
	c.nextRuleChain = froms.New(c.pool)
	return true
}

func (c *TableNameRule) EmitTok() string {
	return c.curTok
}

func (c *TableNameRule) RemainToks() []string {
	return c.remainToks
}

func (c *TableNameRule) ErrorMsg() string {
	return c.errMsg
}

func (c *TableNameRule) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &TableNameRule{
		pool: pool,
	}
}
