package constant

type CommandKeyword = string

const (
	SELECT_KEYWORD CommandKeyword = "select"
	FROM_KEYWORD   CommandKeyword = "from"
	INSERT_KEYWORD CommandKeyword = "insert"
)

var RESERVED_KEYWORD []string = []string{
	SELECT_KEYWORD,
	FROM_KEYWORD,
	INSERT_KEYWORD,
}
