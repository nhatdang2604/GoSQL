package lexer

import (
	lexcom "gosql_client/component/lexer/component"
	rul "gosql_client/component/lexer/rule"
	insrul "gosql_client/component/lexer/rule/inserts/rule_unit"
	iaggr "gosql_client/component/lexer/rule/interfaces"
	selrul "gosql_client/component/lexer/rule/selects/rule_unit"
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
			constant.SELECT_KEYWORD: rul.MakeRuleAggr([]iaggr.Rule{
				&selrul.HasSelectStatement{},
				//TODO:
			}),
			constant.INSERT_KEYWORD: rul.MakeRuleAggr([]iaggr.Rule{
				&insrul.HasKeywordAtFirstRule{},
				//TODO:
			}),
		},
	}
}
