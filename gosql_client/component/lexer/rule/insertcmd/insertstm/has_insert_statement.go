package insertstm

import (
	lexcom "gosql_client/component/lexer/component"
	constant "gosql_client/constant"
)

type HasInsertStatementRule struct{}

func (HasInsertStatementRule) Validate(command lexcom.Command) bool {

	var selectIdx, err = command.FindKeyword(constant.INSERT_KEYWORD)

	if nil != err {
		return false
	}

	if selectIdx != 0 {
		return false
	}

	return true
}
