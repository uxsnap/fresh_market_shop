package authSubrouter

import (
	"context"
	"log"
	"net/http"

	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *AuthSubrouter) RefreshJwt(w http.ResponseWriter, r *http.Request) {
	refreshCookie, err := r.Cookie(refreshJwtCookieName)

	if err != nil {
		log.Printf("failed to get refresh token from request: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	ctx := context.Background()

	accessJwt, refreshJwt, err := h.AuthService.Refresh(ctx, refreshCookie.Value)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	http.SetCookie(w, httpUtils.NewCookie(accessJwtCookieName, accessJwt))
	http.SetCookie(w, httpUtils.NewCookie(refreshJwtCookieName, refreshJwt))

	w.WriteHeader(http.StatusOK)
}
