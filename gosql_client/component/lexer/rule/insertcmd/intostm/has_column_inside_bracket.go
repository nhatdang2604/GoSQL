package intostm

import (
	lexcom "gosql_client/component/lexer/component"
)

type HasColumnInsideBracketRule struct{}

func (HasColumnInsideBracketRule) Validate(command lexcom.Command) bool {

	var openBracketTokenIdx, err1 = FindIdxOfTokenHasOpenBracket(command)
	var closedBracketTokenIdx, err2 = FindIdxOfTokenHasClosedBracket(command)

	if nil != err1 || nil != err2 {
		return false
	}

	if (closedBracketTokenIdx - openBracketTokenIdx) > 1 {
		return true
	}

	return false
}
