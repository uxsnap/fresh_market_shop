package usersSubrouter

import (
	"context"
	"encoding/json"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *UsersSubrouter) createUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var user httpEntity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	uid, err := h.UsersService.CreateUser(ctx, httpEntity.UserToEntity(user))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{
		Uid: uid,
	})
}
