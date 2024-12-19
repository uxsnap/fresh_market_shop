package authSubrouter

import (
	"context"
	"log"
	"net/http"

	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *AuthSubrouter) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var req CreateAdminRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := context.Background()

	fullName, nameError := GetUserName(req.Name)
	if nameError != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nameError)
		return
	}

	token := httpUtils.GetBearerToken(r)

	uid, err := h.AuthService.CreateAdmin(ctx, req.Email, req.Password, token)
	if err != nil {
		log.Printf("failed to create admin user: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
		return
	}

	if _, err := h.UsersService.CreateUser(ctx, entity.User{
		Uid:       uid,
		Email:     req.Email,
		FirstName: fullName.firstName,
		LastName:  fullName.lastName,
	}); err != nil {
		log.Printf("failed to create admin user %s in gw: %v", uid, err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.UserNameError, err.Error(),
		))
		return
	}

	httpUtils.WriteResponseJson(w, RegisterResponse{
		Uid: uid.String(),
	})
}

type CreateAdminRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAdminResponse struct {
	Uid string `json:"uid"`
}
