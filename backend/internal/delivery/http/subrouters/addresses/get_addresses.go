package addressesSubrouter

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
	"github.com/uxsnap/fresh_market_shop/backend/internal/consts"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *AddressesSubrouter) GetAddresses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	urlValues := h.handleUrlValues(r)

	qFilters, err := entity.NewQueryFiltersParser().
		WithRequired(entity.QueryFieldName).
		ParseQuery(urlValues)

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "ошибка парсинга параметров",
		))
		return
	}

	addresses, err := h.AddressesService.GetAddresses(ctx, qFilters)

	if err != nil {
		log.Printf("failed to get addresses: %v", err)
		httpUtils.WriteErrorResponse(
			w, http.StatusInternalServerError,
			errorWrapper.NewError(errorWrapper.AddressesError, "не удалось получить адреса"))
		return
	}

	resp := make([]httpEntity.Address, 0, len(addresses))
	for _, address := range addresses {
		resp = append(resp, httpEntity.AddressFromEntity(address))
	}

	httpUtils.WriteResponseJson(w, resp)
}

func (h *AddressesSubrouter) handleUrlValues(r *http.Request) url.Values {
	urlValues := r.URL.Query()

	cityUid := chi.URLParam(r, "city_uid")

	urlValues.Set(entity.QueryFieldCityUid, cityUid)
	urlValues.Set(entity.QueryFieldLimit, fmt.Sprint(consts.DEFAULT_LIMIT))

	name := urlValues.Get(entity.QueryFieldName)

	splittedName := strings.Split(name, " ")

	if len(splittedName) == 0 {
		return urlValues
	}

	urlValues.Set(entity.QueryFieldName, splittedName[0])

	if len(splittedName) > 1 && splittedName[0] != splittedName[len(splittedName)-1] {
		urlValues.Set(entity.QueryFieldHouseNumber, splittedName[len(splittedName)-1])
	}

	return urlValues
}
