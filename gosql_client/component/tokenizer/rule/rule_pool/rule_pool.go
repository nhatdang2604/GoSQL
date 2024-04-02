package rule_pool

import (
	"gosql_client/component/tokenizer/alias"
	"gosql_client/component/tokenizer/rule/rule_unit"
	"gosql_client/component/tokenizer/rule/rule_unit/common"
	"gosql_client/component/tokenizer/rule/rule_unit/selectcmd"
)

// alias
type Rule = rule_unit.Rule
type RuleKey = alias.RuleKey

type RulePool interface {
	Get(key RuleKey) Rule
}

type RulePoolImpl struct {
	pool map[RuleKey]Rule
}

func (rp RulePoolImpl) Get(key RuleKey) Rule {
	return rp.pool[key]
}

func New() RulePool {
	var rules []Rule = []Rule{
		&common.IsStartRule{},
		&selectcmd.IsSelectRule{},
		&common.IsInsertRule{},
	}

	var pool map[RuleKey]Rule = map[RuleKey]Rule{}
	for _, rule := range rules {
		pool[rule.Key()] = rule
	}

	return RulePoolImpl{pool: pool}
}
