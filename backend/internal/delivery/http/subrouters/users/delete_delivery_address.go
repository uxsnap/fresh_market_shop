package usersSubrouter

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *UsersSubrouter) deleteDeliveryAddress(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	addressUid, err := uuid.FromString(chi.URLParam(r, "address_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}
	userUid, err := uuid.FromString(chi.URLParam(r, "user_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if !uuid.Equal(userInfo.UserUid, userUid) {
		// TODO: check role
		httpUtils.WriteErrorResponse(w, http.StatusForbidden, nil)
		return
	}

	if err := h.UsersService.DeleteDeliveryAddress(ctx, addressUid); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}
