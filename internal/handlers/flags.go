package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/placer14/gof-server/internal/database"
)

func StringFlagHandler(w http.ResponseWriter, r *http.Request) {
	flagKey := r.PathValue("flagKey")
	variation, err := database.GetCurrentFlagVariation(flagKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(variation.Value))
}

func IntFlagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}

func FloatFlagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}

func BoolFlagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}

func GetFlag(w http.ResponseWriter, r *http.Request) {
	flagKeyPathValue := r.PathValue("flagKey")
	flagKey, err := database.GetFlag(flagKeyPathValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	responseJson, err := json.Marshal(flagKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}

func GetFlags(w http.ResponseWriter, r *http.Request) {
	flagVariations, err := database.GetFlags()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	responseJson, err := json.Marshal(flagVariations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}

func GetVariations(w http.ResponseWriter, r *http.Request) {
	flagKeyPathValue := r.PathValue("flagKey")
	flagVariations, err := database.GetVariations(flagKeyPathValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	responseJson, err := json.Marshal(flagVariations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}

func SetFlagValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("%s not allowed", r.Method), http.StatusBadRequest)
		return
	}
	// define custom type
	type Input struct {
		VariationUUID string `json:"variationUUID"`
	}
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	flagKeyPathValue := r.PathValue("flagKey")

	activeVariationUUID := input.VariationUUID

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

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}
