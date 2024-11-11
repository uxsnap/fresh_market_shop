package httpEntity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type Delivery struct {
	Uid           uuid.UUID `json:"uid"`
	OrderUid      uuid.UUID `json:"orderUid"`
	FromLongitude float64   `json:"fromLongitude"`
	FromLatitude  float64   `json:"fromLatitude"`
	ToLongitude   float64   `json:"toLongitude"`
	ToLatitude    float64   `json:"toLatitude"`
	Address       string    `json:"address"`
	Receiver      string    `json:"receiver"`
	Time          int64     `json:"time"`
	Price         int64     `json:"price"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func DeliveryFromEntity(delivery entity.Delivery) Delivery {
	return Delivery{
		Uid:           delivery.Uid,
		OrderUid:      delivery.OrderUid,
		FromLongitude: delivery.FromLongitude,
		FromLatitude:  delivery.FromLatitude,
		ToLongitude:   delivery.ToLongitude,
		ToLatitude:    delivery.ToLatitude,
		Address:       delivery.Address,
		Receiver:      delivery.Receiver,
		Time:          delivery.Time,
		Price:         delivery.Price,
		CreatedAt:     delivery.CreatedAt,
		UpdatedAt:     delivery.UpdatedAt,
	}
}

func DeliveryToEntity(delivery Delivery) entity.Delivery {
	return entity.Delivery{
		Uid:           delivery.Uid,
		OrderUid:      delivery.OrderUid,
		FromLongitude: delivery.FromLongitude,
		FromLatitude:  delivery.FromLatitude,
		ToLongitude:   delivery.ToLongitude,
		ToLatitude:    delivery.ToLatitude,
		Address:       delivery.Address,
		Receiver:      delivery.Receiver,
		Time:          delivery.Time,
		Price:         delivery.Price,
		CreatedAt:     delivery.CreatedAt,
		UpdatedAt:     delivery.UpdatedAt,
	}
}
