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

func (h *OrdersSubrouter) GetHistory(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userUid, err := uuid.FromString(chi.URLParam(r, "user_uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось найти юзера",
		))
		return
	}

	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if userUid != uuid.UUID(userInfo.UserUid) {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	qFilters, err := entity.NewQueryFiltersParser().
		WithAllowed(
			entity.QueryFieldPage,
			entity.QueryFieldLimit,
		).
		ParseQuery(r.URL.Query())

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	orders, err := h.OrdersService.GetOrderHistory(ctx, userUid, qFilters)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.OrderHistoryError, "не удалось получить историю заказов"))
		return
	}

	ows := make([]httpEntity.OrderWithProducts, len(orders))

	for ind, v := range orders {
		ows[ind] = httpEntity.OrderWithProductsFromEntity(v)
	}

	httpUtils.WriteResponseJson(w, ows)
}
