package intostm

import (
	lexcom "gosql_client/component/lexer/component"
)

type HasOpenBracketAfterTableNameRule struct{}

func (HasOpenBracketAfterTableNameRule) Validate(command lexcom.Command) bool {

	var _, err = FindIdxOfTokenHasOpenBracket(command)

	if err != nil {
		return false
	}

	return true
}
