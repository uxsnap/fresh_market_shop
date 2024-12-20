package authSubrouter

import (
	"context"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (h *AuthSubrouter) GetAdmins(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	users, err := h.AuthService.GetAdmins(ctx)
	if err != nil {
		log.Printf("failed to get admin users: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, errorWrapper.NewError(
			errorWrapper.JsonParsingError, "не удалось получить пользователей",
		))
		return
	}

	res := make([]AdminResponse, len(users))
	for ind, user := range users {
		res[ind] = AdminResponse{
			Uid:   user.Uid,
			Email: user.Email,
		}
	}

	httpUtils.WriteResponseJson(w, GetAdminsResponse{
		Admins: res,
	})
}

type AdminResponse struct {
	Uid   uuid.UUID `json:"uid"`
	Email string    `json:"email"`
}

type GetAdminsResponse struct {
	Admins []AdminResponse `json:"admins"`
}
