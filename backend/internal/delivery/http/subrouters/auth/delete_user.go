package authSubrouter

import (
	"context"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *AuthSubrouter) DeleteAuthUser(w http.ResponseWriter, r *http.Request) {
	var req DeleteUserRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessCookie, err := r.Cookie(accessJwtCookieName)
	if err != nil {
		log.Printf("failed to get access token from request: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusUnauthorized, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось получить access_token",
		))
		return
	}

	ctx := context.Background()

	if err := h.AuthService.DeleteAuthUser(ctx, accessCookie.Value, req.Uid); err != nil {
		log.Printf("failed to delete user %s: %v", req.Uid, err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось удалить пользователя",
		))
		return
	}

	if err := h.UsersService.DeleteUser(ctx, req.Uid); err != nil {
		log.Printf("failed to delete user %s in gw: %v", req.Uid, err)
	}

	w.WriteHeader(http.StatusOK)
}

type DeleteUserRequest struct {
	Uid uuid.UUID `json:"uid"`
}
