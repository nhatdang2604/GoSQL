package table_name

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/common/is_semicolon"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/columns/column_with_dot"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/columns/column_with_star"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/columns/column_without_dot"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_chain/start/selects/columns/common/is_comma"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_pool"
	"strings"
)

type TableNameChain struct {
	nextRuleChain rule_chain.RuleChain
	pool          rule_pool.RulePool
	curTok        string
	errMsg        string
	remainToks    []string
}

func (c *TableNameChain) Exec(toks []string) bool {

	var isSuccess bool = false
	var firstTok string = toks[0]
	if c.hasSemicolon(firstTok) {
		isSuccess = c.setAsSemicolon(toks)
	} else {
		isSuccess = c.setAsAs(toks)
	}

	return isSuccess
}

func (c *TableNameChain) hasSemicolon(tok string) bool {
	hasSemicolonRule := c.pool.Get(constants.RULE_HAS_SEMICOLON)
	var hasSemicolon bool = hasSemicolonRule.Validate(rule_input.SingleTok{Tok: tok})
	return hasSemicolon
}

func (c *TableNameChain) setAsSemicolon(toks []string) bool {
	var tok string = toks[0]
	var semicolon string = string(constants.SYMBOL_SEMICOLON)
	var splits []string = strings.Split(tok, semicolon)

	//More than 1 semicolon check
	if len(splits) > 1 {
		c.errMsg = fmt.Sprintf("More than 1 semicolon in '%s'", tok)
		return false
	}

	var tableName string = splits[0]

	toks[0] = semicolon // transform token 'b;' into ';'
	c.curTok = tableName
	c.remainToks = toks
	c.nextRuleChain = is_semicolon.New(c.pool)
	return true
}

func (c *TableNameChain) setAsAs(toks []string) bool {
	//TODO:
}

func (c *TableNameChain) EmitTok() string {
	return c.curTok
}

func (c *TableNameChain) RemainToks() []string {
	return c.remainToks
}

func (c *TableNameChain) ErrorMsg() string {
	return c.errMsg
}

func (c *TableNameChain) NextRuleChain() rule_chain.RuleChain {
	return c.nextRuleChain
}

func New(pool rule_pool.RulePool) rule_chain.RuleChain {

	//Init IsCommaRuleChain
	initIsCommaRuleChain(pool)

	return &TableNameChain{
		pool: pool,
	}
}

func initIsCommaRuleChain(pool rule_pool.RulePool) {
	is_comma.SharedIsCommaChain = is_comma.New(pool)
	is_comma.SharedIsCommaChain.AddColumnRuleChain(column_with_star.New(pool))
	is_comma.SharedIsCommaChain.AddColumnRuleChain(column_with_dot.New(pool))
	is_comma.SharedIsCommaChain.AddColumnRuleChain(column_without_dot.New(pool))
}
