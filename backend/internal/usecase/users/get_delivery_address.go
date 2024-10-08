package useCaseUsers

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseUsers) GetDeliveryAddress(ctx context.Context, uid uuid.UUID) (entity.DeliveryAddress, bool, error) {
	log.Printf("ucUsers.GetDeliveryAddress: uid %s", uid)

	address, isFound, err := uc.usersRepository.GetDeliveryAddressByUid(ctx, uid)
	if err != nil {
		log.Printf("failed to get delivery address by uid %s: %v", uid, err)
		return entity.DeliveryAddress{}, false, errors.WithStack(err)
	}

	return address, isFound, nil
}

func (uc *UseCaseUsers) GetUserDeliveryAddresses(ctx context.Context, userUid uuid.UUID) ([]entity.DeliveryAddress, error) {
	log.Printf("ucUsers.GetUserDeliveryAddresses: user uid %s", userUid)

	addresses, err := uc.usersRepository.GetDeliveryAddressesByUserUid(ctx, userUid)
	if err != nil {
		log.Printf("failed to get delivery addresses by user uid %s: %v", userUid, err)
		return []entity.DeliveryAddress{}, errors.WithStack(err)
	}

	return addresses, nil
}
