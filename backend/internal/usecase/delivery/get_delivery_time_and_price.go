package useCaseDelivery

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseDelivery) GetDeliveryTimeAndPriceForOrder(ctx context.Context, orderUid uuid.UUID) (deliveryTime int64, deliveryPrice int64, err error) {
	log.Printf("usecaseDelivery.GetDeliveryTimeAndPriceForOrder: %s", orderUid)

	delivery, isFound, err := uc.deliveryRepository.GetDeliveryByOrderUid(ctx, orderUid)
	if err != nil {
		log.Printf("failed to get delivery time and price for order: failed to get delivery by order uid %s: %v", orderUid, err)
		return 0, 0, errors.WithStack(err)
	}
	if isFound {
		return delivery.Time, delivery.Price, nil
	}

	//TODO: получать заказ и считать для него
	log.Printf("failed to get delivery time and price for order: delivery for order %s not found", orderUid)
	return 0, 0, errors.Errorf("delivery for order %s not found", orderUid)
}
