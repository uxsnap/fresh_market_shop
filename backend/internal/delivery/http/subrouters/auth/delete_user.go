package authSubrouter

import (
	"context"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *AuthSubrouter) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var req DeleteUserRequest
	if err := httpUtils.EncodeRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessCookie, err := r.Cookie(accessJwtCookieName)
	if err != nil {
		log.Printf("failed to get access token from request: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	ctx := context.Background()

	if err := h.AuthService.DeleteUser(ctx, accessCookie.Value, req.Uid); err != nil {
		log.Printf("failed to delete user: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type DeleteUserRequest struct {
	Uid uuid.UUID `json:"uid"`
}
