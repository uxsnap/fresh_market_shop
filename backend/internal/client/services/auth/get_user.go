package clientAuthService

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/pkg/auth_v1"
	"google.golang.org/grpc/metadata"
)

func (c *AuthClient) GetUser(ctx context.Context, accessJwt string, uid uuid.UUID, email string) (entity.User, error) {
	log.Printf("authClient.GetUser: uid '%s' email '%s'", uid, email)

	ctx = metadata.AppendToOutgoingContext(ctx, accessJwtKey, accessJwt)
	
	resp, err := c.client.GetUser(ctx, &auth_v1.GetUserRequest{
		Uid:   uid.String(),
		Email: email,
	})
	if err != nil {
		log.Printf("failed to get user with uid '%s' email '%s': %v", uid, email, err)
		return entity.User{}, err
	}

	return entity.User{
		Uid:         uuid.FromStringOrNil(resp.GetUid()),
		Email:       resp.GetEmail(),
		Role:        entity.UserRole(resp.GetRole().String()),
		Permissions: entity.PermissionsFromStrings(resp.GetPermissions()),
		CreatedAt:   resp.GetCreatedAt().AsTime(),
		UpdatedAt:   resp.GetUpdatedAt().AsTime(),
	}, nil
}
