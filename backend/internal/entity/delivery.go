package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Delivery struct {
	Uid           uuid.UUID
	OrderUid      uuid.UUID
	FromLongitude float64
	FromLatitude  float64
	ToLongitude   float64
	ToLatitude    float64
	Address       string
	Receiver      string
	Time          int64
	Price         int64
	Status        DeliveryStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type DeliveryStatus string

const (
	DeliveryStatusCalculated DeliveryStatus = "calculated"
	DeliveryStatusNew        DeliveryStatus = "new"
	DeliveryStatusInProgress DeliveryStatus = "in_progress"
	DeliveryStatusDeliveried DeliveryStatus = "deliveried"
	DeliveryStatusFailed     DeliveryStatus = "failed"
)
