package test

import (
	lexcom "gosql_client/component/lexer/component"
	irul "gosql_client/component/lexer/rule/interfaces"
	rul "gosql_client/component/lexer/rule/selects"
	"strings"
	"testing"
)

func TestSuccessCaseWithUppercaseKeyword(t *testing.T) {
	var rule irul.Rule = rul.HasKeywordAtFirstRule{}
	var tokens []string = []string{"SELECT"}
	var command *lexcom.Command = lexcom.Parse(tokens)
	var expected = true

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")
		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}

func TestSuccessCaseWithLowercaseKeyword(t *testing.T) {
	var rule irul.Rule = rul.HasKeywordAtFirstRule{}
	var tokens []string = []string{"select"}
	var command *lexcom.Command = lexcom.Parse(tokens)
	var expected = true

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")
		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}

func TestSuccessCaseWithRandomcaseKeyword(t *testing.T) {
	var rule irul.Rule = rul.HasKeywordAtFirstRule{}
	var tokens []string = []string{"sElEcT"}
	var command *lexcom.Command = lexcom.Parse(tokens)
	var expected = true

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")

		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}

func TestSuccessCaseWithUntrimmedKeyword(t *testing.T) {
	var rule irul.Rule = rul.HasKeywordAtFirstRule{}
	var tokens []string = []string{"  sElEcT  "}
	var command *lexcom.Command = lexcom.Parse(tokens)
	var expected = true

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")

		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}

func TestFailedCaseWithoutKeyword(t *testing.T) {
	var rule irul.Rule = rul.HasKeywordAtFirstRule{}
	var tokens []string = []string{"FROM", "test.sample_db"}
	var command *lexcom.Command = lexcom.Parse(tokens)
	var expected = false

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")
		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}

func TestFailedCaseWithoutAnything(t *testing.T) {
	var rule irul.Rule = rul.HasKeywordAtFirstRule{}
	var tokens []string = []string{}
	var command *lexcom.Command = lexcom.Parse(tokens)
	var expected = false

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")
		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}

func TestFailedCaseWithKeywordAtSecondPlace(t *testing.T) {
	var rule irul.Rule = rul.HasKeywordAtFirstRule{}
	var tokens []string = []string{"first-place", "SELECT"}
	var command *lexcom.Command = lexcom.Parse(tokens)
	var expected = false

	if actual := rule.Validate(*command); expected != actual {
		var input string = strings.Join(tokens, ",")
		t.Errorf("tokens = %s; expect %t; actual: %v", input, expected, actual)
	}
}
