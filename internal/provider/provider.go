package provider

import (
	"context"

	"github.com/open-feature/go-sdk/openfeature"
)

type MDUProvider openfeature.FeatureProvider

type MDUProviderImpl struct{}

func (MDUProviderImpl) Metadata() openfeature.Metadata { return openfeature.Metadata{} }
func (MDUProviderImpl) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	return openfeature.BoolResolutionDetail{}
}
func (MDUProviderImpl) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	return openfeature.StringResolutionDetail{}
}
func (MDUProviderImpl) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	return openfeature.FloatResolutionDetail{}
}
func (MDUProviderImpl) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	return openfeature.IntResolutionDetail{}
}
func (MDUProviderImpl) ObjectEvaluation(ctx context.Context, flag string, defaultValue interface{}, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	return openfeature.InterfaceResolutionDetail{}
}
func (MDUProviderImpl) Hooks() []openfeature.Hook { return []openfeature.Hook{} }

func NewProvider() MDUProviderImpl {
	return MDUProviderImpl{}
}
