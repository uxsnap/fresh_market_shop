package deliveryHttp

import (
	"context"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	var req LogoutRequest
	if err := EncodeRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessCookie, err := r.Cookie(accessJwtCookieName)
	if err != nil {
		log.Printf("failed to get access token from request: %v", err)
		WriteErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	ctx := context.Background()

	if err := h.authService.Logout(ctx, accessCookie.Value, req.Uid); err != nil {
		log.Printf("failed to logout user: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type LogoutRequest struct {
	Uid uuid.UUID `json:"uid"`
}
