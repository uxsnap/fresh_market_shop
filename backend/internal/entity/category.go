package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Category struct {
	Uid        uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
