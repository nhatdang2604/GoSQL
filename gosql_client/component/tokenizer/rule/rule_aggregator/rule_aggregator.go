package rule_aggregator

import (
	"gosql_client/component/tokenizer/rule/rule_chain"
)

type RuleAggregator interface {
	NextRuleChain() rule_chain.RuleChain
}
