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

func (h *UsersSubrouter) getDeliveryAddress(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	uid, err := uuid.FromString(chi.URLParam(r, "address_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	//TODO: проверять что вызывает тот юзер чей юзер uid записан в адресе

	address, isFound, err := h.UsersService.GetDeliveryAddress(ctx, uid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	if !isFound {
		httpUtils.WriteErrorResponse(w, http.StatusNotFound, errors.New("address not found"))
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.DeliveryAddressFromEntity(address))
}
