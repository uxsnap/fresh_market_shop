package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	Uid       uuid.UUID
	Num       int64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderProducts struct {
	OrderUid   uuid.UUID
	ProductUid uuid.UUID
	Count      int64
}
