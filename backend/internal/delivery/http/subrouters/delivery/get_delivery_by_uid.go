package deliverySubrouter

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *DeliverySubrouter) GetDeliveryByUid(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	uid, err := uuid.FromString(chi.URLParam(r, "uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	delivery, isFound, err := h.DeliveryService.GetDeliveryByUid(ctx, uid)
	if err != nil {
		log.Printf("failed to get delivery: %v", err)
		httpUtils.WriteErrorResponse(
			w, http.StatusInternalServerError,
			errorWrapper.NewError(errorWrapper.InternalError, "не удалось получить информацию о доставке"))
		return
	}
	if !isFound {
		log.Printf("delivery not found")
		httpUtils.WriteErrorResponse(w, http.StatusNotFound, errorWrapper.NewError(errorWrapper.InternalError, "информация о доставке не найдена"))
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.DeliveryFromEntity(delivery))
}
