package useCasePayments

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCasePayments) CreatePayment(ctx context.Context, payment entity.Payment) (uuid.UUID, error) {
	log.Printf("usecasePayments.CreatePayment: order uid %s", payment.OrderUid)

	if uuid.Equal(uuid.UUID{}, payment.OrderUid) {
		log.Printf("failed to create payment: order uid is empty")
		return uuid.UUID{}, errors.New("order uid is empty")
	}
	if uuid.Equal(uuid.UUID{}, payment.UserUid) {
		log.Printf("failed to create payment: user uid is empty")
		return uuid.UUID{}, errors.New("user uid is empty")
	}
	if uuid.Equal(uuid.UUID{}, payment.CardUid) {
		log.Printf("failed to create payment: card uid is empty")
		return uuid.UUID{}, errors.New("card uid is empty")
	}
	if payment.Sum == 0 {
		log.Printf("failed to create payment: sum is empty")
		return uuid.UUID{}, errors.New("sum is empty")
	}
	if len(payment.Currency) == 0 {
		log.Printf("failed to create payment: currency is empty")
		return uuid.UUID{}, errors.New("currency is empty")
	}
	if payment.PaymentTime.Unix() == 0 {
		payment.PaymentTime = time.Now().UTC()
	}

	_, userExists, err := uc.usersService.GetUser(ctx, payment.UserUid)
	if err != nil {
		log.Printf("failed to create payment: failed to get user %s", payment.UserUid)
		return uuid.UUID{}, errors.WithStack(err)
	}
	if !userExists {
		log.Printf("failed to create payment: user %s not found", payment.UserUid)
		return uuid.UUID{}, errors.New("user not found")
	}

	card, cardExists, err := uc.GetUserPaymentCardByUid(ctx, payment.CardUid)
	if err != nil {
		log.Printf("failed to create payment: failed to get card %s", payment.CardUid)
		return uuid.UUID{}, errors.WithStack(err)
	}
	if !cardExists {
		log.Printf("failed to create payment: card %s not found", payment.CardUid)
		return uuid.UUID{}, errors.New("card not found")
	}
	if err := validateCardExpired(card.Expired); err != nil {
		log.Printf("failed to create payment: card %s expired is invalid: %v", payment.CardUid, err)
		return uuid.UUID{}, errors.Errorf("card expired is invalid: %v", err)
	}

	payment.Uid = uuid.NewV4()
	if err := uc.paymentsRepository.CreatePayment(ctx, payment); err != nil {
		log.Printf("failed to create payment for order %s: %v", payment.OrderUid, err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	return payment.Uid, nil
}
