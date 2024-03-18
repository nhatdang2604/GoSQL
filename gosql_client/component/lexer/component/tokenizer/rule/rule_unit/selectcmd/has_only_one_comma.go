package selectcmd

import (
	"fmt"
	"gosql_client/component/lexer/component/tokenizer/alias"
	"gosql_client/component/lexer/component/tokenizer/constants"
	"gosql_client/component/lexer/component/tokenizer/rule/rule_input"
	"strings"
)

type HasOnlyOneCommaRule struct {
	errMsg string
}

func (r *HasOnlyOneCommaRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for HasOnlyOneCommaRule.Validate()"
		return false
	}

	var tok string = st.Tok
	var cnt int = r.count(constants.SYMBOL_COMMA, tok)
	if 1 == cnt {
		return true
	}

	r.errMsg = fmt.Sprintf("Number of occurences of `%b` must be 1", constants.SYMBOL_COMMA)
	return false
}

func (r *HasOnlyOneCommaRule) count(needle byte, haystack string) int {
	var needleToStr string = string(needle)
	return strings.Count(haystack, needleToStr)
}

func (r *HasOnlyOneCommaRule) Key() alias.RuleKey {
	return constants.RULE_HAS_ONLY_ONE_COMMA
}

func (r *HasOnlyOneCommaRule) ErrorMsg() string {
	return r.errMsg
}
