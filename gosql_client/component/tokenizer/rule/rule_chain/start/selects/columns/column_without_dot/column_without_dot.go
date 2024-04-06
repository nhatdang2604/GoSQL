package column_without_dot

import (
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/columns/common/column_name"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type ColumnWithoutDotChain = column_name.ColumnNameChain

func New(pool rule_pool.RulePool) rule_chain.RuleChain {

	var chain *ColumnWithoutDotChain = &ColumnWithoutDotChain{}
	chain.SetRulePool(pool)

	return chain
}
