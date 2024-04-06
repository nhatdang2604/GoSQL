package select_aggregator

import (
	"errors"
	"gosql_client/component/tokenizer/rule/rule_aggregator"
	"gosql_client/component/tokenizer/rule/rule_chain"
	"gosql_client/component/tokenizer/rule/rule_chain/start"
	"gosql_client/component/tokenizer/rule/rule_chain/start/common/is_semicolon"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/columns/column_with_dot"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/columns/column_with_dot/dot_at_first"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/columns/column_with_star"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/columns/column_without_dot"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/columns/common/column_name"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/columns/common/is_comma"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/froms"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/froms/table_name"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/froms/table_name/as"
	"gosql_client/component/tokenizer/rule/rule_chain/start/selects/froms/table_name/as/alias"
	"gosql_client/component/tokenizer/rule/rule_pool"
)

type SelectAggregator struct {
	toks         []string
	rulePool     rule_pool.RulePool
	curRuleChain rule_chain.RuleChain
}

func (a *SelectAggregator) init() error {
	var err error = nil
	a.rulePool = rule_pool.New()
	a.curRuleChain, err = a.buildFromSelectChain(a.toks)
	return err
}

func (a *SelectAggregator) buildFromSelectChain(remainToks []string) (rule_chain.RuleChain, error) {
	var err error = nil
	var entrypoint rule_chain.RuleChain = start.New(a.rulePool)
	if entrypoint.Exec(remainToks) {
		var remainToks []string = entrypoint.RemainToks()
		var selectChain rule_chain.RuleChain = selects.New(a.rulePool)
		if selectChain.Exec(remainToks) {
			entrypoint.SetNextRuleChain(selectChain)
			remainToks = selectChain.RemainToks()
			var columnWithDotChain rule_chain.RuleChain = column_with_dot.New(a.rulePool)
			var columnWithStarChain rule_chain.RuleChain = column_with_star.New(a.rulePool)
			var columnWithoutDotChain rule_chain.RuleChain = column_without_dot.New(a.rulePool)

			if columnWithDotChain.Validate(remainToks) {
				columnWithDotChain, err = a.buildFromColumnWithDotChain(remainToks)
				selectChain.SetNextRuleChain(columnWithDotChain)
			} else if columnWithStarChain.Exec(remainToks) {
				columnWithStarChain, err = a.buildFromColumnWithStarChain(remainToks)
				selectChain.SetNextRuleChain(columnWithStarChain)
			} else if columnWithoutDotChain.Exec(remainToks) {
				columnWithoutDotChain, err = a.buildFromColumnWithoutDotChain(remainToks)
				selectChain.SetNextRuleChain(columnWithoutDotChain)
			}
		} else {
			err = errors.New(*selectChain.ErrorMsg())
		}
	} else {
		err = errors.New(*entrypoint.ErrorMsg())
	}

	return entrypoint, err
}

func (a *SelectAggregator) buildFromColumnWithDotChain(remainToks []string) (rule_chain.RuleChain, error) {
	var entrypoint rule_chain.RuleChain = column_with_dot.New(a.rulePool)
	var err error = nil
	var columnWithDotChain rule_chain.RuleChain = entrypoint

	for {
		if columnWithDotChain.Exec(remainToks) {
			remainToks = columnWithDotChain.RemainToks()
			var dotAtFirstChain rule_chain.RuleChain = dot_at_first.New(a.rulePool)
			if dotAtFirstChain.Exec(remainToks) {
				columnWithDotChain.SetNextRuleChain(dotAtFirstChain)
				remainToks = dotAtFirstChain.RemainToks()
				var columnNameChain rule_chain.RuleChain = column_name.New(a.rulePool)
				if columnNameChain.Exec(remainToks) {
					dotAtFirstChain.SetNextRuleChain(columnNameChain)
					remainToks = columnNameChain.RemainToks()
					var commaChain rule_chain.RuleChain = is_comma.New(a.rulePool)
					var fromChain rule_chain.RuleChain = froms.New(a.rulePool)
					if commaChain.Exec(remainToks) {
						columnNameChain.SetNextRuleChain(commaChain)
						remainToks = commaChain.RemainToks()
						columnWithDotChain = column_with_dot.New(a.rulePool)
						if columnWithDotChain.Validate(remainToks) {
							commaChain.SetNextRuleChain(columnWithDotChain)
						}
					} else if fromChain.Exec(remainToks) {
						fromChain, err = a.buildFromFromChain(remainToks)
						if nil != err {
							columnNameChain.SetNextRuleChain(fromChain)
						}
						break
					} else {
						err = errors.New(*fromChain.ErrorMsg())
						break
					}
				} else {
					err = errors.New(*columnNameChain.ErrorMsg())
					break
				}
			} else {
				err = errors.New(*dotAtFirstChain.ErrorMsg())
				break
			}
		} else {
			err = errors.New(*columnWithDotChain.ErrorMsg())
			break
		}
	}

	return entrypoint, err
}

func (a *SelectAggregator) buildFromColumnWithStarChain(remainToks []string) (rule_chain.RuleChain, error) {
	var entrypoint rule_chain.RuleChain = column_with_star.New(a.rulePool)
	var err error = nil
	var columnWithStarChain rule_chain.RuleChain = entrypoint

	if columnWithStarChain.Exec(remainToks) {
		remainToks = columnWithStarChain.RemainToks()
		var fromChain rule_chain.RuleChain = froms.New(a.rulePool)
		if fromChain.Exec(remainToks) {
			fromChain, err = a.buildFromFromChain(remainToks)
			if nil != err {
				columnWithStarChain.SetNextRuleChain(fromChain)
			}
		} else {
			err = errors.New(*fromChain.ErrorMsg())
		}
	} else {
		err = errors.New(*columnWithStarChain.ErrorMsg())
	}

	return entrypoint, err
}

func (a *SelectAggregator) buildFromColumnWithoutDotChain(remainToks []string) (rule_chain.RuleChain, error) {
	var entrypoint rule_chain.RuleChain = column_without_dot.New(a.rulePool)
	var err error = nil
	var columnWithoutDotChain rule_chain.RuleChain = entrypoint

	for {
		if columnWithoutDotChain.Exec(remainToks) {
			remainToks = columnWithoutDotChain.RemainToks()
			var commaChain rule_chain.RuleChain = is_comma.New(a.rulePool)
			var fromChain rule_chain.RuleChain = froms.New(a.rulePool)
			if commaChain.Exec(remainToks) {
				columnWithoutDotChain.SetNextRuleChain(commaChain)
				remainToks = commaChain.RemainToks()
				columnWithoutDotChain = column_without_dot.New(a.rulePool)
				if columnWithoutDotChain.Validate(remainToks) {
					commaChain.SetNextRuleChain(columnWithoutDotChain)
				}
			} else if fromChain.Exec(remainToks) {
				fromChain, err = a.buildFromFromChain(remainToks)
				if nil != err {
					columnWithoutDotChain.SetNextRuleChain(fromChain)
				}
				break
			} else {
				err = errors.New(*fromChain.ErrorMsg())
				break
			}
		} else {
			err = errors.New(*columnWithoutDotChain.ErrorMsg())
			break
		}
	}

	return entrypoint, err
}

func (a *SelectAggregator) buildFromFromChain(remainToks []string) (rule_chain.RuleChain, error) {
	var entrypoint rule_chain.RuleChain = froms.New(a.rulePool)
	var err error = nil
	if entrypoint.Exec(remainToks) {
		remainToks = entrypoint.RemainToks()
		var tableNameChain rule_chain.RuleChain = table_name.New(a.rulePool)
		if tableNameChain.Exec(remainToks) {
			remainToks = tableNameChain.RemainToks()
			entrypoint.SetNextRuleChain(tableNameChain)
			var asChain rule_chain.RuleChain = as.New(a.rulePool)
			var semicolonChain rule_chain.RuleChain = is_semicolon.New(a.rulePool)

			if asChain.Exec(remainToks) {
				remainToks = asChain.RemainToks()
				tableNameChain.SetNextRuleChain(asChain)
				var aliasChain rule_chain.RuleChain = alias.New(a.rulePool)
				if aliasChain.Exec(remainToks) {
					remainToks = aliasChain.RemainToks()
					asChain.SetNextRuleChain(aliasChain)
					semicolonChain = is_semicolon.New(a.rulePool)
					if semicolonChain.Exec(remainToks) {
						aliasChain.SetNextRuleChain(semicolonChain)
					} else {
						err = errors.New(*semicolonChain.ErrorMsg())
					}
				} else {
					err = errors.New(*aliasChain.ErrorMsg())
				}
			} else if semicolonChain.Exec(remainToks) {
				tableNameChain.SetNextRuleChain(semicolonChain)
			} else {
				err = errors.New(*semicolonChain.ErrorMsg())
			}
		} else {
			err = errors.New(*tableNameChain.ErrorMsg())
		}
	} else {
		err = errors.New(*entrypoint.ErrorMsg())
	}

	return entrypoint, err
}

func (a *SelectAggregator) NextRuleChain() rule_chain.RuleChain {

	var curRuleChain rule_chain.RuleChain = a.curRuleChain
	a.curRuleChain = a.curRuleChain.NextRuleChain()

	return curRuleChain
}

func New(toks []string) rule_aggregator.RuleAggregator {
	var rulePool rule_pool.RulePool = rule_pool.New()
	var selectAggregator rule_aggregator.RuleAggregator = &SelectAggregator{
		toks:     toks,
		rulePool: rulePool,
	}

	return selectAggregator
}
