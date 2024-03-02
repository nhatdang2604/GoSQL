package rule_unit

import (
	lexcom "gosql_client/component/lexer/component"
	"gosql_client/constant"
)

type HasFromStatementRule struct{}

func (HasFromStatementRule) Validate(command lexcom.Command) bool {
	var fromIdx, err = command.FindKeyword(constant.FROM_KEYWORD)

	if nil != err {
		return false
	}

	if fromIdx != 0 {
		return false
	}

	return true
}
