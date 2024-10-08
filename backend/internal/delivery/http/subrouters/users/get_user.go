package usersSubrouter

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *UsersSubrouter) getUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	uid, err := uuid.FromString(chi.URLParam(r, "user_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	user, isFound, err := h.UsersService.GetUser(ctx, uid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	if !isFound {
		httpUtils.WriteErrorResponse(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UserFromEntity(user))
}
