package httpEntity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type UserPaymentCard struct {
	Uid         uuid.UUID `json:"uid"`
	UserUid     uuid.UUID `json:"userUid"`
	ExternalUid uuid.UUID `json:"externalUid"`
	Number      string    `json:"number"`
	Expired     string    `json:"expired"`
}

func UserPaymentCardFromEntity(card entity.UserPaymentCard) UserPaymentCard {
	return UserPaymentCard{
		Uid:         card.Uid,
		UserUid:     card.UserUid,
		ExternalUid: card.ExternalUid,
		Number:      card.Number,
		Expired:     card.Expired,
	}
}

func UserPaymentCardToEntity(card UserPaymentCard) entity.UserPaymentCard {
	return entity.UserPaymentCard{
		Uid:         card.Uid,
		UserUid:     card.UserUid,
		ExternalUid: card.ExternalUid,
		Number:      card.Number,
		Expired:     card.Expired,
	}
}

type UserFullPaymentCard struct {
	Uid     uuid.UUID `json:"uid"`
	UserUid uuid.UUID `json:"userUid"`
	Number  string    `json:"number"`
	Expired string    `json:"expired"`
	CVV     string    `json:"cvv"`
}

func UserFullPaymentCardFromEntity(card entity.UserFullPaymentCard) UserFullPaymentCard {
	return UserFullPaymentCard{
		Uid:     card.Uid,
		UserUid: card.UserUid,
		Number:  card.Number,
		Expired: card.Expired,
		CVV:     card.CVV,
	}
}

func UserFullPaymentCardToEntity(card UserFullPaymentCard) entity.UserFullPaymentCard {
	return entity.UserFullPaymentCard{
		Uid:     card.Uid,
		UserUid: card.UserUid,
		Number:  card.Number,
		Expired: card.Expired,
		CVV:     card.CVV,
	}
}

type Payment struct {
	Uid         uuid.UUID `json:"uid"`
	UserUid     uuid.UUID `json:"userUid"`
	OrderUid    uuid.UUID `json:"orderUid"`
	CardUid     uuid.UUID `json:"cardUid"`
	Sum         int64     `json:"sum"`
	Currency    string    `json:"currency"`
	PaymentTime time.Time `json:"time"`
}

func PaymentFromEntity(payment entity.Payment) Payment {
	return Payment{
		Uid:         payment.Uid,
		UserUid:     payment.UserUid,
		OrderUid:    payment.OrderUid,
		CardUid:     payment.CardUid,
		Sum:         payment.Sum,
		Currency:    payment.Currency,
		PaymentTime: payment.PaymentTime,
	}
}

func PaymentToEntity(payment Payment) entity.Payment {
	return entity.Payment{
		Uid:         payment.Uid,
		UserUid:     payment.UserUid,
		OrderUid:    payment.OrderUid,
		CardUid:     payment.CardUid,
		Sum:         payment.Sum,
		Currency:    payment.Currency,
		PaymentTime: payment.PaymentTime,
	}
}
