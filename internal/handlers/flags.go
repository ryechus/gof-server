package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/placer14/gof-server/internal/storage"
)

func GetStringValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	// responseJson, err := json.Marshal(responseType{Value: provider.StringFlagValues[flagKey].FlagValue})

	ctx := r.Context()
	ctx_storage := ctx.Value(storage.KeyInMemoryStorage)
	storage := ctx_storage.(*storage.InMemoryStorage)
	value, _ := storage.GetString(flagKey)
	responseJson, err := json.Marshal(responseType{Value: value})
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
	ctx_storage := ctx.Value(storage.KeyInMemoryStorage)
	storage := ctx_storage.(*storage.InMemoryStorage)

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

	storage.SetString(flagKey, input.FlagValue)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}

func GetFloatValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	ctx_storage := ctx.Value(storage.KeyInMemoryStorage)
	storage := ctx_storage.(*storage.InMemoryStorage)
	value, _ := storage.GetFloat(flagKey)

	responseJson, err := json.Marshal(responseType{Value: value})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}

func SetFloatValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	ctx_storage := ctx.Value(storage.KeyInMemoryStorage)
	storage := ctx_storage.(*storage.InMemoryStorage)

	// define custom type
	type Input struct {
		FlagValue float64 `json:"value"`
	}
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storage.SetFloat(flagKey, input.FlagValue)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}

func GetIntValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	ctx_storage := ctx.Value(storage.KeyInMemoryStorage)
	storage := ctx_storage.(*storage.InMemoryStorage)
	value, _ := storage.GetInt(flagKey)

	responseJson, err := json.Marshal(responseType{Value: int64(value)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}

func SetIntValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	ctx_storage := ctx.Value(storage.KeyInMemoryStorage)
	storage := ctx_storage.(*storage.InMemoryStorage)

	// define custom type
	type Input struct {
		FlagValue int64 `json:"value"`
	}
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storage.SetInt(flagKey, input.FlagValue)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}

func GetBoolValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	ctx_storage := ctx.Value(storage.KeyDBStorage)
	storage := ctx_storage.(*storage.DBStorage)
	value, err := storage.GetBool(flagKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	responseJson, err := json.Marshal(responseType{Value: value})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}

func SetBoolValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	ctx_storage := ctx.Value(storage.KeyDBStorage)
	storage := ctx_storage.(*storage.DBStorage)

	// define custom type
	type Input struct {
		FlagValue bool `json:"value"`
	}
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storage.SetBoolVariations(flagKey, []bool{true, false})

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}
