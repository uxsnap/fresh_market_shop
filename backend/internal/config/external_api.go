package config

import "os"

type ConfigExternalApi struct {
	authServiceGrpcHost string
	authServiceGrpcPort string
}

func New() *ConfigExternalApi {
	return &ConfigExternalApi{
		authServiceGrpcHost: os.Getenv("AUTH_SERVICE_GRPC_HOST"),
		authServiceGrpcPort: os.Getenv("AUTH_SERVICE_GRPC_PORT"),
	}
}
