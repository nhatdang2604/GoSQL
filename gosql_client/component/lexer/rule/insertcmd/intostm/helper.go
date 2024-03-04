package intostm

import (
	"errors"
	lexcom "gosql_client/component/lexer/component"
	"gosql_client/constant"
)

func FindIdxOfTokenHasOpenBracket(insertCommand lexcom.Command) (int, error) {
	var intoIdx, err1 = insertCommand.FindKeyword(constant.INTO_KEYWORD)

	if nil != err1 {
		return -1, err1
	}

	var tokenHasOpenBracketIdx int = intoIdx + 2
	var tokenHasOpenBracket, err2 = insertCommand.GetTokenAt(tokenHasOpenBracketIdx)

	if nil != err2 {
		return -1, err2
	}

	if tokenHasOpenBracket[0] != '(' {
		return -1, errors.New("Doesn't contain '(' character in 'Into' statement")
	}

	return tokenHasOpenBracketIdx, nil
}

func FindIdxOfTokenHasClosedBracket(insertCommand lexcom.Command) (int, error) {
	var valuesIdx, err1 = insertCommand.FindKeyword(constant.VALUES_KEYWORD)

	if nil != err1 {
		return -1, err1
	}

	var tokenHasClosedBracketIdx int = valuesIdx - 1
	var tokenHasClosedBracket, err2 = insertCommand.GetTokenAt(tokenHasClosedBracketIdx)

	if nil != err2 {
		return -1, err2
	}

	var perhapClosedBracket = tokenHasClosedBracket[len(tokenHasClosedBracket)-1]

	if perhapClosedBracket != ')' {
		return -1, errors.New("Doesn't contain ')' character in 'Into' statement")
	}

	return tokenHasClosedBracketIdx, nil
}
