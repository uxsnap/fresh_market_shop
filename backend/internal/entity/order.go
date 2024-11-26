package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	Uid       uuid.UUID
	UserUid   uuid.UUID
	Num       int64
	Sum       int64
	Status    OrderStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderStatus string

const (
	OrderStatusNew        OrderStatus = "new"
	OrderStatusPaid       OrderStatus = "paid"
	OrderStatusInProgress OrderStatus = "in_progress"
	OrderStatusDone       OrderStatus = "done"
)

type OrderProducts struct {
	OrderUid   uuid.UUID
	ProductUid uuid.UUID
	Count      int64
}

type OrderWithProducts struct {
	Order
	Products []OrderProducts
}
