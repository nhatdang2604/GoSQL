package selects

import (
	lexcom "gosql_client/component/lexer/component"
	irul "gosql_client/component/lexer/rule/interfaces"
	"strings"
	"testing"
)

func TestHappyCase(t *testing.T) {
	var rule irul.Rule = HasKeywordRule{}
	var tokens []string = []string{"SELECT"}
	var command *lexcom.Command = &lexcom.Command{Toks: tokens}

	if actual := rule.Validate(*command); !actual {
		var input string = strings.Join(tokens, ",")
		var expect = true
		t.Errorf("tokens = %s; expect %t; actual: %v", input, expect, actual)
	}
}
