package selectcmd

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/alias"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
	"strings"
)

type HasCommaAtLastRule struct {
	errMsg string
}

func (r *HasCommaAtLastRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for HasCommaAtLastRule.Validate()"
		return false
	}

	var tok string = st.Tok
	var perhapComma byte = tok[len(tok)-1]
	var comma byte = constants.SYMBOL_COMMA

	if perhapComma == comma {
		return true
	}

	r.errMsg = fmt.Sprintf("Expect the last-idx is `%b`, but found: `%b`", constants.SYMBOL_COMMA, perhapComma)
	return false
}

func (r *HasCommaAtLastRule) count(needle byte, haystack string) int {
	var needleToStr string = string(needle)
	return strings.Count(haystack, needleToStr)
}

func (r *HasCommaAtLastRule) Key() alias.RuleKey {
	return constants.RULE_HAS_COMMA_AT_LAST
}

func (r *HasCommaAtLastRule) ErrorMsg() string {
	return r.errMsg
}
