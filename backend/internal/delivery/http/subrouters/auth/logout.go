package authSubrouter

import (
	"context"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *AuthSubrouter) Logout(w http.ResponseWriter, r *http.Request) {
	var req LogoutRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bearer := httpUtils.GetBearerToken(r)

	if bearer == "" {
		log.Printf("failed to get access token from request")
		httpUtils.WriteErrorResponse(w, http.StatusUnauthorized, nil)
		return
	}

	ctx := context.Background()

	if err := h.AuthService.Logout(ctx, bearer, req.Uid); err != nil {
		log.Printf("failed to logout user: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type LogoutRequest struct {
	Uid uuid.UUID `json:"uid"`
}
