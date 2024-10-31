package clientAuthService

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/pkg/auth_v1"
	"google.golang.org/grpc/metadata"
)

func (c *AuthClient) UpdateAuthUser(
	ctx context.Context,
	accessToken string,
	uid uuid.UUID,
	email string,
	password string,
) (accessJwt string, refreshJwt string, err error) {
	log.Printf("authClient.UpdateUser: uid %s", uid)

	ctx = metadata.AppendToOutgoingContext(ctx, accessJwtKey, accessToken)

	var resp *auth_v1.JwtResponse
	resp, err = c.client.UpdateUser(ctx, &auth_v1.UpdateUserRequest{
		Uid:      uid.String(),
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Printf("failed to update user %s: %v", uid, err)
		return
	}

	accessJwt = resp.GetAccessJwt()
	refreshJwt = resp.GetRefreshJwt()
	return
}
