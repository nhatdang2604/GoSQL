package common

import (
	"fmt"
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_helper"
	"gosql_client/component/tokenizer/rule/rule_input"
)

type IsSemicolonRule struct {
	errMsg string
}

func (r *IsSemicolonRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for IsSemicolonRule.Validate()"
		return false
	}

	var tok string = st.Tok
	var semiconlonAsStr string = string(constants.SYMBOL_SEMICOLON)
	if rule_helper.AreTokEqual(tok, semiconlonAsStr) {
		return true
	}

	r.errMsg = fmt.Sprintf("expect `%b` but found :%s", constants.SYMBOL_SEMICOLON, tok)
	return false
}

func (r *IsSemicolonRule) Key() alias.RuleKey {
	return constants.RULE_IS_SEMICOLON
}

func (r *IsSemicolonRule) ErrorMsg() string {
	return r.errMsg
}
