package config

import "os"

type ConfigExternalApi struct {
	authServiceGrpcHost string
	authServiceGrpcPort string
}

func NewConfigExternalApi() *ConfigExternalApi {
	return &ConfigExternalApi{
		authServiceGrpcHost: os.Getenv("AUTH_SERVICE_GRPC_HOST"),
		authServiceGrpcPort: os.Getenv("AUTH_SERVICE_GRPC_PORT"),
	}
}

func (c *ConfigExternalApi) AuthServiceGrpcHost() string {
	return c.authServiceGrpcHost
}

func (c *ConfigExternalApi) AuthServiceGrpcPort() string {
	return c.authServiceGrpcPort
}
