package column_without_dot

import (
	"fmt"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/columns/common/is_comma"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/froms"
	"gosql_client/component/tokenizer/rule/rule_pool"
	"gosql_client/component/tokenizer/rule/rule_unit"
	"strings"
)

type ColumnWithoutDotChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *ColumnWithoutDotChain) Exec(toks []string) bool {

	var isSuccess bool = false
	var firstTok string = toks[0]
	if c.hasComma(firstTok) {
		isSuccess = c.setAsComma(toks)
	} else {
		isSuccess = c.setAsFrom(toks)
	}

	return isSuccess
}

func (c *ColumnWithoutDotChain) hasComma(tok string) bool {
	var hasCommaRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_COMMA)
	var hasComma bool = hasCommaRule.Validate(tok)

	return hasComma
}

func (c *ColumnWithoutDotChain) setAsComma(toks []string) bool {
	var tok string = toks[0]
	var comma string = string(constants.SYMBOL_COMMA)
	var splits []string = strings.Split(tok, comma)

	//More than 1 comma check
	if len(splits) > 1 {
		c.errMsg = fmt.Sprintf("More than 1 comma in '%s'", tok)
		return false
	}

	var tableName string = splits[0]

	toks[0] = comma // transform token 'b,' into ','
	c.curTok = tableName
	c.remainToks = toks
	c.nextRuleChain = is_comma.SharedIsCommaChain
	return true
}

func (c *ColumnWithoutDotChain) setAsFrom(toks []string) bool {
	c.curTok = toks[0]
	c.remainToks = toks[1:]
	c.nextRuleChain = froms.New(c.pool)
	return true
}

func (c *ColumnWithoutDotChain) IsValid(tok string) bool {
	return true
}

func (c *ColumnWithoutDotChain) ToRuleChain() rule_chain.RuleChain {
	return c
}

func (c *ColumnWithoutDotChain) EmitTok() string {
	return c.curTok
}

func (c *ColumnWithoutDotChain) RemainToks() []string {
	return c.remainToks
}

func (c *ColumnWithoutDotChain) ErrorMsg() string {
	return c.errMsg
}

func (c *ColumnWithoutDotChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) *ColumnWithoutDotChain {
	return &ColumnWithoutDotChain{
		pool: pool,
	}
}
