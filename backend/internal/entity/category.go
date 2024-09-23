package entity

import "github.com/google/uuid"

type Category struct {
	Uuid        uuid.UUID
	Name        string
	Description string
}
