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
	Weight      int32     `json:"weight"`
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
		Weight:      product.Weight,
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
		Weight:      product.Weight,
	}
}

type ProductsWithExtra struct {
	Products []ProductWithExtra `json:"products"`
	Total    int64              `json:"total"`
}

type ProductWithExtra struct {
	Product Product        `json:"product"`
	Count   int64          `json:"count,omitempty"`
	Photos  []ProductPhoto `json:"photos,omitempty"`
}

type ProductPhoto struct {
	Uid  uuid.UUID `json:"uid"`
	Path string    `json:"path"`
}

func ProductPhotosFromEntity(photos []entity.ProductPhoto) []ProductPhoto {
	res := make([]ProductPhoto, len(photos))
	for i := 0; i < len(photos); i++ {
		res[i] = ProductPhoto{
			Uid:  photos[i].Uid,
			Path: photos[i].FilePath,
		}
	}
	return res
}

type CountResponse struct {
	Count int64 `json:"count"`
}

type ProductCount struct {
	ProductUid uuid.UUID `json:"productUid"`
	Count      int64     `json:"count"`
}

type ProductsCounts struct {
	Products []ProductCount `json:"products"`
}

func ProductsCountsToEntity(order ProductsCounts) entity.ProductsCounts {
	products := make([]entity.ProductCount, len(order.Products))

	for ind, p := range order.Products {
		products[ind] = entity.ProductCount{
			ProductUid: p.ProductUid,
			Count:      p.Count,
		}
	}

	return entity.ProductsCounts{
		Products: products,
	}
}
