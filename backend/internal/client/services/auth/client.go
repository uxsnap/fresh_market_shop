package clientAuthService

import (
	"context"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	config Config

	conn   *grpc.ClientConn
	client auth_v1.AuthClient
}

type Config interface {
	AuthServiceGrpcHost() string
	AuthServiceGrpcPort() string
}

func New(ctx context.Context, cfg Config) (*AuthClient, error) {
	authGrpcAddress := cfg.AuthServiceGrpcHost() + ":" + cfg.AuthServiceGrpcPort()

	cc, err := grpc.NewClient(authGrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &AuthClient{
		config: cfg,
		conn:   cc,
		client: auth_v1.NewAuthClient(cc),
	}, nil
}
