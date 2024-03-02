package rule_unit

import (
	lexcom "gosql_client/component/lexer/component"
	"gosql_client/constant"
	"strings"
)

type HasOneDotRule struct{}

func (HasOneDotRule) Validate(command lexcom.Command) bool {
	var fromIdx, err1 = command.FindKeyword(constant.FROM_KEYWORD)

	if nil != err1 {
		return false
	}

	var fromValue, err2 = command.GetTokenAt(fromIdx + 1)

	if nil != err2 {
		return false
	}

	var dotCnt int = strings.Count(fromValue, ".")

	return dotCnt == 1
}
