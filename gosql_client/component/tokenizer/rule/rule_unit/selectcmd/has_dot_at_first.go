package selectcmd

import (
	"fmt"
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_input"
	"strings"
)

type HasDotAtFirstRule struct {
	errMsg string
}

func (r *HasDotAtFirstRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for HasDotAtFirstRule.Validate()"
		return false
	}

	var tok string = st.Tok
	var perhapDot byte = tok[0]
	var dot byte = constants.SYMBOL_DOT
	if perhapDot == dot {
		return true
	}

	r.errMsg = fmt.Sprintf("Expect the 0-idx is `%b`, but found: `%b`", constants.SYMBOL_DOT, perhapDot)
	return false
}

func (r *HasDotAtFirstRule) count(needle byte, haystack string) int {
	var needleToStr string = string(needle)
	return strings.Count(haystack, needleToStr)
}

func (r *HasDotAtFirstRule) Key() alias.RuleKey {
	return constants.RULE_HAS_DOT_AT_FIRST
}

func (r *HasDotAtFirstRule) ErrorMsg() string {
	return r.errMsg
}
