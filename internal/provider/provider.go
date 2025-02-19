package provider

import (
	"context"
	// "io"
	// "net/http"

	"github.com/open-feature/go-sdk/openfeature"
)

type MDUProvider openfeature.FeatureProvider

type MDUProviderImpl struct{}

func (MDUProviderImpl) Metadata() openfeature.Metadata { return openfeature.Metadata{} }
func (MDUProviderImpl) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	return openfeature.BoolResolutionDetail{Value: BoolFlagValues[flag].FlagValue}
}
func (MDUProviderImpl) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	// url := "http://127.0.0.1:23456/get_string_value"
	// client := &http.Client{}
	// resp, err := client.Get(url)
	// if err != nil {
	// 	return openfeature.StringResolutionDetail{Value: "goodbye, world"}
	// }
	// defer resp.Body.Close()
	// body, _ := io.ReadAll(resp.Body)
	// FlagValues["dataplane_generation"] = StringFlagValue{FlagKey: "dataplane_generation", FlagValue: "metal.v1"}
	return openfeature.StringResolutionDetail{Value: StringFlagValues[flag].FlagValue}
}
func (MDUProviderImpl) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	return openfeature.FloatResolutionDetail{Value: FloatFlagValues[flag].FlagValue}
}
func (MDUProviderImpl) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	return openfeature.IntResolutionDetail{Value: IntFlagValues[flag].FlagValue}
}
func (MDUProviderImpl) ObjectEvaluation(ctx context.Context, flag string, defaultValue interface{}, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	return openfeature.InterfaceResolutionDetail{}
}
func (MDUProviderImpl) Hooks() []openfeature.Hook { return []openfeature.Hook{} }

func NewProvider() MDUProviderImpl {
	PopulateFlagValues()
	return MDUProviderImpl{}
}
