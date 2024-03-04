package intostm

import (
	lexcom "gosql_client/component/lexer/component"
	"gosql_client/constant"
	"strings"
)

type HasDbNameBeforeDotRule struct{}

func (HasDbNameBeforeDotRule) Validate(command lexcom.Command) bool {
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

	var dbName string = intoValue[0:dotIdx]
	dbName = strings.TrimSpace(dbName)
	if dbName == "" {
		return false
	}

	return true
}
