package rule_unit

import (
	lexcom "gosql_client/component/lexer/component"
	"gosql_client/constant"
)

type HasFromStatementRule struct{}

func (HasFromStatementRule) Validate(command lexcom.Command) bool {
	var selectIdx, err = command.FindKeyword(constant.FROM_KEYWORD)

	if nil != err {
		return false
	}

	if selectIdx != 0 {
		return false
	}

	return true
}
