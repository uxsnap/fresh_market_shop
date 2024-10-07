package authSubrouter

import (
	"context"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

// TODO: rename to getUserSSO or getAuthUser
func (h *AuthSubrouter) GetAuthUser(w http.ResponseWriter, r *http.Request) {
	var req GetUserRequest
	if err := httpUtils.EncodeRequest(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessCookie, err := r.Cookie(accessJwtCookieName)
	if err != nil {
		log.Printf("failed to get access token from request: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	ctx := context.Background()

	user, err := h.AuthService.GetAuthUser(ctx, accessCookie.Value, req.Uid, req.Email)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteResponseJson(w, GetUserResponse{
		Uid:         user.Uid,
		Email:       user.Email,
		Role:        string(user.Role),
		Permissions: user.PermissionsStrings(),
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	})
}

type GetUserRequest struct {
	Uid   uuid.UUID `json:"uid"`
	Email string    `json:"email"`
}

type GetUserResponse struct {
	Uid         uuid.UUID `json:"uid"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	Permissions []string  `json:"permissions"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
