package clientAuthService

import (
	"context"
	"log"

	"github.com/uxsnap/fresh_market_shop/backend/pkg/auth_v1"
)

func (c *AuthClient) Refresh(ctx context.Context, refreshToken string) (accessJwt string, refreshJwt string, err error) {
	log.Printf("authClient.Refresh: token %s", refreshToken)

	var resp *auth_v1.JwtResponse
	resp, err = c.client.Refresh(ctx, &auth_v1.RefreshRequest{
		RefreshJwt: refreshToken,
	})
	if err != nil {
		log.Printf("failed to refresh (token %s): %v", refreshToken, err)
		return
	}

	accessJwt = resp.GetAccessJwt()
	refreshJwt = resp.GetRefreshJwt()
	return
}
