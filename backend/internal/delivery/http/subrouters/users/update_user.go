package usersSubrouter

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *UsersSubrouter) updateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	var user httpEntity.User
	if err := httpUtils.DecodeJsonRequest(r, &user); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if !uuid.Equal(userInfo.UserUid, user.Uid) {
		// TODO: check role
		httpUtils.WriteErrorResponse(w, http.StatusForbidden, errorWrapper.NewError(
			errorWrapper.UserInfoError, "неправильный uid пользователя",
		))
		return
	}

	if err := h.UsersService.UpdateUser(ctx, httpEntity.UserToEntity(user)); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
