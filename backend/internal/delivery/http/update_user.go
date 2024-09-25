package deliveryHttp

import (
	"context"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// TODO: rename to updateUserSSO or updateAuthUser
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req UpdateUserRequest
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

	accessJwt, refreshJwt, err := h.authService.UpdateUser(ctx, accessCookie.Value, req.Uid, req.Email, req.Password)
	if err != nil {
		log.Printf("failed to update user: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	http.SetCookie(w, NewCookie(accessJwtCookieName, accessJwt))
	http.SetCookie(w, NewCookie(refreshJwtCookieName, refreshJwt))
}

type UpdateUserRequest struct {
	Uid      uuid.UUID `json:"uid"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
