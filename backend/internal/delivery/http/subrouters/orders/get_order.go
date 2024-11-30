package ordersSubrouter

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *OrdersSubrouter) GetOrder(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JwtAuthMiddleware, err.Error(),
		))
		return
	}

	userUid, err := uuid.FromString(chi.URLParam(r, "user_uid"))

	if err != nil || userInfo.UserUid != userUid {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.OrderGetError, "не удалось получить пользователя",
		))
		return
	}

	orderUid, err := uuid.FromString(chi.URLParam(r, "order_uid"))

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.OrderGetError, "не удалось получить заказ",
		))
		return
	}

	order, isFound, err := h.OrdersService.GetOrder(ctx, entity.QueryFilters{
		OrderUid: orderUid,
		UserUid:  userUid,
	})

	if !isFound {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.OrderGetError, "не удалось найти заказ",
		))
		return
	}

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.OrderGetError, "не удалось получить заказ",
		))
		return
	}

	httpUtils.WriteResponseJson(w, order)
}
