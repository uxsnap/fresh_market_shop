package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Product struct {
	Uid           uuid.UUID
	CategoryUid   uuid.UUID
	Name          string
	Description   string
	Ccal          int32
	Price         int64
	StockQuantity int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ProductWithStockQuantity struct {
	Product
	StockQuantity int64
}
