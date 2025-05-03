package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/placer14/gof-server/internal/config"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"gopkg.in/go-playground/validator.v8"
)

var ValidatorConfig = &validator.Config{TagName: "validate"}

func GetFlagWithVariations(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)

	response, err := storageType.GetFlagWithVariations(flagKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	responseJson, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}

func EvaluateFlag(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")

	ctx := r.Context()
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)

	validate := validator.New(ValidatorConfig)
	var input payloads.EvaluateFlag
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

	w.Header().Set("content-type", "application/json")
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
	var input payloads.CreateFlag
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

	err = storageType.CreateFlag(input.Key, input.FlagType, input.Variations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}
