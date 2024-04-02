package selectcmd

import (
	"fmt"
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_input"
	"strings"
)

type HasDotRule struct {
	errMsg string
}

func (r *HasDotRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for HasDotRule.Validate()"
		return false
	}

	var tok string = st.Tok
	if dotIdx := strings.IndexByte(tok, constants.SYMBOL_DOT); -1 != dotIdx {
		return true
	}

	r.errMsg = fmt.Sprintf("not found `%b` in :%s", constants.SYMBOL_STAR, tok)
	return false
}

func (r *HasDotRule) Key() alias.RuleKey {
	return constants.RULE_HAS_DOT
}

func (r *HasDotRule) ErrorMsg() string {
	return r.errMsg
}
