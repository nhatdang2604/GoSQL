package selects

import (
	lexcom "gosql_client/component/lexer/component"
	irul "gosql_client/component/lexer/rule/interfaces"
	frmstm "gosql_client/component/lexer/rule/selectcmd/fromstm"
	selstm "gosql_client/component/lexer/rule/selectcmd/selectstm"
)

type SelectCommandRuleAggregator struct {
	RuleAggregators []irul.RuleAggregatable
}

func (a SelectCommandRuleAggregator) ValidateMultipleRules(command lexcom.Command) bool {

	for _, ruleAggregator := range a.RuleAggregators {
		if !ruleAggregator.ValidateMultipleRules(command) {
			return false
		}
	}

	return true
}

func MakeRuleAggr() irul.RuleAggregatable {

	return SelectCommandRuleAggregator{
		RuleAggregators: []irul.RuleAggregatable{
			selstm.MakeRuleAggr(),
			frmstm.MakeRuleAggr(),
		},
	}
}
