package selectcmd

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/alias"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_helper"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
)

type IsDotRule struct {
	errMsg string
}

func (r *IsDotRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for IsDotRule.Validate()"
		return false
	}

	var tok string = st.Tok
	var dotAsStr string = string(constants.SYMBOL_DOT)
	if rule_helper.AreTokEqual(tok, dotAsStr) {
		return true
	}

	r.errMsg = fmt.Sprintf("expect `%b` but found :%s", constants.SYMBOL_STAR, tok)
	return false
}

func (r *IsDotRule) Key() alias.RuleKey {
	return constants.RULE_IS_DOT
}

func (r *IsDotRule) ErrorMsg() string {
	return r.errMsg
}
