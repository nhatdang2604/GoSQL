package intostm

import (
	lexcom "gosql_client/component/lexer/component"
	"gosql_client/constant"
	"strings"
)

type HasTableNameAfterDotRule struct{}

func (HasTableNameAfterDotRule) Validate(command lexcom.Command) bool {
	var intoIdx, err1 = command.FindKeyword(constant.INTO_KEYWORD)

	if nil != err1 {
		return false
	}

	var intoValue, err2 = command.GetTokenAt(intoIdx + 1)

	if nil != err2 {
		return false
	}

	var dotIdx int = strings.Index(intoValue, ".")
	if dotIdx == -1 {
		return false
	}

	var tableName string = intoValue[dotIdx:]
	tableName = strings.TrimSpace(tableName)
	if tableName == "" {
		return false
	}

	return true
}
