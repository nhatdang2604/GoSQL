package intostm

import (
	lexcom "gosql_client/component/lexer/component"
	irul "gosql_client/component/lexer/rule/interfaces"
)

type IntoStatementRuleAggregator struct {
	Rules []irul.Rule
}

func (a IntoStatementRuleAggregator) ValidateMultipleRules(command lexcom.Command) bool {

	for _, rule := range a.Rules {
		if !rule.Validate(command) {
			return false
		}
	}

	return true
}

func MakeRuleAggr() irul.RuleAggregatable {

	return IntoStatementRuleAggregator{
		Rules: []irul.Rule{
			HasIntoStatementRule{},
			HasOneDotOnTableNameRule{},
			HasDbNameBeforeDotRule{},
			HasTableNameAfterDotRule{},
			HasOpenBracketAfterTableNameRule{},
			HasClosedBracketBeforeValuesStatementRule{},
			HasColumnInsideBracketRule{},
			AllColumnNotHasReservedKeywordRule{},
			AllColumnSurroundedByQuoteRule{},
		},
	}
}
