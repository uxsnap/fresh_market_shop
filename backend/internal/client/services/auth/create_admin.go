package clientAuthService

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/pkg/auth_v1"
)

func (c *AuthClient) CreateAdmin(ctx context.Context, email string, password string, token string) (uuid.UUID, error) {
	log.Printf("authClient.CreateAdmin: email %s", email)

	resp, err := c.client.CreateAdmin(ctx, &auth_v1.AdminCreateRequest{
		Email:    email,
		Password: password,
		Jwt:      token,
	})
	if err != nil {
		log.Printf("failed to register user: %v", err)
		return uuid.UUID{}, err
	}

	userUid, err := uuid.FromString(resp.Uid)

	if err != nil {
		log.Printf("register returns uncorrect uid: (%s), error: %v", resp.Uid, err)
		return uuid.UUID{}, err
	}

	return userUid, nil
}
