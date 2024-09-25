package clientAuthService

import (
	"context"
	"log"

	"github.com/uxsnap/fresh_market_shop/backend/pkg/auth_v1"
)

func (c *AuthClient) VerifyJwt(ctx context.Context, token string) error {
	log.Printf("authClient.VerifyJwt: token %s", token)

	if _, err := c.client.Verify(ctx, &auth_v1.VerifyRequest{
		Jwt: token,
	}); err != nil {
		log.Printf("failed to verify token %s", token)
		return err
	}
	return nil
}
