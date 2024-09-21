package config

type HttpConfig struct {
	httpPort string
	httpHost string
}

func NewHttpConfig() *HttpConfig {
	return &HttpConfig{}
}

func (c *HttpConfig) HttpHost() string {
	return ""
}

func (c *HttpConfig) HttpPort() string {
	return "8080"
}
