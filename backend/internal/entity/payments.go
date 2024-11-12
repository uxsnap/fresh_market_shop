package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserPaymentCard struct {
	Uid         uuid.UUID
	UserUid     uuid.UUID
	ExternalUid uuid.UUID
	Number      string
	Expired     string
}

type UserFullPaymentCard struct {
	Uid     uuid.UUID
	UserUid uuid.UUID
	Number  string
	Expired string
	CVV     string
}

type Payment struct {
	Uid         uuid.UUID
	UserUid     uuid.UUID
	OrderUid    uuid.UUID
	CardUid     uuid.UUID
	Sum         int64
	Currency    string
	PaymentTime time.Time
}
