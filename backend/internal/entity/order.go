package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	Uid       uuid.UUID
	Num       int64
	Sum       int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderProduct struct {
	Uid   uuid.UUID
	Count int64
}

type OrderProducts struct {
	Products []OrderProduct
}
