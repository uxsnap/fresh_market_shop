package httpUtils

import (
	"encoding/json"
	"log"
	"net/http"

	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

type ErrorResponse struct {
	Error errorWrapper.Error `json:"error"`
}

func WriteErrorResponse(w http.ResponseWriter, status int, err *errorWrapper.Error) {
	SetContentTypeApplicationJson(w)

	if err == nil {
		err = errorWrapper.NewError(errorWrapper.InternalError, "ошибка")
	}

	w.WriteHeader(status)

	encodeErr := json.NewEncoder(w).Encode(ErrorResponse{Error: *err})
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

		wrappedError := errorWrapper.NewError(errorWrapper.JsonParsingError, "не удалось распарсить тело ответа")

		WriteErrorResponse(w, http.StatusInternalServerError, wrappedError)
		return
	}
}
