package constants

import "gosql_client/component/tokenizer/alias"

const (
	RULE_IS_START                  alias.RuleKey = "start_to_tokenize"
	RULE_IS_SELECT                 alias.RuleKey = "is_select"
	RULE_IS_FROM                   alias.RuleKey = "is_from"
	RULE_IS_AS                     alias.RuleKey = "is_as"
	RULE_IS_INSERT                 alias.RuleKey = "is_insert"
	RULE_IS_INTO                   alias.RuleKey = "is_into"
	RULE_IS_VALUES                 alias.RuleKey = "is_values"
	RULE_IS_STAR                   alias.RuleKey = "is_star"
	RULE_HAS_DOT                   alias.RuleKey = "has_dot"
	RULE_HAS_ONLY_ONE_DOT          alias.RuleKey = "has_only_one_dot"
	RULE_HAS_DOT_AT_FIRST          alias.RuleKey = "has_dot_at_first"
	RULE_HAS_COMMA                 alias.RuleKey = "has_comma"
	RULE_HAS_SEMICOLON             alias.RuleKey = "has_semicolon"
	RULE_HAS_ONLY_ONE_COMMA        alias.RuleKey = "has_only_one_comma"
	RULE_HAS_COMMA_AT_LAST         alias.RuleKey = "has_comma_at_last"
	RULE_HAS_SEMICOLON_AT_LAST     alias.RuleKey = "has_semicolon_at_last"
	RULE_IS_DOT                    alias.RuleKey = "is_dot"
	RULE_IS_COMMA                  alias.RuleKey = "is_comma"
	RULE_IS_SEMICOLON              alias.RuleKey = "is_semicolon"
	RULE_IS_OPEN_BRACKET           alias.RuleKey = "is_open_bracket"
	RULE_HAS_OPEN_BRACKET_AT_FIRST alias.RuleKey = "has_open_bracket_at_first"
)
