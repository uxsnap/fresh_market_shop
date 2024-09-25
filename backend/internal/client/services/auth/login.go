package clientAuthService

import (
	"context"
	"log"

	"github.com/uxsnap/fresh_market_shop/backend/pkg/auth_v1"
)

func (c *AuthClient) Login(ctx context.Context, email string, password string) (accessJwt string, refreshJwt string, err error) {
	log.Printf("authClient.Login: email %s", email)

	var resp *auth_v1.JwtResponse
	resp, err = c.client.Login(ctx, &auth_v1.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Printf("failed to login user %s: %v", email, err)
		return
	}

	accessJwt = resp.AccessJwt
	refreshJwt = resp.RefreshJwt
	return
}
