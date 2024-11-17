package addressesSubrouter

import (
	"log"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *AddressesSubrouter) GetCities(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cities, err := h.AddressesService.GetCities(ctx)

	if err != nil {
		log.Printf("failed to get cities: %v", err)
		httpUtils.WriteErrorResponse(
			w, http.StatusInternalServerError,
			errorWrapper.NewError(errorWrapper.AddressesError, "не удалось получить города"))
		return
	}

	resp := make([]httpEntity.City, 0, len(cities))
	for _, city := range cities {
		resp = append(resp, httpEntity.CityFromEntity(city))
	}

	httpUtils.WriteResponseJson(w, resp)
}
