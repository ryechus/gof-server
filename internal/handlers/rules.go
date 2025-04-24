package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/placer14/gof-server/internal/config"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"gopkg.in/go-playground/validator.v8"
)

func PutRule(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx_storage := ctx.Value(config.KeyVariable)
	storageType := ctx_storage.(*config.FlagStorageType)
	validate := validator.New(ValidatorConfig)

	var input payloads.PutRule
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

	err = storageType.PutRule(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}
