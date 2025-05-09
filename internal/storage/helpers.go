package storage

import (
	"strings"

	"github.com/placer14/gof-server/internal/handlers/payloads"
	"github.com/placer14/gof-server/internal/utils"
)

func doComparison(ctx payloads.RuleContext, value any) bool {
	switch strings.ToUpper(ctx.Operator) {
	case utils.IN, "":
		newList := utils.MakeNewList[string](ctx.Values)
		return utils.SliceContains(newList, value.(string))
	case utils.NOTIN, utils.NOT_IN:
		newList := utils.MakeNewList[string](ctx.Values)
		return !utils.SliceContains(newList, value.(string))
	case utils.CONTAINS:
		newList := utils.MakeNewList[string](ctx.Values)
		return utils.StringsContains(newList, value.(string))
	case utils.NOTCONTAINS:
		newList := utils.MakeNewList[string](ctx.Values)
		return !utils.StringsContains(newList, value.(string))
	case utils.GT:
		newList := utils.MakeNewList[float64](ctx.Values)
		minVal := utils.Min(newList...)
		return utils.IsGreaterThan(value.(float64), minVal)
	case utils.GTE:
		newList := utils.MakeNewList[float64](ctx.Values)
		minVal := utils.Min(newList...)
		return utils.IsGreaterThanEqual(value.(float64), minVal)
	case utils.LT:
		newList := utils.MakeNewList[float64](ctx.Values)
		maxVal := utils.Max(newList...)
		return utils.IsLessThan(value.(float64), maxVal)
	case utils.LTE:
		newList := utils.MakeNewList[float64](ctx.Values)
		maxVal := utils.Max(newList...)
		return utils.IsLessThanEqual(value.(float64), maxVal)
	default:
		return false
	}
}

func hasMatchingRules(flagRuleContexts []payloads.RuleContext, attributes map[string]any) bool {
	outcomes := make([]bool, len(flagRuleContexts))
	for flagIdx, ctx := range flagRuleContexts {
		for key, value := range attributes {
			if ctx.Attribute == key {
				outcomes[flagIdx] = doComparison(ctx, value)
			}
		}
	}

	for _, outcome := range outcomes {
		if !outcome {
			return false
		}
	}
	return true
}
