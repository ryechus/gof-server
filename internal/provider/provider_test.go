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

	evalCtx := openfeature.NewEvaluationContext("blah1234", map[string]any{
		"kind":            "organization-id",
		"organization-id": "blah1234",
		"redpanda-id":     "redpanda-blah12342",
		"key":             "redpanda-blah12343",
		"cloud-provider":  "aws",
		"anonymous":       true,
	})
	client := openfeature.NewClient("stringEvalTests")
	generation, err := client.StringValue(context.Background(), "dataplane_generation", "k8s.v1", evalCtx)
	assert.NoError(t, err)
	assert.Equal(t, "metal.v1", generation)
}
