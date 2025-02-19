package provider_test

import (
	"context"
	"testing"

	"github.com/open-feature/go-sdk/openfeature"
	"github.com/placer14/gof-server/internal/provider"
	"github.com/stretchr/testify/assert"
)

func TestStringEvaluation(t *testing.T) {
	err := openfeature.SetProvider(provider.NewProvider())
	assert.NoError(t, err)

	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
	client := openfeature.NewClient("stringEvalTests")
	generation, err := client.StringValue(context.Background(), "dataplane_generation", "k8s.v1", evalCtx)
	assert.NoError(t, err)
	assert.Equal(t, "metal.v1", generation)
}

func TestBoolEvaluation(t *testing.T) {
	err := openfeature.SetProvider(provider.NewProvider())
	assert.NoError(t, err)

	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
	client := openfeature.NewClient("boolEvalTests")

	generation, err := client.BooleanValue(context.Background(), "grant_soil_access", true, evalCtx)
	assert.NoError(t, err)
	assert.Equal(t, false, generation)
}

func TestFloatEvaluation(t *testing.T) {
	err := openfeature.SetProvider(provider.NewProvider())
	assert.NoError(t, err)

	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
	client := openfeature.NewClient("floatEvalTests")

	generation, err := client.FloatValue(context.Background(), "special_ability_buff_perc", 0.0, evalCtx)

	assert.NoError(t, err)
	assert.Equal(t, 0.23456, generation)
}

func TestIntEvaluation(t *testing.T) {
	err := openfeature.SetProvider(provider.NewProvider())
	assert.NoError(t, err)

	evalCtx := openfeature.NewEvaluationContext("", map[string]any{})
	client := openfeature.NewClient("intEvalTests")

	generation, err := client.IntValue(context.Background(), "num_of_special_abilities", 0, evalCtx)

	assert.NoError(t, err)
	assert.Equal(t, int64(12), generation)
}
