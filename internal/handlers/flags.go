package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/placer14/gof-server/internal/config"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"github.com/placer14/gof-server/internal/storage"
	"gopkg.in/go-playground/validator.v8"
)

func GetStringValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	// responseJson, err := json.Marshal(responseType{Value: provider.StringFlagValues[flagKey].FlagValue})

	ctx := r.Context()
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)
	value, _ := storageType.GetString(flagKey)
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
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)

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

	storageType.SetString(flagKey, input.FlagValue)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}

func GetFloatValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)
	value, _ := storageType.GetFloat(flagKey)

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
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)

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

	storageType.SetFloat(flagKey, input.FlagValue)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}

func GetIntValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)
	value, _ := storageType.GetInt(flagKey)

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
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)

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

	storageType.SetInt(flagKey, input.FlagValue)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}

func GetBoolValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)
	value, err := storageType.GetBool(flagKey)
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
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)

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

	storageType.SetBool(flagKey, input.FlagValue)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}

func CreateFlag(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// ctx_storage := ctx.Value(config.KeyVariable)
	// storageType := ctx_storage.(*config.FlagStorageType)

	config := &validator.Config{TagName: "validate"}

	validate := validator.New(config)
	var input createFlagPayload
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate.Struct(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch input.FlagType {
	case "bool":
		variations := createVariations[bool](input.Variations)
		storage.CreateFlag[bool](input.Key, variations)
	case "string":
		variations := createVariations[string](input.Variations)
		storage.CreateFlag[string](input.Key, variations)
	case "float":
		variations := createVariations[float64](input.Variations)
		storage.CreateFlag[float64](input.Key, variations)
	case "int":
		variations := createVariations[int64](input.Variations)
		storage.CreateFlag[int64](input.Key, variations)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}

func createVariations[T comparable](variations []flagVariation) []payloads.FlagVariation {
	var castedVariations []payloads.FlagVariation
	for _, variation := range variations {
		as_bool, ok := variation.Value.(T)
		if !ok {
			// http.Error(w, "something went wrong", http.StatusBadRequest)
			panic("there was a problem")
		}
		castedVariations = append(castedVariations, payloads.FlagVariation{Value: as_bool, Name: variation.Name})
	}
	return castedVariations
}
