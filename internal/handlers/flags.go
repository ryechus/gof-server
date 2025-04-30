package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/placer14/gof-server/internal/config"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"gopkg.in/go-playground/validator.v8"
)

var ValidatorConfig = &validator.Config{TagName: "validate"}

func GetStringValue(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

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
	value, _ := storageType.GetFloat(flagKey)

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

func GetFlag(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)

	validate := validator.New(ValidatorConfig)
	var input payloads.GetFlag
	d := json.NewDecoder(r.Body)
	err := d.Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = validate.Struct(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	value, _err := storageType.EvaluateFlag(flagKey, input)

	if _err != nil {
		http.Error(w, _err.Error(), http.StatusInternalServerError)
		return
	}

	responseJson, err := json.Marshal(responseType{Value: value})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}

func UpdateFlag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)

	validate := validator.New(ValidatorConfig)
	var input payloads.UpdateFlag
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

	err = storageType.UpdateFlag(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}

func CreateFlag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)
	validate := validator.New(ValidatorConfig)
	var input createFlagPayload
	d := json.NewDecoder(r.Body)

	err := d.Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate.Struct(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storageType.CreateFlag(input.Key, input.FlagType, input.Variations)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}
