package valuestm

import (
	lexcom "gosql_client/component/lexer/component"
	"gosql_client/constant"
)

type HasOpenBracketAfterValue struct{}

func (HasOpenBracketAfterValue) Validate(command lexcom.Command) bool {
	var valuesIdx, err1 = command.FindKeyword(constant.VALUES_KEYWORD)

	if nil != err1 {
		return false
	}

	var firstValueIdx int = valuesIdx + 1
	var firstValueToken, err2 = command.GetTokenAt(firstValueIdx)

	if nil != err2 {
		return false
	}

	var perhapOpenBracket = firstValueToken[0]

	if perhapOpenBracket != '(' {
		return false
	}

	return true
}
