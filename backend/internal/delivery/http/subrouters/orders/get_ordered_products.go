package ordersSubrouter

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/uxsnap/fresh_market_shop/backend/internal/consts"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *OrdersSubrouter) handelOrderedProductsURLParams(r *http.Request) (entity.QueryFilters, error) {
	userInfo, err := httpEntity.AuthUserInfoFromContext(r.Context())

	if err != nil {
		return entity.QueryFilters{}, err
	}

	// Не из объекта реквеста, чтобы нельзя было прокинуть необязательные параметры
	urlValues := url.Values{}

	fmt.Println(userInfo)

	urlValues.Set(entity.QueryFieldUserUidForOrder, userInfo.UserUid.String())
	urlValues.Set(entity.QueryFieldLimit, fmt.Sprint(consts.DEFAULT_LIMIT))

	qFilters, err := entity.NewQueryFiltersParser().
		ParseQuery(urlValues)

	return qFilters, err
}

func (h *OrdersSubrouter) GetOrderedProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	qFilters, err := h.handelOrderedProductsURLParams(r)

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "ошибка парсинга параметров",
		))
		return
	}

	products, err := h.OrdersService.GetOrderedProducts(ctx, qFilters)

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp := make([]httpEntity.ProductWithExtra, 0, len(products))

	for _, product := range products {
		resp = append(resp, httpEntity.ProductWithExtra{
			Product: httpEntity.ProductFromEntity(product.Product),
			Count:   product.StockQuantity,
			Photos:  httpEntity.ProductPhotosFromEntity(product.Photos),
		})
	}

	httpUtils.WriteResponseJson(w, resp)
}
