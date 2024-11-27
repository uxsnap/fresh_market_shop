package httpEntity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type Order struct {
	Uid       uuid.UUID `json:"uid"`
	UserUid   uuid.UUID `json:"userUid"`
	Num       int64     `json:"num"`
	Sum       int64     `json:"sum"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func OrderFromEntity(order entity.Order) Order {
	return Order{
		Uid:       order.Uid,
		UserUid:   order.UserUid,
		Num:       order.Num,
		Sum:       order.Sum,
		Status:    string(order.Status),
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}

func OrderToEntity(order Order) entity.Order {
	return entity.Order{
		Uid:       order.Uid,
		UserUid:   order.UserUid,
		Num:       order.Num,
		Sum:       order.Sum,
		Status:    entity.OrderStatus(order.Status),
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}

type OrderProducts struct {
	OrderUid   uuid.UUID      `json:"orderUid"`
	ProductUid uuid.UUID      `json:"productUid"`
	Count      int64          `json:"count"`
	Photos     []ProductPhoto `json:"photos"`
}

func OrderProductsFromEntity(op entity.OrderProducts) OrderProducts {
	return OrderProducts{
		OrderUid:   op.OrderUid,
		ProductUid: op.ProductUid,
		Count:      op.Count,
		Photos:     ProductPhotosFromEntity(op.Photos),
	}
}

type OrderWithProducts struct {
	Order    `json:"order"`
	Products []OrderProducts `json:"products"`
}

func OrderWithProductsFromEntity(owp entity.OrderWithProducts) OrderWithProducts {
	res := OrderWithProducts{
		Order:    OrderFromEntity(owp.Order),
		Products: make([]OrderProducts, len(owp.Products)),
	}

	for ind, v := range owp.Products {
		res.Products[ind] = OrderProductsFromEntity(v)
	}

	return res
}
