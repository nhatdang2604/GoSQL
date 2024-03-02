package component

import (
	"errors"
	"fmt"
	"gosql_client/helper"
	"strings"
)

type HasToken interface {
	Tokens() []string
}

type Command struct {
	Toks []string
}

func (c *Command) Tokens() []string {
	return c.Toks
}

func (c *Command) GetFirstToken() (string, error) {
	if c.IsEmpty() {
		return "", errors.New("command has empty tokens")
	}

	var firstToken, err = c.GetTokenAt(0)

	return firstToken, err
}

func (c *Command) IsEmpty() bool {
	var isEmpty bool = (len(c.Toks) == 0)
	return isEmpty
}

func (c *Command) GetTokenAt(i int) (string, error) {
	if len(c.Toks) <= i {
		return "", errors.New("index out of bound")
	}

	var trimmedTok string = strings.TrimSpace(c.Toks[i])

	return trimmedTok, nil
}

func (c *Command) FindTokenIgnoringCaseSensitive(needle string) []int {
	var indexes []int = []int{}

	for index, tok := range c.Toks {
		if helper.IsTokenEqualIgnoringCase(needle, tok) {
			indexes = append(indexes, index)
		}
	}

	return indexes
}

func (c *Command) FindKeyword(keyword string) (int, error) {

	//Check valid keyword
	var isKeyword bool = helper.IsKeyword(keyword)
	if !isKeyword {
		return -1, errors.New(fmt.Sprintf("'%s' is not a keyword", keyword))
	}

	//Find keyword
	for idx, tok := range c.Toks {
		if helper.IsTokenEqualIgnoringCase(keyword, tok) {
			return idx, nil
		}
	}

	return -1, errors.New(fmt.Sprintf("Command doesn't contain keyword = '%s'", keyword))
}

func MakeCommand(tokens []string) *Command {
	return &Command{Toks: tokens}
}
