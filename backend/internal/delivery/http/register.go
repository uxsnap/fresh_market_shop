package deliveryHttp

import (
	"context"
	"log"
	"net/http"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := EncodeRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := context.Background()

	uid, err := h.authService.Register(ctx, req.Email, req.Password)
	if err != nil {
		log.Printf("failed to register user: %v", err)
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	WriteResponseJson(w, RegisterResponse{
		Uid: uid.String(),
	})
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Uid string `json:"uid"`
}
