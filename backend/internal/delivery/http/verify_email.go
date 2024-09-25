package deliveryHttp

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	ctx := context.Background()

	if err := h.authService.VerifyEmail(ctx, token); err != nil {
		log.Printf("failed to verify email: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
