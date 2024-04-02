package common

import (
	"fmt"
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_helper"
	"gosql_client/component/tokenizer/rule/rule_input"
)

type IsInsertRule struct {
	errMsg string
}

func (r *IsInsertRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for IsInsertRule.Validate()"
		return false
	}

	var tok string = st.Tok
	if rule_helper.AreTokEqual(tok, constants.KEYWORD_INSERT) {
		return true
	}

	r.errMsg = fmt.Sprintf("expect `insert` but found :%s", tok)
	return false
}

func (r *IsInsertRule) Key() alias.RuleKey {
	return constants.RULE_IS_INSERT
}

func (r *IsInsertRule) ErrorMsg() string {
	return r.errMsg
}
