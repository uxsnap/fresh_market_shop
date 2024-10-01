package authSubrouter

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h *AuthSubrouter) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	ctx := context.Background()

	if err := h.AuthService.VerifyEmail(ctx, token); err != nil {
		log.Printf("failed to verify email: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
