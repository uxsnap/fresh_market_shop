package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Product struct {
	Uid         uuid.UUID
	CategoryUid uuid.UUID
	Name        string
	Description string
	Ccal        int32
	Price       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Weight      int32
}

type ProductsWithExtra struct {
	Products []ProductWithExtra
	Total    int64
}

type ProductWithExtra struct {
	Product
	StockQuantity int64
	Photos        []ProductPhoto
}

type ProductPhoto struct {
	Uid        uuid.UUID
	ProductUid uuid.UUID
	FilePath   string
}

type ProductCount struct {
	ProductUid uuid.UUID
	Count      int64
}

type ProductsCounts struct {
	Products []ProductCount
}
