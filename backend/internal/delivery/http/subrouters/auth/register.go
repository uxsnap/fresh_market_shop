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

	nameSlice := strings.Split(req.Name, " ")
	firstName := nameSlice[0]
	lastName := ""
	if len(nameSlice) == 2 {
		lastName = nameSlice[1]
	}

	if len(firstName) < 2 {
		log.Printf("failed to validate register user")
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.UserNameError, "длина имени пользователя должна быть больше 1",
		))
		return
	}

	uid, err := h.AuthService.Register(ctx, req.Email, req.Password)
	if err != nil {
		log.Printf("failed to register user: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
		return
	}

	if _, err := h.UsersService.CreateUser(ctx, entity.User{
		Uid:       uid,
		Email:     req.Email,
		FirstName: firstName,
		LastName:  lastName,
	}); err != nil {
		log.Printf("failed to create user %s in gw: %v", uid, err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.UserNameError, err.Error(),
		))
		return
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
