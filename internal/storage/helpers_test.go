package storage_test

import (
	"testing"

	"github.com/placer14/gof-server/internal/handlers/payloads"
	"github.com/placer14/gof-server/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestDoComparison(t *testing.T) {
	testCases := map[string]struct {
		ctx      payloads.RuleContext
		value    any
		expected bool
	}{
		"in-true": {
			ctx: payloads.RuleContext{
				Operator: "IN",
				Values:   []any{"hello", "world"},
			},
			value:    "world",
			expected: true,
		},
		"in-false": {
			ctx: payloads.RuleContext{
				Operator: "IN",
				Values:   []any{"hello", "world"},
			},
			value:    "goodbye",
			expected: false,
		},
		"not in-true": {
			ctx: payloads.RuleContext{
				Operator: "NOT_IN",
				Values:   []any{"hello", "world"},
			},
			value:    "goodbye",
			expected: true,
		},
		"not in-false": {
			ctx: payloads.RuleContext{
				Operator: "NOT_IN",
				Values:   []any{"hello", "world"},
			},
			value:    "world",
			expected: false,
		},
		"containts-true": {
			ctx: payloads.RuleContext{
				Operator: "CONTAINS",
				Values:   []any{"hello", "world"},
			},
			value:    "all around the world",
			expected: true,
		},
		"containts-false": {
			ctx: payloads.RuleContext{
				Operator: "CONTAINS",
				Values:   []any{"hello", "world"},
			},
			value:    "at the bottom of the well",
			expected: false,
		},
		"not contains-false": {
			ctx: payloads.RuleContext{
				Operator: "NOT_CONTAINS",
				Values:   []any{"hello", "world"},
			},
			value:    "all around the world",
			expected: false,
		},
		"not contains-true": {
			ctx: payloads.RuleContext{
				Operator: "NOT_CONTAINS",
				Values:   []any{"hello", "world"},
			},
			value:    "at the bottom of the well",
			expected: true,
		},
		"gt-true": {
			ctx: payloads.RuleContext{
				Operator: "GT",
				Values:   []any{float64(101), float64(60)},
			},
			value:    float64(61),
			expected: true,
		},
		"gt-false": {
			ctx: payloads.RuleContext{
				Operator: "GT",
				Values:   []any{float64(101), float64(60)},
			},
			value:    float64(59),
			expected: false,
		},
		"gte-true-1": {
			ctx: payloads.RuleContext{
				Operator: "GTE",
				Values:   []any{float64(101), float64(60)},
			},
			value:    float64(60),
			expected: true,
		},
		"gte-true-2": {
			ctx: payloads.RuleContext{
				Operator: "GTE",
				Values:   []any{float64(101), float64(60)},
			},
			value:    float64(61),
			expected: true,
		},
		"gte-false": {
			ctx: payloads.RuleContext{
				Operator: "GTE",
				Values:   []any{float64(101), float64(60)},
			},
			value:    float64(59),
			expected: false,
		},
		"LTE-true-1": {
			ctx: payloads.RuleContext{
				Operator: "LTE",
				Values:   []any{float64(101), float64(60)},
			},
			value:    float64(100),
			expected: true,
		},
		"LTE-true-2": {
			ctx: payloads.RuleContext{
				Operator: "LTE",
				Values:   []any{float64(101), float64(60)},
			},
			value:    float64(101),
			expected: true,
		},
		"LTE-false": {
			ctx: payloads.RuleContext{
				Operator: "LTE",
				Values:   []any{float64(101), float64(60)},
			},
			value:    float64(102),
			expected: false,
		},
		"unknown-false": {
			ctx: payloads.RuleContext{
				Operator: "DNE",
				Values:   []any{float64(101), float64(60)},
			},
			value:    float64(102),
			expected: false,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, storage.DoComparison(test.ctx, test.value), test.expected)
		})
	}
}

func TestHasMatchingRules(t *testing.T) {
	testEmail := "example@email.com"
	flagRuleContexts := []payloads.RuleContext{
		{
			ContextKind: "user",
			Attribute:   "email",
			Values:      []any{testEmail},
			Operator:    "IN",
		},
	}
	attributes := map[string]any{
		"zipcode": "12345",
		"email":   testEmail,
	}
	testCases := map[string]struct {
		ruleContexts []payloads.RuleContext
		attributes   map[string]any
		expected     bool
	}{
		"one-rule-context-true": {
			[]payloads.RuleContext{
				{
					ContextKind: "user",
					Attribute:   "email",
					Values:      []any{testEmail},
					Operator:    "IN",
				},
			},
			map[string]any{
				"zipcode": "12345",
				"email":   testEmail,
			},
			true,
		},
		"multiple-rule-context-true": {
			[]payloads.RuleContext{
				{
					ContextKind: "user",
					Attribute:   "email",
					Values:      []any{testEmail},
					Operator:    "IN",
				},
				{
					ContextKind: "user",
					Attribute:   "zipcode",
					Values:      []any{"12345"},
					Operator:    "IN",
				},
			},
			map[string]any{
				"zipcode": "12345",
				"email":   testEmail,
			},
			true,
		},
		"multiple-rule-context-false": {
			[]payloads.RuleContext{
				{
					ContextKind: "user",
					Attribute:   "email",
					Values:      []any{testEmail},
					Operator:    "IN",
				},
				{
					ContextKind: "user",
					Attribute:   "zipcode",
					Values:      []any{"12345"},
					Operator:    "IN",
				},
			},
			map[string]any{
				"zipcode": "12346",
				"email":   testEmail,
			},
			false,
		},
		"multiple-rule-context-gte-true": {
			[]payloads.RuleContext{
				{
					ContextKind: "user",
					Attribute:   "email",
					Values:      []any{testEmail},
					Operator:    "IN",
				},
				{
					ContextKind: "user",
					Attribute:   "birth_year",
					Values:      []any{float64(1990)},
					Operator:    "GTE",
				},
			},
			map[string]any{
				"birth_year": float64(1990),
				"email":      testEmail,
			},
			true,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, storage.HasMatchingRules(testCase.ruleContexts, testCase.attributes))
		})
	}
	assert.True(t, storage.HasMatchingRules(flagRuleContexts, attributes))
}
