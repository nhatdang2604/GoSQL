package lexer

import (
	lexcom "gosql_client/component/lexer/component"
	rul "gosql_client/component/lexer/rule"
	iaggr "gosql_client/component/lexer/rule/interfaces"
	urul "gosql_client/component/lexer/rule/selects/rule_unit"
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
			constant.SelectKeyword: rul.MakeRuleAggr([]iaggr.Rule{
				&urul.HasKeywordAtFirstRule{},
				//TODO:
			}),
			constant.InsertKeyword: rul.MakeRuleAggr([]iaggr.Rule{
				//TODO:
			}),
		},
	}
}
