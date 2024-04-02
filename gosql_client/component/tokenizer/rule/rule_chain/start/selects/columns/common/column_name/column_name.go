package column_name

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

type ColumnNameChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *ColumnNameChain) Exec(toks []string) bool {

	var isSuccess bool = false
	var firstTok string = toks[0]
	if c.hasComma(firstTok) {
		isSuccess = c.setAsComma(toks)
	} else {
		isSuccess = c.setAsFrom(toks)
	}

	return isSuccess
}

func (c *ColumnNameChain) hasComma(tok string) bool {
	var hasCommaRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_COMMA)
	var hasComma bool = hasCommaRule.Validate(tok)

	return hasComma
}

func (c *ColumnNameChain) setAsComma(toks []string) bool {
	var tok string = toks[0]
	var comma string = string(constants.SYMBOL_COMMA)
	var splits []string = strings.Split(tok, comma)

	//More than 1 comma check
	if len(splits) > 1 {
		c.errMsg = fmt.Sprintf("More than 1 comma in '%s'", tok)
		return false
	}

	var columnName string = splits[0]

	toks[0] = comma // transform token 'b,' into ','
	c.curTok = columnName
	c.remainToks = toks
	c.nextRuleChain = is_comma.SharedIsCommaChain
	return true
}

func (c *ColumnNameChain) setAsFrom(toks []string) bool {
	c.curTok = toks[0]
	c.remainToks = toks[1:]
	c.nextRuleChain = froms.New(c.pool)
	return true
}

func (c *ColumnNameChain) EmitTok() string {
	return c.curTok
}

func (c *ColumnNameChain) RemainToks() []string {
	return c.remainToks
}

func (c *ColumnNameChain) ErrorMsg() string {
	return c.errMsg
}

func (c *ColumnNameChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &ColumnNameChain{
		pool: pool,
	}
}
