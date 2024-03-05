package intostm

import (
	lexcom "gosql_client/component/lexer/component"
)

type AllColumnSurroundedByQuoteRule struct{}

func (AllColumnSurroundedByQuoteRule) Validate(command lexcom.Command) bool {

	var openBracketTokenIdx, err1 = FindIdxOfTokenHasOpenBracket(command)
	var closedBracketTokenIdx, err2 = FindIdxOfTokenHasClosedBracket(command)

	if nil != err1 || nil != err2 {
		return false
	}

	for columnIdx := openBracketTokenIdx + 1; columnIdx < closedBracketTokenIdx; columnIdx++ {
		var columnName, _ = command.GetTokenAt(columnIdx)

		if columnName[0] != '\'' || columnName[len(columnName)-1] != '\'' {
			return false
		}
	}

	return true
}
