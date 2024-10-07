package clientAuthService

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/pkg/auth_v1"
	"google.golang.org/grpc/metadata"
)

func (c *AuthClient) DeleteAuthUser(ctx context.Context, accessJwt string, uid uuid.UUID) error {
	log.Printf("authClient.DeleteUser: uid '%s'", uid)

	ctx = metadata.AppendToOutgoingContext(ctx, accessJwtKey, accessJwt)

	if _, err := c.client.DeleteUser(ctx, &auth_v1.DeleteUserRequest{
		Uid: uid.String(),
	}); err != nil {
		log.Printf("failed to delete user with uid %s", uid)
		return err
	}

	return nil
}
