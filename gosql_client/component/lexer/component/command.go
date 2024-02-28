package component

import (
	"errors"
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

func Parse(tokens []string) *Command {
	return &Command{Toks: tokens}
}
