package selectcmd

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/alias"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_helper"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
)

type IsCommaRule struct {
	errMsg string
}

func (r *IsCommaRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for IsCommaRule.Validate()"
		return false
	}

	var tok string = st.Tok
	var commaAsStr string = string(constants.SYMBOL_COMMA)
	if rule_helper.AreTokEqual(tok, commaAsStr) {
		return true
	}

	r.errMsg = fmt.Sprintf("expect `%b` but found :%s", constants.SYMBOL_COMMA, tok)
	return false
}

func (r *IsCommaRule) Key() alias.RuleKey {
	return constants.RULE_IS_COMMA
}

func (r *IsCommaRule) ErrorMsg() string {
	return r.errMsg
}
