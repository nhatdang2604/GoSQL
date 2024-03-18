package column_with_star

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/froms"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_pool"
)

type ColumnWithStarChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *ColumnWithStarChain) Exec(toks []string) bool {
	c.remainToks = toks
	var isSuccess bool = false
	var firstTok string = toks[0]
	if c.isStar(firstTok) {
		c.curTok = firstTok
		c.remainToks = toks[1:]
		isSuccess = c.setNextRule(c.remainToks)
	}

	if !isSuccess && c.errMsg == "" {
		c.errMsg = fmt.Sprintf("Expected `%s` after `%b`", constants.KEYWORD_FROM, constants.SYMBOL_STAR)
	}

	return isSuccess
}

func (c *ColumnWithStarChain) isStar(tok string) bool {
	isStarRule := c.pool.Get(constants.RULE_IS_STAR)
	var isStar bool = isStarRule.Validate(rule_input.SingleTok{Tok: tok})
	return isStar
}

func (c *ColumnWithStarChain) setNextRule(toks []string) bool {
	var isSuccess bool = false
	var nextTok string = toks[0]
	if c.isNextRuleFrom(nextTok) {
		c.nextRuleChain = froms.New(c.pool)
		isSuccess = true
	}

	return isSuccess
}

func (c *ColumnWithStarChain) isNextRuleFrom(tok string) bool {
	fromRule := c.pool.Get(constants.RULE_IS_FROM)
	var isFrom bool = fromRule.Validate(rule_input.SingleTok{Tok: tok})
	return isFrom
}

func (c *ColumnWithStarChain) EmitTok() string {
	return c.curTok
}

func (c *ColumnWithStarChain) RemainToks() []string {
	return c.remainToks
}

func (c *ColumnWithStarChain) ErrorMsg() string {
	return c.errMsg
}

func (c *ColumnWithStarChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {
	return &ColumnWithStarChain{
		pool: pool,
	}
}
