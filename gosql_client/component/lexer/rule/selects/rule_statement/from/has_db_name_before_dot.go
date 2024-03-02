package rule_unit

import (
	lexcom "gosql_client/component/lexer/component"
	"gosql_client/constant"
	"strings"
)

type HasDbNameBeforeDot struct{}

func (HasDbNameBeforeDot) Validate(command lexcom.Command) bool {
	var fromIdx, err1 = command.FindKeyword(constant.FROM_KEYWORD)

	if nil != err1 {
		return false
	}

	var fromValue, err2 = command.GetTokenAt(fromIdx + 1)

	if nil != err2 {
		return false
	}

	var dotIdx int = strings.Index(fromValue, ".")
	if dotIdx == -1 {
		return false
	}

	var dbName string = fromValue[0:dotIdx]
	dbName = strings.TrimSpace(dbName)
	if dbName == "" {
		return false
	}

	return true
}
