package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/placer14/gof-server/internal/provider"
)

func GetStringValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	flagKey := r.PathValue("flagKey")

	responseJson, err := json.Marshal(responseType{Value: provider.StringFlagValues[flagKey].FlagValue})

	// ctx := r.Context()
	// m := ctx.Value(provider.KeyFlagStore)
	// mImpl := m.(*provider.MDUProviderMock)
	// responseJson, err := json.Marshal(responseType{Value: mImpl.GetString(flagKey)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}

func GetFloatValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	flagKey := r.PathValue("flagKey")

	responseJson, err := json.Marshal(responseType{Value: provider.BoolFlagValues[flagKey].FlagValue})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}
