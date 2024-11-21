package recommendationsSubrouter

import (
	"context"
	"net/http"
	"net/url"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *RecommendationsSubrouter) getRecommendations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	urlValues := h.handleUrlValues(ctx, r.URL.Query())

	qFilters, err := entity.NewQueryFiltersParser().ParseQuery(urlValues)

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	products, err := h.ProductsService.GetProductsWithExtra(ctx, qFilters)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError,
			errorWrapper.NewError(errorWrapper.RecommendationsError, "не удалось получить продукты рекомендаций"),
		)
		return
	}

	resp := make([]httpEntity.ProductWithExtra, 0, len(products))
	for _, product := range products {
		resp = append(resp, httpEntity.ProductWithExtra{
			Product: httpEntity.ProductFromEntity(product.Product),
			Photos:  httpEntity.ProductPhotosFromEntity(product.Photos),
		})
	}

	httpUtils.WriteResponseJson(w, resp)
}

func (h *RecommendationsSubrouter) handleUrlValues(ctx context.Context, urlValues url.Values) url.Values {
	userInfo, err := httpEntity.AuthUserInfoFromContext(ctx)
	if err == nil {
		urlValues.Set(entity.QueryFieldUserUid, userInfo.UserUid.String())
	}

	urlValues.Set(entity.QueryFieldWithRandom, "true")

	categoryUids, err := h.ProductsService.GetCategoriesByUserOrders(
		ctx, userInfo.UserUid,
	)
	if err == nil && len(categoryUids) != 0 {
		for _, uid := range categoryUids {
			urlValues.Add(entity.QueryFieldCategoryUids, uid.String())
		}
	}

	return urlValues
}
