package usersSubrouter

import (
	"context"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/consts"
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

	accessCookie := httpUtils.GetBearerToken(r)

	if err := h.UsersService.UpdateUser(ctx, httpEntity.UserToEntity(user)); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	accessJwt, refreshJwt, err := h.AuthService.UpdateAuthUser(ctx, accessCookie, user.Uid, user.Email, "")
	if err != nil {
		log.Printf("failed to update user %s: %v", user.Uid, err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	http.SetCookie(w, httpUtils.NewCookie(consts.ACCESS_JWT_COOKIE_NAME, accessJwt))
	http.SetCookie(w, httpUtils.NewCookie(consts.REFRESH_JWT_COOKIE_NAME, refreshJwt))

	w.WriteHeader(http.StatusOK)
}
