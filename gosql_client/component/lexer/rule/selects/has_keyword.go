package selects

import (
	lexcom "gosql_client/component/lexer/component"
	constant "gosql_client/constant"
	"strings"
)

type HasKeywordRule struct{}

func (HasKeywordRule) Validate(command lexcom.Command) bool {

	var firstToken, err = command.GetFirstToken()

	if nil != err {
		return false
	}

	if constant.SelectKeyword != strings.ToLower(firstToken) {
		return false
	}

	return true
}
