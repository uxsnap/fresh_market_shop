package usersSubrouter

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *UsersSubrouter) getUserDeliveryAddresses(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	uid, err := uuid.FromString(chi.URLParam(r, "user_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	addresses, err := h.UsersService.GetUserDeliveryAddresses(ctx, uid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp := make([]httpEntity.DeliveryAddress, 0, len(addresses))
	for _, address := range addresses {
		resp = append(resp, httpEntity.DeliveryAddressFromEntity(address))
	}

	httpUtils.WriteResponseJson(w, resp)
}
