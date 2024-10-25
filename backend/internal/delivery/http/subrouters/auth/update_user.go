package authSubrouter

import (
	"context"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (h *AuthSubrouter) UpdateAuthUser(w http.ResponseWriter, r *http.Request) {
	var req UpdateUserRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessCookie, err := r.Cookie(accessJwtCookieName)
	if err != nil {
		log.Printf("failed to get access token from request: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusUnauthorized, nil)
		return
	}

	ctx := context.Background()

	accessJwt, refreshJwt, err := h.AuthService.UpdateAuthUser(ctx, accessCookie.Value, req.Uid, req.Email, req.Password)
	if err != nil {
		log.Printf("failed to update user %s: %v", req.Uid, err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	if err := h.UsersService.UpdateUser(ctx, entity.User{
		Uid:   req.Uid,
		Email: req.Email,
	}); err != nil {
		log.Printf("failed to update user %s in gw: %v", req.Uid, err)
	}

	w.WriteHeader(http.StatusOK)
	http.SetCookie(w, httpUtils.NewCookie(accessJwtCookieName, accessJwt))
	http.SetCookie(w, httpUtils.NewCookie(refreshJwtCookieName, refreshJwt))
}

type UpdateUserRequest struct {
	Uid      uuid.UUID `json:"uid"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
