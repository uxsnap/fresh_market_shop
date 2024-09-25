package deliveryHttp

import (
	"context"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var req DeleteUserRequest
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

	if err := h.authService.DeleteUser(ctx, accessCookie.Value, req.Uid); err != nil {
		log.Printf("failed to delete user: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type DeleteUserRequest struct {
	Uid uuid.UUID `json:"uid"`
}
