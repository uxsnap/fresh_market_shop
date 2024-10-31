package usersSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

// unused
func (h *UsersSubrouter) createUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var user httpEntity.User
	if err := httpUtils.DecodeJsonRequest(r, &user); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	uid, err := h.UsersService.CreateUser(ctx, httpEntity.UserToEntity(user))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{
		Uid: uid,
	})
}
