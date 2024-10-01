package authSubrouter

import (
	"context"
	"log"
	"net/http"

	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *AuthSubrouter) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := httpUtils.EncodeRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := context.Background()

	uid, err := h.AuthService.Register(ctx, req.Email, req.Password)
	if err != nil {
		log.Printf("failed to register user: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteResponseJson(w, RegisterResponse{
		Uid: uid.String(),
	})
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Uid string `json:"uid"`
}
