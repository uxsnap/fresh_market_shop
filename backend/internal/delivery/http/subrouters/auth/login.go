package authSubrouter

import (
	"context"
	"net/http"

	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *AuthSubrouter) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := context.Background()

	accessJwt, refreshJwt, err := h.AuthService.Login(ctx, req.Email, req.Password)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	http.SetCookie(w, httpUtils.NewCookie(accessJwtCookieName, accessJwt))
	http.SetCookie(w, httpUtils.NewCookie(refreshJwtCookieName, refreshJwt))
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
