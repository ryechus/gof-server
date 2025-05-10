package handlers

import (
	"encoding/json"
	"github.com/placer14/gof-server/internal/config"
	"github.com/placer14/gof-server/internal/handlers/payloads"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
)

func CreateRollout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctxStorage := ctx.Value(config.KeyVariable)
	storageType := ctxStorage.(*config.FlagStorageType)

	validate := validator.New(ValidatorConfig)

	var input payloads.PutRolloutRule
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

	err = storageType.SetupRolloutRule(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//responseJson, err := json.Marshal(scalarResponse{Value: "pong"})
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(""))
}
