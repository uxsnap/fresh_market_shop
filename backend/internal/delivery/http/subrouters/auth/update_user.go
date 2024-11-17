package authSubrouter

import (
	"context"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/consts"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *AuthSubrouter) UpdateAuthUser(w http.ResponseWriter, r *http.Request) {
	var req UpdateUserRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessCookie := httpUtils.GetBearerToken(r)
	if accessCookie == "" {
		log.Printf("failed to get access token from request")
		httpUtils.WriteErrorResponse(w, http.StatusUnauthorized, errorWrapper.NewError("not authorized", "отсутствует токен"))
		return
	}

	ctx := context.Background()

	if err := h.UsersService.UpdateUser(ctx, entity.User{
		Uid:   req.Uid,
		Email: req.Email,
	}); err != nil {
		log.Printf("failed to update user %s in gw: %v", req.Uid, err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.UserNameError, err.Error(),
		))
		return
	}

	accessJwt, refreshJwt, err := h.AuthService.UpdateAuthUser(ctx, accessCookie, req.Uid, req.Email, req.Password)
	if err != nil {
		log.Printf("failed to update user %s: %v", req.Uid, err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
		return
	}

	http.SetCookie(w, httpUtils.NewCookie(consts.ACCESS_JWT_COOKIE_NAME, accessJwt))
	http.SetCookie(w, httpUtils.NewCookie(consts.REFRESH_JWT_COOKIE_NAME, refreshJwt))

	w.WriteHeader(http.StatusOK)
}

type UpdateUserRequest struct {
	Uid      uuid.UUID `json:"uid"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
