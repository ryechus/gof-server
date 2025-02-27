package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/placer14/gof-server/internal/database"
)

func StringFlagHandler(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")
	variation, err := database.GetCurrentFlagVariation(flagKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(variation.Value))
}

func IntFlagHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}

func FloatFlagHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}

func BoolFlagHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}

func SetFlagValue(w http.ResponseWriter, r *http.Request) {
	flagKeyPathValue := r.PathValue("flagKey")
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	activeVariationUUID := params.Get("activeVariationUUID")

	flagKey, err := database.SetCurrentFlagVariation(flagKeyPathValue, activeVariationUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseJson, err := json.Marshal(flagKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}
