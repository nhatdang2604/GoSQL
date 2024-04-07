package insertcmd

import (
	"fmt"
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_helper"
	"gosql_client/component/tokenizer/rule/rule_input"
)

type IsValueRule struct {
	errMsg string
}

func (r *IsValueRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for IsValueRule.Validate()"
		return false
	}

	var tok string = st.Tok
	if rule_helper.AreTokEqual(tok, constants.KEYWORD_VALUES) {
		return true
	}

	r.errMsg = fmt.Sprintf("expect `values` but found :%s", tok)
	return false
}

func (r *IsValueRule) Key() alias.RuleKey {
	return constants.RULE_IS_VALUES
}

func (r *IsValueRule) ErrorMsg() string {
	return r.errMsg
}
