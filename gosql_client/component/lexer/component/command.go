package component

import (
	"errors"
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
	if len(c.Toks) == 0 {
		return "", errors.New("Command has empty tokens")
	}

	return c.Toks[0], nil
}

func Parse(tokens []string) *Command {
	return &Command{Toks: tokens}
}
