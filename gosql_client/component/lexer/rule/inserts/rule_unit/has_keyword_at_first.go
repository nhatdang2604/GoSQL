package rule_unit

import (
	lexcom "gosql_client/component/lexer/component"
	constant "gosql_client/constant"
	"gosql_client/helper"
)

type HasKeywordAtFirstRule struct{}

func (HasKeywordAtFirstRule) Validate(command lexcom.Command) bool {

	var firstToken, err = command.GetFirstToken()

	if nil != err {
		return false
	}

	if !helper.IsTokenEqualIgnoringCase(constant.INSERT_KEYWORD, firstToken) {
		return false
	}

	return true
}
