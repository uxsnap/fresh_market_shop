package httpEntity

import uuid "github.com/satori/go.uuid"

type UUID struct {
	Uid uuid.UUID `json:"uid"`
}
