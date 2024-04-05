package column_with_dot

import (
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
	"gosql_client/component/tokenizer/rule/rule_unit"
	"strings"
)

type ColumnWithDotChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        *string
	errMsg        *string
	remainToks    []string
}

func (c *ColumnWithDotChain) Exec(toks []string) bool {
	var isSuccess bool = false

	c.curTok = nil

	if isSuccess = c.Validate(toks); isSuccess {
		var firstTok string = toks[0]
		var dbName string = c.getDbName(firstTok)
		c.curTok = &dbName
		toks[0] = c.rmDbNameFromTok(firstTok) //replace the "a.b" with ".b"
		c.remainToks = toks
	}

	return isSuccess
}

func (c *ColumnWithDotChain) Validate(toks []string) bool {
	var tok string = toks[0]
	var hasDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_DOT)
	var hasOnlyOneDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_ONLY_ONE_DOT)

	var hasDot = hasDotRule.Validate(rule_input.SingleTok{Tok: tok})
	var hasOnlyOneDot = hasOnlyOneDotRule.Validate(rule_input.SingleTok{Tok: tok})

	if !hasDot {
		var msg string = hasDotRule.ErrorMsg()
		c.errMsg = &msg
	} else if !hasOnlyOneDot {
		var msg string = hasOnlyOneDotRule.ErrorMsg()
		c.errMsg = &msg
	}

	return hasDot && hasOnlyOneDot
}

func (c *ColumnWithDotChain) getDbName(tok string) string {
	var splits = c.splitDbNameWithTableName(tok)
	var dbName = splits[0]
	return dbName
}

func (c *ColumnWithDotChain) splitDbNameWithTableName(tok string) []string {
	var dot string = string(constants.SYMBOL_DOT)
	var splits []string = strings.Split(tok, dot)
	return splits
}

func (c *ColumnWithDotChain) rmDbNameFromTok(tok string) string {
	var dot string = string(constants.SYMBOL_DOT)
	var splits []string = c.splitDbNameWithTableName(tok)
	var tableName string = splits[1]
	return dot + tableName
}

func (c *ColumnWithDotChain) EmitTok() *string {
	return c.curTok
}

func (c *ColumnWithDotChain) TokType() alias.TokType {
	return constants.TOKTYPE_IDENTIFIER
}

func (c *ColumnWithDotChain) RemainToks() []string {
	return c.remainToks
}

func (c *ColumnWithDotChain) ErrorMsg() *string {
	return c.errMsg
}

func (c *ColumnWithDotChain) SetNextRuleChain(nextRuleChain rule_chain.RuleChain) {
	c.nextRuleChain = nextRuleChain
}

func (c *ColumnWithDotChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &ColumnWithDotChain{
		pool: pool,
	}
}
