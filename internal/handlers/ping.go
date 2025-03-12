package handlers

import (
	"encoding/json"
	"net/http"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type responseType struct {
		Value string `json:"value"`
	}
	responseJson, err := json.Marshal(responseType{Value: "pong"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(responseJson))
}
