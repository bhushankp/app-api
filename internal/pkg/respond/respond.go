package respond

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error     string `json:"error"`
	RequestID string `json:"request_id,omitempty"`
}

func JSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func Error(w http.ResponseWriter, status int, errMsg string, reqID string) {
	w.Header().Set("Content-Type", "application/json")
	if reqID != "" {
		w.Header().Set("X-Request-ID", reqID)
	}
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Error: errMsg, RequestID: reqID})
}
