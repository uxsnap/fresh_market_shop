package clientAuthService

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

func (c *AuthClient) CreateAdmin(ctx context.Context, email string, password string, token string) (uuid.UUID, error) {
	// log.Printf("authClient.CreateAdmin: email %s", email)

	// resp, err := c.client.CreateAdmin(ctx, &auth_v1.CreateAdmin{
	// 	Email:    email,
	// 	Password: password,
	// })
	// if err != nil {
	// 	log.Printf("failed to register user: %v", err)
	// 	return uuid.UUID{}, err
	// }

	// userUid, err := uuid.FromString(resp.Uid)

	// if err != nil {
	// 	log.Printf("register returns uncorrect uid: (%s), error: %v", resp.Uid, err)
	// 	return uuid.UUID{}, err
	// }

	return uuid.UUID{}, nil
}
