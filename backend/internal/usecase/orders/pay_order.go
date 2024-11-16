package useCaseOrders

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseOrders) PayOrder(ctx context.Context, userUid uuid.UUID, orderUid uuid.UUID, cardUid uuid.UUID) (uuid.UUID, error) {
	log.Printf("usecasePayments.PayOrder: order uid %s, card uid %s", orderUid, cardUid)

	order, isFound, err := uc.ordersRepository.GetOrderByUid(ctx, orderUid)
	if err != nil {
		log.Printf("failed to get order by uid %s: %v", orderUid, err)
		return uuid.UUID{}, errors.WithStack(err)
	}
	if !isFound {
		log.Printf("order with uid %s not found: %v", orderUid, err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	if !uuid.Equal(userUid, order.UserUid) {
		log.Printf("order.userUid mismatch userUid")
		return uuid.UUID{}, errors.New("cant pay order of another user")
	}

	return uuid.UUID{}, nil
}
