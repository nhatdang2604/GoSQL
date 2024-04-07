package insertcmd

import (
	"fmt"
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/constants"
	"gosql_client/component/tokenizer/rule/rule_input"
)

type HasOpenBracketAtFirstRule struct {
	errMsg string
}

func (r *HasOpenBracketAtFirstRule) Validate(i interface{}) bool {
	var st, ok = i.(rule_input.SingleTok)

	if !ok {
		r.errMsg = "unexpected input type for HasOpenBracketAtFirstRule.Validate()"
		return false
	}

	var tok string = st.Tok
	var perhapOpenBracket byte = tok[0]
	var openBracket byte = constants.SYMBOL_OPEN_BRACKET
	if perhapOpenBracket == openBracket {
		return true
	}

	r.errMsg = fmt.Sprintf("Expect the 0-idx is `%b`, but found: `%b`", constants.SYMBOL_OPEN_BRACKET, perhapOpenBracket)
	return false
}

func (r *HasOpenBracketAtFirstRule) Key() alias.RuleKey {
	return constants.RULE_HAS_OPEN_BRACKET_AT_FIRST
}

func (r *HasOpenBracketAtFirstRule) ErrorMsg() string {
	return r.errMsg
}
