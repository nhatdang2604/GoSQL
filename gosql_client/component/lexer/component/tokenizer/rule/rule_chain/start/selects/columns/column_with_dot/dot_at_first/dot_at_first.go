package dot_at_first

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/columns/column_with_dot/table_name"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_pool"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_unit"
)

type DotAtFirstChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *DotAtFirstChain) Exec(toks []string) bool {
	c.remainToks = toks
	var isSuccess bool = false
	var firstTok string = toks[0]

	if c.hasOnlyOneDotAtFirst(firstTok) {
		c.curTok = string(constants.SYMBOL_DOT)
		toks[0] = c.rmDotFromTok(firstTok) //replace the ".b" with "b"
		c.remainToks = toks
		isSuccess = c.setNextRule(c.remainToks)
	}
	return isSuccess
}

func (c *DotAtFirstChain) hasOnlyOneDotAtFirst(tok string) bool {
	var hasDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_DOT)
	var hasOnlyOneDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_ONLY_ONE_DOT)
	var hasDotAtFirstRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_DOT_AT_FIRST)

	var hasDot = hasDotRule.Validate(rule_input.SingleTok{Tok: tok})
	var hasOnlyOneDot = hasOnlyOneDotRule.Validate(rule_input.SingleTok{Tok: tok})
	var hasDotAtFirst = hasDotAtFirstRule.Validate(rule_input.SingleTok{Tok: tok})

	return hasDot && hasOnlyOneDot && hasDotAtFirst
}

func (c *DotAtFirstChain) rmDotFromTok(tok string) string {
	return tok[1:]
}

func (c *DotAtFirstChain) setNextRule(toks []string) bool {
	var isSuccess bool = false
	var nextTok string = toks[0] // Because the '.b' tok now become 'b' tok

	if c.isNextRuleHasCommaAtLast(nextTok) || c.isNextRuleDoesntHaveComma(nextTok) {
		c.nextRuleChain = table_name.New(c.pool)
		isSuccess = true
	}

	if !isSuccess {
		c.errMsg = fmt.Sprintf("Expected column name after `%b`", constants.SYMBOL_DOT)
	}

	return isSuccess
}

func (c *DotAtFirstChain) isNextRuleHasCommaAtLast(tok string) bool {
	var hasCommaRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_COMMA)
	var hasOnlyOneCommaRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_ONLY_ONE_COMMA)
	var HasCommaAtLastRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_COMMA_AT_LAST)

	var hasComma bool = hasCommaRule.Validate(tok)
	var hasOnlyOneComma bool = hasOnlyOneCommaRule.Validate(tok)
	var HasCommaAtLast bool = HasCommaAtLastRule.Validate(tok)

	return hasComma && hasOnlyOneComma && HasCommaAtLast
}

func (c *DotAtFirstChain) isNextRuleDoesntHaveComma(tok string) bool {
	var hasCommaRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_COMMA)
	var notHaveComma bool = !hasCommaRule.Validate(tok)

	return notHaveComma
}

func (c *DotAtFirstChain) isNextRuleColumnNotHavingComma(tok string) bool {
	var hasCommaRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_COMMA)

	var notHaveComma bool = !hasCommaRule.Validate(tok)
	var hasOnlyOneComma bool = hasOnlyOneCommaRule.Validate(tok)
	var HasCommaAtLast bool = HasCommaAtLastRule.Validate(tok)

	return hasComma && hasOnlyOneComma && HasCommaAtLast
}

func (c *DotAtFirstChain) EmitTok() string {
	return c.curTok
}

func (c *DotAtFirstChain) RemainToks() []string {
	return c.remainToks
}

func (c *DotAtFirstChain) ErrorMsg() string {
	return c.errMsg
}

func (c *DotAtFirstChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &DotAtFirstChain{
		pool: pool,
	}
}
