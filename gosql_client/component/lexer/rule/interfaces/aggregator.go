package interfaces

import (
	lexcom "gosql_client/component/lexer/component"
)

type RuleAggregatable interface {
	ValidateMultipleRules(command lexcom.Command) bool
}
