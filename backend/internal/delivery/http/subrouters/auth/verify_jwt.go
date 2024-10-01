package authSubrouter

import (
	"context"
	"net/http"

	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

// TODO: убрать этот метод отсюда, пока чисто для тестов
func (h *AuthSubrouter) VerifyJwt(w http.ResponseWriter, r *http.Request) {
	var req VerifyJwtRequest
	if err := httpUtils.EncodeRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	resp := VerifyJwtResponse{}
	if err := h.AuthService.VerifyJwt(ctx, req.Jwt); err != nil {
		resp.Valid = false
		resp.Message = err.Error()
	} else {
		resp.Valid = true
	}

	httpUtils.WriteResponseJson(w, resp)
}

type VerifyJwtRequest struct {
	Jwt string `json:"jwt"`
}

type VerifyJwtResponse struct {
	Valid   bool   `json:"isValid"`
	Message string `json:"message"`
}
