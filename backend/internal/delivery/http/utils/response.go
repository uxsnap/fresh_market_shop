package httpUtils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteErrorResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	encodeErr := json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
	if encodeErr != nil {
		log.Printf("failed to encode error response")
		w.Write([]byte(encodeErr.Error()))
		return
	}
}

func SetContentTypeApplicationJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func WriteResponseJson(w http.ResponseWriter, resp interface{}) {
	SetContentTypeApplicationJson(w)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode response: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}
