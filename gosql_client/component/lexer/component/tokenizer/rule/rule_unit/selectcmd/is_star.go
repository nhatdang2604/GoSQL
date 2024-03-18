package selectcmd

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/alias"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_helper"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
)

type IsStarRule struct {
	errMsg string
}

func (r *IsStarRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for IsStarRule.Validate()"
		return false
	}

	var tok string = st.Tok
	var starAsStr string = string(constants.SYMBOL_STAR)
	if rule_helper.AreTokEqual(tok, starAsStr) {
		return true
	}

	r.errMsg = fmt.Sprintf("expect `%b` but found :%s", constants.SYMBOL_STAR, tok)
	return false
}

func (r *IsStarRule) Key() alias.RuleKey {
	return constants.RULE_IS_STAR
}

func (r *IsStarRule) ErrorMsg() string {
	return r.errMsg
}
