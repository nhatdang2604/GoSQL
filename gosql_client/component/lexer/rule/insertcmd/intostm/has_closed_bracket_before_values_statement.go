package intostm

import (
	lexcom "gosql_client/component/lexer/component"
)

type HasClosedBracketBeforeValuesStatementRule struct{}

func (HasClosedBracketBeforeValuesStatementRule) Validate(command lexcom.Command) bool {

	var _, err = FindIdxOfTokenHasClosedBracket(command)

	if err != nil {
		return false
	}

	return true
}
