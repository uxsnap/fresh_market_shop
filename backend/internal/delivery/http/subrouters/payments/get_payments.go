package paymentsSubrouter

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *PaymentsSubrouter) GetPayments(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	paramUidName, uid, err := getUidParamFromPaymentsRequest(r)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, errorWrapper.NewError("bad request", err.Error()))
		return
	}

	var payments []entity.Payment

	switch paramUidName {
	case entity.QueryFieldUserUid:
		payments, err = h.PaymentsService.GetUserPayments(ctx, uid)
	case entity.QueryFieldCardUid:
		payments, err = h.PaymentsService.GetCardPayments(ctx, uid)
	}

	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(errorWrapper.InternalError, err.Error()))
		return
	}

	resp := make([]httpEntity.Payment, len(payments))
	for i := 0; i < len(payments); i++ {
		resp[i] = httpEntity.PaymentFromEntity(payments[i])
	}

	httpUtils.WriteResponseJson(w, resp)
}

func getUidParamFromPaymentsRequest(r *http.Request) (string, uuid.UUID, error) {
	qFilters, err := entity.NewQueryFiltersParser().
		WithAllowed(
			entity.QueryFieldUserUid,
			entity.QueryFieldCardUid,
		).
		ParseQuery(r.URL.Query())
	if err != nil {
		return "", uuid.UUID{}, err
	}

	paramName := ""
	paramValue := uuid.UUID{}

	params := map[string]uuid.UUID{
		entity.QueryFieldUserUid: qFilters.UserUid,
		entity.QueryFieldCardUid: qFilters.CardUid,
	}

	for name, v := range params {
		if uuid.Equal(uuid.UUID{}, v) {
			continue
		}
		if !uuid.Equal(uuid.UUID{}, paramValue) {
			return "", uuid.UUID{}, errors.New("only one url parameter can be sent")
		}
		paramName = name
		paramValue = v
	}

	if uuid.Equal(uuid.UUID{}, paramValue) {
		return "", uuid.UUID{}, errors.New("required url parameter")
	}

	return paramName, paramValue, nil
}
