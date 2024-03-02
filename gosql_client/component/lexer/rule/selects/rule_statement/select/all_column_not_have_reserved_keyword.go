package rule_unit

import (
	lexcom "gosql_client/component/lexer/component"
	constant "gosql_client/constant"
	"gosql_client/helper"
)

type ColumnNotHasReservedKeywordRule struct{}

func (ColumnNotHasReservedKeywordRule) Validate(command lexcom.Command) bool {

	var selectIndex, err1 = command.FindKeyword(constant.SELECT_KEYWORD)
	var fromIndex, err2 = command.FindKeyword(constant.FROM_KEYWORD)

	if nil != err1 || nil != err2 {
		return false
	}

	for columnIdx := selectIndex + 1; columnIdx < fromIndex; columnIdx++ {
		var columnName, _ = command.GetTokenAt(columnIdx)

		for _, keyword := range constant.RESERVED_KEYWORD {
			if helper.IsTokenEqualIgnoringCase(columnName, keyword) {
				return false
			}
		}
	}

	return true
}
