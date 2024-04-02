package column_with_dot

import (
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/columns/column_with_dot/dot_at_first"
	"gosql_client/component/tokenizer/rule/rule_input"
	"gosql_client/component/tokenizer/rule/rule_pool"
	"gosql_client/component/tokenizer/rule/rule_unit"
	"strings"
)

type ColumnWithDotChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *ColumnWithDotChain) Exec(toks []string) bool {
	c.remainToks = toks
	var isSuccess bool = false
	var firstTok string = toks[0]

	if c.hasOneAndOnlyOneDot(firstTok) {
		c.curTok = c.getDbName(firstTok)
		toks[0] = c.rmDbNameFromTok(firstTok) //replace the "a.b" with ".b"
		c.remainToks = toks
		isSuccess = c.setNextRule(c.remainToks)
	}
	return isSuccess
}

func (c *ColumnWithDotChain) hasOneAndOnlyOneDot(tok string) bool {
	var hasDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_DOT)
	var hasOnlyOneDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_ONLY_ONE_DOT)

	var hasDot = hasDotRule.Validate(rule_input.SingleTok{Tok: tok})
	var hasOnlyOneDot = hasOnlyOneDotRule.Validate(rule_input.SingleTok{Tok: tok})

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

func (c *ColumnWithDotChain) setNextRule(toks []string) bool {
	var isSuccess bool = false
	var nextTok string = toks[0] // Because the 'a.b' tok now become '.b' tok

	if c.isNextRuleDotAtFirst(nextTok) {
		c.nextRuleChain = dot_at_first.New(c.pool)
		isSuccess = true
	}

	return isSuccess
}

func (c *ColumnWithDotChain) isNextRuleDotAtFirst(tok string) bool {

	var hasDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_DOT)
	var hasOnlyOneDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_ONLY_ONE_DOT)
	var hasDotAtFirstRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_DOT_AT_FIRST)

	var hasDot = hasDotRule.Validate(rule_input.SingleTok{Tok: tok})
	var hasOnlyOneDot = hasOnlyOneDotRule.Validate(rule_input.SingleTok{Tok: tok})
	var hasDotAtFirst = hasDotAtFirstRule.Validate(rule_input.SingleTok{Tok: tok})

	return hasDot && hasOnlyOneDot && hasDotAtFirst
}

func (c *ColumnWithDotChain) IsValid(tok string) bool {
	return c.hasOneAndOnlyOneDot(tok)
}

func (c *ColumnWithDotChain) ToRuleChain() rule_chain.RuleChain {
	return c
}

func (c *ColumnWithDotChain) EmitTok() string {
	return c.curTok
}

func (c *ColumnWithDotChain) RemainToks() []string {
	return c.remainToks
}

func (c *ColumnWithDotChain) ErrorMsg() string {
	return c.errMsg
}

func (c *ColumnWithDotChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) *ColumnWithDotChain {
	return &ColumnWithDotChain{
		pool: pool,
	}
}
