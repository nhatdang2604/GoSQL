package column_name

import (
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_pool"
	"gosql_client/component/tokenizer/rule/rule_unit"
	"strings"
)

type ColumnNameChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *ColumnNameChain) Exec(toks []string) bool {

	c.remainToks = toks
	c.curTok = nil

	var isSuccess bool = false
	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		if c.hasComma(firstTok) {
			c.setAsNextRuleChainIsComma(toks)
		} else {
			c.setAsNextRuleChainIsFrom(toks)
		}
	}

	return isSuccess
}

// Always return true in this case
func (c *ColumnNameChain) Validate(toks []string) bool {
	return true
}

func (c *ColumnNameChain) hasComma(tok string) bool {
	var hasCommaRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_COMMA)
	var hasComma bool = hasCommaRule.Validate(tok)

	return hasComma
}

func (c *ColumnNameChain) setAsNextRuleChainIsComma(toks []string) bool {
	var tok string = toks[0]
	var comma string = string(constants.SYMBOL_COMMA)
	var splits []string = strings.Split(tok, comma)
	var columnName string = splits[0]
	var rmColumnNameFromSplits = splits[1:]
	var afterRmColumName = strings.Join(rmColumnNameFromSplits, comma)

	toks[0] = afterRmColumName
	c.curTok = &columnName
	c.remainToks = toks
	return true
}

func (c *ColumnNameChain) setAsNextRuleChainIsFrom(toks []string) bool {
	var firstTok string = toks[0]
	c.curTok = &firstTok
	c.remainToks = toks[1:]
	return true
}

func (c *ColumnNameChain) EmitTok() *string {
	return c.curTok
}

func (c *ColumnNameChain) TokType() alias.TokType {
	return constants.TOKTYPE_IDENTIFIER
}

func (c *ColumnNameChain) RemainToks() []string {
	return c.remainToks
}

func (c *ColumnNameChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *ColumnNameChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *ColumnNameChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func (c *ColumnNameChain) SetRulePool(pool rule_pool.RulePool) {
	c.pool = pool
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &ColumnNameChain{
		pool: pool,
	}
}
