package authSubrouter

import (
	"context"
	"log"
	"net/http"

	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (h *AuthSubrouter) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := context.Background()

	uid, err := h.AuthService.Register(ctx, req.Email, req.Password)
	if err != nil {
		log.Printf("failed to register user: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	if _, err := h.UsersService.CreateUser(ctx, entity.User{
		Uid:   uid,
		Email: req.Email,
	}); err != nil {
		log.Printf("failed to create user %s in gw: %v", uid, err)
	}

	httpUtils.WriteResponseJson(w, RegisterResponse{
		Uid: uid.String(),
	})
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Uid string `json:"uid"`
}
