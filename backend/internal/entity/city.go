package entity

import (
	uuid "github.com/satori/go.uuid"
)

type City struct {
	Uid  uuid.UUID
	Name string
}
