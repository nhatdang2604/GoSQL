package valuestm

import (
	lexcom "gosql_client/component/lexer/component"
	irul "gosql_client/component/lexer/rule/interfaces"
)

type ValuesStatementRuleAggregator struct {
	Rules []irul.Rule
}

func (a ValuesStatementRuleAggregator) ValidateMultipleRules(command lexcom.Command) bool {

	for _, rule := range a.Rules {
		if !rule.Validate(command) {
			return false
		}
	}

	return true
}

func MakeRuleAggr() irul.RuleAggregatable {

	return ValuesStatementRuleAggregator{
		Rules: []irul.Rule{
			HasValuesStatementRule{},
		},
	}
}
