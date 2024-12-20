package clientAuthService

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *AuthClient) GetAdmins(ctx context.Context, token string) ([]entity.AuthUser, error) {
	log.Printf("authClient.CreateAdmin: token %s", token)

	ctx = metadata.AppendToOutgoingContext(ctx, accessJwtKey, token)

	resp, err := c.client.GetAdmins(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("failed to get admin users: %v", err)
		return []entity.AuthUser{}, err
	}

	result := make([]entity.AuthUser, len(resp.Users))
	for ind, user := range resp.Users {
		result[ind] = entity.AuthUser{
			Uid:   uuid.FromStringOrNil(user.GetUid()),
			Email: user.GetEmail(),
		}
	}

	return result, nil
}
