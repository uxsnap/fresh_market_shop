package usersSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *UsersSubrouter) addDeliveryAddress(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var address httpEntity.DeliveryAddress
	if err := httpUtils.DecodeJsonRequest(r, &address); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	uid, err := h.UsersService.AddDeliveryAddress(ctx, httpEntity.DeliveryAddressToEntity(address))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.UUID{
		Uid: uid,
	})
}
