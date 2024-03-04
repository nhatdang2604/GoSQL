package insertstm

import (
	lexcom "gosql_client/component/lexer/component"
	irul "gosql_client/component/lexer/rule/interfaces"
)

type InsertStatementRuleAggregator struct {
	Rules []irul.Rule
}

func (a InsertStatementRuleAggregator) ValidateMultipleRules(command lexcom.Command) bool {

	for _, rule := range a.Rules {
		if !rule.Validate(command) {
			return false
		}
	}

	return true
}

func MakeRuleAggr() irul.RuleAggregatable {

	return InsertStatementRuleAggregator{
		Rules: []irul.Rule{
			HasInsertStatementRule{},
		},
	}
}
