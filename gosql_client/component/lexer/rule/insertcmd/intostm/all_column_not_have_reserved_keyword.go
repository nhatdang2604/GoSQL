package intostm

import (
	lexcom "gosql_client/component/lexer/component"
	constant "gosql_client/constant"
	"gosql_client/helper"
)

type AllColumnNotHasReservedKeywordRule struct{}

func (AllColumnNotHasReservedKeywordRule) Validate(command lexcom.Command) bool {

	var openBracketTokenIdx, err1 = FindIdxOfTokenHasOpenBracket(command)
	var closedBracketTokenIdx, err2 = FindIdxOfTokenHasClosedBracket(command)

	if nil != err1 || nil != err2 {
		return false
	}

	for columnIdx := openBracketTokenIdx + 1; columnIdx < closedBracketTokenIdx; columnIdx++ {
		var columnName, _ = command.GetTokenAt(columnIdx)

		for _, keyword := range constant.RESERVED_KEYWORD {
			if helper.IsTokenEqualIgnoringCase(columnName, keyword) {
				return false
			}
		}
	}

	return true
}
