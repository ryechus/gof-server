package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/placer14/gof-server/internal/provider"
)

func GetStringValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	// responseJson, err := json.Marshal(responseType{Value: provider.StringFlagValues[flagKey].FlagValue})

	ctx := r.Context()
	m := ctx.Value(provider.KeyFlagStore)
	mImpl := m.(*provider.MDUProviderMock)
	responseJson, err := json.Marshal(responseType{Value: mImpl.GetString(flagKey)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}

func SetStringvalue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	m := ctx.Value(provider.KeyFlagStore)
	mImpl := m.(*provider.MDUProviderMock)

	// define custom type
	type Input struct {
		FlagValue string `json:"value"`
	}
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mImpl.SetString(flagKey, input.FlagValue)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}

func GetFloatValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	responseJson, err := json.Marshal(responseType{Value: provider.FloatFlagValues[flagKey].FlagValue})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}

func GetIntValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	responseJson, err := json.Marshal(responseType{Value: int64(provider.IntFlagValues[flagKey].FlagValue)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}

func GetBoolValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	responseJson, err := json.Marshal(responseType{Value: provider.BoolFlagValues[flagKey].FlagValue})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}
