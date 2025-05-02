package provider_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/open-feature/go-sdk/openfeature"
// 	"github.com/placer14/gof-server/internal/provider"
// 	"github.com/placer14/gof-server/internal/storage"
// 	"github.com/stretchr/testify/assert"
// )

// func TestStringEvaluation(t *testing.T) {
// 	store := storage.NewInMemoryStorage()
// 	subject := provider.NewProvider(store)

// 	flagKey := "dataplane-generation"

// 	err := openfeature.SetProvider(subject)
// 	assert.NoError(t, err)

// 	expectedStr := "metal.v1"
// 	err = store.SetString(flagKey, expectedStr)
// 	assert.NoError(t, err)

// 	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
// 	client := openfeature.NewClient("stringEvalTests")
// 	generation, err := client.StringValue(context.Background(), flagKey, "k8s.v1", evalCtx)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedStr, generation)
// }

// func TestBoolEvaluation(t *testing.T) {
// 	store := storage.NewInMemoryStorage()
// 	subject := provider.NewProvider(store)

// 	flagKey := "grant-soil-access"

// 	err := openfeature.SetProvider(subject)
// 	assert.NoError(t, err)

// 	value := true
// 	store.SetBool(flagKey, value)

// 	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
// 	client := openfeature.NewClient("boolEvalTests")

// 	generation, err := client.BooleanValue(context.Background(), flagKey, false, evalCtx)
// 	assert.NoError(t, err)
// 	assert.Equal(t, value, generation)
// }

// func TestFloatEvaluation(t *testing.T) {
// 	store := storage.NewInMemoryStorage()
// 	subject := provider.NewProvider(store)

// 	flagKey := "percent-failure-allowed"
// 	percentFailAllowed := float64(0.5)

// 	store.SetFloat(flagKey, percentFailAllowed)
// 	err := openfeature.SetProvider(subject)
// 	assert.NoError(t, err)

// 	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
// 	client := openfeature.NewClient("floatEvalTests")

// 	generation, err := client.FloatValue(context.Background(), flagKey, 0.0, evalCtx)

// 	assert.NoError(t, err)
// 	assert.Equal(t, percentFailAllowed, generation)
// }

// func TestIntEvaluation(t *testing.T) {
// 	store := storage.NewInMemoryStorage()
// 	subject := provider.NewProvider(store)

// 	flagKey := "num-workers"
// 	numWorkers := 5

// 	store.SetInt(flagKey, int64(numWorkers))
// 	err := openfeature.SetProvider(subject)
// 	assert.NoError(t, err)

// 	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
// 	client := openfeature.NewClient("intEvalTests")

// 	generation, err := client.IntValue(context.Background(), flagKey, 0, evalCtx)

// 	assert.NoError(t, err)
// 	assert.Equal(t, int64(numWorkers), generation)
// }
