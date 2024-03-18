package selects

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/column_with_dot"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/columns/column_with_dot"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/columns/column_with_star"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_pool"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_unit"
)

type SelectChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *SelectChain) Exec(toks []string) bool {

	var isSuccess bool = false
	var firstTok string = toks[0]
	if c.isSelect(firstTok) {
		c.curTok = firstTok
		c.remainToks = c.remainToks[1:]
		if len(toks) > 1 {
			isSuccess = c.setNextRule(c.remainToks)
		}
	}

	if !isSuccess && c.errMsg != "" {
		c.errMsg = fmt.Sprintf("Unexpected value after '%s' keyword", constants.KEYWORD_SELECT)
	}

	return isSuccess
}

func (c *SelectChain) isSelect(tok string) bool {
	selectRule := c.pool.Get(constants.RULE_IS_SELECT)
	var isSelect bool = selectRule.Validate(rule_input.SingleTok{Tok: tok})
	return isSelect
}

func (c *SelectChain) setNextRule(toks []string) bool {
	var isSuccess bool = false
	var nextTok string = toks[0]
	if c.isNextRuleColumnWithStar(nextTok) {
		c.nextRuleChain = column_with_star.New(c.pool)
		isSuccess = true
	} else if c.isNextRuleColumnWithDot(nextTok) {
		c.nextRuleChain = column_with_dot.New(c.pool)
		isSuccess = true
	} else if c.isNextRuleColumnWithoutDot(nextTok) {
		c.nextRuleChain = column_without_dot.New(c.pool)
		isSuccess = true
	}

	return isSuccess
}

func (c *SelectChain) isNextRuleColumnWithStar(tok string) bool {
	var isStarRule rule_unit.Rule = c.pool.Get(constants.RULE_IS_STAR)
	return isStarRule.Validate(rule_input.SingleTok{Tok: tok})

}

func (c *SelectChain) isNextRuleColumnWithDot(tok string) bool {
	var hasDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_DOT)
	var hasOnlyOneDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_ONLY_ONE_DOT)

	var hasDot = hasDotRule.Validate(rule_input.SingleTok{Tok: tok})
	var hasOnlyOneDot = hasOnlyOneDotRule.Validate(rule_input.SingleTok{Tok: tok})

	return hasDot && hasOnlyOneDot
}

func (c *SelectChain) isNextRuleColumnWithoutDot(tok string) bool {
	var hasDotRule rule_unit.Rule = c.pool.Get(constants.RULE_HAS_DOT)
	var hasDot = hasDotRule.Validate(rule_input.SingleTok{Tok: tok})
	var notHasDot = !hasDot

	return notHasDot
}

func (c *SelectChain) EmitTok() string {
	return c.curTok
}

func (c *SelectChain) RemainToks() []string {
	return c.remainToks
}

func (c *SelectChain) ErrorMsg() string {
	return c.errMsg
}

func (c *SelectChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &SelectChain{
		pool: pool,
	}
}
