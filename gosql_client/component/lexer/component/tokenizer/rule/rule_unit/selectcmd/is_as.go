package selectcmd

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/alias"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_helper"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
)

type IsAsRule struct {
	errMsg string
}

func (r *IsAsRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for IsAsRule.Validate()"
		return false
	}

	var tok string = st.Tok
	if rule_helper.AreTokEqual(tok, constants.KEYWORD_AS) {
		return true
	}

	r.errMsg = fmt.Sprintf("expect `%s` but found :%s", constants.KEYWORD_AS, tok)
	return false
}

func (r *IsAsRule) Key() alias.RuleKey {
	return constants.RULE_IS_AS
}

func (r *IsAsRule) ErrorMsg() string {
	return r.errMsg
}
