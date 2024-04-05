package is_comma

import (
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/columns/interfaces/column_rule_chain"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

var (
	SharedIsCommaChain *IsCommaChain
)

func (c *IsCommaChain) AddColumnRuleChain(columnRuleChain column_rule_chain.ColumnRuleChain) {
	c.columnRuleChainPool = append(c.columnRuleChainPool, columnRuleChain)
}

type IsCommaChain struct {
	nextRuleChain       rule_chain.RuleChain
	pool                rule_pool.RulePool
	curTok              *string
	errMsg              *string
	remainToks          []string
	columnRuleChainPool []column_rule_chain.ColumnRuleChain
}

func (c *IsCommaChain) Exec(toks []string) bool {
	c.curTok = nil

	var isSuccess bool = false
	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		c.curTok = &firstTok
		c.remainToks = toks[1:]
	}
	return isSuccess
}

func (c *IsCommaChain) Validate(toks []string) bool {
	var tok string = toks[0]
	var isCommaRule rule_pool.Rule = c.pool.Get(constants.RULE_IS_COMMA)
	var isComma bool = isCommaRule.Validate(tok)

	if !isComma {
		var msg string = isCommaRule.ErrorMsg()
		c.errMsg = &msg
	}

	return isComma
}

func (c *IsCommaChain) EmitTok() *string {
	return c.curTok
}

func (c *IsCommaChain) RemainToks() []string {
	return c.remainToks
}

func (c *IsCommaChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *IsCommaChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *IsCommaChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) *IsCommaChain {
	return &IsCommaChain{
		columnRuleChainPool: []column_rule_chain.ColumnRuleChain{},
		pool:                pool,
	}
}
