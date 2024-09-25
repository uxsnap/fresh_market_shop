package clientAuthService

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/pkg/auth_v1"
	"google.golang.org/grpc/metadata"
)

func (c *AuthClient) Logout(ctx context.Context, accessJwt string, uid uuid.UUID) error {
	log.Printf("authClient.Logout: uid %s", uid)

	ctx = metadata.AppendToOutgoingContext(ctx, accessJwtKey, accessJwt)

	if _, err := c.client.Logout(ctx, &auth_v1.LogoutRequest{
		Uid: uid.String(),
	}); err != nil {
		log.Printf("failed to logout user with uid %s", uid)
		return err
	}

	return nil
}
