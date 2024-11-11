package deliverySubrouter

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *DeliverySubrouter) GetDeliveryByOrderUid(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	orderUid, err := uuid.FromString(chi.URLParam(r, "order_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	delivery, isFound, err := h.DeliveryService.GetDeliveryByOrderUid(ctx, orderUid)
	if err != nil {
		log.Printf("failed to get delivery by order uid: %v", err)
		httpUtils.WriteErrorResponse(
			w, http.StatusInternalServerError,
			errorWrapper.NewError(errorWrapper.InternalError, fmt.Sprintf("не удалось получить информацию о доставке по заказу: %s", err.Error())))
		return
	}
	if !isFound {
		log.Printf("delivery not found")
		httpUtils.WriteErrorResponse(w, http.StatusNotFound, errorWrapper.NewError(errorWrapper.InternalError, "информация о доставке по заказу не найдена"))
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.DeliveryFromEntity(delivery))
}
