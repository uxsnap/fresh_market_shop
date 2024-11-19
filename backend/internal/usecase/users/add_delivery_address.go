package useCaseUsers

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseUsers) AddDeliveryAddress(ctx context.Context, address entity.DeliveryAddress) (uuid.UUID, error) {
	log.Printf(
		"ucUsers.AddDeliveryAddress: user uid %s, address: %s %s",
		address.UserUid, address.CityName, address.StreetName,
	)

	if err := validateDeliveryAddress(address); err != nil {
		log.Printf("failed to add delivery address for user %s: %v", address.UserUid, err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	_, userIsFound, err := uc.usersRepository.GetUserByUid(ctx, address.UserUid)
	if err != nil {
		log.Printf("failed to add delivery address for user %s: %v", address.UserUid, err)
		return uuid.UUID{}, errors.WithStack(err)
	}
	if !userIsFound {
		log.Printf("failed to add delivery address. user %s not found", address.UserUid)
		return uuid.UUID{}, errors.New("user not found")
	}

	address.Uid = uuid.NewV4()
	address.CreatedAt = time.Now().UTC()

	if err := uc.usersRepository.CreateDeliveryAddress(ctx, address); err != nil {
		log.Printf("failed to add delivery address for user %s: %v", address.UserUid, err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	return address.Uid, nil
}
