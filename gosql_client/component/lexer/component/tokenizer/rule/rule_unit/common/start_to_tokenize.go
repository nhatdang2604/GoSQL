package common

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/alias"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
)

type StartToTokenizeRule struct {
	errMsg string
}

func (r *StartToTokenizeRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for StartToTokenizeRule.Validate()"
		return false
	}

	var tok string = st.Tok
	if tok == "" {
		return true
	}

	r.errMsg = fmt.Sprintf("expect \"\" but found :%s", tok)
	return false
}

func (r *StartToTokenizeRule) Key() alias.RuleKey {
	return constants.RULE_START_TO_TOKENIZE
}

func (r *StartToTokenizeRule) ErrorMsg() string {
	return r.errMsg
}
