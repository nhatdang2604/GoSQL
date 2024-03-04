package intostm

import (
	lexcom "gosql_client/component/lexer/component"
	"gosql_client/constant"
	"strings"
)

type HasOneDotOnTableNameRule struct{}

func (HasOneDotOnTableNameRule) Validate(command lexcom.Command) bool {
	var intoIdx, err1 = command.FindKeyword(constant.INTO_KEYWORD)

	if nil != err1 {
		return false
	}

	var intoValue, err2 = command.GetTokenAt(intoIdx + 1)

	if nil != err2 {
		return false
	}

	var dotCnt int = strings.Count(intoValue, ".")

	return dotCnt == 1
}
