package lexer

import (
	lexcom "gosql_client/component/lexer/component"
	iaggr "gosql_client/component/lexer/rule/interfaces"
	selrul "gosql_client/component/lexer/rule/selectcmd"
	constant "gosql_client/constant"
)

type Lexer struct {
	tokenizer     lexcom.Tokenizable
	aggregatorMap map[string]iaggr.RuleAggregatable
}

func MakeLexer() *Lexer {
	return &Lexer{
		tokenizer: lexcom.MakeTokenizable(),
		aggregatorMap: map[string]iaggr.RuleAggregatable{
			constant.SELECT_KEYWORD: selrul.MakeRuleAggr(),
			// constant.INSERT_KEYWORD: rul.MakeRuleAggr([]iaggr.Rule{
			// 	&insrul.HasKeywordAtFirstRule{},
			// 	//TODO:
			// }),
		},
	}
}
