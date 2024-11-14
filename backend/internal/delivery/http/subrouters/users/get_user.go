package usersSubrouter

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *UsersSubrouter) getUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	uid, err := uuid.FromString(chi.URLParam(r, "user_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	user, isFound, userErr := h.UsersService.GetUser(ctx, uid)

	if userErr != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if !isFound {
		httpUtils.WriteErrorResponse(w, http.StatusNotFound, errorWrapper.NewError(
			errorWrapper.UserNotFoundError, "пользователь не найден",
		))
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UserFromEntity(user))
}
