package intostm

import (
	lexcom "gosql_client/component/lexer/component"
	"gosql_client/constant"
)

type HasIntoStatementRule struct{}

func (HasIntoStatementRule) Validate(command lexcom.Command) bool {
	var intoIdx, err = command.FindKeyword(constant.INTO_KEYWORD)

	if nil != err {
		return false
	}

	if intoIdx != 1 {
		return false
	}

	return true
}
