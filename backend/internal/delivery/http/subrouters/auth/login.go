package authSubrouter

import (
	"context"
	"net/http"

	"github.com/uxsnap/fresh_market_shop/backend/internal/consts"
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
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	http.SetCookie(w, httpUtils.NewCookie(consts.ACCESS_JWT_COOKIE_NAME, accessJwt))
	http.SetCookie(w, httpUtils.NewCookie(consts.REFRESH_JWT_COOKIE_NAME, refreshJwt))

	w.WriteHeader(http.StatusOK)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
