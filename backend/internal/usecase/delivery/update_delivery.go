package useCaseDelivery

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseDelivery) UpdateDelivery(ctx context.Context, delivery entity.Delivery) error {
	log.Printf("usecaseDelivery.CreateDelivery: order uid %s", delivery.OrderUid)

	if uuid.Equal(delivery.Uid, uuid.UUID{}) {
		log.Printf("invalid delivery: uid is empty")
		return errors.New("uid is empty")
	}
	if delivery.FromLatitude == 0.0 || delivery.FromLongitude == 0.0 {
		log.Printf("invalid delivery: invalid store coordinates")
		return errors.New("invalid store coordinates")
	}
	if delivery.ToLatitude == 0.0 || delivery.ToLongitude == 0.0 {
		log.Printf("invalid delivery: invalid delivery address coordinates")
		return errors.New("invalid delivery address coordinates")
	}
	if len(delivery.Address) == 0 {
		log.Printf("invalid delivery: invalid address description")
		return errors.New("invalid address description")
	}
	if len(delivery.Receiver) == 0 {
		log.Printf("invalid delivery: invalid delivery receiver description")
		return errors.New("invalid delivery receiver description")
	}

	savedDelivery, isFound, err := uc.deliveryRepository.GetDeliveryByUid(ctx, delivery.Uid)
	if err != nil {
		log.Printf("failed to update delivery: failed to get delivery by uid %s: %v", delivery.Uid, err)
		return errors.WithStack(err)
	}

	if !isFound {
		log.Printf("failed to update delivery: delivery by uid %s not found", delivery.Uid)
		return errors.New("delivery not found")
	}

	// добавить сравнение через дельту
	if delivery.FromLatitude != savedDelivery.FromLatitude {
		log.Printf("failed to update delivery: cant update field `from_latitude`")
		return errors.New("cant update field `from_latitude`")
	}
	if delivery.FromLongitude != savedDelivery.FromLongitude {
		log.Printf("failed to update delivery: cant update field `from_longitude`")
		return errors.New("cant update field `from_longitude`")
	}
	if !uuid.Equal(delivery.OrderUid, savedDelivery.OrderUid) {
		log.Printf("failed to update delivery: cant update field `order_uid`")
		return errors.New("cant update field `order_uid`")
	}
	if delivery.Price <= 0 {
		log.Printf("failed to update delivery: invalid delivery price")
		return errors.New("invalid delivery price")
	}

	if delivery.ToLatitude != savedDelivery.ToLatitude {
		savedDelivery.ToLatitude = delivery.ToLatitude
	}
	if delivery.ToLongitude != savedDelivery.ToLongitude {
		savedDelivery.ToLongitude = delivery.ToLongitude
	}
	if delivery.Address != savedDelivery.Address {
		savedDelivery.Address = delivery.Address
	}
	if delivery.Receiver != savedDelivery.Receiver {
		savedDelivery.Receiver = delivery.Receiver
	}
	if delivery.Price != savedDelivery.Price {
		savedDelivery.Price = delivery.Price
	}
	if delivery.Status != savedDelivery.Status {
		savedDelivery.Status = delivery.Status
	}

	savedDelivery.UpdatedAt = time.Now().UTC()

	if err := uc.deliveryRepository.UpdateDelivery(ctx, savedDelivery); err != nil {
		log.Printf("failed to update delivery %s: %v", delivery.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}
