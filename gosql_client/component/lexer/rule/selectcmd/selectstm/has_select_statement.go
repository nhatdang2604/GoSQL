package selectstm

import (
	lexcom "gosql_client/component/lexer/component"
	constant "gosql_client/constant"
)

type HasSelectStatementRule struct{}

func (HasSelectStatementRule) Validate(command lexcom.Command) bool {

	var selectIdx, err = command.FindKeyword(constant.SELECT_KEYWORD)

	if nil != err {
		return false
	}

	if selectIdx != 0 {
		return false
	}

	return true
}
