package useCaseDelivery

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseDelivery) GetDeliveryByOrderUid(ctx context.Context, orderUid uuid.UUID) (entity.Delivery, bool, error) {
	log.Printf("usecaseDelivery.GetDeliveryByOrderUid: %s", orderUid)

	delivery, isFound, err := uc.deliveryRepository.GetDeliveryByOrderUid(ctx, orderUid)
	if err != nil {
		log.Printf("failed to get delivery by order uid %s: %v", orderUid, err)
		return entity.Delivery{}, false, errors.WithStack(err)
	}
	return delivery, isFound, nil
}
