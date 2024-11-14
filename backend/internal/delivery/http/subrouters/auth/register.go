package authSubrouter

import (
	"context"
	"log"
	"net/http"
	"strings"

	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
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

	nameSlice := strings.Split(req.Name, " ")

	if len(nameSlice) == 0 || len(nameSlice[0]) < 2 {
		log.Printf("failed to validate register user")
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.UserNameError, "длина имени пользователя должна быть больше 1",
		))
		return
	}

	if _, err := h.UsersService.CreateUser(ctx, entity.User{
		Uid:       uid,
		Email:     req.Email,
		FirstName: nameSlice[0],
		LastName:  nameSlice[1],
	}); err != nil {
		log.Printf("failed to create user %s in gw: %v", uid, err)
	}

	httpUtils.WriteResponseJson(w, RegisterResponse{
		Uid: uid.String(),
	})
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Uid string `json:"uid"`
}
