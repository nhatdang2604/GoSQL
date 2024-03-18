package selectcmd

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/alias"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_helper"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
)

type IsFromRule struct {
	errMsg string
}

func (r *IsFromRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for IsFromRule.Validate()"
		return false
	}

	var tok string = st.Tok
	if rule_helper.AreTokEqual(tok, constants.KEYWORD_FROM) {
		return true
	}

	r.errMsg = fmt.Sprintf("expect `%s` but found :%s", constants.KEYWORD_FROM, tok)
	return false
}

func (r *IsFromRule) Key() alias.RuleKey {
	return constants.RULE_IS_FROM
}

func (r *IsFromRule) ErrorMsg() string {
	return r.errMsg
}
