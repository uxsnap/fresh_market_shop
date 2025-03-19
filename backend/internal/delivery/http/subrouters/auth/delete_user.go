package authSubrouter

import (
	"context"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *AuthSubrouter) DeleteAuthUser(w http.ResponseWriter, r *http.Request) {
	var req DeleteUserRequest
	if err := httpUtils.DecodeJsonRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenCookie := httpUtils.GetBearerToken(r)

	if tokenCookie == "" {
		log.Printf("failed to get access token from request")
		httpUtils.WriteErrorResponse(w, http.StatusUnauthorized, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось получить access_token",
		))
		return
	}

	ctx := context.Background()

	// проверить есть ли оплаченные недставленные заказы
	orders, err := h.OrdersService.GetOrderHistory(ctx, req.Uid, entity.QueryFilters{})
	if err != nil {
		log.Printf("failed to check active orders: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusUnauthorized, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось проверить наличие активных заказов",
		))
		return
	}

	for _, order := range orders {
		if order.Status == entity.OrderStatusPaid {
			log.Printf("user have paid, not deliveried order : %v", err)
			httpUtils.WriteErrorResponse(w, http.StatusUnauthorized, errorWrapper.NewError(
				errorWrapper.JsonParsingError, "есть оплаченные, недоставленные заказы. Нельзя удалить аккаунт пока заказы не будут доставлены",
			))
			return
		}
	}

	if err := h.AuthService.DeleteAuthUser(ctx, tokenCookie, req.Uid); err != nil {
		log.Printf("failed to delete user %s: %v", req.Uid, err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось удалить пользователя",
		))
		return
	}

	if err := h.UsersService.DeleteUser(ctx, req.Uid); err != nil {
		log.Printf("failed to delete user %s in gw: %v", req.Uid, err)
	}

	w.WriteHeader(http.StatusOK)
}

type DeleteUserRequest struct {
	Uid uuid.UUID `json:"uid"`
}
