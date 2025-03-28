package provider

import (
	"context"

	"github.com/open-feature/go-sdk/openfeature"
)

type MDUProvider openfeature.FeatureProvider

type MDUProviderImpl struct {
	store Storageable
}

func (p *MDUProviderImpl) Metadata() openfeature.Metadata { return openfeature.Metadata{} }
func (p *MDUProviderImpl) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	value, err := p.store.GetBool(flag)
	if err != nil {
		detail := openfeature.ProviderResolutionDetail{
			ResolutionError: openfeature.NewFlagNotFoundResolutionError(err.Error()),
			Reason:          openfeature.ErrorReason,
			Variant:         "",
			FlagMetadata:    openfeature.FlagMetadata{},
		}
		return openfeature.BoolResolutionDetail{Value: value, ProviderResolutionDetail: detail}
	}
	return openfeature.BoolResolutionDetail{Value: value}
}
func (p *MDUProviderImpl) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	value, err := p.store.GetString(flag)
	if err != nil {
		detail := openfeature.ProviderResolutionDetail{
			ResolutionError: openfeature.NewFlagNotFoundResolutionError(err.Error()),
			Reason:          openfeature.ErrorReason,
			Variant:         "",
			FlagMetadata:    openfeature.FlagMetadata{},
		}
		return openfeature.StringResolutionDetail{Value: value, ProviderResolutionDetail: detail}
	}

	return openfeature.StringResolutionDetail{Value: value}
}

func (p *MDUProviderImpl) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	return openfeature.FloatResolutionDetail{Value: FloatFlagValues[flag].FlagValue}
}
func (p *MDUProviderImpl) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	return openfeature.IntResolutionDetail{Value: IntFlagValues[flag].FlagValue}
}
func (p *MDUProviderImpl) ObjectEvaluation(ctx context.Context, flag string, defaultValue interface{}, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	return openfeature.InterfaceResolutionDetail{}
}
func (p *MDUProviderImpl) Hooks() []openfeature.Hook { return []openfeature.Hook{} }

func NewProvider(store Storageable) *MDUProviderImpl {
	// PopulateFlagValues()
	return &MDUProviderImpl{store: store}
}
