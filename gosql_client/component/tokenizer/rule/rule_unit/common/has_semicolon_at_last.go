package common

import (
	"fmt"
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_input"
)

type HasSemicolonAtLastRule struct {
	errMsg string
}

func (r *HasSemicolonAtLastRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for HasSemicolonAtLastRule.Validate()"
		return false
	}

	var tok string = st.Tok
	var perhapSemicolon byte = tok[len(tok)-1]
	var semicolon byte = constants.SYMBOL_SEMICOLON

	if perhapSemicolon == semicolon {
		return true
	}

	r.errMsg = fmt.Sprintf("Expect the last-idx is `%b`, but found: `%b`", constants.SYMBOL_SEMICOLON, perhapSemicolon)
	return false
}

func (r *HasSemicolonAtLastRule) Key() alias.RuleKey {
	return constants.RULE_HAS_SEMICOLON_AT_LAST
}

func (r *HasSemicolonAtLastRule) ErrorMsg() string {
	return r.errMsg
}
