package test

import (
	lexcom "gosql_client/component/lexer/component"
	urul "gosql_client/component/lexer/rule/insertcmd"
	irul "gosql_client/component/lexer/rule/interfaces"
	"strings"
	"testing"
)

func TestSuccessCaseWithUppercaseKeyword(t *testing.T) {
	var rule irul.Rule = urul.HasKeywordAtFirstRule{}
	var tokens []string = []string{"INSERT"}
	var command *lexcom.Command = lexcom.MakeCommand(tokens)
	var expected = true

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")
		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}

func TestSuccessCaseWithLowercaseKeyword(t *testing.T) {
	var rule irul.Rule = urul.HasKeywordAtFirstRule{}
	var tokens []string = []string{"insert"}
	var command *lexcom.Command = lexcom.MakeCommand(tokens)
	var expected = true

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")
		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}

func TestSuccessCaseWithRandomcaseKeyword(t *testing.T) {
	var rule irul.Rule = urul.HasKeywordAtFirstRule{}
	var tokens []string = []string{"iNsErT"}
	var command *lexcom.Command = lexcom.MakeCommand(tokens)
	var expected = true

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")

		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}

func TestSuccessCaseWithUntrimmedKeyword(t *testing.T) {
	var rule irul.Rule = urul.HasKeywordAtFirstRule{}
	var tokens []string = []string{"  iNsErT  "}
	var command *lexcom.Command = lexcom.MakeCommand(tokens)
	var expected = true

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")

		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}

func TestFailedCaseWithoutKeyword(t *testing.T) {
	var rule irul.Rule = urul.HasKeywordAtFirstRule{}
	var tokens []string = []string{"INTO", "test.sample_db"}
	var command *lexcom.Command = lexcom.MakeCommand(tokens)
	var expected = false

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")
		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}

func TestFailedCaseWithoutAnything(t *testing.T) {
	var rule irul.Rule = urul.HasKeywordAtFirstRule{}
	var tokens []string = []string{}
	var command *lexcom.Command = lexcom.MakeCommand(tokens)
	var expected = false

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")
		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}

func TestFailedCaseWithKeywordAtSecondPlace(t *testing.T) {
	var rule irul.Rule = urul.HasKeywordAtFirstRule{}
	var tokens []string = []string{"first-place", "INSERT"}
	var command *lexcom.Command = lexcom.MakeCommand(tokens)
	var expected = false

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")
		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}
