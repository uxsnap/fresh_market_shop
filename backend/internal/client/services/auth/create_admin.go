package clientAuthService

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/pkg/auth_v1"
	"google.golang.org/grpc/metadata"
)

func (c *AuthClient) CreateAdmin(ctx context.Context, email string, password string, token string) (uuid.UUID, error) {
	log.Printf("authClient.CreateAdmin: email %s, token %s", email, token)

	ctx = metadata.AppendToOutgoingContext(ctx, accessJwtKey, token)

	resp, err := c.client.CreateAdmin(ctx, &auth_v1.AdminCreateRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Printf("failed to create admin user: %v", err)
		return uuid.UUID{}, err
	}

	userUid, err := uuid.FromString(resp.Uid)

	if err != nil {
		log.Printf("create admin returns uncorrect uid: (%s), error: %v", resp.Uid, err)
		return uuid.UUID{}, err
	}

	return userUid, nil
}
