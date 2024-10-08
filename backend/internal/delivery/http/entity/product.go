package httpEntity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type Product struct {
	Uid         uuid.UUID `json:"uid"`
	CategoryUid uuid.UUID `json:"categoryUid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Ccal        int32     `json:"ccal"`
	Price       int64     `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ProductFromEntity(product entity.Product) Product {
	return Product{
		Uid:         product.Uid,
		CategoryUid: product.CategoryUid,
		Name:        product.Name,
		Description: product.Description,
		Ccal:        product.Ccal,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ProductToEntity(product Product) entity.Product {
	return entity.Product{
		Uid:         product.Uid,
		CategoryUid: product.CategoryUid,
		Name:        product.Name,
		Description: product.Description,
		Ccal:        product.Ccal,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

type ProductWithCount struct {
	Product Product `json:"product"`
	Count   int64   `json:"count"`
}

type CountResponse struct {
	Count int64 `json:"count"`
}

type ProductCount struct {
	ProductUid uuid.UUID `json:"productUid"`
	Count      int64     `json:"count"`
}
