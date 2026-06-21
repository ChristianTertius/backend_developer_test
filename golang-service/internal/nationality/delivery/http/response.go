package http

import (
	"encoding/json"
	"net/http"
)

type apiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func writeJSON(w http.ResponseWriter, code int, payload apiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(payload)
}

func respondData(w http.ResponseWriter, code int, data interface{}) {
	writeJSON(w, code, apiResponse{Status: "success", Data: data})
}

func respondError(w http.ResponseWriter, code int, message string) {
	writeJSON(w, code, apiResponse{Status: "error", Message: message})
}
