package usersSubrouter

import (
	"context"
	"encoding/json"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *UsersSubrouter) updateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var user httpEntity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := h.UsersService.UpdateUser(ctx, httpEntity.UserToEntity(user)); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
