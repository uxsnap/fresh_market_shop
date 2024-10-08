package useCaseUsers

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
	UpdateUser(ctx context.Context, user entity.User) error
	GetUserByUid(ctx context.Context, uid uuid.UUID) (entity.User, bool, error)
	DeleteUserByUid(ctx context.Context, uid uuid.UUID) error

	CreateDeliveryAddress(ctx context.Context, address entity.DeliveryAddress) error
	UpdateDeliveryAddress(ctx context.Context, address entity.DeliveryAddress) error
	DeleteDeliveryAddressByUid(ctx context.Context, uid uuid.UUID) error
	DeleteDeliveryAddressesByUserUid(ctx context.Context, userUid uuid.UUID) error
	GetDeliveryAddressByUid(ctx context.Context, uid uuid.UUID) (entity.DeliveryAddress, bool, error)
	GetDeliveryAddressesByUserUid(ctx context.Context, userUid uuid.UUID) ([]entity.DeliveryAddress, error)
}
