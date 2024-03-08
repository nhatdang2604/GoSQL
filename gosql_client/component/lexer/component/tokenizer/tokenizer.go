package tokenizer

import (
	"gosql_client/component/lexer/component/tokenizer/alias"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/interfaces"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_helper"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_pool"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_unit"
	"strings"
)

//select * from db.test;
//select a from db.test;
//select t.a from db.test;

//insert into db.test values (1, 'b', true);
//insert into db.test values (1, 'b', true), (2, 'c', false);

// Type alias
const (
	Rule     = rule_unit.Rule
	RuleKey  = alias.RuleKey
	RulePool = rule_pool.RulePool
	TokType  = alias.TokType
)

type Tokenizer struct {
	input              string
	uncleanedToks      []string
	hasMoreTokens      bool
	curUncleanedTokIdx int
	curTok             string
	curTokType         TokType
	rulePool           RulePool
}

func (t Tokenizer) HasMoreTokens() bool {
	if t.SymbolVal() == ';' {
		t.hasMoreTokens = false
	}

	return t.hasMoreTokens
}

func (t Tokenizer) initialize() {
	t.uncleanedToks = strings.Split(t.input, " ")
	t.hasMoreTokens = true
	t.curTokType = constants.TOKTYPE_UNK_CONST
	t.curTok = ""
	t.rulePool = rule_pool.New()
	t.sanitizeInput()
	t.Advance()
}

func (t Tokenizer) sanitizeInput() {
	for idx, tok := range t.uncleanedToks {
		t.uncleanedToks[idx] = rule_helper.SanitizeTok(tok)
	}
}

func (t Tokenizer) Advance() {
	//TODO:
}

func (t Tokenizer) isStart() bool {
	return t.curTok == ""
}

func NewTokenizer(input string) interfaces.Tokenizable {
	var tokenizer Tokenizer = Tokenizer{input: input}
	tokenizer.initialize()
	return tokenizer
}
