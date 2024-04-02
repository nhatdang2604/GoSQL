package selectcmd

import (
	"fmt"
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_input"
	"strings"
)

type HasCommaRule struct {
	errMsg string
}

func (r *HasCommaRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for HasCommaRule.Validate()"
		return false
	}

	var tok string = st.Tok
	if commaIdx := strings.IndexByte(tok, constants.SYMBOL_COMMA); -1 != commaIdx {
		return true
	}

	r.errMsg = fmt.Sprintf("not found `%b` in :%s", constants.SYMBOL_COMMA, tok)
	return false
}

func (r *HasCommaRule) Key() alias.RuleKey {
	return constants.RULE_HAS_COMMA
}

func (r *HasCommaRule) ErrorMsg() string {
	return r.errMsg
}
