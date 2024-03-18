package common

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/alias"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_helper"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
)

type IsStartRule struct {
	errMsg string
}

func (r *IsStartRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for IsStartRule.Validate()"
		return false
	}

	var tok string = st.Tok
	if rule_helper.AreTokEqual(tok, constants.KEYWORD_START) {
		return true
	}

	r.errMsg = fmt.Sprintf("expect \"\" but found :%s", tok)
	return false
}

func (r *IsStartRule) Key() alias.RuleKey {
	return constants.RULE_IS_START
}

func (r *IsStartRule) ErrorMsg() string {
	return r.errMsg
}
