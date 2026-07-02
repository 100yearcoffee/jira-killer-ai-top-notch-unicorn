package http

import (
	"encoding/json"
	nethttp "net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w nethttp.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(value)
}

func writeError(w nethttp.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Error: message})
}
