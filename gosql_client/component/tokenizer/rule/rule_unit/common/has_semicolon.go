package common

import (
	"fmt"
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_input"
	"strings"
)

type HasSemicolonRule struct {
	errMsg string
}

func (r *HasSemicolonRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for HasSemicolonRule.Validate()"
		return false
	}

	var tok string = st.Tok
	if semicolonIdx := strings.IndexByte(tok, constants.SYMBOL_SEMICOLON); -1 != semicolonIdx {
		return true
	}

	r.errMsg = fmt.Sprintf("not found `%b` in :%s", constants.SYMBOL_SEMICOLON, tok)
	return false
}

func (r *HasSemicolonRule) Key() alias.RuleKey {
	return constants.RULE_HAS_SEMICOLON
}

func (r *HasSemicolonRule) ErrorMsg() string {
	return r.errMsg
}
