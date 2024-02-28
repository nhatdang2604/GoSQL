package interfaces

import (
	lexcom "gosql_client/component/lexer/component"
)

type Rule interface {
	Validate(command lexcom.Command) bool
}
