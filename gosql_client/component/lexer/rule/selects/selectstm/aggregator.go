package selectstm

import (
	lexcom "gosql_client/component/lexer/component"
	irul "gosql_client/component/lexer/rule/interfaces"
)

type SelectStatementRuleAggregator struct {
	Rules []irul.Rule
}

func (a SelectStatementRuleAggregator) ValidateMultipleRules(command lexcom.Command) bool {

	for _, rule := range a.Rules {
		if !rule.Validate(command) {
			return false
		}
	}

	return true
}

func MakeRuleAggr() irul.RuleAggregatable {

	return SelectStatementRuleAggregator{
		Rules: []irul.Rule{
			HasSelectStatementRule{},
			HasColumnAfterSelectRule{},
			AllColumnNotHasReservedKeywordRule{},
		},
	}
}
