package rule

import (
	lexcom "gosql_client/component/lexer/component"
	irul "gosql_client/component/lexer/rule/interfaces"
)

type RuleAggregator struct {
	Rules []irul.Rule
}

func (a *RuleAggregator) ValidateMultipleRules(command lexcom.Command) bool {

	for _, rule := range a.Rules {
		if !rule.Validate(command) {
			return false
		}
	}

	return true
}

func MakeRuleAggr(rules []irul.Rule) irul.RuleAggregatable {
	return &RuleAggregator{Rules: rules}
}
