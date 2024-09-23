package config

import "os"

type ConfigHTTP struct {
	port string
	host string
}

func NewConfigHTTP() *ConfigHTTP {
	return &ConfigHTTP{
		host: os.Getenv("SERVICE_HOST"),
		port: os.Getenv("SERVICE_PORT"),
	}
}

func (c *ConfigHTTP) ServiceHost() string {
	return c.host
}

func (c *ConfigHTTP) ServicePort() string {
	return c.port
}
