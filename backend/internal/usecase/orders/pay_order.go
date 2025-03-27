package useCaseOrders

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseOrders) PayOrder(ctx context.Context, userUid uuid.UUID, orderUid uuid.UUID, cardUid uuid.UUID, deliveryUid uuid.UUID) (uuid.UUID, error) {
	log.Printf("usecasePayments.PayOrder: order uid %s, card uid %s", orderUid, cardUid)

	order, isFound, err := uc.ordersRepository.GetOrder(ctx, entity.QueryFilters{
		OrderUid: orderUid,
	})
	if err != nil {
		log.Printf("failed to get order by uid %s: %v", orderUid, err)
		return uuid.UUID{}, errors.WithStack(err)
	}
	if !isFound {
		log.Printf("order with uid %s not found: %v", orderUid, err)
		return uuid.UUID{}, errors.New("order not found")
	}

	if !uuid.Equal(userUid, order.UserUid) {
		log.Printf("order.userUid mismatch userUid")
		return uuid.UUID{}, errors.New("cant pay order of another user")
	}

	delivery, isFound, err := uc.deliveryService.GetDeliveryByUid(ctx, deliveryUid)
	if err != nil {
		log.Printf("failed to get delivery by uid %s: %v", orderUid, err)
		return uuid.UUID{}, errors.WithStack(err)
	}
	if !isFound {
		log.Printf("delivery with order uid %s not found: %v", orderUid, err)
		return uuid.UUID{}, errors.New("delivery not found")
	}
	delivery.Status = entity.DeliveryStatusNew
	delivery.UpdatedAt = time.Now().UTC()

	payment := entity.Payment{
		UserUid:  userUid,
		OrderUid: orderUid,
		CardUid:  cardUid,
		Sum:      order.Sum + delivery.Price,
		Currency: "RUB",
	}
	order.Status = entity.OrderStatusPaid
	order.Sum = order.Sum + 10 + delivery.Price
	order.UpdatedAt = time.Now().UTC()

	var uid uuid.UUID

	if err := uc.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {
		var err error
		if err = uc.ordersRepository.UpdateOrder(ctx, order); err != nil {
			log.Printf("failed to update order status %s: %v", orderUid, err)
			return errors.WithStack(err)
		}

		if err = uc.deliveryService.UpdateDelivery(ctx, delivery); err != nil {
			log.Printf("failed to update delivery status: %v", err)
			return errors.WithStack(err)
		}

		uid, err = uc.paymentsService.CreatePayment(ctx, payment)
		if err != nil {
			log.Printf("failed to create payment for order %s: %v", orderUid, err)
			return errors.WithStack(err)
		}
		return nil
	}); err != nil {
		return uuid.UUID{}, errors.WithStack(err)
	}

	return uid, err
}
