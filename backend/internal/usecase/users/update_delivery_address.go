package useCaseUsers

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseUsers) UpdateDeliveryAddress(ctx context.Context, address entity.DeliveryAddress) error {
	log.Printf(
		"ucUsers.UpdateDeliveryAddress: user uid %s, address: %s %s %d",
		address.UserUid, address.CityName, address.StreetName, address.Building,
	)

	if err := validateDeliveryAddress(address); err != nil {
		log.Printf("failed to update delivery address %s: %v", address.Uid, err)
		return errors.WithStack(err)
	}

	savedAddress, isFound, err := uc.usersRepository.GetDeliveryAddressByUid(ctx, address.Uid)
	if err != nil {
		log.Printf("failed to update delivery address %s: %v", address.Uid, err)
		return errors.WithStack(err)
	}
	if !isFound {
		log.Printf("failed to update delivery address %s: address not found", address.Uid)
		return errors.WithStack(err)
	}

	address.CreatedAt = savedAddress.CreatedAt
	address.UpdatedAt = time.Now().UTC()

	if err := uc.usersRepository.UpdateDeliveryAddress(ctx, address); err != nil {
		log.Printf("failed to update delivery address %s: %v", address.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}
