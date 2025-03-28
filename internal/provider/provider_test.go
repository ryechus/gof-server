package provider_test

import (
	"context"
	"testing"

	"github.com/open-feature/go-sdk/openfeature"
	"github.com/placer14/gof-server/internal/provider"
	"github.com/stretchr/testify/assert"
)

func TestStringEvaluation(t *testing.T) {
	store := provider.NewStorage()
	subject := provider.NewProvider(store)
	err := openfeature.SetProvider(subject)
	assert.NoError(t, err)
	// m := provider.NewProviderMock()
	// assert.NoError(t, openfeature.SetProvider(m))

	flagKey := "dataplane-generation"
	expectedStr := "metal.v1"
	assert.NoError(t, store.SetString(flagKey, expectedStr))

	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
	client := openfeature.NewClient("stringEvalTests")
	generation, err := client.StringValue(context.Background(), flagKey, "k8s.v1", evalCtx)
	assert.NoError(t, err)
	assert.Equal(t, expectedStr, generation)
}

// func TestBoolEvaluation(t *testing.T) {
// 	m := provider.NewProviderMock()
// 	flagKey := "grant-soil-access"

// 	err := openfeature.SetProvider(m)
// 	assert.NoError(t, err)

// 	value := true
// 	m.SetBool(flagKey, value)

// 	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
// 	client := openfeature.NewClient("boolEvalTests")

// 	generation, err := client.BooleanValue(context.Background(), flagKey, false, evalCtx)
// 	assert.NoError(t, err)
// 	assert.Equal(t, value, generation)
// }

// func TestFloatEvaluation(t *testing.T) {
// 	m := provider.NewProviderMock()
// 	flagKey := "percent-failure-allowed"
// 	percentFailAllowed := float64(0.5)

// 	m.SetFloat(flagKey, percentFailAllowed)
// 	err := openfeature.SetProvider(m)
// 	assert.NoError(t, err)

// 	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
// 	client := openfeature.NewClient("floatEvalTests")

// 	generation, err := client.FloatValue(context.Background(), flagKey, 0.0, evalCtx)

// 	assert.NoError(t, err)
// 	assert.Equal(t, percentFailAllowed, generation)
// }

// func TestIntEvaluation(t *testing.T) {
// 	m := provider.NewProviderMock()
// 	flagKey := "num-workers"
// 	numWorkers := 5

// 	m.SetInt(flagKey, int64(numWorkers))
// 	err := openfeature.SetProvider(m)
// 	assert.NoError(t, err)

// 	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
// 	client := openfeature.NewClient("intEvalTests")

// 	generation, err := client.IntValue(context.Background(), flagKey, 0, evalCtx)

// 	assert.NoError(t, err)
// 	assert.Equal(t, int64(numWorkers), generation)
// }
