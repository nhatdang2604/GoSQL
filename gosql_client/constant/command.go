package constant

type CommandKeyword = string

const (
	SELECT_KEYWORD CommandKeyword = "select"
	FROM_KEYWORD   CommandKeyword = "from"

	INSERT_KEYWORD CommandKeyword = "insert"
	INTO_KEYWORD   CommandKeyword = "into"
	VALUES_KEYWORD CommandKeyword = "values"
)

var RESERVED_KEYWORD []string = []string{
	SELECT_KEYWORD,
	FROM_KEYWORD,

	INSERT_KEYWORD,
	INTO_KEYWORD,
	VALUES_KEYWORD,
}
