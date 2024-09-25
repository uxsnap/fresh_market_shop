package deliveryHttp

import (
	"context"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := EncodeRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := context.Background()

	accessJwt, refreshJwt, err := h.authService.Login(ctx, req.Email, req.Password)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	http.SetCookie(w, NewCookie(accessJwtCookieName, accessJwt))
	http.SetCookie(w, NewCookie(refreshJwtCookieName, refreshJwt))
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
