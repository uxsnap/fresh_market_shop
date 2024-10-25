package httpEntity

import (
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type OrderProduct struct {
	Uid   uuid.UUID `json:"uid"`
	Count int64     `json:"count"`
}

type OrderProducts struct {
	Products []OrderProduct `json:"products"`
	Sum      int64          `json:"sum"`
}

func OrderProductsToEntity(order OrderProducts) entity.OrderProducts {
	products := make([]entity.OrderProduct, len(order.Products))

	for ind, p := range order.Products {
		products[ind] = entity.OrderProduct{
			Uid:   p.Uid,
			Count: p.Count,
		}
	}

	return entity.OrderProducts{
		Products: products,
	}
}
