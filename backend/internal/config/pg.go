package config

import (
	"fmt"
	"os"
)

type ConfigPG struct {
	dbName     string
	pgUser     string
	pgPassword string
	pgPort     string
	pgHost     string
}

func NewConfigPG() *ConfigPG {
	return &ConfigPG{
		dbName:     os.Getenv("PG_DB"),
		pgUser:     os.Getenv("PG_USER"),
		pgPassword: os.Getenv("PG_PASSWORD"),
		pgPort:     os.Getenv("PG_PORT"),
		pgHost:     os.Getenv("PG_HOST"),
	}
}

func (c *ConfigPG) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		c.pgHost, c.pgPort, c.dbName, c.pgUser, c.pgPassword,
	)
}
