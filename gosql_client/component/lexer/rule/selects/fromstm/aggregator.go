package fromstm

import (
	lexcom "gosql_client/component/lexer/component"
	irul "gosql_client/component/lexer/rule/interfaces"
)

type FromStatementRuleAggregator struct {
	Rules []irul.Rule
}

func (a FromStatementRuleAggregator) ValidateMultipleRules(command lexcom.Command) bool {

	for _, rule := range a.Rules {
		if !rule.Validate(command) {
			return false
		}
	}

	return true
}

func MakeRuleAggr() irul.RuleAggregatable {

	return FromStatementRuleAggregator{
		Rules: []irul.Rule{
			HasFromStatementRule{},
			HasOneDotRule{},
			HasDbNameBeforeDotRule{},
			HasTableNameAfterDotRule{},
		},
	}
}
