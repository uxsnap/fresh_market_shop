package usersSubrouter

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *UsersSubrouter) updateDeliveryAddress(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var address httpEntity.DeliveryAddress
	if err := httpUtils.DecodeJsonRequest(r, &address); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := h.UsersService.UpdateDeliveryAddress(ctx, httpEntity.DeliveryAddressToEntity(address)); err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
