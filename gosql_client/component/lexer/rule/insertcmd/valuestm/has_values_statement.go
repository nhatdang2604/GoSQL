package valuestm

import (
	lexcom "gosql_client/component/lexer/component"
	"gosql_client/constant"
)

type HasValuesStatementRule struct{}

func (HasValuesStatementRule) Validate(command lexcom.Command) bool {
	var valuesIdx, err = command.FindKeyword(constant.VALUES_KEYWORD)

	if nil != err {
		return false
	}

	if valuesIdx < 3 {
		return false
	}

	return true
}
