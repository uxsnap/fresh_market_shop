package useCasePayments

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type PaymentsRepository interface {
	CreateUserPaymentCard(ctx context.Context, card entity.UserPaymentCard) error
	GetUserPaymentCardByUid(ctx context.Context, uid uuid.UUID) (entity.UserPaymentCard, bool, error)
	GetUserPaymentCards(ctx context.Context, userUid uuid.UUID) ([]entity.UserPaymentCard, error)
	DeleteUserPaymentCardByUid(ctx context.Context, cardUid uuid.UUID) error
	DeleteUserPaymentCards(ctx context.Context, userUid uuid.UUID) error

	CreatePayment(ctx context.Context, payment entity.Payment) error
	GetPaymentByUid(ctx context.Context, uid uuid.UUID) (entity.Payment, bool, error)
	GetPaymentByOrderUid(ctx context.Context, orderUid uuid.UUID) (entity.Payment, bool, error)
	GetPaymentsByUserUid(ctx context.Context, userUid uuid.UUID) ([]entity.Payment, error)
	GetPaymentsByCardUid(ctx context.Context, cardUid uuid.UUID) ([]entity.Payment, error)
}

type UsersService interface {
	GetUser(ctx context.Context, uid uuid.UUID) (entity.User, bool, error)
}

type OrdersService interface {
	GetOrder(ctx context.Context, orderUid uuid.UUID) (entity.Order, bool, error)
}
