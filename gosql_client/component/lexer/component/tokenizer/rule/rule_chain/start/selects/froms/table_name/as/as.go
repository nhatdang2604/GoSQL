package as

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/columns/column_with_dot"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/columns/column_with_star"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/columns/column_without_dot"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/columns/common/is_comma"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/froms/table_name/as/alias"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_pool"
)

type AsChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *AsChain) Exec(toks []string) bool {

	var isSuccess bool = false
	var firstTok string = toks[0]
	if c.isAs(firstTok) {
		c.curTok = firstTok
		c.remainToks = toks[1:]
		isSuccess = c.setNextRule(c.remainToks)
	}

	if !isSuccess && c.errMsg != "" {
		c.errMsg = fmt.Sprintf("Expected '%s' keyword, found: %s", constants.KEYWORD_AS, firstTok)
	}

	return isSuccess
}

func (c *AsChain) isAs(tok string) bool {
	asRule := c.pool.Get(constants.RULE_IS_AS)
	var isAs bool = asRule.Validate(rule_input.SingleTok{Tok: tok})
	return isAs
}

func (c *AsChain) setNextRule(toks []string) bool {
	c.nextRuleChain = alias.New(c.pool)
	return true
}

func (c *AsChain) EmitTok() string {
	return c.curTok
}

func (c *AsChain) RemainToks() []string {
	return c.remainToks
}

func (c *AsChain) ErrorMsg() string {
	return c.errMsg
}

func (c *AsChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {

	//Init IsCommaRuleChain
	initIsCommaRuleChain(pool)

	return &AsChain{
		pool: pool,
	}
}

func initIsCommaRuleChain(pool rule_pool.RulePool) {
	is_comma.SharedIsCommaChain = is_comma.New(pool)
	is_comma.SharedIsCommaChain.AddColumnRuleChain(column_with_star.New(pool))
	is_comma.SharedIsCommaChain.AddColumnRuleChain(column_with_dot.New(pool))
	is_comma.SharedIsCommaChain.AddColumnRuleChain(column_without_dot.New(pool))
}
