package entity

import uuid "github.com/satori/go.uuid"

type Category struct {
	Uuid        uuid.UUID
	Name        string
	Description string
}
