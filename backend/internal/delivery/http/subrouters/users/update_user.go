package usersSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *UsersSubrouter) updateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var user httpEntity.User
	if err := httpUtils.DecodeJsonRequest(r, &user); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if err := h.UsersService.UpdateUser(ctx, httpEntity.UserToEntity(user)); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}
