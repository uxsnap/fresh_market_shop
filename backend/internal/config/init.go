package config

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func Init(path string) error {
	return errors.WithStack(godotenv.Load(path))
}
