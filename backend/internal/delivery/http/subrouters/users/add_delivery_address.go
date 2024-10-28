package usersSubrouter

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *UsersSubrouter) addDeliveryAddress(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	var address httpEntity.DeliveryAddress
	if err := httpUtils.DecodeJsonRequest(r, &address); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	if !uuid.Equal(userInfo.UserUid, address.UserUid) {
		// TODO: check role
		httpUtils.WriteErrorResponse(w, http.StatusForbidden, nil)
		return
	}

	uid, err := h.UsersService.AddDeliveryAddress(ctx, httpEntity.DeliveryAddressToEntity(address))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{
		Uid: uid,
	})
}
