package deliveryHttp

import (
	"context"
	"net/http"
)

// TODO: убрать этот метод отсюда, пока чисто для тестов
func (h *Handler) VerifyJwt(w http.ResponseWriter, r *http.Request) {
	var req VerifyJwtRequest
	if err := EncodeRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	resp := VerifyJwtResponse{}
	if err := h.authService.VerifyJwt(ctx, req.Jwt); err != nil {
		resp.Valid = false
		resp.Message = err.Error()
	} else {
		resp.Valid = true
	}

	WriteResponseJson(w, resp)
}

type VerifyJwtRequest struct {
	Jwt string `json:"jwt"`
}

type VerifyJwtResponse struct {
	Valid   bool   `json:"isValid"`
	Message string `json:"message"`
}
