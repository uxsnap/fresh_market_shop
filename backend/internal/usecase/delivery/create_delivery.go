package useCaseDelivery

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseDelivery) CreateDelivery(ctx context.Context, delivery entity.Delivery, orderPrice int64) (uuid.UUID, error) {
	log.Printf("usecaseDelivery.CreateDelivery: order uid %s", delivery.OrderUid)

	if uuid.Equal(delivery.OrderUid, uuid.UUID{}) {
		log.Printf("invalid delivery: order uid is empty")
		return uuid.UUID{}, errors.New("order uid is empty")
	}
	if delivery.ToLatitude == 0.0 || delivery.ToLongitude == 0.0 {
		log.Printf("invalid delivery: invalid delivery address coordinates")
		return uuid.UUID{}, errors.New("invalid delivery address coordinates")
	}
	if len(delivery.Address) == 0 {
		log.Printf("invalid delivery: invalid address description")
		return uuid.UUID{}, errors.New("invalid address description")
	}
	if len(delivery.Receiver) == 0 {
		log.Printf("invalid delivery: invalid delivery receiver description")
		return uuid.UUID{}, errors.New("invalid delivery receiver description")
	}

	deliveryPrice, deliveryTime, err := uc.CalculateDelivery(ctx, delivery.OrderUid, orderPrice, delivery.ToLongitude, delivery.ToLatitude)
	if err != nil {
		log.Printf("failed to calculate delivery: %v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	delivery.Uid = uuid.NewV4()
	delivery.Price = deliveryPrice
	delivery.Time = int64(deliveryTime)
	delivery.FromLatitude = fromLatitude
	delivery.FromLongitude = fromLongitude
	delivery.CreatedAt = time.Now().UTC()

	if err := uc.deliveryRepository.CreateDelivery(ctx, delivery); err != nil {
		log.Printf("failed to create delivery: %v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	return delivery.Uid, nil
}
