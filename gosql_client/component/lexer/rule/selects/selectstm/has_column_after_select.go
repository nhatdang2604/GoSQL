package selectstm

import (
	lexcom "gosql_client/component/lexer/component"
	"gosql_client/constant"
)

type HasColumnAfterSelectRule struct{}

func (HasColumnAfterSelectRule) Validate(command lexcom.Command) bool {

	var selectIndex, err1 = command.FindKeyword(constant.SELECT_KEYWORD)
	var fromIndex, err2 = command.FindKeyword(constant.FROM_KEYWORD)

	if nil != err1 || nil != err2 {
		return false
	}

	if (fromIndex - selectIndex) > 1 {
		return true
	}

	return false
}
