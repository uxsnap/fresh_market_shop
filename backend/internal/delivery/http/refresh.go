package deliveryHttp

import (
	"context"
	"log"
	"net/http"
)

func (h *Handler) RefreshJwt(w http.ResponseWriter, r *http.Request) {
	refreshCookie, err := r.Cookie(refreshJwtCookieName)
	if err != nil {
		log.Printf("failed to get refresh token from request: %v", err)
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()

	accessJwt, refreshJwt, err := h.authService.Refresh(ctx, refreshCookie.Value)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	http.SetCookie(w, NewCookie(accessJwtCookieName, accessJwt))
	http.SetCookie(w, NewCookie(refreshJwtCookieName, refreshJwt))
}
