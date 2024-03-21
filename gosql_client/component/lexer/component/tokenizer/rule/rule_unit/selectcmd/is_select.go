package selectcmd

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/alias"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_helper"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
)

type IsSelectRule struct {
	errMsg string
}

func (r *IsSelectRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for IsSelectRule.Validate()"
		return false
	}

	var tok string = st.Tok
	if rule_helper.AreTokEqual(tok, constants.KEYWORD_SELECT) {
		return true
	}

	r.errMsg = fmt.Sprintf("expect `select` but found :%s", tok)
	return false
}

func (r *IsSelectRule) Key() alias.RuleKey {
	return constants.RULE_IS_SELECT
}

func (r *IsSelectRule) ErrorMsg() string {
	return r.errMsg
}
