package insertcmd

import (
	"fmt"
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_helper"
	"gosql_client/component/tokenizer/rule/rule_input"
)

type IsIntoRule struct {
	errMsg string
}

func (r *IsIntoRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for IsIntoRule.Validate()"
		return false
	}

	var tok string = st.Tok
	if rule_helper.AreTokEqual(tok, constants.KEYWORD_INTO) {
		return true
	}

	r.errMsg = fmt.Sprintf("expect `into` but found :%s", tok)
	return false
}

func (r *IsIntoRule) Key() alias.RuleKey {
	return constants.RULE_IS_INTO
}

func (r *IsIntoRule) ErrorMsg() string {
	return r.errMsg
}
